package infrastructure

import (
	"encoding/json"
	"log"
	"marketplace/stocks-api/src/products/models"

	config "marketplace/stocks-api/src/shared"

	zmq "gopkg.in/zeromq/goczmq.v4"
)

type MQconnector struct {
	service models.IProductService
	socket *zmq.Sock
}

func InitializeBroker(productService models.IProductService){
	socket, err := zmq.NewPull(config.GetConfig().MqServerUrl)
	if err != nil {
        log.Fatal(err.Error())
    }
	log.Println("MQ router has been created")
	mq := &MQconnector{service: productService,socket: socket}
	go mq.reciveMessage()
}

func (mq *MQconnector) reciveMessage(){
	for {
		log.Print("Listening new message...")
		request, err := mq.socket.RecvMessage()
		if err != nil {
            log.Fatal(err)
        }
		products := []models.ProductOnCart{}
		err = json.Unmarshal(request[0], &products)
		if err != nil {
			log.Printf("Cannot unmarshall items %v",err.Error());
			continue
		}
		if len(products) == 0 {
			continue
		}
		mq.service.ReciveCheckout(&products)
	}
}
