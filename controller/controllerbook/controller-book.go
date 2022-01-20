package controllerbook

import (
	"github.com/gin-gonic/gin"
	"github.com/haidityara/mtix/helper"
	"github.com/haidityara/mtix/model/modelbook"
	"github.com/haidityara/mtix/service/servicebook"
	"net/http"
)

type ControllerBook interface {
	Create(ctx *gin.Context)
	GetByID(ctx *gin.Context)
}

type controller struct {
	srv servicebook.ServiceBook
}

func (c *controller) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	book, err := c.srv.GetByID(id)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, book, nil))
	return
}

func (c *controller) Create(ctx *gin.Context) {
	request := modelbook.Request{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}
	uID := ctx.MustGet("user_id")
	request.UserID = uID.(uint)
	create, err := c.srv.Create(request)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.NewResponse(http.StatusCreated, create, nil))
	return
}

func New(srv servicebook.ServiceBook) ControllerBook {
	return &controller{srv: srv}
}
