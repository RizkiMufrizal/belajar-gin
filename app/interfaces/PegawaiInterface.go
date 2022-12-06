package interfaces

import "github.com/gin-gonic/gin"

type PegawaiController interface {
	HelloWorld(c *gin.Context)
}
