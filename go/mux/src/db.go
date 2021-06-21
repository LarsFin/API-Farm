package apifarm

type DB interface {
	GetVideoGame(uint) *VideoGame
	GetAllVideoGames() []VideoGame
	AddVideoGame(VideoGame) VideoGame
}

type InMemory struct {
	idCounter  uint
	videoGames *[]VideoGame
}

func NewInMemory() *InMemory {
	return &InMemory{
		0,
		&([]VideoGame{}),
	}
}

func NewInMemoryForTests() (*InMemory, **[]VideoGame) {
	im := &InMemory{
		0,
		&([]VideoGame{}),
	}

	vgs := &im.videoGames

	return im, vgs
}

func (db *InMemory) GetVideoGame(id uint) *VideoGame {
	for _, vg := range *db.videoGames {
		if vg.ID == id {
			return &vg
		}
	}

	return nil
}

func (db *InMemory) GetAllVideoGames() []VideoGame {
	return *db.videoGames
}

func (db *InMemory) AddVideoGame(vg VideoGame) VideoGame {
	db.idCounter++
	vg.ID = db.idCounter

	vgs := append(*db.videoGames, vg)
	db.videoGames = &vgs
	return vg
}
