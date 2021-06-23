package apifarm

type DB interface {
	GetVideoGame(uint) *VideoGame
	GetAllVideoGames() []VideoGame
	AddVideoGame(VideoGame) VideoGame
	UpdateVideoGame(VideoGame) *VideoGame
	Reset()
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

func (db *InMemory) UpdateVideoGame(uvg VideoGame) *VideoGame {
	for i, vg := range *db.videoGames {
		if vg.ID == uvg.ID {
			(*db.videoGames)[i] = uvg
			return &uvg
		}
	}

	return nil
}

func (db *InMemory) Reset() {
	db.idCounter = 0
	db.videoGames = &[]VideoGame{}
}
