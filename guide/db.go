package guide

import (
	"fmt"
	"time"

	"encoding/xml"
	"io/ioutil"
	"runtime"

	"github.com/Unknwon/goconfig"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	confPath    = "./conf/app.conf"
	db          *xorm.Engine
	Conf        *goconfig.ConfigFile
	err         error
	syncCfgPath = "./conf/sync.xml"
	SyncCfg     SyncConfig
	logPath     = "./data/base.log"
	Logger      *logrus.Logger
)

func init() {
	writer, _ := rotatelogs.New(logPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(60400)*time.Second))
	Logger = logrus.New()
	Logger.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  writer,
			logrus.DebugLevel: writer,
			logrus.ErrorLevel: writer,
			logrus.FatalLevel: writer,
			logrus.WarnLevel:  writer,
			logrus.PanicLevel: writer,
		},
		&logrus.JSONFormatter{},
	))
	Conf, err = goconfig.LoadConfigFile(confPath)
	if err != nil {
		GetLogEntry().Fatalf("load config: %s", err)
	}
	initDB()
	initSC()
}

// 初始化数据库信息
func initDB() {
	db_driver := Conf.MustValue("db", "db_driver")
	db_server := Conf.MustValue("db", "db_server")
	db_port := Conf.MustValue("db", "db_port")
	db_datebase := Conf.MustValue("db", "db_datebase")
	db_username := Conf.MustValue("db", "db_username")
	db_password := Conf.MustValue("db", "db_password")
	db_driver_connstr := Conf.MustValue("db", "db_driver_connstr")
	dbconnstr := fmt.Sprintf(db_driver_connstr, db_server, db_username, db_password, db_port, db_datebase)
	if db, err = xorm.NewEngine(db_driver, dbconnstr); err != nil {
		GetLogEntry().Fatalf("new db engine: %s", err)
	}
	db.TZLocation = time.Local
	db.SetMaxIdleConns(3000)
	db.SetMaxOpenConns(1500)
	db.ShowExecTime(true)
	db.SetMapper(&core.SameMapper{})
	db.SetColumnMapper(&core.SameMapper{})
	db.ShowSQL(true)
}

func ClearDB() {
	if db != nil {
		db.Close()
	}
}

// 初始化同步数据配置
func initSC() {
	data, err := ioutil.ReadFile(syncCfgPath)
	if err != nil {
		GetLogEntry().Fatalf("load sync config: %s", err)
	}
	if err = xml.Unmarshal(data, &SyncCfg); err != nil {
		GetLogEntry().Fatalf("xml Unmarshal syncCfg: %s", err)
	}
}
func GetDB() *xorm.Engine {
	if db == nil {
		initDB()
	}
	db.Ping()
	return db
}

// 获取日志实例
func GetLogEntry() *logrus.Entry {
	entry := logrus.NewEntry(Logger)
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

// 清空指定数据表
func TruncateTable(table string) error {
	session := GetDB().NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	dml := fmt.Sprintf(`TRUNCATE TABLE "%s"`, table)
	if _, err := session.Exec(dml); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	if err := session.Commit(); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	return nil
}

func BackUpTable(table string) error {
	session := GetDB().NewSession()
	defer session.Close()
	key := time.Now().Unix()
	table_bak := fmt.Sprintf(`%s_%d`, table, key)
	dml := fmt.Sprintf(`select * into "%s" from "%s"`, table_bak, table)
	if err := session.Begin(); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	if _, err := session.Exec(dml); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	if err := session.Commit(); err != nil {
		GetLogEntry().Errorln(err)
		return err
	}
	return nil
}
