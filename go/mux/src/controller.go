package apifarm

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
		res.OkJson(query.Result)
		break
	case 500:
		res.Error(query.Error)
		break
	}
}
