package apifarm

import (
	"time"
)

type Service interface {
	GetAll() Query
	Add([]byte) Query
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

	return s.qf.BuildResult(b, uint(0))
}

func (s *VideoGameService) Add(data []byte) Query {
	vg, err := s.json.DeserializeVideoGame(data)

	if err != nil {
		switch err.(type) {
		case *time.ParseError:
			msg := VideoGameInvalidDate(err.(*time.ParseError).Value)
			return s.qf.BuildMessage(msg, uint(400))
		default:
			return s.qf.Error(err)
		}
	}

	if len(vg.Name) == 0 {
		return s.qf.BuildMessage(VideoGameNameRequired, uint(400))
	}

	dt := time.Time{}
	if vg.DateReleased.Time == dt {
		return s.qf.BuildMessage(VideoGameDateRequired, uint(400))
	}

	svg := s.db.AddVideoGame(*vg)

	b, err := s.json.Serialize(svg)

	if err != nil {
		return s.qf.Error(err)
	}

	return s.qf.BuildResult(b, uint(0))
}
