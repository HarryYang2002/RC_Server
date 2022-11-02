package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	//account为账号，password为密码，用于登录rabbitMQ的web端,port为端口
	if err != nil {
		panic(any(err))
	}

	//对rabbitMQ的操作都是以channel来进行的
	//上面的conn是物理上的connect，下面的channel是在物理层下的虚拟channel
	ch, err := conn.Channel()
	if err != nil {
		panic(any(err))
	}

	//建立队列，若已经建立，则不新建（upset）
	q, err := ch.QueueDeclare(
		"go_q1",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(any(err))
	}

	go consume("c1", conn, q.Name)
	go consume("c2", conn, q.Name)

	i := 0
	for {
		i++
		err := ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				Body: []byte(fmt.Sprintf("message %d", i)),
			},
		)
		if err != nil {
			fmt.Println(err.Error())
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func consume(consumer string, conn *amqp.Connection, q string) {
	ch, err := conn.Channel()
	if err != nil {
		panic(any(err))
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ch)

	msgs, err := ch.Consume(
		q,
		consumer,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(any(err))
	}

	for msg := range msgs {
		fmt.Printf("%s: %s\n", consumer, msg.Body)
	}
}
