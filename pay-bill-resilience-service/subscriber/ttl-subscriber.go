package subscriber

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
)

func TTLListener() {
	conn, err := redis.Dial("tcp", "localhost:6389")
	if err != nil {
		log.Printf("ERROR: fail init redis pool: %s", err.Error())
		os.Exit(1)
	}

	psc := redis.PubSubConn{Conn: conn}

	// Set up subscriptions
	psc.Subscribe("__keyevent@0__:expired")

	// While not a permanent error on the connection.
	for conn.Err() == nil {
		switch v := psc.Receive().(type) {
		case redis.Message:
			log.Println(v.Channel, ": message: ", string(v.Data), " ")
			if string(v.Data) == "health-matrix" {
				go func() {
					log.Println("No health matrix call came... Sending email... Triggering alarm...")
					// send email to the corresponding API
					// This email is fire and forget
				}()
			}
		case redis.Subscription:
			log.Println(v.Channel, ": ", v.Kind, " ", v.Count)
		case error:
			log.Println("error ")
		}
	}
}
