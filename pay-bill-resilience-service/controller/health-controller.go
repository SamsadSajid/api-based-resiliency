package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
)

type Matrix struct {
	Biller1 string `json:"biller_1"`
	Biller2 string `json:"biller_3"`
	Biller3 string `json:"biller_2"`
}

func HealthController(c *gin.Context) {
	var hlthreq Matrix

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

	// pipelining two operations
	conn.Send("SET", "health-matrix", hlthreq)
	conn.Send("EXPIRE", "health-matrix", 10)

	_, err := conn.Do("")
	if err != nil {
		log.Println("ERROR: fail to set for key |health-matrix| for val |", hlthreq, "|, error ", err.Error(),
			"with ttl 10s")
	}
}
