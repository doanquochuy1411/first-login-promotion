package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmptyResponse struct{}

type Response struct {
	Status  string      `json:"status" example:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty" example:"This is a message of response"`
	Code    int         `json:"code,omitempty" example:"200"`
}

func RespondWithJSON(c *gin.Context, code int, status string, data interface{}, message string) {
	resp := Response{
		Status:  status,
		Data:    data,
		Message: message,
		Code:    code,
	}
	c.JSON(code, resp)
}

func RespondWithError(c *gin.Context, code int, message string) {
	RespondWithJSON(c, code, "error", nil, message)
}

func RespondWithExcel(c *gin.Context, fileName string, data []byte) {
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", data)
}
