package main

import (
	"github.com/LingNovo/syncdata/guide"
	"github.com/LingNovo/syncdata/logic"
)

func main() {
	defer guide.ClearDB()
	// 初始化数据
	if guide.SyncCfg.IsSyncTable {
		if err := logic.SyncTable(); err != nil {
			logic.GetLogEntry().Errorln(err)
			return
		}
	}
	// Sys_Organize 字段 ShortName 数据类型需要改，长度不足
	// 在Sys_User,Sys_Organize 新增 保存部门全部路径的字段
	// 同步部门时根据全路径作为唯一标识
	// 同步用户时根据用户名称加所在部门全路径为唯一标识
	// 后续此方式废弃
	if guide.SyncCfg.IsPaddingDept {
		if err := logic.PaddingDeptShortName(guide.SyncCfg.RootOrgId); err != nil {
			logic.GetLogEntry().Errorln(err)
			return
		}
	}
	if guide.SyncCfg.IsPullData { //是否重新从钉钉平台获取数据
		tables := guide.SyncCfg.Tables
		for _, table := range tables {
			if table.IsClear {
				// 清除缓存表
				if err := guide.TruncateTable(table.Source); err != nil {
					logic.GetLogEntry().Errorln(err)
					return
				}
			}
			if table.IsBackUp {
				// 备份目标表
				if err := guide.BackUpTable(table.Target); err != nil {
					logic.GetLogEntry().Errorln(err)
					return
				}
			}

		}
		// 同步部门、用户数据
		if err := logic.SyncDepartmentList(1); err != nil {
			logic.GetLogEntry().Errorln(err)
			return
		}
		// 同步角色、用户角色数据
		//if err := logic.SyncRoleList(); err != nil {
		//	logic.GetLogEntry().Errorln(err)
		//	return
		//}
	}
	if guide.SyncCfg.IsSyncData {
		// 根据缓存表与目标表配置信息同步数据
		if err := logic.SyncDataByConfig(&guide.SyncCfg); err != nil {
			logic.GetLogEntry().Errorln(err)
		}
	}
}
