package mysql

import "errors"

const (
	TableName      string = "wow.user"
	Insert         string = "insert into " + TableName + "(name, pass, email)values(?, ?, ?)"
	Select         string = "select * from " + TableName + " "
	Delete         string = "delete from " + TableName + " "
	Update         string = "update " + TableName + " set "
	CreateDatabase string = "create database wow if not exists "
	CreateTable    string = `
        create table if not exists ` + TableName + `(
		uid      int unsigned  auto_increment,
		name     varchar(20)   not null,
		pass     varchar(64)   not null,
		email    varchar(30)   not null,
        primary key(uid))
		engine=InnoDB default charset=utf8`
)

var (
	InvalidUID      error = errors.New("Invalid UID !")
	InvalidName     error = errors.New("Invalid Name !")
	InvalidEmail    error = errors.New("Invalid Email !")
	InvalidPassword error = errors.New("Invalid Password !")
)
