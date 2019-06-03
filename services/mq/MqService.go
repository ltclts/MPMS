package mq

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type Response map[string]interface{}

type queueConfigInfo struct {
	queue    string
	exchange string
	url      string
}

type queueInsInfo struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

const PackagePush = "package_push"

var queueConfigs = map[string]queueConfigInfo{
	PackagePush: {"*", "*", "*"},
}

var queuesIns = map[string]queueInsInfo{}

/**
连接mq
*/
func connect(queue string) Response {
	if _, ok := queueConfigs[queue]; !ok {
		return Response{"code": 1, "msg": queue + " config not found"}
	}
	queueConfig := queueConfigs[queue]
	var err error
	var conn *amqp.Connection
	var channel *amqp.Channel
	conn, err = amqp.Dial(queueConfig.url)
	if err != nil {
		return Response{"code": 2, "msg": "failed to connect tp mq"}
	}

	channel, err = conn.Channel()
	if err != nil {
		return Response{"code": 3, "msg": "failed to open a channel"}
	}
	queuesIns[queue] = queueInsInfo{conn, channel}
	return Response{"code": 0, "msg": "ok"}
}

/**
发送消息
*/
func Send(queue string, info map[string]interface{}) Response {
	msgJson, err := json.Marshal(info)
	if err != nil {
		return Response{"code": 1, "msg": "error to marshal json"}
	}

	if _, ok := queueConfigs[queue]; !ok { //检查是否有队列的配置信息
		return Response{"code": 2, "msg": "没有找到" + queue + "的配置"}
	}

	if _, ok := queuesIns[queue]; !ok { //检查是否有队列的实例化信息
		connect(queue)
	}

	queueConfig := queueConfigs[queue]
	queueIns := queuesIns[queue]
	queueIns.channel.Publish(queueConfig.exchange, queueConfig.queue, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        msgJson})

	return Response{"code": 0, "msg": "ok"}
}

/**
消费
*/
func Consume(queue string, deal func(msgs <-chan amqp.Delivery)) {

	if _, ok := queueConfigs[queue]; !ok { //检查是否有队列的配置信息
		panic(fmt.Sprintf("%s", "没有找到"+queue+"的配置"))
	}

	if _, ok := queuesIns[queue]; !ok { //检查是否有队列的实例化信息
		connect(queue)
	}

	queueConfig := queueConfigs[queue]
	queueIns := queuesIns[queue]
	msgs, err := queueIns.channel.Consume(queueConfig.queue, "", false, false, false, false, nil)
	failOnErr(err, "failed to open a channel")
	forever := make(chan bool)
	go deal(msgs)
	fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C\n")
	<-forever
}

/**
如果错误抛异常
*/
func failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}
