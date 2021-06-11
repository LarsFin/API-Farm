package apifarm

type Controller struct {
}

func (c *Controller) HandlePing(res Response) {
	res.OkText("pong")
}
