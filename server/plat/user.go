package plat

import (
	"errors"
	"strconv"
)

// User user
type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	inWhichRoom *Room
}

// AddUser add user
func (p *Platform) AddUser(name, pass string) (map[string]string, error) {
	mp, err := p.database.Get(name)
	if err != nil {
		return nil, err
	}
	if len(mp) == 0 {
		return nil, errors.New("No Such User")
	}

	if !checkPassword(mp["pass"], pass) {
		return nil, errors.New("Invalid Password")
	}

	id, err := strconv.Atoi(mp["id"])
	if err != nil {
		return nil, err
	}

	if _, ok := p.UserTable[id]; !ok {
		u := User{
			ID:          id,
			Name:        name,
			inWhichRoom: nil,
		}

		p.UserTable[id] = &u
	}

	// TODO: AddConn
	if c, ok := p.ConnPool[id]; !ok || c == nil {
		p.waitConn = append(p.waitConn, id)
	}

	return mp, nil
}

// DelUser delete user
func (p *Platform) DelUser(id int) {
	delete(p.UserTable, id)
	delete(p.ConnPool, id)
}

func checkPassword(saved, pass string) bool {
	return saved == pass
}
