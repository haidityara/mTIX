package controllershow

import (
	"github.com/gin-gonic/gin"
	"github.com/haidityara/mtix/helper"
	"github.com/haidityara/mtix/model/modelshow"
	"github.com/haidityara/mtix/service/serviceshow"
	"net/http"
)

type ControllerShow interface {
	Create(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetActiveShow(ctx *gin.Context)
}

type controller struct {
	srv serviceshow.ServiceShow
}

func (c *controller) Create(ctx *gin.Context) {
	request := modelshow.Request{}
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
	show, err := c.srv.GetByID(id)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, show, nil))
	return
}

func (c *controller) GetActiveShow(ctx *gin.Context) {
	date := ctx.Param("date")
	show, err := c.srv.GetActiveShow(date)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, show, nil))
	return
}

func New(srv serviceshow.ServiceShow) ControllerShow {
	return &controller{srv: srv}
}
