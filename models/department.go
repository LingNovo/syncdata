package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"reflect"

	"github.com/LingNovo/syncdata/guide"
	"github.com/satori/go.uuid"
)

// 钉钉部门信息
type Ding_Department struct {
	// 部门id
	Id int64 `xorm:"BigInt" pk notnull"`
	// 部门名称
	Name string `xorm:"Varchar(255)`
	// 父部门id，根部门为1
	Parentid int64 `xorm:"BigInt"`
	// 在父部门中的次序值
	Order int64 `xorm:"BigInt"`
	// 是否同步创建一个关联此部门的企业群, true表示是, false表示不是
	CreateDeptGroup bool `xorm:"Bool`
	// 当群已经创建后，是否有新人加入部门会自动加入该群, true表示是, false表示不是
	AutoAddUser bool `xorm:"Bool`
	// 是否隐藏部门, true表示隐藏, false表示显示
	DeptHiding bool `xorm:"Bool`
	// 可以查看指定隐藏部门的其他部门列表，如果部门隐藏，则此值生效，取值为其他的部门id组成的的字符串，使用|符号进行分割
	DeptPermits string `xorm:"Varchar(255)`
	// 可以查看指定隐藏部门的其他人员列表，如果部门隐藏，则此值生效，取值为其他的人员userid组成的的字符串，使用|符号进行分割
	UserPermits string `xorm:"Varchar(255)`
	// 是否本部门的员工仅可见员工自己, 为true时，本部门员工默认只能看到员工自己
	OuterDept bool `xorm:"Bool`
	// 本部门的员工仅可见员工自己为true时，可以配置额外可见部门，值为部门id组成的的字符串，使用|符号进行分割
	OuterPermitDepts string `xorm:"Varchar(255)`
	// 本部门的员工仅可见员工自己为true时，可以配置额外可见人员，值为userid组成的的字符串，使用| 符号进行分割
	OuterPermitUsers string `xorm:"Varchar(255)`
	// 企业群群主
	OrgDeptOwner string `xorm:"Varchar(255)`
	// 部门的主管列表,取值为由主管的userid组成的字符串，不同的userid使用|符号进行分割
	DeptManagerUseridList string `xorm:"Varchar(255)`
	// 部门标识字段，开发者可用该字段来唯一标识一个部门，并与钉钉外部通讯录里的部门做映射
	SourceIdentifier string `xorm:"Varchar(255)`
	// 部门群是否包含子部门
	GroupContainSubDept bool `xorm:"Varchar(255)`
	//全路径
	ShortName string `xorm:"Text`
}

func (m *Ding_Department) String() string {
	return "Ding_Department"
}

func (m *Ding_Department) ClearPk() {
	m.Id = 0
}

func (m *Ding_Department) ToMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["Id"] = m.Id
	resp["Name"] = strings.TrimSpace(m.Name)
	resp["Parentid"] = m.Parentid
	resp["Order"] = m.Order
	resp["CreateDeptGroup"] = m.CreateDeptGroup
	resp["AutoAddUser"] = m.AutoAddUser
	resp["DeptHiding"] = m.DeptHiding
	resp["DeptPermits"] = strings.TrimSpace(m.DeptPermits)
	resp["UserPermits"] = strings.TrimSpace(m.UserPermits)
	resp["OuterDept"] = m.OuterDept
	resp["OuterPermitDepts"] = strings.TrimSpace(m.OuterPermitDepts)
	resp["OuterPermitUsers"] = strings.TrimSpace(m.OuterPermitUsers)
	resp["OrgDeptOwner"] = strings.TrimSpace(m.OrgDeptOwner)
	resp["DeptManagerUseridList"] = strings.TrimSpace(m.DeptManagerUseridList)
	resp["SourceIdentifier"] = strings.TrimSpace(m.SourceIdentifier)
	resp["GroupContainSubDept"] = m.GroupContainSubDept
	resp["ShortName"] = strings.TrimSpace(m.ShortName)
	return resp
}

func (m *Ding_Department) Parse(in map[string]interface{}, columns []guide.Column) {
	var err error
	if value, ok := in["Id"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.Id, err = strconv.ParseInt(fmt.Sprint(value), 10, 64); err != nil {
				GetLogEntry().Warningln(err)
			}
		}

	}
	if value, ok := in["Name"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.Name = fmt.Sprint(value)
		}
	}
	if value, ok := in["Parentid"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.Parentid, err = strconv.ParseInt(fmt.Sprint(value), 10, 64); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["Order"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.Order, err = strconv.ParseInt(fmt.Sprint(value), 10, 64); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["CreateDeptGroup"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.CreateDeptGroup, err = strconv.ParseBool(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["AutoAddUser"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.AutoAddUser, err = strconv.ParseBool(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["DeptHiding"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.DeptHiding, err = strconv.ParseBool(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["DeptPermits"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.DeptPermits = fmt.Sprint(value)
		}
	}
	if value, ok := in["UserPermits"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.UserPermits = fmt.Sprint(value)
		}
	}
	if value, ok := in["OuterDept"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.OuterDept, err = strconv.ParseBool(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["OuterPermitDepts"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.OuterPermitDepts = fmt.Sprint(value)
		}
	}
	if value, ok := in["OuterPermitUsers"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.OuterPermitUsers = fmt.Sprint(value)
		}
	}
	if value, ok := in["OrgDeptOwner"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.OrgDeptOwner = fmt.Sprint(value)
		}
	}
	if value, ok := in["DeptManagerUseridList"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.DeptManagerUseridList = fmt.Sprint(value)
		}
	}
	if value, ok := in["SourceIdentifier"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.SourceIdentifier = fmt.Sprint(value)
		}
	}
	if value, ok := in["GroupContainSubDept"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.GroupContainSubDept, err = strconv.ParseBool(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["ShortName"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.ShortName = fmt.Sprint(value)
		}
	}
}

func (m *Ding_Department) Match(source DataHandler, columns []guide.Column) {
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

// 目标库组织机构信息
type Sys_Organize struct {
	// 组织主键
	F_Id      string `xorm:Varchar(50) pk notnull"`
	F_UnionId int64  `xorm:"BigInt"`
	// 父级
	F_ParentId string `xorm:Varchar(50)`
	// 层次
	F_Layers int `xorm:Int`
	// 编码
	F_EnCode string `xorm:Varchar(50)`
	// 名称
	F_FullName string `xorm:Varchar(50)`
	// 简称
	F_ShortName string `xorm:Varchar(50)`
	// 分类
	F_CategoryId string `xorm:Varchar(50)`
	// 负责人
	F_ManagerId string `xorm:Varchar(50)`
	// 电话
	F_TelePhone string `xorm:Varchar(20)`
	// 手机
	F_MobilePhone string `xorm:Varchar(20)`
	// 微信
	F_WeChat string `xorm:Varchar(50)`
	// 传真
	F_Fax string `xorm:Varchar(20)`
	// 邮箱
	F_Email string `xorm:Varchar(50)`
	// 归属区域
	F_AreaId string `xorm:Varchar(50)`
	// 联系地址
	F_Address string `xorm:Varchar(500)`
	// 允许编辑
	F_AllowEdit bool `xorm:"Bool`
	// 允许删除
	F_AllowDelete bool `xorm:"Bool`
	// 排序码
	F_SortCode int `xorm:Int`
	// 删除标志
	F_DeleteMark bool `xorm:"Bool`
	// 有效标志
	F_EnabledMark bool `xorm:"Bool`
	// 描述
	F_Description string `xorm:Varchar(500)`
	// 创建时间
	F_CreatorTime time.Time `xorm:DateTime`
	// 创建用户
	F_CreatorUserId string `xorm:Varchar(255)`
	// 最后修改时间
	F_LastModifyTime time.Time `xorm:DateTime`
	// 最后修改用户
	F_LastModifyUserId string `xorm:Varchar(255)`
	// 删除时间
	F_DeleteTime time.Time `xorm:DateTime`
	// 删除用户
	F_DeleteUserId string `xorm:Varchar(255)`
	isNew          bool   `xorm:"-"`
}

func (m *Sys_Organize) String() string {
	return "Sys_Organize"
}

func (m *Sys_Organize) ClearPk() {
	m.F_Id = ""
}

func (m *Sys_Organize) ToMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["F_Id"] = strings.TrimSpace(m.F_Id)
	resp["F_UnionId"] = m.F_UnionId
	resp["F_ParentId"] = strings.TrimSpace(m.F_ParentId)
	resp["F_Layers"] = m.F_Layers
	resp["F_EnCode"] = strings.TrimSpace(m.F_EnCode)
	resp["F_FullName"] = strings.TrimSpace(m.F_FullName)
	resp["F_ShortName"] = strings.TrimSpace(m.F_ShortName)
	resp["F_CategoryId"] = strings.TrimSpace(m.F_CategoryId)
	resp["F_ManagerId"] = strings.TrimSpace(m.F_ManagerId)
	resp["F_TelePhone"] = strings.TrimSpace(m.F_TelePhone)
	resp["F_MobilePhone"] = strings.TrimSpace(m.F_MobilePhone)
	resp["F_WeChat"] = strings.TrimSpace(m.F_WeChat)
	resp["F_Fax"] = strings.TrimSpace(m.F_Fax)
	resp["F_Email"] = strings.TrimSpace(m.F_Email)
	resp["F_AreaId"] = strings.TrimSpace(m.F_AreaId)
	resp["F_Address"] = strings.TrimSpace(m.F_Address)
	resp["F_AllowEdit"] = m.F_AllowEdit
	resp["F_AllowDelete"] = m.F_AllowDelete
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
	return resp
}

func (m *Sys_Organize) Parse(in map[string]interface{}, columns []guide.Column) {
	var err error
	if value, ok := in["F_Id"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_Id = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_UnionId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.F_UnionId, err = strconv.ParseInt(fmt.Sprint(value), 10, 64); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["F_ParentId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_ParentId = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_Layers"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.F_Layers, err = strconv.Atoi(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["F_EnCode"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_EnCode = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_FullName"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_FullName = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_ShortName"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_ShortName = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_CategoryId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_CategoryId = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_ManagerId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_ManagerId = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_TelePhone"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_TelePhone = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_MobilePhone"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_MobilePhone = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_WeChat"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_WeChat = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_Fax"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_Fax = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_Email"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_Email = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_AreaId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_AreaId = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_Address"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_Address = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_AllowEdit"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.F_AllowEdit, err = strconv.ParseBool(fmt.Sprint(value)); err != nil {
				GetLogEntry().Warningln(err)
			}
		}
	}
	if value, ok := in["F_AllowDelete"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.F_AllowDelete, err = strconv.ParseBool(fmt.Sprint(value)); err != nil {
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
}

func (m *Sys_Organize) Match(source DataHandler, columns []guide.Column) {
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

func Duplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}

func (m *Sys_Organize) Sync(table *guide.Table) error {
	session := guide.GetDB().NewSession()
	defer session.Close()
	targets := make([]*Sys_Organize, 0)
	if err := session.Find(&targets); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	sources := make([]*Ding_Department, 0)
	if err := session.Find(&sources); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}

	if len(targets) == 0 {
		in_array := make([]*Sys_Organize, 0)
		for _, source := range sources {
			target := &Sys_Organize{}
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
	unique_uped_maps := make(map[string]*Ding_Department)
	up_maps := make(map[*Sys_Organize]*Sys_Organize)
	session = guide.GetDB().NewSession()
	for _, source := range sources {
		sMap := source.ToMap()
		for _, target := range targets {
			isOk := false
			tMap := target.ToMap()
			cond := &Sys_Organize{}
			condMap := cond.ToMap()
			cond_s := &Ding_Department{}
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
				cond_s_array := make([]*Ding_Department, 0)
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
				up_target := &Sys_Organize{}
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
	in_array := make([]*Sys_Organize, 0)
	session = guide.GetDB().NewSession()
	defer session.Close()
	for _, source := range sources {
		sMap := source.ToMap()
		cond := &Ding_Department{}
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
		cond_array := make([]*Ding_Department, 0)
		if err := session.Find(&cond_array, cond); err != nil {
			GetLogEntry().Errorln(err)
			return err
		}
		if len(cond_array) > 1 {
			continue
		}
		unique_in_key := strings.Join(unique_in_key_array, "_")
		if _, ok := unique_uped_maps[unique_in_key]; !ok {
			target := &Sys_Organize{}
			target.isNew = true
			target.Match(source, table.Columns)
			uid, _ := uuid.NewV4()
			target.F_Id = uid.String()
			target.F_CreatorTime = time.Now().Local()
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

func (m *Sys_Organize) GetRecord() (*guide.Record, error) {
	session := guide.GetDB().NewSession()
	defer session.Close()
	cond := &Sys_Organize{}
	cond.F_Id = m.F_Id
	ok, err := session.Get(cond)
	if err != nil {
		GetLogEntry().Errorln(err)
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	if cond.F_UnionId == 0 {
		data_json, _ := json.Marshal(&guide.Record{
			Id:       cond.F_Id,
			Name:     cond.F_FullName,
			FullName: cond.F_ShortName,
		})
		guide.Logger.Infof("Sys_Organize:%s", string(data_json))
	}
	resp := &guide.Record{}
	resp.Id = cond.F_Id
	resp.Name = cond.F_FullName
	resp.FullName = cond.F_ShortName
	resp_target_cond := &Ding_Department{}
	resp_target_cond.Id = cond.F_UnionId
	ok, err = session.Get(resp_target_cond)
	if err != nil {
		GetLogEntry().Errorln(err)
		return nil, err
	}
	resp.IsValid = ok
	if !ok || len(strings.TrimSpace(resp.FullName)) == 0 {
		data_json, _ := json.Marshal(&guide.Record{
			Id:       resp.Id,
			Name:     resp.Name,
			FullName: resp.FullName,
		})
		guide.Logger.Infof("Sys_Organize :%s", string(data_json))
	}
	data_cond := &Sys_User{}
	data_cond.F_DepartmentId = resp.Id
	data_array := make([]*Sys_User, 0)
	if err := session.Find(&data_array, data_cond); err != nil {
		GetLogEntry().Errorln(err)
		return nil, err
	}
	if len(data_array) > 0 {
		resp.Data = make([]*guide.Record, 0)
		for _, data := range data_array {
			if len(strings.TrimSpace(data.F_UnionId)) == 0 {
				data_json, _ := json.Marshal(&guide.Record{
					Id:       data.F_Id,
					Name:     data.F_RealName,
					FullName: data.F_DeptShortName + "-" + data.F_RealName,
				})
				guide.Logger.Infof("Sys_User :%s", string(data_json))
			}
			data_target_cond := &Ding_User{}
			data_target_cond.Unionid = data.F_UnionId
			ok, err := session.Get(data_target_cond)
			if err != nil {
				GetLogEntry().Errorln(err)
				return nil, err
			}
			data_info := &guide.Record{}
			data_info.Id = data.F_UnionId
			data_info.Name = data.F_RealName
			data_info.FullName = data.F_DeptShortName + "-" + data.F_RealName
			data_info.IsValid = ok
			resp.Data = append(resp.Data, data_info)
			if !ok {
				data_json, _ := json.Marshal(&guide.Record{
					Id:       data_info.Id,
					Name:     data_info.Name,
					FullName: data_info.FullName,
				})
				guide.Logger.Infof("Sys_User :%s", string(data_json))
			}
		}
	}
	child_cond := &Sys_Organize{}
	child_cond.F_ParentId = resp.Id
	child_array := make([]*Sys_Organize, 0)
	if err := session.Find(&child_array, child_cond); err != nil {
		GetLogEntry().Errorln(err)
		return nil, err
	}
	if len(child_array) > 0 {
		resp.Itmes = make([]*guide.Record, 0)
		for _, item := range child_array {
			child, err := item.GetRecord()
			if err != nil {
				GetLogEntry().Errorln(err)
				return nil, err
			}
			resp.Itmes = append(resp.Itmes, child)
		}
	}
	return resp, nil
}
