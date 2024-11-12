package controller

import (
	"be-awards/lib"
	"be-awards/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllNominasi(c *gin.Context) {
	data, err := service.GetAllNominasi()
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		lib.ReturnToJson(c, http.StatusInternalServerError, "500", "Error : "+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, http.StatusOK, "200", "Berhasil menampilkan data", data)
}
