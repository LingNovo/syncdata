package logic

import (
	"runtime"

	"github.com/LingNovo/syncdata/guide"
	m "github.com/LingNovo/syncdata/models"
	"github.com/sirupsen/logrus"
)

// 同步表结构
func SyncTable() error {
	return guide.GetDB().Sync(
		new(m.Ding_User),
		//new(m.Ding_UserRole),
		new(m.Sys_User),
		new(m.Sys_UserLogOn),
		//new(m.Ding_Role),
		//new(m.Sys_Role),
		new(m.Ding_Department),
		new(m.Sys_Organize),
	)
}

// 获取日志实例
func GetLogEntry() *logrus.Entry {
	entry := logrus.NewEntry(guide.Logger)
	funcName, file, line, ok := runtime.Caller(1)
	if ok {
		return entry.WithFields(logrus.Fields{
			"FuncName": runtime.FuncForPC(funcName).Name(),
			"Line":     line,
			"File":     file,
		})
	}
	return entry
}

// 数据同步处理接口
type SyncProcessor interface {
	// 同步数据
	Sync(table *guide.Table) error
}

// 数据同步处理对象集合
var syncProcess = map[string]SyncProcessor{
	"Sys_User":      new(m.Sys_User),
	"Sys_UserLogOn": new(m.Sys_UserLogOn),
	"Sys_Role":      new(m.Sys_Role),
	"Sys_Organize":  new(m.Sys_Organize),
}

// 同步数据
// cfg 同步配置对象
func SyncDataByConfig(cfg *guide.SyncConfig) error {
	for _, table := range cfg.Tables {
		if value, ok := syncProcess[table.Target]; ok {
			if err := value.Sync(&table); err != nil {
				return err
			}
		}
	}
	return nil
}
