package models

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/LingNovo/syncdata/guide"
	"github.com/satori/go.uuid"
)

// 钉钉用户信息
type Ding_User struct {
	// 员工唯一标识ID（不可修改）
	Userid string `json:"userid" xorm:"Varchar(255) pk notnull"`
	// 表示人员在此部门中的排序，列表是按order的倒序排列输出的，即从大到小排列输出的（OA后台里面调整了顺序的话order才有值）
	Order int64 `json:"order" xorm:"BigInt"`
	// 在当前isv全局范围内唯一标识一个用户的身份,用户无法修改
	Unionid string `json:"unionid" xorm:"Varchar(255)`
	// 手机号（ISV不可见）
	Mobile string `json:"mobile" xorm:"Varchar(255)`
	// 分机号（ISV不可见）
	Tel string `json:"tel" xorm:"Varchar(255)`
	// 办公地点（ISV不可见）
	WorkPlace string `json:"workPlace" xorm:"Varchar(255)`
	// 备注（ISV不可见）
	Remark string `json:"remark" xorm:"Varchar(255)`
	// 是否是企业的管理员, true表示是, false表示不是
	IsAdmin bool `json:"isAdmin" xorm:"Bool`
	// 是否为企业的老板, true表示是, false表示不是 （不能通过接口设置,可以通过OA后台设置）
	IsBoss bool `json:"isBoss" xorm:"Bool`
	// 是否隐藏号码, true表示是, false表示不是
	IsHide bool `json:"isHide" xorm:"Bool`
	// 是否是部门的主管, true表示是, false表示不是
	IsLeader bool `json:"isLeader" xorm:"Bool`
	// 成员名称
	Name string `json:"name" xorm:"Varchar(255)`
	// 表示该用户是否激活了钉钉
	Active bool `json:"active" xorm:"Bool`
	// 成员所属部门id列表
	Department []int64 `json:"department" xorm:"Text`
	// 职位信息
	Position string `json:"position" xorm:"Varchar(255)`
	// 员工的邮箱
	Email string `json:"email" xorm:"Varchar(255)`
	// 员工的企业邮箱，如果员工的企业邮箱没有开通，返回信息中不包含
	OrgEmail string `json:"orgEmail" xorm:"Varchar(255)`
	// 头像url
	Avatar string `json:"avatar" xorm:"Varchar(255)`
	// 员工工号
	Jobnumber string `json:"jobnumber" xorm:"Varchar(255)`
	// 入职时间
	HiredDate int64 `json:"hiredDate" xorm:"Varchar(255)`
	// 扩展属性，可以设置多种属性(但手机上最多只能显示10个扩展属性，具体显示哪些属性，请到OA管理后台->设置->通讯录信息设置和OA管理后台->设置->手机端显示信息设置)
	Extattr map[string]string `json:"extattr" xorm:"Text`
	// 在扫码登录和企业通讯录里面的openid不一致，该ID由userId加密生成，员工在当前企业的唯一标识ID，如果该员工重新入职openid不变，该ID不能用来查询用户信息
	OpenId string `json:"openId" xorm:"Varchar(255)"`
	// 在钉钉全局范围内唯一标识用户的身份
	DingId string `json:"dingId" xorm:"Varchar(255)"`
	// 用户所在部门全路径
	DeptShortName string `xorm:"Text"`
}

func (m *Ding_User) String() string {
	return "Ding_User"
}

func (m *Ding_User) ClearPk() {
	m.Userid = ""
}

func (m *Ding_User) ToMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["Userid"] = strings.TrimSpace(m.Userid)
	resp["Order"] = m.Order
	resp["Unionid"] = strings.TrimSpace(m.Unionid)
	resp["Mobile"] = strings.TrimSpace(m.Mobile)
	resp["Tel"] = strings.TrimSpace(m.Tel)
	resp["WorkPlace"] = strings.TrimSpace(m.WorkPlace)
	resp["Remark"] = strings.TrimSpace(m.Remark)
	resp["IsAdmin"] = m.IsAdmin
	resp["IsBoss"] = m.IsBoss
	resp["IsHide"] = m.IsHide
	resp["IsLeader"] = m.IsLeader
	resp["Name"] = strings.TrimSpace(m.Name)
	resp["Active"] = m.Active
	resp["Department"] = m.Department
	resp["Position"] = strings.TrimSpace(m.Position)
	resp["Email"] = strings.TrimSpace(m.Email)
	resp["OrgEmail"] = strings.TrimSpace(m.OrgEmail)
	resp["Avatar"] = strings.TrimSpace(m.Avatar)
	resp["Jobnumber"] = strings.TrimSpace(m.Jobnumber)
	resp["HiredDate"] = m.HiredDate
	resp["Extattr"] = m.Extattr
	resp["OpenId"] = strings.TrimSpace(m.OpenId)
	resp["DingId"] = strings.TrimSpace(m.DingId)
	resp["DeptShortName"] = strings.TrimSpace(m.DeptShortName)
	return resp
}

func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		GetLogEntry().Warningln("to slice arr not slice")
		return nil
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}

func (m *Ding_User) Parse(in map[string]interface{}, columns []guide.Column) {
	var err error
	if value, ok := in["Userid"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.Userid = fmt.Sprint(value)
		}
	}
	if value, ok := in["Order"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.Order, err = strconv.ParseInt(fmt.Sprint(value), 10, 64); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["Unionid"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.Unionid = fmt.Sprint(value)
		}
	}
	if value, ok := in["Mobile"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.Mobile = fmt.Sprint(value)
		}
	}
	if value, ok := in["Tel"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.Tel = fmt.Sprint(value)
		}
	}
	if value, ok := in["WorkPlace"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.WorkPlace = fmt.Sprint(value)
		}
	}
	if value, ok := in["Remark"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.Remark = fmt.Sprint(value)
		}
	}
	if value, ok := in["IsAdmin"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.IsAdmin, err = strconv.ParseBool(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["IsBoss"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.IsBoss, err = strconv.ParseBool(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["IsHide"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.IsHide, err = strconv.ParseBool(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["IsLeader"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.IsLeader, err = strconv.ParseBool(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["Name"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.Name = fmt.Sprint(value)
		}
	}
	if value, ok := in["Active"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.Active, err = strconv.ParseBool(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["Department"]; ok {
		array := ToSlice(value)
		for _, v := range array {
			department, err := strconv.ParseInt(fmt.Sprint(v), 10, 64)
			if err != nil {
				GetLogEntry().Warningln(err)
			}
			m.Department = append(m.Department, department)
		}
	}
	if value, ok := in["Position"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.Position = fmt.Sprint(value)
		}
	}
	if value, ok := in["Email"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.Email = fmt.Sprint(value)
		}
	}
	if value, ok := in["OrgEmail"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.OrgEmail = fmt.Sprint(value)
		}
	}
	if value, ok := in["Avatar"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.Avatar = fmt.Sprint(value)
		}
	}
	if value, ok := in["Jobnumber"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.Jobnumber = fmt.Sprint(value)
		}
	}
	if value, ok := in["HiredDate"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.HiredDate, err = strconv.ParseInt(fmt.Sprint(value), 10, 64); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["Extattr"]; ok {
		m.Extattr = value.(map[string]string)
	}
	if value, ok := in["OpenId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.OpenId = fmt.Sprint(value)
		}
	}
	if value, ok := in["DingId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.DingId = fmt.Sprint(value)
		}
	}
}

func (m *Ding_User) Match(source DataHandler, columns []guide.Column) {
	tMap := m.ToMap()
	sMap := source.ToMap()
	for _, col := range columns {
		if len(strings.TrimSpace(col.DefaultValue)) > 0 {
			sMap[col.Source] = col.DefaultValue
		} else {
			sMap[col.Source] = tMap[col.Target]
		}
		if col.DefaultType == "fn" {
			if fn, ok := internalFns[col.DefaultValue]; ok {
				value, err := fn(source.ToMap())
				if err != nil {
					GetLogEntry().Warningln(err)
					continue
				}
				sMap[col.Source] = value
			}
		}
	}
	m.Parse(sMap, columns)
}

// 钉钉用户角色关联信息
type Ding_UserRole struct {
	Id string `xorm:"Varchar(255)" pk notnull`
	// 员工唯一标识ID（不可修改）
	UserId string `xorm:"Varchar(255)"`
	// 角色id（ISV不可见）
	RoleId int64 `xorm:"BigInt`
	// 角色名称（ISV不可见）
	RoleName string `xorm:"Varchar(255)`
	// 角色分组id
	GroupId int64 `xorm:"BigInt`
	// 角色分组名称（ISV不可见）
	GroupName string `xorm:"Varchar(255)`
}

func (m *Ding_UserRole) String() string {
	return "Ding_UserRole"
}

func (m *Ding_UserRole) ClearPk() {
	m.Id = ""
}

func (m *Ding_UserRole) ToMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["Id"] = strings.TrimSpace(m.Id)
	resp["UserId"] = strings.TrimSpace(m.UserId)
	resp["RoleId"] = m.RoleId
	resp["RoleName"] = strings.TrimSpace(m.RoleName)
	resp["GroupId"] = m.GroupId
	resp["GroupName"] = strings.TrimSpace(m.GroupName)
	return resp
}

func (m *Ding_UserRole) Parse(in map[string]interface{}, columns []guide.Column) {
	var err error
	if value, ok := in["UserId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.UserId = fmt.Sprint(value)
		}
	}
	if value, ok := in["RoleId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.RoleId, err = strconv.ParseInt(fmt.Sprint(value), 10, 64); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["RoleName"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.RoleName = fmt.Sprint(value)
		}
	}
	if value, ok := in["GroupId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.GroupId, err = strconv.ParseInt(fmt.Sprint(value), 10, 64); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["GroupName"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.GroupName = fmt.Sprint(value)
		}
	}
}

func (m *Ding_UserRole) Match(source DataHandler, columns []guide.Column) {
	tMap := m.ToMap()
	sMap := source.ToMap()
	for _, col := range columns {
		if len(strings.TrimSpace(col.DefaultValue)) > 0 {
			sMap[col.Source] = col.DefaultValue
		} else {
			sMap[col.Source] = tMap[col.Target]
		}
		if col.DefaultType == "fn" {
			if fn, ok := internalFns[col.DefaultValue]; ok {
				value, err := fn(source.ToMap())
				if err != nil {
					GetLogEntry().Warningln(err)
					continue
				}
				sMap[col.Source] = value
			}
		}
	}
	m.Parse(sMap, columns)
}

// 目标库用户信息
type Sys_User struct {
	// 用户主键
	F_Id      string `xorm:"Varchar(50) pk notnull"`
	F_UnionId string `xorm:"Varchar(100)"`
	// 账户
	F_Account    string `xorm:"Varchar(50)"`
	WeChatNumber string `xorm:"Varchar(100)"`
	OpenId       string `xorm:"Varchar(100)"`
	// 姓名
	F_RealName string `xorm:"Varchar(50)"`
	// 工号
	F_NickName string `xorm:"Varchar(50)"`
	// 头像
	F_HeadIcon string `xorm:"Varchar(50)"`
	// 性别
	F_Gender bool `xorm:"Bool"`
	// 生日
	F_RuZhiDate time.Time `xorm:"DateTime"`
	// 手机
	F_MobilePhone string `xorm:"Varchar(20)"`
	// 邮箱
	F_Email string `xorm:"Varchar(50)"`
	// 微信
	F_WeChat string `xorm:"Varchar(50)"`
	// 主管主键
	F_ManagerId string `xorm:"Varchar(50)"`
	// 安全级别
	F_SecurityLevel int `xorm:"Int"`
	// 个性签名
	F_Signature string `xorm:"Varchar(500)"`
	// 组织主键
	F_OrganizeId string `xorm:"Varchar(50)"`
	// 部门主键
	F_DepartmentId string `xorm:"Varchar(500)"`
	// 角色主键
	F_RoleId string `xorm:"Varchar(500)"`
	// 岗位主键
	F_DutyId string `xorm:"Varchar(500)"`
	// 是否管理员
	F_IsAdministrator bool `xorm:"Bool"`
	// 排序码
	F_SortCode int `xorm:"Int"`
	// 删除标志
	F_DeleteMark bool `xorm:"Bool"`
	// 有效标志
	F_EnabledMark bool `xorm:"Bool"`
	// 描述
	F_Description string `xorm:"Varchar(500)"`
	// 创建时间
	F_CreatorTime time.Time `xorm:"DateTime"`
	// 创建用户
	F_CreatorUserId string `xorm:"Varchar(50)"`
	// 最后修改时间
	F_LastModifyTime time.Time `xorm:"DateTime"`
	// 最后修改用户
	F_LastModifyUserId string `xorm:"Varchar(50)"`
	// 删除时间
	F_DeleteTime time.Time `xorm:"DateTime"`
	// 删除用户
	F_DeleteUserId string `xorm:"Varchar(500)"`
	// 是否参与排名
	IsPaiMing bool `xorm:"Bool"`
	// 员工编号
	F_JobNumber string `xorm:"Varchar(500)"`
	// 用户所在部门全路径
	F_DeptShortName string `xorm:"Text"`
	isNew           bool   `xorm:"-"`
}

func (m *Sys_User) String() string {
	return "Sys_User"
}

func (m *Sys_User) ClearPk() {
	m.F_Id = ""
}

func (m *Sys_User) ToMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["F_Id"] = strings.TrimSpace(m.F_Id)
	resp["F_Account"] = strings.TrimSpace(m.F_Account)
	resp["F_UnionId"] = strings.TrimSpace(m.F_UnionId)
	resp["WeChatNumber"] = strings.TrimSpace(m.WeChatNumber)
	resp["OpenId"] = strings.TrimSpace(m.OpenId)
	resp["F_RealName"] = strings.TrimSpace(m.F_RealName)
	resp["F_NickName"] = strings.TrimSpace(m.F_NickName)
	resp["F_HeadIcon"] = strings.TrimSpace(m.F_HeadIcon)
	resp["F_Gender"] = m.F_Gender
	if !m.F_RuZhiDate.IsZero() {
		resp["F_RuZhiDate"] = m.F_RuZhiDate
	}
	resp["F_MobilePhone"] = strings.TrimSpace(m.F_MobilePhone)
	resp["F_Email"] = strings.TrimSpace(m.F_Email)
	resp["F_WeChat"] = strings.TrimSpace(m.F_WeChat)
	resp["F_ManagerId"] = strings.TrimSpace(m.F_ManagerId)
	resp["F_SecurityLevel"] = m.F_SecurityLevel
	resp["F_Signature"] = strings.TrimSpace(m.F_Signature)
	resp["F_OrganizeId"] = strings.TrimSpace(m.F_OrganizeId)
	resp["F_DepartmentId"] = strings.TrimSpace(m.F_DepartmentId)
	resp["F_RoleId"] = strings.TrimSpace(m.F_RoleId)
	resp["F_DutyId"] = strings.TrimSpace(m.F_DutyId)
	resp["F_IsAdministrator"] = m.F_IsAdministrator
	resp["F_SortCode"] = m.F_SortCode
	resp["F_DeleteMark"] = m.F_DeleteMark
	resp["F_EnabledMark"] = m.F_EnabledMark
	resp["F_Description"] = strings.TrimSpace(m.F_Description)
	if !m.F_CreatorTime.IsZero() {
		resp["F_CreatorTime"] = m.F_CreatorTime
	}
	resp["F_CreatorUserId"] = strings.TrimSpace(m.F_CreatorUserId)
	if !m.F_LastModifyTime.IsZero() {
		resp["F_LastModifyTime"] = m.F_LastModifyTime
	}
	resp["F_LastModifyUserId"] = strings.TrimSpace(m.F_LastModifyUserId)
	if !m.F_DeleteTime.IsZero() {
		resp["F_DeleteTime"] = m.F_DeleteTime
	}

	resp["F_DeleteUserId"] = strings.TrimSpace(m.F_DeleteUserId)
	resp["IsPaiMing"] = m.IsPaiMing
	resp["F_JobNumber"] = strings.TrimSpace(m.F_JobNumber)
	resp["F_DeptShortName"] = strings.TrimSpace(m.F_DeptShortName)
	return resp
}

func (m *Sys_User) Parse(in map[string]interface{}, columns []guide.Column) {
	var err error
	if value, ok := in["F_Id"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_Id = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_UnionId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_UnionId = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_Account"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_Account = fmt.Sprint(value)
		}
	}
	if value, ok := in["WeChatNumber"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.WeChatNumber = fmt.Sprint(value)
		}
	}
	if value, ok := in["OpenId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.OpenId = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_RealName"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_RealName = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_NickName"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_NickName = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_HeadIcon"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_HeadIcon = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_Gender"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.F_Gender, err = strconv.ParseBool(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["F_RuZhiDate"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			tm, ok := value.(time.Time)
			if ok {
				m.F_CreatorTime = tm
			}
		}
	}
	if value, ok := in["F_MobilePhone"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_MobilePhone = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_Email"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_Email = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_WeChat"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_WeChat = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_ManagerId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_ManagerId = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_SecurityLevel"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.F_SecurityLevel, err = strconv.Atoi(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["F_Signature"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_Signature = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_OrganizeId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_OrganizeId = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_DepartmentId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_DepartmentId = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_RoleId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_RoleId = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_DutyId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_DutyId = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_IsAdministrator"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.F_IsAdministrator, err = strconv.ParseBool(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["F_SortCode"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.F_SortCode, err = strconv.Atoi(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["F_DeleteMark"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.F_DeleteMark, err = strconv.ParseBool(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["F_EnabledMark"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.F_EnabledMark, err = strconv.ParseBool(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["F_Description"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_Description = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_CreatorTime"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			tm, ok := value.(time.Time)
			if ok {
				m.F_CreatorTime = tm
			}
		}
	}
	if value, ok := in["F_CreatorUserId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_CreatorUserId = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_LastModifyTime"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			tm, ok := value.(time.Time)
			if ok {
				m.F_LastModifyTime = tm
			}
		}
	}
	if value, ok := in["F_LastModifyUserId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_LastModifyUserId = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_DeleteTime"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			tm, ok := value.(time.Time)
			if ok {
				m.F_DeleteTime = tm
			}
		}
	}
	if value, ok := in["F_DeleteUserId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_DeleteUserId = fmt.Sprint(value)
		}
	}
	if value, ok := in["IsPaiMing"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.IsPaiMing, err = strconv.ParseBool(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["F_JobNumber"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_JobNumber = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_DeptShortName"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_DeptShortName = fmt.Sprint(value)
		}
	}
}

func (m *Sys_User) Match(source DataHandler, columns []guide.Column) {
	tMap := m.ToMap()
	for _, col := range columns {
		if len(strings.TrimSpace(col.DefaultValue)) > 0 {
			if col.DefaultType == "fn" {
				if fn, ok := internalFns[col.DefaultValue]; ok {
					value, err := fn(source.ToMap())
					if err != nil {
						GetLogEntry().Warningln(err)
						continue
					}
					tMap[col.Target] = value
				}
			} else {
				tMap[col.Target] = col.DefaultValue
			}
		}
	}
	sMap := source.ToMap()
	for _, col := range columns {
		if len(strings.TrimSpace(col.Source)) > 0 {
			tMap[col.Target] = sMap[col.Source]
		}
	}
	m.Parse(tMap, columns)
}

func (m *Sys_User) Sync(table *guide.Table) error {
	session := guide.GetDB().NewSession()
	defer session.Close()
	targets := make([]*Sys_User, 0)
	if err := session.Find(&targets); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	sources := make([]*Ding_User, 0)
	if err := session.Find(&sources); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}

	if len(targets) == 0 {
		in_array := make([]*Sys_User, 0)
		for _, source := range sources {
			target := &Sys_User{}
			target.isNew = true
			target.Match(source, table.Columns)
			uid, _ := uuid.NewV4()
			target.F_Id = uid.String()
			target.F_CreatorTime = time.Now().Local()
			in_array = append(in_array, target)
		}
		session = guide.GetDB().NewSession()
		defer session.Close()
		if err := session.Begin(); err != nil {
			GetLogEntry().Errorln(err)
			return err
		}
		for _, bean := range in_array {
			if _, err := session.Insert(bean); err != nil {
				session.Rollback()
				GetLogEntry().Errorln(err)
				return err
			}
		}
		if err := session.Commit(); err != nil {
			GetLogEntry().Errorln(err)
			return err
		}

		return nil
	}
	unique_keys := make([]*guide.Column, 0)
	for _, col := range table.Columns {
		if col.Unique {
			unique_keys = append(unique_keys, &guide.Column{
				Source:       col.Source,
				Target:       col.Target,
				Unique:       col.Unique,
				DefaultValue: col.DefaultValue,
				DefaultType:  col.DefaultType,
				Description:  col.Description,
			})
		}
	}
	unique_uped_maps := make(map[string]*Ding_User)
	up_maps := make(map[*Sys_User]*Sys_User)
	session = guide.GetDB().NewSession()
	for _, source := range sources {
		sMap := source.ToMap()
		for _, target := range targets {
			isOk := false
			tMap := target.ToMap()
			cond := &Sys_User{}
			condMap := cond.ToMap()
			cond_s := &Ding_User{}
			condMap_s := cond_s.ToMap()
			unique_uped_key_array := make([]string, 0)
			for _, key := range unique_keys {
				sVal := strings.TrimSpace(fmt.Sprint(sMap[key.Source]))
				tVal := strings.TrimSpace(fmt.Sprint(tMap[key.Target]))
				if sVal != "" && tVal != "" && sVal == tVal {
					condMap[key.Target] = tVal
					condMap_s[key.Source] = sVal
					unique_uped_key_array = append(unique_uped_key_array, tVal)
					isOk = true
				} else {
					isOk = false
				}
			}
			if isOk {
				cond_s.Parse(condMap_s, table.Columns)
				cond_s_array := make([]*Ding_User, 0)
				if err := session.Find(&cond_s_array, cond_s); err != nil {
					GetLogEntry().Errorln(err)
					return err
				}
				if len(cond_s_array) > 1 {
					continue
				}
				cond.Parse(condMap, table.Columns)
				target.Match(source, table.Columns)
				target.ClearPk()
				up_target := &Sys_User{}
				up_target.Parse(target.ToMap(), table.Columns)
				up_target.F_LastModifyTime = time.Now().Local()
				up_maps[up_target] = cond
				unique_uped_key := strings.Join(unique_uped_key_array, "_")
				if _, ok := unique_uped_maps[unique_uped_key]; !ok {
					unique_uped_maps[unique_uped_key] = source
				}
			}
		}
	}
	session.Close()
	if len(up_maps) > 0 {
		session = guide.GetDB().NewSession()
		defer session.Close()
		if err := session.Begin(); err != nil {
			GetLogEntry().Errorln(err)
			return err
		}
		for bean, cond := range up_maps {
			if _, err := session.Update(bean, cond); err != nil {
				session.Rollback()
				GetLogEntry().Errorln(err)
				return err
			}
		}
		if err := session.Commit(); err != nil {
			GetLogEntry().Errorln(err)
			return err
		}

	}
	in_array := make([]*Sys_User, 0)
	session = guide.GetDB().NewSession()
	defer session.Close()
	for _, source := range sources {
		sMap := source.ToMap()
		cond := &Ding_User{}
		condMap := cond.ToMap()
		unique_in_key_array := make([]string, 0)
		for _, key := range unique_keys {
			unique_in_key_val := strings.TrimSpace(fmt.Sprint(sMap[key.Source]))
			if len(unique_in_key_val) > 0 {
				condMap[key.Source] = unique_in_key_val
				unique_in_key_array = append(unique_in_key_array, unique_in_key_val)
			}
		}
		if len(unique_in_key_array) == 0 {
			continue
		}
		cond.Parse(condMap, table.Columns)
		cond_array := make([]*Ding_User, 0)
		if err := session.Find(&cond_array, cond); err != nil {
			GetLogEntry().Errorln(err)
			return err
		}
		if len(cond_array) > 1 {
			continue
		}
		unique_in_key := strings.Join(unique_in_key_array, "_")
		if _, ok := unique_uped_maps[unique_in_key]; !ok {
			target := &Sys_User{}
			target.isNew = true
			target.Match(source, table.Columns)
			uid, _ := uuid.NewV4()
			target.F_Id = uid.String()
			target.F_CreatorTime = time.Now().UTC()
			in_array = append(in_array, target)
		}
	}
	if len(in_array) > 0 {
		session = guide.GetDB().NewSession()
		defer session.Close()
		if err := session.Begin(); err != nil {
			GetLogEntry().Errorln(err)
			return err
		}
		for _, bean := range in_array {
			if _, err := session.Insert(bean); err != nil {
				session.Rollback()
				GetLogEntry().Errorln(err)
				return err
			}
			data, _ := json.Marshal(bean)
			GetLogEntry().Infoln(string(data))
		}
		if err := session.Commit(); err != nil {
			GetLogEntry().Errorln(err)
			return err
		}

	}
	return nil
}

// 目标库用户登录信息
type Sys_UserLogOn struct {
	// 用户登录主键
	F_Id      string `xorm:Varchar(50) pk notnull"`
	F_UnionId string `xorm:"Varchar(100)"`
	// 用户主键
	F_UserId string `xorm:Varchar(50)`
	// 用户密码
	F_UserPassword string `xorm:Varchar(50)`
	// 用户秘钥
	F_UserSecretkey string `xorm:Varchar(50)`
	// 允许登录时间开始
	F_AllowStartTime time.Time `xorm:"DateTime"`
	// 允许登录时间结束
	F_AllowEndTime time.Time `xorm:"DateTime"`
	// 暂停用户开始日期
	F_LockStartDate time.Time `xorm:"DateTime"`
	// 暂停用户结束日期
	F_LockEndDate time.Time `xorm:"DateTime"`
	// 第一次访问时间
	F_FirstVisitTime time.Time `xorm:"DateTime"`
	// 上一次访问时间
	F_PreviousVisitTime time.Time `xorm:"DateTime"`
	// 最后访问时间
	F_LastVisitTime time.Time `xorm:"DateTime"`
	// 最后修改密码日期
	F_ChangePasswordDate time.Time `xorm:"DateTime"`
	// 允许同时有多用户登录
	F_MultiUserLogin bool `xorm:"Bool"`
	// 登录次数
	F_LogOnCount int `xorm:"Int"`
	// 在线状态
	F_UserOnLine bool `xorm:"Bool"`
	// 密码提示问题
	F_Question string `xorm:Varchar(50)`
	// 密码提示答案
	F_AnswerQuestion string `xorm:Varchar(500)`
	// 是否访问限制
	F_CheckIPAddress bool `xorm:"Bool"`
	// 系统语言
	F_Language string `xorm:Varchar(50)`
	// 系统样式
	F_Theme string `xorm:Varchar(50)`
	// 员工编号
	F_JobNumber string `xorm:"Varchar(500)"`
	isNew       bool   `xorm:"-"`
}

func (m *Sys_UserLogOn) String() string {
	return "Sys_UserLogOn"
}

func (m *Sys_UserLogOn) ClearPk() {
	m.F_Id = ""
}

func (m *Sys_UserLogOn) ToMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["F_Id"] = strings.TrimSpace(m.F_Id)
	resp["F_UnionId"] = strings.TrimSpace(m.F_UnionId)
	resp["F_UserId"] = strings.TrimSpace(m.F_UserId)
	resp["F_UserPassword"] = strings.TrimSpace(m.F_UserPassword)
	resp["F_UserSecretkey"] = strings.TrimSpace(m.F_UserSecretkey)
	if !m.F_AllowEndTime.IsZero() {
		resp["F_AllowStartTime"] = m.F_AllowStartTime
	}
	if !m.F_AllowEndTime.IsZero() {
		resp["F_AllowEndTime"] = m.F_AllowEndTime
	}
	if !m.F_LockStartDate.IsZero() {
		resp["F_LockStartDate"] = m.F_LockStartDate
	}
	if !m.F_LockEndDate.IsZero() {
		resp["F_LockEndDate"] = m.F_LockEndDate
	}
	if !m.F_FirstVisitTime.IsZero() {
		resp["F_FirstVisitTime"] = m.F_FirstVisitTime
	}
	if !m.F_PreviousVisitTime.IsZero() {
		resp["F_PreviousVisitTime"] = m.F_PreviousVisitTime
	}
	if !m.F_LastVisitTime.IsZero() {
		resp["F_LastVisitTime"] = m.F_LastVisitTime
	}
	if !m.F_ChangePasswordDate.IsZero() {
		resp["F_ChangePasswordDate"] = m.F_ChangePasswordDate
	}
	resp["F_MultiUserLogin"] = m.F_MultiUserLogin
	resp["F_LogOnCount"] = m.F_LogOnCount
	resp["F_UserOnLine"] = m.F_UserOnLine
	resp["F_Question"] = strings.TrimSpace(m.F_Question)
	resp["F_AnswerQuestion"] = strings.TrimSpace(m.F_AnswerQuestion)
	resp["F_CheckIPAddress"] = m.F_CheckIPAddress
	resp["F_Language"] = strings.TrimSpace(m.F_Language)
	resp["F_Theme"] = strings.TrimSpace(m.F_Theme)
	resp["F_JobNumber"] = strings.TrimSpace(m.F_JobNumber)
	return resp
}

func (m *Sys_UserLogOn) Parse(in map[string]interface{}, columns []guide.Column) {
	var err error
	if value, ok := in["F_Id"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_Id = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_UnionId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_UnionId = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_UserId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_UserId = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_UserPassword"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_UserPassword = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_UserSecretkey"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_UserSecretkey = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_AllowStartTime"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			tm, ok := value.(time.Time)
			if ok {
				m.F_AllowStartTime = tm
			}
		}
	}
	if value, ok := in["F_AllowEndTime"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			tm, ok := value.(time.Time)
			if ok {
				m.F_AllowEndTime = tm
			}
		}
	}
	if value, ok := in["F_LockStartDate"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			tm, ok := value.(time.Time)
			if ok {
				m.F_LockStartDate = tm
			}
		}
	}
	if value, ok := in["F_LockEndDate"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			tm, ok := value.(time.Time)
			if ok {
				m.F_LockEndDate = tm
			}
		}
	}
	if value, ok := in["F_FirstVisitTime"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			tm, ok := value.(time.Time)
			if ok {
				m.F_FirstVisitTime = tm
			}
		}
	}
	if value, ok := in["F_PreviousVisitTime"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			tm, ok := value.(time.Time)
			if ok {
				m.F_PreviousVisitTime = tm
			}
		}
	}
	if value, ok := in["F_LastVisitTime"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			tm, ok := value.(time.Time)
			if ok {
				m.F_LastVisitTime = tm
			}
		}
	}
	if value, ok := in["F_ChangePasswordDate"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			tm, ok := value.(time.Time)
			if ok {
				m.F_ChangePasswordDate = tm
			}
		}
	}
	if value, ok := in["F_MultiUserLogin"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.F_MultiUserLogin, err = strconv.ParseBool(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["F_LogOnCount"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.F_LogOnCount, err = strconv.Atoi(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["F_UserOnLine"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.F_UserOnLine, err = strconv.ParseBool(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["F_Question"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_Question = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_AnswerQuestion"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_AnswerQuestion = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_CheckIPAddress"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.F_CheckIPAddress, err = strconv.ParseBool(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["F_Language"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_Language = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_Theme"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_Theme = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_JobNumber"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_JobNumber = fmt.Sprint(value)
		}
	}
}

func (m *Sys_UserLogOn) Match(source DataHandler, columns []guide.Column) {
	tMap := m.ToMap()
	for _, col := range columns {
		if m.isNew && len(strings.TrimSpace(col.DefaultValue)) > 0 {
			if col.DefaultType == "fn" {
				if fn, ok := internalFns[col.DefaultValue]; ok {
					value, err := fn(source.ToMap())
					if err != nil {
						GetLogEntry().Warningln(err)
						continue
					}
					tMap[col.Target] = value
				}
			} else {
				tMap[col.Target] = col.DefaultValue
			}
		} else if !m.isNew && len(strings.TrimSpace(col.DefaultValue)) > 0 {
			tMap[col.Target] = ""
		}
	}
	sMap := source.ToMap()
	for _, col := range columns {
		if len(strings.TrimSpace(col.Source)) > 0 && len(strings.TrimSpace(col.Target)) > 0 {
			tMap[col.Target] = sMap[col.Source]
		}
	}
	m.Parse(tMap, columns)
}

func (m *Sys_UserLogOn) Sync(table *guide.Table) error {
	session := guide.GetDB().NewSession()
	defer session.Close()
	targets := make([]*Sys_UserLogOn, 0)
	if err := session.Find(&targets); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	sources := make([]*Sys_User, 0)
	if err := session.Find(&sources); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}

	if len(targets) == 0 {
		in_array := make([]*Sys_UserLogOn, 0)
		for _, source := range sources {
			target := &Sys_UserLogOn{}
			target.isNew = true
			target.Match(source, table.Columns)
			uid, _ := uuid.NewV4()
			target.F_Id = uid.String()
			in_array = append(in_array, target)
		}
		session = guide.GetDB().NewSession()
		defer session.Close()
		if err := session.Begin(); err != nil {
			GetLogEntry().Errorln(err)
			return err
		}
		for _, bean := range in_array {
			if _, err := session.Insert(bean); err != nil {
				session.Rollback()
				GetLogEntry().Errorln(err)
				return err
			}
		}
		if err := session.Commit(); err != nil {
			GetLogEntry().Errorln(err)
			return err
		}

		return nil
	}
	unique_keys := make([]*guide.Column, 0)
	for _, col := range table.Columns {
		if col.Unique {
			unique_keys = append(unique_keys, &guide.Column{
				Source:       col.Source,
				Target:       col.Target,
				Unique:       col.Unique,
				DefaultValue: col.DefaultValue,
				DefaultType:  col.DefaultType,
				Description:  col.Description,
			})
		}
	}
	unique_uped_maps := make(map[string]*Sys_User)
	//up_maps := make(map[*Sys_UserLogOn]*Sys_UserLogOn)
	for _, source := range sources {
		sMap := source.ToMap()
		for _, target := range targets {
			isOk := false
			tMap := target.ToMap()
			//cond := &Sys_UserLogOn{}
			//condMap := cond.ToMap()
			unique_uped_key_array := make([]string, 0)
			for _, key := range unique_keys {
				sVal := strings.TrimSpace(fmt.Sprint(sMap[key.Source]))
				tVal := strings.TrimSpace(fmt.Sprint(tMap[key.Target]))
				if sVal != "" && tVal != "" && sVal == tVal {
					//condMap[key.Target] = tVal
					unique_uped_key_array = append(unique_uped_key_array, tVal)
					isOk = true
				} else {
					isOk = false
				}
			}
			if isOk {
				//cond.Parse(condMap, table.Columns)
				//target.Match(source, table.Columns)
				//target.ClearPk()
				//up_target := &Sys_UserLogOn{}
				//up_target.Parse(target.ToMap(), table.Columns)
				//up_maps[up_target] = cond
				unique_uped_key := strings.Join(unique_uped_key_array, "_")
				if _, ok := unique_uped_maps[unique_uped_key]; !ok {
					unique_uped_maps[unique_uped_key] = source
				}
			}
		}
	}
	//	if len(up_maps) > 0 {
	//
	//		if err := session.Begin(); err != nil {
	//			GetLogEntry().Errorln(err)
	//			return err
	//		}
	//		for bean, cond := range up_maps {
	//			if _, err := session.Update(bean, cond); err != nil {
	//				session.Rollback()
	//				GetLogEntry().Errorln(err)
	//				return err
	//			}
	//		}
	//		if err := session.Commit(); err != nil {
	//			GetLogEntry().Errorln(err)
	//			return err
	//		}
	//
	//	}
	in_array := make([]*Sys_UserLogOn, 0)
	for _, source := range sources {
		sMap := source.ToMap()
		unique_in_key_array := make([]string, 0)
		for _, key := range unique_keys {
			unique_in_key_val := strings.TrimSpace(fmt.Sprint(sMap[key.Source]))
			if len(unique_in_key_val) > 0 {
				unique_in_key_array = append(unique_in_key_array, unique_in_key_val)
			}
		}
		if len(unique_in_key_array) == 0 {
			continue
		}
		unique_in_key := strings.Join(unique_in_key_array, "_")
		if _, ok := unique_uped_maps[unique_in_key]; !ok {
			target := &Sys_UserLogOn{}
			target.isNew = true
			target.Match(source, table.Columns)
			uid, _ := uuid.NewV4()
			target.F_Id = uid.String()
			in_array = append(in_array, target)
		}
	}
	if len(in_array) > 0 {
		session = guide.GetDB().NewSession()
		defer session.Close()
		if err := session.Begin(); err != nil {
			GetLogEntry().Errorln(err)
			return err
		}
		for _, bean := range in_array {
			if _, err := session.Insert(bean); err != nil {
				session.Rollback()
				GetLogEntry().Errorln(err)
				return err
			}
		}
		if err := session.Commit(); err != nil {
			GetLogEntry().Errorln(err)
			return err
		}
	}
	return nil
}
