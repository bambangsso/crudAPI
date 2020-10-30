package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
  
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/gomodule/redigo/redis"
)

var MPosDB *sql.DB
var MPosGORM *gorm.DB
var err error
var RedisConn redis.Conn
var pool *redis.Pool


func InitPostgres() {
	//connString := "user=bambang_sso password=suksesmulia dbname=bookmefy sslmode=disable"
	connString := "user=bambang_susilo password=suksesmulia dbname=mpos sslmode=disable"
	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	MPosDB = db
}

////////////////////////////////////PostgreSQL Fuction///////////////////////////////////////////
func InitGormPostgres() {
	MPosGORM, err = gorm.Open("postgres", "user=bambang_susilo dbname=mpos password=suksesmulia sslmode=disable")
	if err != nil {
		panic(err)
	}

	//defer MPosGORM.Close()
}

////////////////////////////////////Redis Fuction///////////////////////////////////////////

func InitRedis() {
	pool = RedisNewPool()
	RedisConn = pool.Get()
	//RedisConn.Close()
}


func RedisNewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle: 80, // Maximum number of idle connections in the pool.
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}
////////////////////////////////////////////////////////////////////////////////////////////


func RedisSet(c redis.Conn, key string, value string) error {
	cc := pool.Get()
	defer cc.Close()
	_, err := cc.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}

func RedisGet(c redis.Conn, key string) string {
	cc := pool.Get()
	defer cc.Close()	
	s, err := redis.String(cc.Do("GET", key))
	if err != nil {
		return "ERR"
	}
  	return s
}

func RedisPing(c redis.Conn) error {
	cc := pool.Get()
	defer cc.Close()		
	s, err := redis.String(cc.Do("PING"))
	if err != nil {
		return err
	}
	fmt.Printf("Redis PING Response = %s\n", s)

	return nil
}

func RedisDelete(c redis.Conn, key string) error {
	cc := pool.Get()
	defer cc.Close()		
	_, err := cc.Do("DEL", key)
	if err != nil {
		return err
	}
	return nil
}
////////////////////////////////////////////////////////////////////////////////////////////
