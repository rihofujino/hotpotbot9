package db

import (
	"database/sql"
)

//ServerInfo ...
type ServerInfo struct {
	User     string
	Password string
	Host     string
}

var (
	instance *ServerInfo
)

func init() {
	instance = &ServerInfo{
		User:     "root",
		Password: "gyozabot",
		Host:     "tcp(mysql:3306)",
	}
}

//OpenMysql ...
func OpenMysql() (*sql.DB, error) {
	dbServerInfo := GetDbServerInfo()
	db, err := sql.Open("mysql", dbServerInfo.User+":"+dbServerInfo.Password+"@"+dbServerInfo.Host+"/"+"?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// GetDbServerInfo ...
func GetDbServerInfo() *ServerInfo {
	return instance
}
