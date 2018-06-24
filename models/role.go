package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/satori/go.uuid"

	"github.com/LingNovo/syncdata/guide"
)

// 钉钉角色信息
type Ding_Role struct {
	RoleId    int64  `xorm:"BigInt pk notnull"`
	RoleName  string `xorm:"Varchar(255)`
	GroupId   int64  `xorm:"BigInt`
	GroupName string `xorm:"Varchar(255)`
	// 是否为新增
	IsNew bool `xorm:"Bool"`
}

func (m *Ding_Role) String() string {
	return "Ding_Role"
}

func (m *Ding_Role) ClearPk() {
	m.RoleId = 0
}

func (m *Ding_Role) ToMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["RoleId"] = m.RoleId
	resp["RoleName"] = strings.TrimSpace(m.RoleName)
	resp["GroupId"] = m.GroupId
	resp["GroupName"] = strings.TrimSpace(m.GroupName)
	return resp
}

func (m *Ding_Role) Parse(in map[string]interface{}, columns []guide.Column) {
	var err error
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

func (m *Ding_Role) Match(source DataHandler, columns []guide.Column) {
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

// 目标库角色信息
type Sys_Role struct {
	// 角色主键
	F_Id      string `xorm:"Varchar(50) pk notnull"`
	F_UnionId int64  `xorm:"BigInt"`
	// 组织主键
	F_OrganizeId string `xorm:"Varchar(50)"`
	// 分类:1-角色2-岗位
	F_Category int `xorm:Int`
	// 编号
	F_EnCode string `xorm:"Varchar(50)"`
	// 名称
	F_FullName string `xorm:"Varchar(50)"`
	// 类型
	F_Type string `xorm:"Varchar(50)"`
	// 允许编辑
	F_AllowEdit bool `xorm:"Bool"`
	// 允许删除
	F_AllowDelete bool `xorm:"Bool"`
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
	// 是否为新增
	isNew bool `xorm:"-"`
}

func (m *Sys_Role) String() string {
	return "Sys_Role"
}

func (m *Sys_Role) ClearPk() {
	m.F_Id = ""
}

func (m *Sys_Role) ToMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["F_Id"] = strings.TrimSpace(m.F_Id)
	resp["F_UnionId"] = m.F_UnionId
	resp["F_OrganizeId"] = strings.TrimSpace(m.F_OrganizeId)
	resp["F_Category"] = m.F_Category
	resp["F_EnCode"] = strings.TrimSpace(m.F_EnCode)
	resp["F_FullName"] = strings.TrimSpace(m.F_FullName)
	resp["F_Type"] = strings.TrimSpace(m.F_Type)
	resp["F_AllowEdit"] = m.F_AllowEdit
	resp["F_AllowDelete"] = m.F_AllowDelete
	resp["F_SortCode"] = m.F_SortCode
	resp["F_DeleteMark"] = m.F_DeleteMark
	resp["F_EnabledMark"] = m.F_EnabledMark
	resp["F_Description"] = strings.TrimSpace(m.F_Description)
	if !m.F_CreatorTime.IsZero() {
		resp["F_CreatorTime"] = m.F_CreatorTime
	}
	resp["F_CreatorUserId"] = m.F_CreatorUserId
	if !m.F_LastModifyTime.IsZero() {
		resp["F_LastModifyTime"] = m.F_LastModifyTime
	}
	resp["F_LastModifyUserId"] = m.F_LastModifyUserId
	if !m.F_DeleteTime.IsZero() {
		resp["F_DeleteTime"] = m.F_DeleteTime
	}
	resp["F_DeleteUserId"] = m.F_DeleteUserId
	return resp
}

func (m *Sys_Role) Parse(in map[string]interface{}, columns []guide.Column) {
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
	if value, ok := in["F_OrganizeId"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_OrganizeId = fmt.Sprint(value)
		}
	}
	if value, ok := in["F_Category"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			if m.F_Category, err = strconv.Atoi(fmt.Sprint(value)); err != nil {
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
	if value, ok := in["F_Type"]; ok {
		if value != nil && len(strings.TrimSpace(fmt.Sprint(value))) > 0 {
			m.F_Type = fmt.Sprint(value)
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

func (m *Sys_Role) Match(source DataHandler, columns []guide.Column) {
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

func (m *Sys_Role) Sync(table *guide.Table) error {
	session := guide.GetDB().NewSession()
	defer session.Close()
	targets := make([]*Sys_Role, 0)
	if err := session.Find(&targets); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	sources := make([]*Ding_Role, 0)
	if err := session.Find(&sources); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}

	if len(targets) == 0 {
		in_array := make([]*Sys_Role, 0)
		for _, source := range sources {
			target := &Sys_Role{}
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
	unique_uped_maps := make(map[string]*Ding_Role)
	up_maps := make(map[*Sys_Role]*Sys_Role)
	session = guide.GetDB().NewSession()
	for _, source := range sources {
		sMap := source.ToMap()
		for _, target := range targets {
			isOk := false
			tMap := target.ToMap()
			cond := &Sys_Role{}
			condMap := cond.ToMap()
			cond_s := &Ding_Role{}
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
				cond_s_array := make([]*Ding_Role, 0)
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
				up_target := &Sys_Role{}
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
	in_array := make([]*Sys_Role, 0)
	session = guide.GetDB().NewSession()
	defer session.Close()
	for _, source := range sources {
		sMap := source.ToMap()
		cond := &Ding_Role{}
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
		cond_array := make([]*Ding_Role, 0)
		if err := session.Find(&cond_array, cond); err != nil {
			GetLogEntry().Errorln(err)
			return err
		}
		if len(cond_array) > 1 {
			continue
		}
		unique_in_key := strings.Join(unique_in_key_array, "_")
		if _, ok := unique_uped_maps[unique_in_key]; !ok {
			target := &Sys_Role{}
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
