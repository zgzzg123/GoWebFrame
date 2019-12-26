package webSocketControl

import (
	"github.com/gorilla/websocket"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/garyburd/redigo/redis"
	"goDoc/library/pRedis"
	"goDoc/library/request"
	"encoding/json"
	"strconv"
)

type msgProfile struct {
	Channel string
	Level   int
	Title   string
	Content string
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var msg []byte
var clientPool []*websocket.Conn

func WsHandler(c *gin.Context) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	clientPool = append(clientPool, conn)

	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if msg != nil {
			for _,v := range clientPool {
				err := v.WriteMessage(t, msg)
				if err != nil {
					fmt.Println(err.Error())
				}
			}

		}

		msg = nil
	}
}

func ListenSub() {
	redisConn := pRedis.Pool.Get()
	defer redisConn.Close()

	psc := redis.PubSubConn{Conn: redisConn}
	psc.Subscribe("test")
	for {

		switch v := psc.Receive().(type) {
		case redis.Subscription:
			fmt.Printf("Subscription: %s: %s %d\n", v.Channel, v.Kind, v.Count)
		case redis.Message: //单个订阅subscribe
			msg = v.Data
			fmt.Printf("Message: %s: message: %s\n", v.Channel, v.Data)
		case redis.PMessage: //模式订阅psubscribe
			fmt.Printf("PMessage: %s %s %s\n", v.Pattern, v.Channel, v.Data)
			msg = v.Data
		case error:
			return
		}
	}

}

func PubMessage(c *gin.Context) {
	redisConn := pRedis.Pool.Get()
	defer redisConn.Close()

	msgJson := ""
	if c.Request.Method == "GET" {
		msgJson = `{"channel":"accountAssign","level":1,"title":"有新的收款","content":"又有30条新的收款信息，请处理哦~","created_at":"2018-05-18"}`
	} else {
		params := request.All(c, msgProfile{})
		fmt.Println(params)

		level, _ := strconv.Atoi(params["Level"].(string))
		message := msgProfile{
			Channel: params["Channel"].(string),
			Level:   level,
			Title:   params["Title"].(string),
			Content: params["Content"].(string),
		}

		msgJsonTmp, _ := json.Marshal(message)
		msgJson = string(msgJsonTmp)
	}

	if msgJson != "" {
		_, err := redisConn.Do("PUBLISH", "test", msgJson)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	c.JSON(200, "ok")
}
