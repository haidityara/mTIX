package server

import (
	"github.com/gin-gonic/gin"
	"github.com/haidityara/mtix/controller/controllerbook"
	"github.com/haidityara/mtix/controller/controllercinema"
	"github.com/haidityara/mtix/controller/controllercinemacity"
	"github.com/haidityara/mtix/controller/controllercinemahall"
	"github.com/haidityara/mtix/controller/controllercinemaseat"
	"github.com/haidityara/mtix/controller/controllermovie"
	"github.com/haidityara/mtix/controller/controllerpayment"
	"github.com/haidityara/mtix/controller/controllershow"
	"github.com/haidityara/mtix/controller/controlleruser"
	"github.com/haidityara/mtix/middleware"
	"github.com/haidityara/mtix/repository/repositorybook"
	"github.com/haidityara/mtix/repository/repositorycinema"
	"github.com/haidityara/mtix/repository/repositorycinemacity"
	"github.com/haidityara/mtix/repository/repositorycinemahall"
	"github.com/haidityara/mtix/repository/repositorycinemaseat"
	"github.com/haidityara/mtix/repository/repositorymovie"
	"github.com/haidityara/mtix/repository/repositorypayment"
	"github.com/haidityara/mtix/repository/repositoryshow"
	"github.com/haidityara/mtix/repository/repositoryuser"
	"github.com/haidityara/mtix/service/sericecinemacity"
	"github.com/haidityara/mtix/service/servicebook"
	"github.com/haidityara/mtix/service/servicecinema"
	"github.com/haidityara/mtix/service/servicecinemahall"
	"github.com/haidityara/mtix/service/servicecinemaseat"
	"github.com/haidityara/mtix/service/servicemovie"
	"github.com/haidityara/mtix/service/servicepayment"
	"github.com/haidityara/mtix/service/serviceshow"
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

	// cinema hall
	repoCinemaHall := repositorycinemahall.New(db)
	srvCinemaHall := servicecinemahall.New(repoCinemaHall)
	ctrlCinemaHall := controllercinemahall.New(srvCinemaHall)

	// cinema seat
	repoCinemaSeat := repositorycinemaseat.New(db)
	srvCinemaSeat := servicecinemaseat.New(repoCinemaSeat)
	ctrlCinemaSeat := controllercinemaseat.New(srvCinemaSeat)

	// route movie
	repoMovie := repositorymovie.New(db)
	srvMovie := servicemovie.New(repoMovie)
	ctrlMovie := controllermovie.New(srvMovie)

	// route show
	repoShow := repositoryshow.New(db)
	srvShow := serviceshow.New(repoShow)
	ctrlShow := controllershow.New(srvShow)

	// route book
	repoBook := repositorybook.New(db)
	srvBook := servicebook.New(repoBook)
	ctrlBook := controllerbook.New(srvBook)

	// route payment
	repoPayment := repositorypayment.New(db)
	srvPayment := servicepayment.New(repoPayment)
	ctrlPayment := controllerpayment.New(srvPayment)

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

	// route cinema hall
	r.POST("/cinema-hall", middleware.Authorization, ctrlCinemaHall.Create)
	r.GET("/cinema-hall/:id", middleware.Authorization, ctrlCinemaHall.GetByID)

	// route cinema seat
	r.POST("/cinema-seat", middleware.Authorization, ctrlCinemaSeat.Create)
	r.GET("/cinema-seat/:id", middleware.Authorization, ctrlCinemaSeat.GetByID)

	// route movie
	r.POST("/movies", middleware.Authorization, ctrlMovie.Create)
	r.GET("/movies/:id", middleware.Authorization, ctrlMovie.GetByID)
	r.GET("/movies", middleware.Authorization, ctrlMovie.GetAll)

	// route show
	r.POST("/shows", middleware.Authorization, ctrlShow.Create)
	r.GET("/shows/:id", middleware.Authorization, ctrlShow.GetByID)
	r.GET("/shows/date/:date", middleware.Authorization, ctrlShow.GetActiveShow)

	//route bok
	r.POST("/books", middleware.Authorization, ctrlBook.Create)
	r.GET("/books/:id", middleware.Authorization, ctrlBook.GetByID)

	// route payment
	r.POST("/payments", middleware.Authorization, ctrlPayment.Create)
	r.GET("/payments/:id", middleware.Authorization, ctrlPayment.GetByID)
}
