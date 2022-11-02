package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

const exchange = "go_ex"

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(any(err))
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(any(err))
	}

	//创建exchange
	err = ch.ExchangeDeclare(
		exchange,
		"fanout",
		true,  // durable
		false, // autoDelete
		false, // internal
		false, // noWait
		nil,   // args
	)
	if err != nil {
		panic(any(err))
	}

	go subscribe(conn, exchange)
	go subscribe(conn, exchange)

	i := 0
	for {
		i++
		err := ch.Publish(
			exchange,
			"",    // key
			false, // mandatory
			false, // immediate
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

// 收
func subscribe(conn *amqp.Connection, ex string) {
	ch, err := conn.Channel()
	if err != nil {
		panic(any(err))
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			fmt.Println(err)
			panic(any(err))
		}
	}(ch)

	q, err := ch.QueueDeclare(
		"",    // name,为空系统会自动分配一个
		false, // durable
		true,  // autoDelete
		false, // exlusive
		false, // noWait
		nil,   // args
	)
	if err != nil {
		panic(any(err))
	}
	defer func(ch *amqp.Channel, name string, ifUnused, ifEmpty, noWait bool) {
		_, err := ch.QueueDelete(name, ifUnused, ifEmpty, noWait)
		if err != nil {
			fmt.Println(err)
			panic(any(err))
		}
	}(ch, q.Name, false, false, false)

	err = ch.QueueBind(
		q.Name,
		"", // key
		ex,
		false, // noWait
		nil,   // args
	)
	if err != nil {
		panic(any(err))
	}

	consume("c", ch, q.Name)
}

func consume(consumer string, ch *amqp.Channel, q string) {
	msgs, err := ch.Consume(
		q,
		consumer, // consumer
		true,     // autoAck
		false,    // exclusive
		false,    // noLocal
		false,    // noWait
		nil,      // args
	)
	if err != nil {
		panic(any(err))
	}

	for msg := range msgs {
		fmt.Printf("%s: %s\n", consumer, msg.Body)
	}
}
