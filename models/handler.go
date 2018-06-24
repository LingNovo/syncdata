package models

import (
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/LingNovo/syncdata/guide"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

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

// 数据接口
type DataHandler interface {
	// 获取数据表名称
	String() string
	// 清除主键值
	ClearPk()
	// 将数据转化为map结构
	ToMap() map[string]interface{}
	// 从map结构里解析数据
	Parse(in map[string]interface{}, columns []guide.Column)
	// 从源数据接口中匹配数据
	Match(source DataHandler, columns []guide.Column)
}

var internalFns = map[string]func(args ...interface{}) (interface{}, error){
	"fn_uuid":            fn_uuid,
	"fn_dept_parentid":   fn_dept_parentid,
	"fn_dept_categoryid": fn_dept_categoryid,
	"fn_getdate":         fn_getdate,
	"fn_user_deptid":     fn_user_deptid,
	"fn_account":         fn_account,
}

func fn_uuid(args ...interface{}) (interface{}, error) {
	uid, _ := uuid.NewV4()
	return uid.String(), nil
}

func fn_dept_parentid(args ...interface{}) (interface{}, error) {
	var (
		du  map[string]interface{}
		ok  bool
		err error
	)
	if du, ok = args[0].(map[string]interface{}); !ok {
		return "", nil
	}
	parentId, _ := strconv.ParseInt(fmt.Sprint(du["Parentid"]), 10, 64)
	if parentId == 1 {
		return "", nil
	}
	session := guide.GetDB().NewSession()
	defer session.Close()
	ding_dept_parent := Ding_Department{Id: parentId}
	if ok, err = session.Get(&ding_dept_parent); !ok {
		return "", err
	}
	dept := Sys_Organize{F_UnionId: ding_dept_parent.Id}
	if ok, err = session.Get(&dept); !ok {
		return "", err
	}
	return dept.F_Id, nil
}

func fn_dept_categoryid(args ...interface{}) (interface{}, error) {
	var (
		du map[string]interface{}
		ok bool
	)
	if du, ok = args[0].(map[string]interface{}); !ok {
		return "", nil
	}
	parentId, _ := strconv.ParseInt(fmt.Sprint(du["Parentid"]), 10, 64)
	groupContainSubDept, _ := strconv.ParseBool(fmt.Sprint(du["GroupContainSubDept"]))
	if parentId == 1 {
		return "Company", nil
	}
	if parentId > 1 && groupContainSubDept {
		return "Department", nil
	}
	return "WorkGroup", nil
}

func fn_getdate(args ...interface{}) (interface{}, error) {
	return time.Now().Local().Format("2006-01-02 15:04:05"), nil
}

func fn_user_deptid(args ...interface{}) (interface{}, error) {
	var (
		du  map[string]interface{}
		ok  bool
		err error
	)
	if du, ok = args[0].(map[string]interface{}); !ok {
		return "", nil
	}
	session := guide.GetDB().NewSession()
	defer session.Close()
	deptShortName := du["DeptShortName"].(string)
	dept := Sys_Organize{F_ShortName: deptShortName}
	if ok, err = session.Get(&dept); !ok {
		return "", err
	}
	return dept.F_Id, nil
}

func fn_account(args ...interface{}) (interface{}, error) {
	var (
		count int64
		err   error
	)
	session := guide.GetDB().NewSession()
	defer session.Close()
	var su Sys_User
	count, err = session.Count(&su)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("HYC-%d", count+1), nil
}
