package rabbitmq

import (
	"fmt"
	"time"
	"testing"

)

func TestDirectProducer(t *testing.T) {
	//log.InitAppLog()

	config := &RabbitMqConfig{
		Host: "120.79.70.137",
		Port: 5672,
		Username: "sunbin",
		Password: "sunbin",
		Vhost: "/",
	}

	m, err := New(GetMqUrl(config)).Open()
	if err != nil {
		panic(err.Error())
	}

	p, err := m.Producer("test-producer")
	if err != nil {
		panic(fmt.Sprintf("Create producer failed, %v", err))
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
	if err = p.SetExchangeBinds(exb).Confirm(true).Open(); err != nil {
		panic(fmt.Sprintf("producer Open failed, %v", err))
	}
	for i := 0; i < 1000; i++ {
		//if i > 0 && i%3 == 0 {
		//	p.CloseChan() // 取样关闭连接，测试是否会自动recover
		//}
		err = p.Publish("exchange.unitest", "route.unitest1", NewPublishMsg([]byte(`{"name":"fwz"}`)))
		err = p.Publish("exchange.unitest", "route.unitest2", NewPublishMsg([]byte(`{"name":"zwf"}`)))
		fmt.Printf("Produce state:%d, err:%v\n", p.State(), err)
		time.Sleep(time.Second)
	}

	p.Close()
	m.Close()

}

func TestFanoutProducer(t *testing.T) {
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

	p, err := m.Producer("test-producer")
	if err != nil {
		panic(fmt.Sprintf("Create producer failed, %v", err))
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
				&Binding{
					RouteKey: "route.fanout.unitest2",
					Queues: []*Queue{
						DefaultQueue("queue.fanout.unitest2"),
					},
				},
			},
		},
	}
	if err = p.SetExchangeBinds(exb).Confirm(true).Open(); err != nil {
		panic(fmt.Sprintf("producer Open failed, %v", err))
	}
	for i := 0; i < 1000; i++ {
		if i > 0 && i%3 == 0 {
			p.CloseChan() // 取样关闭连接，测试是否会自动recover
		}
		err = p.Publish("exchange.fanout.unitest", "", NewPublishMsg([]byte(fmt.Sprintf(`{"name":"fwztttt---%d"}`, i)))) // 在pub/sub模式下，这里的route key要为空
		fmt.Printf("Produce state:%d, err:%v\n", p.State(), err)
		time.Sleep(time.Second)
	}

	p.Close()
	m.Close()

}
