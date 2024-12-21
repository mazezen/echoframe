package vm

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/mazezen/echoframe/utils"
	"github.com/mazezen/itools"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

var err error

type Store struct {
	Log    *zap.Logger
	OrmLog *zap.Logger
	Ldb    *itools.LevelDB
	Db     *xorm.Engine
	Rdb    *redis.Client
	Mongo  *itools.MongoDb
}

var Box *Store

func BootStore() {
	var once sync.Once
	once.Do(func() {
		Box = &Store{
			Log:    newLogger(1),
			OrmLog: newLogger(2),
		}
		Box.newOrm()
		Box.newRedis()
		Box.newLevelDB()
		Box.newMongo()
	})
}

func (s *Store) newLevelDB() {
	pr, err := utils.FindProjectRoot()
	if err != nil {
		panic(err)
	}

	val := itools.Gm.Get("leveldb").(string)
	var p string
	if val == "" {
		p = fmt.Sprintf("%s%s", pr, "./leveldb_data")
	} else {
		p = fmt.Sprintf("%s%s", pr, val)
	}

	s.Ldb, err = itools.CreateLevelDB(p)
	if err != nil {
		panic(err)
	}
}

func (s *Store) newOrm() {
	m := itools.Gm.Get("mysql").(map[string]interface{})
	d, ok := m["dsn"].(string)
	if !ok {
		panic(ok)
	}
	show, ok := m["show"].(bool)
	if !ok {
		panic(ok)
	}
	s.Db, err = itools.NewXrm(d, show)
	if err != nil {
		panic(err)
	}
	s.Db.SetLogger(xorm.NewSimpleLogger(XLog))
}

func (s *Store) newRedis() {
	m := itools.Gm.Get("redis").(map[string]interface{})
	url, ok := m["url"].(string)
	if !ok {
		panic(ok)
	}
	password, ok := m["password"].(int)
	if !ok {
		panic(ok)
	}
	_password := strconv.Itoa(password)

	channel := m["db"].(int)
	s.Rdb, err = itools.NewRedis(url, _password, channel)
	if err != nil {
		panic(err)
	}
}

func (s *Store) newMongo() {
	m := itools.Gm.Get("mongodb").(string)
	s.Mongo, err = itools.NewMongoDb(nil, m)
	if err != nil {
		panic(err)
	}
}

func findProjectRoot() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
			return currentDir, nil
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			return "", fmt.Errorf("could not find go.mod file")
		}

		currentDir = parentDir
	}
}

var XLog *XLogger

type XLogger struct{}

func (*XLogger) Write(p []byte) (n int, err error) {
	Box.OrmLog.Info("数据库操作", zap.String("数据库", string(p)))
	return len(p), nil
}
