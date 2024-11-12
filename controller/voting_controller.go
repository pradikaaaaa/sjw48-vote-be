package controller

import (
	"be-awards/lib"
	"be-awards/model"
	"be-awards/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VoteProcess(c *gin.Context) {
	var requestVote model.VoteProcess
	if err := c.ShouldBindJSON(&requestVote); err != nil {
		lib.ReturnToJson(c, http.StatusBadRequest, "400", "error : "+err.Error(), false)
		return
	}

	_, err := service.VotingProcess(requestVote)
	if err != nil {
		lib.ReturnToJson(c, http.StatusInternalServerError, "500", "error : "+err.Error(), false)
		return
	}

	lib.ReturnToJson(c, http.StatusOK, "200", "Berhasil menambahkan Vote", true)
}
