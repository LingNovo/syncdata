package guide

import (
	"encoding/xml"
)

// 数据表之间同步数据配置信息
type SyncConfig struct {
	XMLName xml.Name `xml:"SyncConfig"`
	// Sys_Organize组织机构根节点Id
	RootOrgId string `xml:"RootOrgId,attr"`
	// 同步数据之前是否先填充部门全路径
	IsPaddingDept bool `xml:"IsPaddingDept,attr"`
	// 是否同步表结构
	IsSyncTable bool `xml:"IsSyncTable,attr"`
	// 同步数据之前是否先从钉钉平台获取数据
	IsPullData bool `xml:"IsPullData,attr"`
	// 是否从缓存表同步数据
	IsSyncData bool `xml:"IsSyncData,attr"`
	// 数据表集合
	Tables []Table `xml:"Table"`
}

// 数据表配置信息
type Table struct {
	XMLName xml.Name `xml:"Table"`
	// 源表名称
	Source string `xml:"Source,attr"`
	// 目标表名称
	Target string `xml:"Target,attr"`
	// 是否清除缓存表
	IsClear bool `xml:"IsClear,attr"`
	// 是否备份表
	IsBackUp bool `xml:"IsBackUp,attr"`
	// 描述信息
	Description string `xml:"Description,attr"`
	// 数据列配置信息
	Columns []Column `xml:"Column"`
}

// 数据列配置信息
type Column struct {
	XMLName xml.Name `xml:"Column"`
	// 源列名称
	Source string `xml:"Source,attr"`
	// 目标列名称
	Target string `xml:"Target,attr"`
	// 唯一标识
	Unique bool `xml:"Unique,attr"`
	// 默认值
	DefaultValue string `xml:"DefaultValue,attr"`
	// 默认值类型
	DefaultType string `xml:"DefaultType,attr"`
	// 描述信息
	Description string `xml:"Description,attr"`
}
