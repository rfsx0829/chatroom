package plat

// User user
type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	inWhichRoom *Room
}

func (u *User) enterRoom(r *Room) {
	u.inWhichRoom = r
	r.inRoom = append(r.inRoom, u)
}

// AddUser add user
func (p *Platform) AddUser(id int, name string) {
	if _, ok := p.UserTable[id]; !ok {
		u := User{
			ID:          id,
			Name:        name,
			inWhichRoom: p.RoomTable[1],
		}

		p.UserTable[id] = &u
	}

	// TODO: AddConn
	if c, ok := p.ConnPool[id]; !ok || c == nil {
		p.waitConn = append(p.waitConn, id)
	}
}

// DelUser delete user
func (p *Platform) DelUser(id int) {
	delete(p.UserTable, id)
	delete(p.ConnPool, id)
}
