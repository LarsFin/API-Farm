package apifarm

type DB interface {
	GetAllVideoGames() []VideoGame
	AddVideoGame(VideoGame) VideoGame
}

type InMemory struct {
	videoGames *[]VideoGame
}

func NewInMemory() *InMemory {
	return &InMemory{
		&([]VideoGame{}),
	}
}

func NewInMemoryForTests() (*InMemory, **[]VideoGame) {
	im := &InMemory{&([]VideoGame{})}
	vgs := &im.videoGames

	return im, vgs
}

func (db *InMemory) GetAllVideoGames() []VideoGame {
	return *db.videoGames
}

func (db *InMemory) AddVideoGame(vg VideoGame) VideoGame {
	vgs := append(*db.videoGames, vg)
	db.videoGames = &vgs
	return vg
}
