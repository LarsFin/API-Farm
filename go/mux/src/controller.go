package apifarm

type Controller struct {
}

func (c *Controller) HandlePing(res response) {
	res.OkText("pong")
}
