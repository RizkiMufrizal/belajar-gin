package app

import (
	"belajar-gin/app/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func initRoutes() http.Handler {
	router := gin.Default()

	pegawaiController := controller.NewPegawayController()
	router.GET("/hello", pegawaiController.HelloWorld)

	return router
}

func initServer() *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: initRoutes(),
	}
}

func StartWebServer() {
	var groupRouter errgroup.Group

	groupRouter.Go(func() error {
		return initServer().ListenAndServe()
	})

	if err := groupRouter.Wait(); err != nil {
		panic(err.Error())
	}
	fmt.Println("Server RUn At Port 8080")
}
