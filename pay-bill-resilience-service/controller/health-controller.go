package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
)

func HealthController(c *gin.Context) {
	var hlthreq map[string]interface{}

	if err := c.ShouldBind(&hlthreq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})

		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	}

	log.Println("Biller health matrix is ", hlthreq)
	connPool := c.MustGet("redis-pool").(*redis.Pool)

	// get conn and put back when exit from method
	conn := connPool.Get()
	defer conn.Close()

	jsonString, _ := json.Marshal(hlthreq)

	// pipelining two operations
	conn.Send("SET", "health-matrix", jsonString)
	conn.Send("EXPIRE", "health-matrix", 21)

	_, err := conn.Do("")
	if err != nil {
		log.Println("ERROR: fail to set for key |health-matrix| for val |", hlthreq, "|, error ", err.Error())
	}
}
