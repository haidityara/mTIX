package server

import (
	"github.com/gin-gonic/gin"
	"github.com/haidityara/mtix/controller/controllercinema"
	"github.com/haidityara/mtix/controller/controllercinemacity"
	"github.com/haidityara/mtix/controller/controllercinemahall"
	"github.com/haidityara/mtix/controller/controlleruser"
	"github.com/haidityara/mtix/middleware"
	"github.com/haidityara/mtix/repository/repositorycinema"
	"github.com/haidityara/mtix/repository/repositorycinemacity"
	"github.com/haidityara/mtix/repository/repositorycinemahall"
	"github.com/haidityara/mtix/repository/repositoryuser"
	"github.com/haidityara/mtix/service/sericecinemacity"
	"github.com/haidityara/mtix/service/servicecinema"
	"github.com/haidityara/mtix/service/servicecinemahall"
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

	// cinema
	repoCinema := repositorycinema.New(db)
	srvCinema := servicecinema.New(repoCinema)
	ctrlCinema := controllercinema.New(srvCinema)

	repoCinemaHall := repositorycinemahall.New(db)
	srvCinemaHall := servicecinemahall.New(repoCinemaHall)
	ctrlCinemaHall := controllercinemahall.New(srvCinemaHall)

	routeUser := r.Group("/users")

	// route user
	routeUser.POST("/register", ctrlUser.Create)
	routeUser.POST("/login", ctrlUser.Login)
	routeUser.PUT("/update-account", middleware.Authorization, ctrlUser.Update)
	routeUser.DELETE("/delete-account", middleware.Authorization, ctrlUser.DeleteByID)

	// route cinemaCity
	r.POST("/cinema-city", middleware.Authorization, ctrlCinemaCity.Create)
	r.GET("/cinema-city/:id", middleware.Authorization, ctrlCinemaCity.GetByID)

	// route cinema
	r.POST("/cinema", middleware.Authorization, ctrlCinema.Create)
	r.GET("/cinema/:id", middleware.Authorization, ctrlCinema.GetByID)

	// router cinema hall
	r.POST("/cinema-hall", middleware.Authorization, ctrlCinemaHall.Create)
	r.GET("/cinema-hall/:id", middleware.Authorization, ctrlCinemaHall.GetByID)
}
