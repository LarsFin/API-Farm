package apifarm

import (
	"net/http"
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

func (s *VideoGameService) Get(id uint) Query {
	storedVideoGame := s.db.GetVideoGame(id)

	b, _ := s.json.Serialize(*storedVideoGame)

	return s.qf.BuildResult(b, uint(0))
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
		switch t := err.(type) {
		case *InvalidAttributeError:
			return s.qf.BuildMessage(t.Error(), http.StatusBadRequest)
		case *time.ParseError:
			msg := VideoGameInvalidDate(t.Value)
			return s.qf.BuildMessage(msg, http.StatusBadRequest)
		default:
			return s.qf.BuildMessage(InvalidJSON, http.StatusBadRequest)
		}
	}

	if len(vg.Name) == 0 {
		return s.qf.BuildMessage(VideoGameNameRequired, http.StatusBadRequest)
	}

	dt := time.Time{}
	if vg.DateReleased.Time == dt {
		return s.qf.BuildMessage(VideoGameDateRequired, http.StatusBadRequest)
	}

	svg := s.db.AddVideoGame(*vg)

	b, err := s.json.Serialize(svg)

	if err != nil {
		return s.qf.Error(err)
	}

	return s.qf.BuildResult(b, uint(0))
}
