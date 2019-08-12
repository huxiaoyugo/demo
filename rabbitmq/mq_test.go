package rabbitmq

import (
	"testing"
)

func Test_String(t *testing.T) {
	go WorkProducer()
	go WorkConsume()


	forerver := make(chan int)
	<-forerver
}

func Test_String2(t *testing.T) {
	 WorkConsume()
}


func TestStrComsume(t* testing.T) {
	StrComsume()
}


func TestPublishProducer(t *testing.T) {
	PublishProducer()
}


func TestPublishConsumer(t *testing.T) {
	PublishConsumer()
}

func TestPublishConsumer2(t *testing.T) {
	PublishConsumer()
}


func TestRouteProducer(t *testing.T) {
	RouteProducer("blue")
}

func TestRouteComsume(t *testing.T) {
	RouteComsume([]string{"red","blue"})
}


func TestRouteComsume2(t *testing.T) {
	RouteComsume([]string{"blue"})
}

func TestRouteComsume3(t *testing.T) {
	RouteComsume([]string{"green"})
}