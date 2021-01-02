package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
)

type Request struct {
	Code string `form:"code"`
}

func GetBillerStatusController(c *gin.Context) {
	var req Request

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})

		return
	}

	connPool := c.MustGet("redis-pool").(*redis.Pool)

	// get conn and put back when exit from method
	conn := connPool.Get()
	defer conn.Close()

	rsp, err := redis.String(conn.Do("GET", "health-matrix"))
	if err != nil {
		log.Println("ERROR: fail to get for key |health-matrix|, error ", err.Error())
	}

	var healthMatrix map[string]interface{}

	err = json.Unmarshal([]byte(rsp), &healthMatrix)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})

		return
	}

	status, ok := healthMatrix[req.Code]
	if !ok {
		status = "UP"
	}

	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}
