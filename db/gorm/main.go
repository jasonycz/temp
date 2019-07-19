package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// User user model
type User struct {
	ID              int64
	Name            string
	Email           string
	EmailVerifiedAt string
	Password        string
	RememberToken   string
	CreatedAt       string
	UpdatedAt       string
}

const (
	defaultDSNFormat = "%s:%s@%s(%s:%d)/%s?%s"
	defaultAnd       = "&"
	defaultFormat    = "%s%s=%s" + defaultAnd
	defaultDrive     = "mysql"
)

// Conf db configuration
type Conf struct {
	Username        string
	Password        string
	Protocol        string
	DefaultDB       string
	Parameters      Parameters
	Port            int
	IP              string
	Driver          string
	ConnMaxLifeTime int64
	MaxIdleConns    int
	MaxOpenConns    int
}

// Parameters parameters detail: https://github.com/go-sql-driver/mysql#dsn-data-source-name
type Parameters map[string]interface{}

// Open provides method to opening a database link with DBConfig struct
func (c *Conf) Open(ping bool) (*gorm.DB, error) {
	db, err := gorm.Open(c.Driver, c.EncodeDSN())
	// defer db.Close()
	if nil != err {
		return db, err
	}

	if c.MaxOpenConns > 0 {
		db.DB().SetMaxOpenConns(c.MaxOpenConns)
	}
	if c.MaxIdleConns > 0 {
		db.DB().SetMaxIdleConns(c.MaxIdleConns)
	}
	if c.ConnMaxLifeTime > 0 {
		db.DB().SetConnMaxLifetime(time.Duration(c.ConnMaxLifeTime) * time.Second)
	}
	if ping {
		err = db.DB().Ping()
	}

	return db, err
}

// EncodeDSN encode Conf sturct to string
func (c *Conf) EncodeDSN() string {
	port := c.Port
	ip := c.IP

	return strings.TrimRight(fmt.Sprintf(
		defaultDSNFormat, c.Username, c.Password,
		c.Protocol, ip, port, c.DefaultDB,
		encodeParamseters(c.Parameters)), "?",
	)
}

//Private: encode Paramseters(map[string]string) to string
func encodeParamseters(p Parameters) string {
	str := ""
	for k, v := range p {
		if v == "" {
			continue
		}
		str = fmt.Sprintf(defaultFormat, str, k, v)
	}

	return strings.TrimRight(str, defaultAnd)
}

var (
	randSigleton = rand.New(rand.NewSource(time.Now().UnixNano()))
	db           *gorm.DB
)

// DBPool db pool
type DBPool struct {
	Topic string `json:"cluster"`
	// 使用sync Map
	SercvicePool map[string]*Service
}

// Service db service
type Service struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
	DB   *gorm.DB
}

var (
	dbPool  DBPool
	keyList []string
)

// pool 的 key为 conf 的 md5
func init() {
	dbPool = DBPool{
		Topic:        "eshop",
		SercvicePool: map[string]*Service{},
	}

	confList := []Conf{
		{
			Username:        "root",
			Password:        "root",
			Protocol:        "tcp",
			DefaultDB:       "eshop",
			Port:            3306,
			IP:              "127.0.0.1",
			Driver:          "mysql",
			ConnMaxLifeTime: 120,
			MaxIdleConns:    12,
			MaxOpenConns:    10,
		},
		{
			Username:        "root",
			Password:        "root",
			Protocol:        "tcp",
			DefaultDB:       "eshop",
			Port:            3307,
			IP:              "127.0.0.1",
			Driver:          "mysql",
			ConnMaxLifeTime: 120,
			MaxIdleConns:    12,
			MaxOpenConns:    10,
		},
	}

	for _, dbConf := range confList {
		db, err := dbConf.Open(true)
		if err != nil {
			panic(err)
		}
		key := fmt.Sprintf("%s_%d", dbConf.IP, dbConf.Port)
		keyList = append(keyList, key)
		dbPool.SercvicePool[key] = &Service{
			IP:   dbConf.IP,
			Port: dbConf.Port,
			DB:   db,
		}
	}

	// gormDB, err := gorm.Open("mysql", "root:root@/eshop?charset=utf8&parseTime=True&loc=Local")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// fmt.Printf("%+v", gormDB)
	// gormDB.DB().Ping()

	// db = gormDB
	// defer gormDB.Close()

	// db.DB().SetMaxOpenConns(10) // 最大连接数为 10
	// db.DB().SetMaxIdleConns(5)  // 最大空闲连接数 5
	// db.DB().SetConnMaxLifetime(1 * time.Minute)
}

func pool(w http.ResponseWriter, r *http.Request) {
	var user User

	db := dbPool.SercvicePool[keyList[randSigleton.Intn(len(keyList))]].DB

	// fmt.Printf("%+v", db)
	db.First(&user, 1) // 找到id为1的产品
	fmt.Printf("%+v", user)

}

func main() {

	startHTTPServer()
	//自动检查 User 结构是否变化，变化则进行迁移
	// db.AutoMigrate(&User{})

	// 增
	// db.Create(&User{Code: "L1212", Price: 1000})

	// 查
	// var user User
	// db.First(&user, 1) // 找到id为1的产品

	// fmt.Printf("%+v", user)
	// db.First(&User, "code = ?", "L1212") // 找出 code 为 l1212 的产品

	// 改 - 更新产品的价格为 2000
	// db.Model(&User).Update("Price", 2000)

	// 删 - 删除产品
	// db.Delete(&User)
}

func startHTTPServer() {
	http.HandleFunc("/pool", pool)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
