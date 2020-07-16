package rabbitmq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"time"

	"redrocksummer3/model"
	"redrocksummer3/tool"
)



func Order(order, person, position string, number int) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")//连接
	tool.HandleError(err, "Can't connect to MQ")//检测错误
	defer conn.Close()//关闭
	amqpChannel, err := conn.Channel()//获取管道信息
	tool.HandleError(err, "Can't create a Channel")//检测错误
	defer amqpChannel.Close()//关闭
	//队列声明并发送数据
	queue, err := amqpChannel.QueueDeclare(
		"goodList",//队列名称
		true,//持久化
		false,//是否自动删除，没有消息就自动删除次队列
		false,//是否具有排他性（true只有自己可以访问，连接端口会自动删除）
		false,//是否阻塞（设置为true时就像没有bufio的channel）
		nil,//额外属性（我也不知道有啥额外属性）
	)
	tool.HandleError(err, "Could not declare queue")//检测错误
	rand.Seed(time.Now().UnixNano())//TODO 随机数种子，这个随机数是发给谁的？
	//电影票结构体（应用在其他项目就是其他结构体）
	good := model.Orders{
		Name:  order,
		Pname: person,
		Where: position,
		Num:   number,
	}
	//movie := model.Movie{
	//	Name:  order,
	//	Num:   number,
	//	Time:  tm,
	//	Where: position,
	//}
	body,  err := json.Marshal(good)//json序列化结构体
	//body1, err := json.Marshal(movie)

	if err != nil {
		tool.HandleError(err, "Error encoding JSON")//检测错误
	}
	//发送消息Publish模式
	err = amqpChannel.Publish(
		"",//交换机名称
		queue.Name,//队列名称
		false,//如为true，根据exchange类型与routkey规则，
		// 如果无法找到符合条件队列，那么就会把发送的消息返回给发送者
		false,//如为true，但exchange发送到消息队列后发现
		//队列上没有绑定消费者，则会把消息发送给发送者
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,//TODO 这个并不太清楚没找到这是啥
			ContentType:  "text/plain",//明文
			Body:         body,//发送的东西
		})
	if err != nil {
		log.Fatalf("Error publishing message: %s", err)//检测错误
	}
	log.Println("AddGood:", string(body))//打印日志
}
