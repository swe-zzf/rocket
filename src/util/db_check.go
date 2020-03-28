package util

import (
	"database/sql"
	"github.com/astaxie/beego/toolbox"
)

type DatabaseCheck struct {
	dbUrl string
}

func (dbc *DatabaseCheck) Check() error {
	_, err := sql.Open("mysql", dbc.dbUrl)
	if err != nil {
		return err
	}
	return nil
}

func RegisterHealthCheck(dbUrl string) {
	toolbox.AddHealthCheck("database", &DatabaseCheck{dbUrl: dbUrl})
}
