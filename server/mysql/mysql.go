package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Conn struct {
	db *sql.DB
}

var Default *Conn = nil

func NewConn(source string) (*Conn, error) {
	con, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}

	return &Conn{
		db: con,
	}, nil
}

func InitDefault(source string) (err error) {
	Default, err = NewConn(source)
	if err != nil {
		return err
	}

	return Default.CreateTable()
}

func (c *Conn) CreateDatabase() error {
	_, err := c.db.Exec(CreateDatabase)
	return err
}

func (c *Conn) CreateTable() error {
	_, err := c.db.Exec(CreateTable)
	return err
}
