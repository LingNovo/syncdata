package logic

import (
	"encoding/json"
	"strings"

	oapi "github.com/LingNovo/dingtalk/protos/oapi"
	"github.com/LingNovo/syncdata/guide"
	m "github.com/LingNovo/syncdata/models"
	empty "github.com/golang/protobuf/ptypes/empty"
	ctx "golang.org/x/net/context"
	"google.golang.org/grpc"
)

// 从钉钉平台同步用户列表信息
func SyncUserList(deptId int64, deptShortName string) error {
	address := guide.Conf.MustValue("server", "address")
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		GetLogEntry().Fatalf("grpc conn: %s", err)
		return err
	}
	defer cc.Close()
	owc := oapi.NewOapiWarpperClient(cc)
	if _, err := owc.RefreshCompanyAccessToken(ctx.Background(), &empty.Empty{}); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	req := &oapi.UserListRequest{}
	req.DepartmentId = deptId
	req.Offset = 0
	req.Size = 100
	resp, err := owc.UserList(ctx.Background(), req)
	if err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	// 推送钉钉用户信息到目标库用户表
	if err := PushUserList(resp.Userlist, deptShortName); err != nil {
		return err
	}
	// 如果还有数据则递归
	if resp.HasMore {
		req.Offset++
		if err := SyncUserList(req.DepartmentId, deptShortName); err != nil {
			return err
		}
	}
	return nil
}

// 推送钉钉用户信息到目标库用户表
func PushUserList(uls []*oapi.UDetailedList, deptShortName string) error {
	session := guide.GetDB().NewSession()
	defer session.Close()
	dus := make([]*m.Ding_User, 0)
	for _, ul := range uls {
		du, err := ConvertUser(ul)
		if err != nil {
			return err
		}
		du.DeptShortName = deptShortName
		cond := m.Ding_User{Userid: ul.Userid}
		ok, err := session.Get(&cond)
		if err != nil {
			return err
		}
		if !ok {
			dus = append(dus, du)
		}
	}
	if err := session.Begin(); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	if _, err := session.InsertMulti(&dus); err != nil {
		session.Rollback()
		GetLogEntry().Errorln(err)
		return err
	}
	if err := session.Commit(); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	return nil
}

// 将钉钉用户信息转换为目标库用户信息
func ConvertUser(ul *oapi.UDetailedList) (*m.Ding_User, error) {
	var resp m.Ding_User
	data, err := json.Marshal(ul)
	if err != nil {
		GetLogEntry().Errorln(err)
		return nil, err
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		GetLogEntry().Errorln(err)
		return nil, err
	}
	return &resp, nil
}

func GetUserCount() int64 {
	address := guide.Conf.MustValue("server", "address")
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		GetLogEntry().Fatalf("grpc conn: %s", err)
		return 0
	}
	defer cc.Close()
	owc := oapi.NewOapiWarpperClient(cc)
	if _, err := owc.RefreshCompanyAccessToken(ctx.Background(), &empty.Empty{}); err != nil {
		GetLogEntry().Errorln(err)
		return 0
	}
	req := &oapi.UserGetOrgUserCountRequest{}
	req.OnlyActive = 0
	resp, err := owc.UserGetOrgUserCount(ctx.Background(), req)
	if err != nil {
		GetLogEntry().Errorln(err)
		return 0
	}
	return resp.Count
}

func GetUserInfo(userId string) *oapi.UserInfoByUserIdResponse {
	address := guide.Conf.MustValue("server", "address")
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		GetLogEntry().Fatalf("grpc conn: %s", err)
		return nil
	}
	defer cc.Close()
	owc := oapi.NewOapiWarpperClient(cc)
	if _, err := owc.RefreshCompanyAccessToken(ctx.Background(), &empty.Empty{}); err != nil {
		GetLogEntry().Errorln(err)
		return nil
	}
	req := &oapi.UserInfoByUserIdRequest{}
	req.Userid = userId
	resp, err := owc.UserInfoByUserId(ctx.Background(), req)
	if err != nil {
		GetLogEntry().Errorln(err)
		return nil
	}
	return resp
}

func UserC() int64 {
	session := guide.GetDB().NewSession()
	defer session.Close()
	var su m.Sys_User
	count, err := session.Count(&su)
	if err != nil {
		GetLogEntry().Errorln(err)
		return 0
	}
	return count
}

func StringArrayReverse(l []string) {
	for i := 0; i < int(len(l)/2); i++ {
		li := len(l) - i - 1
		l[i], l[li] = l[li], l[i]
	}
}

func PaddingDeptShortName(deptId string) error {
	session := guide.GetDB().NewSession()
	defer session.Close()
	cond := m.Sys_Organize{}
	cond.F_Id = deptId
	ok, err := session.Get(&cond)
	if err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	if !ok {
		return nil
	}
	shortName := cond.F_FullName
	if strings.TrimSpace(cond.F_ParentId) != "0" {
		array := make([]string, 0)
		if array, err = GetSysDeptDepth(cond.F_Id, array); err != nil {
			GetLogEntry().Errorln(err)
			return err
		}
		StringArrayReverse(array)
		shortName = strings.Join(array, "-")
	}
	session = guide.GetDB().NewSession()
	defer session.Close()
	up_dept_cond := m.Sys_Organize{}
	up_dept_cond.F_Id = deptId
	if err := session.Begin(); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	up_dept := &m.Sys_Organize{}
	up_dept.F_ShortName = shortName
	if _, err := session.Update(up_dept, &up_dept_cond); err != nil {
		session.Rollback()
		GetLogEntry().Errorln(err)
		return err
	}
	if err := session.Commit(); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	session = guide.GetDB().NewSession()
	defer session.Close()
	user_cond := m.Sys_User{}
	user_cond.F_DepartmentId = cond.F_Id
	user_list := make([]*m.Sys_User, 0)
	if err := session.Find(&user_list, &user_cond); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	if err := session.Begin(); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	for _, user := range user_list {
		up_user_cond := m.Sys_User{}
		up_user_cond.F_Id = user.F_Id
		up_user := m.Sys_User{}
		up_user.F_DeptShortName = shortName
		if _, err := session.Update(&up_user, &up_user_cond); err != nil {
			session.Rollback()
			GetLogEntry().Errorln(err)
			return err
		}
	}
	if err := session.Commit(); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	session = guide.GetDB().NewSession()
	defer session.Close()
	list_cond := m.Sys_Organize{}
	list_cond.F_ParentId = cond.F_Id
	depts := make([]*m.Sys_Organize, 0)
	if err := session.Find(&depts, &list_cond); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	for _, dept := range depts {
		PaddingDeptShortName(dept.F_Id)
	}
	return nil
}

func GetSysDeptDepth(id string, array []string) ([]string, error) {
	session := guide.GetDB().NewSession()
	defer session.Close()
	cond := m.Sys_Organize{}
	cond.F_Id = id
	ok, err := session.Get(&cond)
	if err != nil {
		GetLogEntry().Errorln(err)
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	array = append(array, cond.F_FullName)
	if strings.TrimSpace(cond.F_ParentId) != "0" {
		return GetSysDeptDepth(cond.F_ParentId, array)
	}
	return array, nil
}
