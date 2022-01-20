package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haidityara/mtix/config/configdb"
	"os"
)

func Start() error {
	db, err := configdb.New()
	if err != nil {
		return err
	}

	r := gin.Default()
	NewRouter(r, db)

	r.Use(gin.Recovery())

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8000"
	}

	r.Run(fmt.Sprintf(":%s", port))
	return nil
}
