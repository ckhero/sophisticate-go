package main

import (
	"fmt"
	_ "github.com/micro/go-plugins/broker/kafka/v2"
	_ "main/SLB"
)

func init() {

	fmt.Println("main",1)
}
func main()  {
}

//func pub() {
//	tick := time.NewTicker(time.Second)
//	for _ = range tick.C {
//		msg := &broker.Message{
//			Header: map[string]string{
//				"id": fmt.Sprintf("%d", i),
//			},
//			Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
//		}
//		if err := broker.Publish(topic, msg); err != nil {
//			log.Printf("[pub] failed: %v", err)
//		} else {
//			fmt.Println("[pub] pubbed message:", string(msg.Body))
//		}
//		i++
//	}
//}
//
//func sub() {
//	_, err := broker.Subscribe(topic, func(p broker.Event) error {
//		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
//		return nil
//	})
//	if err != nil {
//		fmt.Println(err)
//	}
//}