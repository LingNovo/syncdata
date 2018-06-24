package logic

import (
	"encoding/json"
	"sort"
	"strings"
	"time"

	oapi "github.com/LingNovo/dingtalk/protos/oapi"
	"github.com/LingNovo/syncdata/guide"
	m "github.com/LingNovo/syncdata/models"
	empty "github.com/golang/protobuf/ptypes/empty"
	ctx "golang.org/x/net/context"
	"google.golang.org/grpc"
)

// 从钉钉平台同步部门列表信息
func SyncDepartmentList(deptId int64) error {
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
	req := &oapi.DepartmentDetailRequest{}
	req.Id = deptId
	resp, err := owc.DepartmentDetail(ctx.Background(), req)
	if err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	shortName := resp.Name
	if deptId > 1 {
		shortName, err = GetDepartmentDepth(deptId)
		if err != nil {
			GetLogEntry().Errorln(err)
			return err
		}
	}
	if err := PushDepartment(resp, shortName); err != nil {
		return err
	}
	if err := SyncUserList(deptId, shortName); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	subReq := &oapi.SubDepartmentListRequest{}
	subReq.Id = deptId
	subResp, err := owc.SubDepartmentList(ctx.Background(), subReq)
	if err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	if len(subResp.SubDeptIdList) > 0 {
		time.Sleep(1 * time.Second)
		for _, subId := range subResp.SubDeptIdList {
			if err := SyncDepartmentList(subId); err != nil {
				return err
			}
		}
	}
	return nil
}

func GetDepartmentDepth(deptId int64) (string, error) {
	address := guide.Conf.MustValue("server", "address")
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		GetLogEntry().Fatalf("grpc conn: %s", err)
		return "", err
	}
	defer cc.Close()
	owc := oapi.NewOapiWarpperClient(cc)
	if _, err := owc.RefreshCompanyAccessToken(ctx.Background(), &empty.Empty{}); err != nil {
		GetLogEntry().Errorln(err)
		return "", err
	}
	req := &oapi.DepartmentListParentDeptsByDeptRequest{}
	req.Id = deptId
	resp, err := owc.DepartmentListParentDeptsByDept(ctx.Background(), req)
	if err != nil {
		GetLogEntry().Errorln(err)
		return "", err
	}
	ids := resp.ParentIds
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
	array := make([]string, 0)
	for _, id := range ids {
		detailReq := &oapi.DepartmentDetailRequest{}
		detailReq.Id = id
		detailResp, err := owc.DepartmentDetail(ctx.Background(), detailReq)
		if err != nil {
			GetLogEntry().Errorln(err)
			return "", err
		}
		array = append(array, detailResp.Name)
	}
	return strings.Join(array, "-"), nil
}

// 推送钉钉部门信息到目标库部门表
func PushDepartment(ddr *oapi.DepartmentDetailResponse, deptShortName string) error {
	session := guide.GetDB().NewSession()
	defer session.Close()
	dept, err := ConvertDepartment(ddr)
	if err != nil {
		return err
	}
	dept.ShortName = deptShortName
	cond := m.Ding_Department{Id: dept.Id}
	ok, err := session.Get(&cond)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	if err := session.Begin(); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	if _, err := session.Insert(dept); err != nil {
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

// 将钉钉部门信息转换为目标库部门信息
func ConvertDepartment(ddr *oapi.DepartmentDetailResponse) (*m.Ding_Department, error) {
	var resp m.Ding_Department
	data, err := json.Marshal(ddr)
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
