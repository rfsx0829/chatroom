package service

type Platform struct {
	OnlineList []*Person
	RoomList   []*Room
}

var Client *Platform

func init() {
	Client = NewPlatform()
}

func NewPlatform() *Platform {
	if Client != nil {
		return nil
	}
	return &Platform{
		OnlineList: make([]*Person, 0, 10),
		RoomList:   make([]*Room, 0, 2),
	}
}

func (p *Platform) getUnUsedId() int {
	var (
		id = len(p.RoomList)
		ok = false
	)

	for !ok {
		for _, e := range p.RoomList {
			if e.Id == id {
				id++
				break
			}
		}

		ok = true
	}

	return id
}

func (p *Platform) GetUserById(uid int) *Person {
	for _, e := range p.OnlineList {
		if e.Info.Uid == uid {
			return e
		}
	}
	return nil
}

func (p *Platform) GetRoomById(rid int) *Room {
	for _, e := range p.RoomList {
		if e.Id == rid {
			return e
		}
	}
	return nil
}
