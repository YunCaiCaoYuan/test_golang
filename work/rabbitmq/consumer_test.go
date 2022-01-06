package rabbitmq

import (
	"fmt"
	"time"
	"testing"

)

func TestDirectConsumer(t *testing.T) {
	//log.InitAppLog()

	config := &RabbitMqConfig{
		Host: "192.168.33.10",
		Port: 5672,
		Username: "guest",
		Password: "guest",
		Vhost: "/test",
	}

	m, err := New(GetMqUrl(config)).Open()
	if err != nil {
		panic(err.Error())
	}

	c, err := m.Consumer("test-consumer")
	if err != nil {
		panic(fmt.Sprintf("Create consumer failed, %v", err))
	}

	exb := []*ExchangeBinds{
		&ExchangeBinds{
			Exch: DefaultExchange("exchange.unitest", ExchangeDirect),
			Bindings: []*Binding{
				&Binding{
					RouteKey: "route.unitest1",
					Queues: []*Queue{
						DefaultQueue("queue.unitest1"),
					},
				},
				&Binding{
					RouteKey: "route.unitest2",
					Queues: []*Queue{
						DefaultQueue("queue.unitest2"),
					},
				},
			},
		},
	}

	msgC := make(chan Delivery, 1)
	defer close(msgC)

	if err = c.SetExchangeBinds(exb).SetMsgCallback(msgC).SetQos(50).Open(); err != nil {
		panic(fmt.Sprintf("consumer Open failed, %v", err))
	}
	fmt.Println("start comsumer:")
	i := 0
	for msg := range msgC {
		i++
		if i%5 == 0 {
			c.CloseChan() // 取样关闭连接，测试是否会自动recover
		}
		// fmt.Println(msg)
		fmt.Printf("Tag(%d) Body: %s\n", msg.DeliveryTag, string(msg.Body))
		msg.Ack(true)
		time.Sleep(time.Second)
	}

	c.Close()
	m.Close()

}

func TestFanoutConsumer(t *testing.T) {
	//log.InitAppLog()

	config := &RabbitMqConfig{
		Host: "192.168.33.10",
		Port: 5672,
		Username: "guest",
		Password: "guest",
		Vhost: "/test",
	}

	m, err := New(GetMqUrl(config)).Open()
	if err != nil {
		panic(err.Error())
	}

	c, err := m.Consumer("test-consumer")
	if err != nil {
		panic(fmt.Sprintf("Create consumer failed, %v", err))
	}

	exb := []*ExchangeBinds{
		&ExchangeBinds{
			Exch: DefaultExchange("exchange.fanout.unitest", ExchangeFanout),
			Bindings: []*Binding{
				&Binding{
					RouteKey: "route.fanout.unitest1",
					Queues: []*Queue{
						DefaultQueue("queue.fanout.unitest1"),
					},
				},
				// &Binding{
				// 	RouteKey: "route.fanout.unitest2",
				// 	Queues: []*Queue{
				// 		DefaultQueue("queue.fanout.unitest2"),
				// 	},
				// },
			},
		},
	}

	msgC := make(chan Delivery, 1)
	defer close(msgC)

	if err = c.SetExchangeBinds(exb).SetMsgCallback(msgC).SetQos(50).Open(); err != nil {
		panic(fmt.Sprintf("consumer Open failed, %v", err))
	}
	fmt.Println("start comsumer:")
	i := 0
	for msg := range msgC {
		i++
		if i%5 == 0 {
			c.CloseChan() // 取样关闭连接，测试是否会自动recover
		}
		// fmt.Println(msg)
		fmt.Printf("Tag(%d) Body: %s\n", msg.DeliveryTag, string(msg.Body))
		msg.Ack(true)
		time.Sleep(time.Second)
	}

	c.Close()
	m.Close()

}
