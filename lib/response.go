package lib

import "github.com/gin-gonic/gin"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ReturnToJson(w *gin.Context, status int, codeStatus string, desc string, data interface{}) {
	var res Response

	res.Status = codeStatus
	res.Message = desc
	res.Data = data

	w.JSON(status, res)
}
