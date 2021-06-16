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
	body, _ := req.GetBody()

	query := c.s.Add(body)

	res.CreatedJSON(query.Result)
}
