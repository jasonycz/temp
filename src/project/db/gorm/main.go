package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
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
	defaultAnd    = "&"
	defaultFormat = "%s%s=%s" + defaultAnd
	defaultDrive  = "mysql"
)

// DBConf db configuration
type DBConf struct {
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
func (c *DBConf) Open(ping bool) (*gorm.DB, error) {
	db, err := gorm.Open(c.Driver, c.EncodeDSN())
	if err != nil {
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

// EncodeDSN encode DBConf sturct to string
func (c *DBConf) EncodeDSN() string {
	return strings.TrimRight(fmt.Sprintf(
		"%s:%s@%s(%s:%d)/%s?%s", c.Username, c.Password,
		c.Protocol, c.IP, c.Port, c.DefaultDB,
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

// Pool db pool
type Pool struct {
	sync.Map
}

var (
	connInfo map[string]*Pool
	pool     Pool
	keyList  []string
	cluster  = "eshop"
)

// pool 的 key为 DBConf 的 md5
func init() {
	connInfo = map[string]*Pool{}
	confList := []DBConf{
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
			// continue
			panic(err)
		}
		key := fmt.Sprintf("%s_%d", dbConf.IP, dbConf.Port)
		keyList = append(keyList, key)
		pool.Store(key, db)
	}

	connInfo[cluster] = &pool
}

func poolTest(w http.ResponseWriter, r *http.Request) {
	var user User
	var db *gorm.DB

	value, hi := connInfo[cluster].Load("127.0.0.1_3306")
	fmt.Printf("%v============hi %v", value, hi)

	confList := []DBConf{
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
	keyListTmp := []string{}
	for _, dbConf := range confList {
		db, err := dbConf.Open(true)
		key := fmt.Sprintf("%s_%d", dbConf.IP, dbConf.Port)
		if err != nil {
			connInfo[cluster].Delete(key)
			continue
		}
		keyListTmp = append(keyListTmp, key)
		actual, loaded := connInfo[cluster].LoadOrStore(key, db)
		// loaded == true then need to close db
		if loaded {
			db.Close()
		}
		fmt.Printf("actual is %v, loaded is %t", actual, loaded)
	}
	keyList = keyListTmp

	if _, ok := connInfo[cluster]; !ok {
		panic(fmt.Sprintf("cluster %s is not exits", cluster))
	}

	fmt.Printf("||||\n%+v\n||||", connInfo[cluster])

	poolVal, ok := connInfo[cluster].Load(keyList[randSigleton.Intn(len(keyList))])
	if !ok {
		panic("not exits")
	}
	db = poolVal.(*gorm.DB)
	db.First(&user, 1)
	// db.First(&user, "name = ?", "3307")
	text, _ := json.Marshal(user)
	w.Write(text)

}

func main() {
	startHTTPServer()
}

func startHTTPServer() {
	http.HandleFunc("/pool", poolTest)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
