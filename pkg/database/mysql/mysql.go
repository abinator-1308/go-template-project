package mysql

import "github.com/devlibx/gox-base/util"
import _ "github.com/go-sql-driver/mysql"

type MySQLConfig struct {
	Host         string `json:"host" yaml:"host"`
	Port         int    `json:"port" yaml:"port"`
	User         string `json:"user" yaml:"user"`
	Password     string `json:"password" yaml:"password"`
	Db           string `json:"db" yaml:"db"`
	TablePrefix  string
	TablePostfix string
}

func (m *MySQLConfig) SetupDefaults() {
	if util.IsStringEmpty(m.Host) {
		m.Host = "localhost"
	}
	if m.Port <= 0 {
		m.Port = 3306
	}
	if util.IsStringEmpty(m.User) {
		m.User = "test"
	}
	if util.IsStringEmpty(m.Password) {
		m.Password = "test"
	}
	if util.IsStringEmpty(m.Db) {
		m.Db = "test"
	}
}
