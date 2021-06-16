package apifarm

type Service interface {
	GetAll() Query
}

type VideoGameService struct {
	db   DB
	json DataUtils
	qf   QueryFactory
}

func NewVideoGameService(db DB) *VideoGameService {
	return &VideoGameService{
		db,
		&JSON{},
		&queryFactory{},
	}
}

func NewVideoGameServiceWithUtils(db DB, json DataUtils, qf QueryFactory) *VideoGameService {
	return &VideoGameService{
		db,
		json,
		qf,
	}
}

func (s *VideoGameService) GetAll() Query {
	storedVideoGames := s.db.GetAllVideoGames()

	b, err := s.json.Serialize(storedVideoGames)

	if err != nil {
		return s.qf.Error(err)
	}

	return s.qf.Build(b, uint(0))
}

func (s *VideoGameService) Add(data []byte) Query {
	vg, _ := s.json.DeserializeVideoGame(data)

	svg := s.db.AddVideoGame(*vg)

	b, _ := s.json.Serialize(svg)

	return s.qf.Build(b, uint(0))
}
