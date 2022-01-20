package controllercinema

import (
	"github.com/gin-gonic/gin"
	"github.com/haidityara/mtix/helper"
	"github.com/haidityara/mtix/model/modelcinema"
	"github.com/haidityara/mtix/service/servicecinema"
	"net/http"
)

type ControllerCinema interface {
	Create(ctx *gin.Context)
	GetByID(ctx *gin.Context)
}

type controller struct {
	srv servicecinema.ServiceCinema
}

func (c *controller) Create(ctx *gin.Context) {
	request := modelcinema.Request{}
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

func New(srv servicecinema.ServiceCinema) ControllerCinema {
	return &controller{srv: srv}
}
