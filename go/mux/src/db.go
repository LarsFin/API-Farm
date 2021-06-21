package apifarm

type DB interface {
	GetVideoGame(uint) *VideoGame
	GetAllVideoGames() []VideoGame
	AddVideoGame(VideoGame) VideoGame
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

	setEmptySlices(&vg)

	vgs := append(*db.videoGames, vg)
	db.videoGames = &vgs
	return vg
}

func (db *InMemory) Reset() {
	db.idCounter = 0
	db.videoGames = &[]VideoGame{}
}

func setEmptySlices(vg *VideoGame) {
	if vg.Developers == nil {
		vg.Developers = []string{}
	}
	if vg.Publishers == nil {
		vg.Publishers = []string{}
	}
	if vg.Directors == nil {
		vg.Directors = []string{}
	}
	if vg.Producers == nil {
		vg.Producers = []string{}
	}
	if vg.Designers == nil {
		vg.Designers = []string{}
	}
	if vg.Programmers == nil {
		vg.Programmers = []string{}
	}
	if vg.Artists == nil {
		vg.Artists = []string{}
	}
	if vg.Composers == nil {
		vg.Composers = []string{}
	}
	if vg.Platforms == nil {
		vg.Platforms = []string{}
	}
}
