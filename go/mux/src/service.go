package apifarm

import (
	"net/http"
	"time"
)

type Service interface {
	Get(uint) Query
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

	if storedVideoGame == nil {
		return s.qf.BuildMessage(VideoGameNotFound(id), http.StatusNotFound)
	}

	b, err := s.json.Serialize(storedVideoGame)

	if err != nil {
		return s.qf.Error(err)
	}

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

	setEmptySlices(vg)

	svg := s.db.AddVideoGame(*vg)

	b, err := s.json.Serialize(&svg)

	if err != nil {
		return s.qf.Error(err)
	}

	return s.qf.BuildResult(b, uint(0))
}

func (s *VideoGameService) Update(id uint, data []byte) Query {
	vg := s.db.GetVideoGame(id)

	if vg == nil {
		return s.qf.BuildMessage(VideoGameNotFound(id), http.StatusNotFound)
	}

	vgu, err := s.json.DeserializeVideoGame(data)

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

	uvg := updateVideoGameFields(*vg, *vgu)

	suvg := s.db.UpdateVideoGame(uvg)

	b, _ := s.json.Serialize(suvg)

	return s.qf.BuildResult(b, uint(0))
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

func updateVideoGameFields(vg, vgu VideoGame) VideoGame {
	uvg := VideoGame{ID: vg.ID}

	if len(vgu.Name) > 0 {
		uvg.Name = vgu.Name
	} else {
		uvg.Name = vg.Name
	}

	if vgu.Developers != nil {
		uvg.Developers = vgu.Developers
	} else {
		uvg.Developers = vg.Developers
	}

	if vgu.Publishers != nil {
		uvg.Publishers = vgu.Publishers
	} else {
		uvg.Publishers = vg.Publishers
	}

	if vgu.Directors != nil {
		uvg.Directors = vgu.Directors
	} else {
		uvg.Directors = vg.Directors
	}

	if vgu.Producers != nil {
		uvg.Producers = vgu.Producers
	} else {
		uvg.Producers = vg.Producers
	}

	if vgu.Designers != nil {
		uvg.Designers = vgu.Designers
	} else {
		uvg.Designers = vg.Designers
	}

	if vgu.Programmers != nil {
		uvg.Programmers = vgu.Programmers
	} else {
		uvg.Programmers = vg.Programmers
	}

	if vgu.Artists != nil {
		uvg.Artists = vgu.Artists
	} else {
		uvg.Artists = vg.Artists
	}

	if vgu.Composers != nil {
		uvg.Composers = vgu.Composers
	} else {
		uvg.Composers = vg.Composers
	}

	if vgu.Platforms != nil {
		uvg.Platforms = vgu.Platforms
	} else {
		uvg.Platforms = vg.Platforms
	}

	dt := time.Time{}
	if vgu.DateReleased.Time != dt {
		uvg.DateReleased = vgu.DateReleased
	} else {
		uvg.DateReleased = vg.DateReleased
	}

	return uvg
}
