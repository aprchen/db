package db

import (
	"database/sql"
	"github.com/containous/traefik/safe"
	"github.com/didi/gendry/manager"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"time"
)

var (
	mysqlOnce     sync.Once
	mysqlInstance *MysqlService
)
// before get data need use function LoadConfiguration to link mysql
// e.g: err := db.Mysql().LoadConfiguration(db.MysqlMessageFromEnv());
// then get data like this:
// rows, err := db.Mysql().Master.DB().Query(cond, vals...)
func Mysql() *MysqlService {
	mysqlOnce.Do(func() {
		mysqlInstance = &MysqlService{
			Master: &MysqlClient{
				db: &safe.Safe{},
			},
		}
	})
	return mysqlInstance
}

func (ms *MysqlService) LoadConfiguration(configMsg MysqlMessage) error {
	return ms.Master.init(configMsg)
}

type MysqlService struct {
	Master *MysqlClient
}

type MysqlClient struct {
	db *safe.Safe
}

func (mc *MysqlClient) DB() *sql.DB {
	if v, ok := mc.db.Get().(*sql.DB); ok {
		return v
	}
	panic("mysql need load configuration")
}

func (mc *MysqlClient) init(configMsg MysqlMessage) error {
	if err := configMsg.Check(); err != nil {
		return err
	}
	m, err := manager.New(configMsg.Name, configMsg.User, configMsg.Password, configMsg.Host).
		Set(
			manager.SetCharset("utf8"),
			manager.SetAllowCleartextPasswords(true),
			manager.SetInterpolateParams(true),
			manager.SetTimeout(1*time.Second),
			manager.SetReadTimeout(3*time.Second),
		).Port(configMsg.Port).Open(true)
	if err != nil {
		return err
	}
	mc.db.Set(m)
	return nil
}
