package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"

	"redrocksummer3/rabbitmq"
)

func Buy(c *gin.Context)  {
	//接受数据
	order       := c.PostForm("order")
	person      := c.PostForm("person")
	num         := c.PostForm("num")

	position    := c.PostForm("position")
	num_id, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("err:", err)
	}
	rabbitmq.Order(order, person, position, num_id)//应该是传输进消息队列
}
