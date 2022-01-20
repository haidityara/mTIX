package server

import (
	"github.com/gin-gonic/gin"
	"github.com/haidityara/mtix/controller/controllercinemacity"
	"github.com/haidityara/mtix/controller/controlleruser"
	"github.com/haidityara/mtix/middleware"
	"github.com/haidityara/mtix/repository/repositorycinemacity"
	"github.com/haidityara/mtix/repository/repositoryuser"
	"github.com/haidityara/mtix/service/sericecinemacity"
	"github.com/haidityara/mtix/service/serviceuser"
	"gorm.io/gorm"
)

func NewRouter(r *gin.Engine, db *gorm.DB) {
	repoUser := repositoryuser.New(db)
	srvUser := serviceuser.New(repoUser)
	ctrlUser := controlleruser.New(srvUser)

	// cinemaCity
	repoCinemaCity := repositorycinemacity.New(db)
	srvCinemaCity := sericecinemacity.New(repoCinemaCity)
	ctrlCinemaCity := controllercinemacity.New(srvCinemaCity)

	routeUser := r.Group("/users")

	// route user
	routeUser.POST("/register", ctrlUser.Create)
	routeUser.POST("/login", ctrlUser.Login)
	routeUser.PUT("/update-account", middleware.Authorization, ctrlUser.Update)
	routeUser.DELETE("/delete-account", middleware.Authorization, ctrlUser.DeleteByID)

	// route cinemaCity
	r.POST("/cinema-city", middleware.Authorization, ctrlCinemaCity.Create)
	r.GET("/cinema-city/:id", middleware.Authorization, ctrlCinemaCity.GetByID)
}
