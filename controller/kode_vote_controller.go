package controller

import (
	"be-awards/lib"
	model "be-awards/model"
	service "be-awards/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateVote(c *gin.Context) {
	var requestVote model.RequestVote
	if err := c.ShouldBindJSON(&requestVote); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		lib.ReturnToJson(c, http.StatusBadRequest, "400", "error : "+err.Error(), false)
		return
	}

	_, err := service.AddKoteVote(requestVote)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		lib.ReturnToJson(c, http.StatusInternalServerError, "500", "error : "+err.Error(), false)
		return
	}
	// c.JSON(http.StatusOK, true)
	lib.ReturnToJson(c, http.StatusOK, "200", "Berhasil menambahkan Kode Vote", true)
}

func GetJumlahVote(c *gin.Context) {
	var requestKodeVote model.RequestKodeVote
	if err := c.ShouldBindJSON(&requestKodeVote); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		lib.ReturnToJson(c, http.StatusBadRequest, "400", "error : "+err.Error(), false)
		return
	}

	data, err := service.GetJumlahVote(requestKodeVote)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		lib.ReturnToJson(c, http.StatusInternalServerError, "500", "error : "+err.Error(), false)
		return
	}
	// c.JSON(http.StatusOK, data)
	lib.ReturnToJson(c, http.StatusOK, "200", "Berhasil menampilkan data", data)
}
