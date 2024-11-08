package conf

import (
	"fmt"
	"gitee.com/qiaogy91/K8sGenie/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Config struct {
	Http    *Http    `toml:"http" json:"http"`
	Grpc    *Grpc    `toml:"grpc" json:"grpc"`
	Mysql   *Mysql   `toml:"mysql" json:"mysql"`
	Rancher *Rancher `toml:"rancher" json:"rancher"`
	QARobot *QARobot `toml:"QARobot" json:"QARobot"`
}

type QARobot struct {
	AppID     string `json:"appID" toml:"appID"`
	AppSecret string `json:"appSecret" toml:"appSecret"`
}

type Rancher struct {
	Addr     string `toml:"addr" json:"addr" env:"RANCHER_GRPC_ADDR"`
	Port     int    `toml:"port" json:"port" env:"RANCHER_GRPC_PORT"`
	Token    string `toml:"token" json:"token" env:"RANCHER_GRPC_TOKEN"`
	KubeFile string `toml:"kube_file" json:"kube_file" env:"KUBE_FILE"`
}

type Http struct {
	Addr         string `toml:"addr" json:"addr" env:"HTTP_ADDR"`
	Port         int    `toml:"port" json:"port" env:"HTTP_PORT"`
	ReadTimeout  int    `toml:"read_timeout" json:"read_timeout" env:"HTTP_READ_TIMEOUT"`
	WriteTimeout int    `toml:"write_timeout" json:"write_timeout" env:"HTTP_WRITE_TIMEOUT"`
}

type Grpc struct {
	Addr  string `toml:"addr" json:"addr" env:"GRPC_ADDR"`
	Port  int    `toml:"port" json:"port" env:"GRPC_PORT"`
	Token string `toml:"token" json:"token" env:"GRPC_TOKEN"`
}

type Mysql struct {
	Addr     string `toml:"addr" json:"addr" env:"MYSQL_ADDR"`
	Port     int    `toml:"port" json:"port" env:"MYSQL_PORT"`
	Database string `toml:"database" json:"database" env:"MYSQL_DATABASE"`
	Username string `toml:"username" json:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" json:"password" env:"MYSQL_PASSWORD"`

	// 连接池参数
	MaxOpenConn int `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	MaxIdleConn int `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	MaxIdleTime int `toml:"max_idle_time" env:"MYSQL_MAX_IDLE_TIME"`

	lock sync.Mutex
	db   *gorm.DB
}

func (c *Config) RancherAddr() string { return fmt.Sprintf("%s:%d", c.Rancher.Addr, c.Rancher.Port) }

func (c *Config) HttpAddr() string {
	return fmt.Sprintf("%s:%d", c.Http.Addr, c.Http.Port)
}

func (c *Config) GrpcAddr() string {
	return fmt.Sprintf("%s:%d", c.Grpc.Addr, c.Grpc.Port)
}

func (c *Config) GetMysqlPool() (db *gorm.DB) {
	c.Mysql.lock.Lock()
	defer c.Mysql.lock.Unlock()

	if c.Mysql.db != nil {
		return c.Mysql.db
	}

	var err error
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		c.Mysql.Username,
		c.Mysql.Password,
		c.Mysql.Addr,
		c.Mysql.Port,
		c.Mysql.Database,
	)

	// 链接数据库
	c.Mysql.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true, // 执行任何 SQL 时都创建并缓存预编译语句，可以提高后续的调用速度
		SkipDefaultTransaction: true, // 对于写操作，默认Gorm 为了数据的完整性将其封装在事务中运行。如果没有这方面要求可关闭，性能会提升30%
	})

	if err != nil {
		panic("链接数据库实例失败::" + err.Error())
	}

	// 2.设置连接池：Gorm 将连接池对象设置为了DB对象的一个属性
	sqlDB, err := c.Mysql.db.DB()
	if err != nil {
		panic("设置数据库连接池失败::" + err.Error())
	}

	sqlDB.SetMaxOpenConns(c.Mysql.MaxOpenConn)
	sqlDB.SetMaxIdleConns(c.Mysql.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Hour * time.Duration(c.Mysql.MaxLifeTime))
	sqlDB.SetConnMaxIdleTime(time.Hour * time.Duration(c.Mysql.MaxIdleTime))
	common.L().Info().Msgf("Mysql 数据库链接已建立")

	return c.Mysql.db
}

func (c *Config) closeMysqlPool() error {
	if c.Mysql.db != nil {
		sqlDB, _ := c.Mysql.db.DB()
		if err := sqlDB.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (c *Config) CloseConnection() {
	if err := c.closeMysqlPool(); err != nil {
		common.L().Error().Msgf("close mysql err: %v", err)
	}
}

func NewConfig() *Config {
	return &Config{
		Http:    new(Http),
		Grpc:    new(Grpc),
		Mysql:   new(Mysql),
		Rancher: new(Rancher),
	}
}
