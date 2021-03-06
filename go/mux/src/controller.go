package apifarm

import (
	"net/http"
	"strconv"
)

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

func (c *Controller) HandleGet(req Request, res Response) {
	strID := req.GetParam("id")

	id, err := strconv.Atoi(strID)

	if err != nil {
		res.BadRequestText(ParamInvalidID(strID))
		return
	}

	query := c.s.Get(uint(id))

	switch query.Code {
	case 0:
		res.OkJSON(query.Result)
	case http.StatusNotFound:
		res.NotFoundText(query.Message)
	case http.StatusInternalServerError:
		res.Error(query.Error)
	}
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

func (c *Controller) HandlePut(req Request, res Response) {
	strID := req.GetParam("id")

	id, err := strconv.Atoi(strID)

	if err != nil {
		res.BadRequestText(ParamInvalidID(strID))
		return
	}

	body, err := req.GetBody()

	if err != nil {
		res.Error(err)
		return
	}

	query := c.s.Update(uint(id), body)

	switch query.Code {
	case 0:
		res.OkJSON(query.Result)
	case http.StatusBadRequest:
		res.BadRequestText(query.Message)
	case http.StatusNotFound:
		res.NotFoundText(query.Message)
	case http.StatusInternalServerError:
		res.Error(query.Error)
	}
}

func (c *Controller) HandleDelete(req Request, res Response) {
	strID := req.GetParam("id")

	id, err := strconv.Atoi(strID)

	if err != nil {
		res.BadRequestText(ParamInvalidID(strID))
		return
	}

	query := c.s.Delete(uint(id))

	switch query.Code {
	case 0:
		res.OkText(query.Message)
	case http.StatusNotFound:
		res.NotFoundText(query.Message)
	}
}

type APITestingController struct {
	dl DataLoader
}

const SampleDataPath = "./data.json"

func NewAPITestingController(dl DataLoader) *APITestingController {
	return &APITestingController{
		dl,
	}
}

func (c *APITestingController) HandleTestSetup(res Response) {
	query := c.dl.Load(SampleDataPath)

	if query.Code == http.StatusInternalServerError {
		res.Error(query.Error)
		return
	}

	res.OkText(query.Message)
}
