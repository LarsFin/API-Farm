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

func NewInMemoryWithVideoGames(videoGames *[]VideoGame) *InMemory {
	return &InMemory{
		videoGames,
	}
}

func (db *InMemory) GetAllVideoGames() []VideoGame {
	return *db.videoGames
}

func (db *InMemory) AddVideoGame(vg VideoGame) VideoGame {
	return VideoGame{}
}
