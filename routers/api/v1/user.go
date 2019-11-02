package v1

import (
	"Analysis-statistics/db"
	"Analysis-statistics/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserPortraitStatistics 用户画像统计
func UserPortraitStatistics(c *gin.Context) {
	userID := c.Param("userID")
	result := e.SUCCESS
	res := map[string]int{
		"COMPILE_ERROR":      0,
		"WRONG_ANSWER":       0,
		"ACCEPTE":            0,
		"PARTIALLY_ACCEPTED": 0,
		"SUBMIT_COUNT":       0,
		"ACCEPT_COUNT":       0,
	}
	res, err1 := db.GetUserSubmitResultCount(userID, res)
	_, err2 := db.GetProfileSubmitCount(userID, res)
	if err1 != nil && err2 != nil {
		result = e.ERROR
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
		"data":   res,
	},
	)
}
