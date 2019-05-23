package plat

// Room room
type Room struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Nums int    `json:"nums"`

	inRoom []*User
}

func (r *Room) removeUser(uid int) {
	for i, e := range r.inRoom {
		if e.ID == uid {
			r.inRoom = append(r.inRoom[:i], r.inRoom[i+1:]...)
			r.Nums--
		}
	}
}

// GetRoomList return rmlist
func (p *Platform) GetRoomList() []*Room {
	list := make([]*Room, 0, len(p.RoomTable))
	for _, v := range p.RoomTable {
		list = append(list, v)
	}
	return list
}

// CreateRoom create room
func (p *Platform) CreateRoom(name string) {
	id := p.getUnusedID()
	if id == -1 {
		return
	}

	r := Room{
		ID:     id,
		Name:   name,
		inRoom: make([]*User, 0, 5),
		Nums:   0,
	}

	p.RoomTable[r.ID] = &r
}

// DeleteRoom delete room
func (p *Platform) DeleteRoom(id int) {
	if r, ok := p.RoomTable[id]; ok {
		for _, e := range r.inRoom {
			e.inWhichRoom = p.RoomTable[1]
		}
		delete(p.RoomTable, id)
	}
}

// Enter room
func (p *Platform) Enter(uid, rid int) {
	if u, ok := p.UserTable[uid]; ok {
		if r, ok := p.RoomTable[rid]; ok {
			u.inWhichRoom = r
			r.inRoom = append(r.inRoom, u)
			r.Nums++
		}
	}
}

// Leave room
func (p *Platform) Leave(uid int) {
	if u, ok := p.UserTable[uid]; ok {
		u.inWhichRoom = p.RoomTable[1]
	}
}
