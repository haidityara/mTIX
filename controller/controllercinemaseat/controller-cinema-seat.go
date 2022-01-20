package controllercinemaseat

import (
	"github.com/gin-gonic/gin"
	"github.com/haidityara/mtix/helper"
	"github.com/haidityara/mtix/model/modelcinemaseat"
	"github.com/haidityara/mtix/service/servicecinemaseat"
	"net/http"
)

type ControllerCinemaSeat interface {
	Create(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetAvailableSeats(ctx *gin.Context)
}

type controller struct {
	srv servicecinemaseat.ServiceCinemaSeat
}

func (c *controller) GetAvailableSeats(ctx *gin.Context) {
	request := modelcinemaseat.RequestAvailableSeat{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}
	seats, err := c.srv.GetAvailableSeats(request.ShowID, request.CinemaHallID)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, seats, nil))
}

func (c *controller) Create(ctx *gin.Context) {
	request := modelcinemaseat.Request{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}
	create, err := c.srv.Create(request)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.NewResponse(http.StatusCreated, create, nil))
	return
}

func (c *controller) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	get, err := c.srv.GetByID(id)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, get, nil))
	return
}

func New(srv servicecinemaseat.ServiceCinemaSeat) ControllerCinemaSeat {
	return &controller{srv: srv}
}
