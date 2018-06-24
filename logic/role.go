package logic

import (
	"strings"

	oapi "github.com/LingNovo/dingtalk/protos/oapi"
	tapi "github.com/LingNovo/dingtalk/protos/tapi"
	"github.com/LingNovo/syncdata/guide"
	m "github.com/LingNovo/syncdata/models"
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/satori/go.uuid"
	ctx "golang.org/x/net/context"
	"google.golang.org/grpc"
)

// 从钉钉平台同步角色列表信息
func SyncRoleList() error {
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
	twc := tapi.NewTapiWarpperClient(cc)
	req := &tapi.RoleListRequest{}
	req.Offset = 1
	req.Size = 100
	resp, err := twc.RoleList(ctx.Background(), req)
	if err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	// 推送钉钉角色信息到目标库角色表
	drs, err := PushRoleList(resp.DingtalkCorpRoleListResponse.Result.List.RoleGroups)
	if err != nil {
		return err
	}
	for _, dr := range drs {
		if err := SyncUserRoles(dr); err != nil {
			return err
		}
	}
	if strings.ToLower(resp.DingtalkCorpRoleListResponse.Result.HasMore) == "true" {
		req.Offset++
		if err := SyncRoleList(); err != nil {
			return err
		}
	}
	return nil
}

// 推送钉钉角色信息到目标库角色表
func PushRoleList(rls []*tapi.RoleListResultListRoleGroups) ([]*m.Ding_Role, error) {
	session := guide.GetDB().NewSession()
	defer session.Close()
	drs := make([]*m.Ding_Role, 0)
	for _, rl := range rls {
		for _, rb := range rl.Roles.Roles {
			dr, err := ConvertRole(rb.Id)
			if err != nil {
				return nil, err
			}
			if dr != nil {
				cond := m.Ding_Role{RoleId: dr.RoleId}
				ok, err := session.Get(&cond)
				if err != nil {
					return nil, err
				}
				if !ok {
					drs = append(drs, dr)
				}
			}
		}
	}
	if err := session.Begin(); err != nil {
		GetLogEntry().Errorln(err)
		return nil, err
	}
	if _, err := session.InsertMulti(drs); err != nil {
		session.Rollback()
		GetLogEntry().Errorln(err)
		return nil, err
	}
	if err := session.Commit(); err != nil {
		GetLogEntry().Errorln(err)
		return nil, err
	}
	return drs, nil
}

// 将钉钉角色信息转换为目标库角色信息
func ConvertRole(roleId int64) (*m.Ding_Role, error) {
	address := guide.Conf.MustValue("server", "address")
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		GetLogEntry().Fatalf("grpc conn: %s", err)
		return nil, err
	}
	defer cc.Close()
	owc := oapi.NewOapiWarpperClient(cc)
	if _, err := owc.RefreshCompanyAccessToken(ctx.Background(), &empty.Empty{}); err != nil {
		GetLogEntry().Errorln(err)
		return nil, err
	}
	twc := tapi.NewTapiWarpperClient(cc)
	req := &tapi.GetRoleRequest{}
	req.RoleId = roleId
	resp, err := twc.GetRole(ctx.Background(), req)
	if err != nil {
		GetLogEntry().Errorln(err)
		return nil, err
	}
	gReq := &tapi.GetRoleGroupRequest{}
	gReq.GroupId = resp.DingtalkOapiRoleGetroleResponse.Role.GroupId
	gResp, err := twc.GetRoleGroup(ctx.Background(), gReq)
	if err != nil {
		GetLogEntry().Errorln(err)
		return nil, err
	}
	role := &m.Ding_Role{}
	role.RoleId = roleId
	role.RoleName = resp.DingtalkOapiRoleGetroleResponse.Role.Name
	role.GroupId = resp.DingtalkOapiRoleGetroleResponse.Role.GroupId
	role.GroupName = gResp.DingtalkCorpRoleGetrolegroupResponse.Result.RoleGroup.GroupName
	return role, nil
}

// 从钉钉平台同步用用户角色表信息
func SyncUserRoles(role *m.Ding_Role) error {
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
	twc := tapi.NewTapiWarpperClient(cc)
	req := &tapi.RoleSimplelistRequest{}
	req.RoleId = role.RoleId
	req.Offset = 1
	req.Size = 100
	resp, err := twc.RoleSimplelist(ctx.Background(), req)
	if err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	// 推送钉钉用户角色信息到目标库用户角色表
	simpleList := resp.DingtalkCorpRoleSimplelistResponse.Result.List.EmpSimpleList
	if err := PushUserRoles(simpleList, role); err != nil {
		return err
	}
	if resp.DingtalkCorpRoleSimplelistResponse.Result.HasMore {
		req.Offset++
		if err := SyncUserRoles(role); err != nil {
			return err
		}
	}
	return nil
}

// 推送钉钉用户角色信息到目标库用户角色表
func PushUserRoles(uls []*tapi.RoleSimplelistResponseResultListEmpSimpleList, role *m.Ding_Role) error {
	session := guide.GetDB().NewSession()
	defer session.Close()
	urs := make([]*m.Ding_UserRole, 0)
	for _, ul := range uls {
		cond := m.Ding_UserRole{
			UserId: ul.Userid,
			RoleId: role.RoleId,
		}
		ok, err := session.Get(&cond)
		if err != nil {
			return err
		}
		if !ok {
			uid, _ := uuid.NewV4()
			urs = append(urs, &m.Ding_UserRole{
				Id:        uid.String(),
				UserId:    ul.Userid,
				RoleId:    role.RoleId,
				RoleName:  role.RoleName,
				GroupId:   role.GroupId,
				GroupName: role.GroupName,
			})
		}
	}
	if err := session.Begin(); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	if _, err := session.InsertMulti(&urs); err != nil {
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
