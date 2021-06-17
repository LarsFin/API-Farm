package apifarm

import "net/http"

type Controller struct {
	s Service
}

func NewController(s Service) *Controller {
	return &Controller{
		s,
	}
}

func (c *Controller) HandlePing(res Response) {
	res.OkText("pong")
}

func (c *Controller) HandleGetAll(res Response) {
	query := c.s.GetAll()

	switch query.Code {
	case 0:
		res.OkJSON(query.Result)
	case http.StatusInternalServerError:
		res.Error(query.Error)
	}
}

func (c *Controller) HandlePost(req Request, res Response) {
	body, err := req.GetBody()

	if err != nil {
		res.Error(err)
		return
	}

	query := c.s.Add(body)

	switch query.Code {
	case 0:
		res.CreatedJSON(query.Result)
	case http.StatusBadRequest:
		res.BadRequestText(query.Message)
	case http.StatusInternalServerError:
		res.Error(query.Error)
	}
}
