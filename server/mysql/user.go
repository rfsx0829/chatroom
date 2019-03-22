package mysql

import (
	"errors"
	"fmt"
)

type Profile struct {
	Uid   int    `json:"uid"`
	Name  string `json:"name"`
	Pass  string `json:"pass"`
	Email string `json:"email"`
}

func (c *Conn) AddUser(pf *Profile) error {
	_, err := c.db.Exec(Insert, pf.Name, pf.Pass, pf.Email)
	return err
}

func (c *Conn) DeleteUser(uid int) error {
	_, err := c.db.Exec(Delete + fmt.Sprintf("where uid=%d", uid))
	return err
}

func (c *Conn) CheckUserName(name string) (bool, error) {
	rows, err := c.db.Query(Select + fmt.Sprintf("where name='%s'", name))
	if err != nil {
		return false, err
	}

	return rows.Next(), nil
}

func (c *Conn) GetUserInfo(name, email string) (*Profile, error) {
	var (
		pf  Profile
		err error
		str = Select
	)

	if name == "" {
		str += fmt.Sprintf("email='%s'", email)
	} else {
		str += fmt.Sprintf("name='%s'", name)
	}

	rows, err := c.db.Query(str)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&pf.Uid, &pf.Name, &pf.Pass, &pf.Email)
		return &pf, err
	}

	return nil, errors.New("Invalid name or email !")
}
