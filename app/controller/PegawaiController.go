package controller

import (
	"belajar-gin/app/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

type pegawaiControllerstruct struct {
}

func NewPegawayController() interfaces.PegawaiController {
	return pegawaiControllerstruct{}
}

func (p pegawaiControllerstruct) HelloWorld(c *gin.Context) {
	response := map[string]string{
		"message": "Success",
		"code":    "200",
	}

	c.JSON(http.StatusOK, response)
}
