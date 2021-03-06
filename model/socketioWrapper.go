package model

import (
	gosocketio "github.com/graarh/golang-socketio"
	"log"
	"time"
)

func SendJoin(c *gosocketio.Client, qq string) bool {
	log.Println("获取QQ号连接")
	result, err := c.Ack("GetWebConn", qq, time.Second*5)
	if err != nil {
		log.Println("GetWebConn返回错误:", err)
		return false
	} else if result != "\"OK\"" {
		log.Println("result返回不是OK,尝试重新链接中:", result)
		return false
	} else {
		log.Println("emit", result)
		return true
	}
}
