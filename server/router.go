package server

import (
	"github.com/gin-gonic/gin"
	"github.com/haidityara/mtix/controller/controlleruser"
	"github.com/haidityara/mtix/middleware"
	"github.com/haidityara/mtix/repository/repositoryuser"
	"github.com/haidityara/mtix/service/serviceuser"
	"gorm.io/gorm"
)

func NewRouter(r *gin.Engine, db *gorm.DB) {
	repoUser := repositoryuser.New(db)
	srvUser := serviceuser.New(repoUser)
	ctrlUser := controlleruser.New(srvUser)

	routeUser := r.Group("/users")

	// route user
	routeUser.POST("/register", ctrlUser.Create)
	routeUser.POST("/login", ctrlUser.Login)
	routeUser.PUT("/update-account", middleware.Authorization, ctrlUser.Update)
	routeUser.DELETE("/delete-account", middleware.Authorization, ctrlUser.DeleteByID)
}
