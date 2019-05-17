package main

import (
	. "./util"
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)


var subHandler MQTT.MessageHandler = func(client MQTT.Client, message MQTT.Message) {
	topic := message.Topic()
	msg := ConvertByteToString(message.Payload(), GB18030)
	Log.Printf("Receive from TOPIC: %s\t", topic)
	Log.Printf("MSG: %s\n", msg)

	actionParam := GetActionParam(msg)
	if actionParam.Action == "" {
		return
	}
	//log.Println(actionParam)

	jsonData, err := json.Marshal(actionParam)
	if err != nil {
		Log.Println(err)
		return
	}
	Log.Printf("Send to TOPIC: %s\t MSG: %s\n", TopicMap[topic], string(jsonData))
	client.Publish(TopicMap[topic], 0, false, string(jsonData)).Wait()

}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	server := "tcp://"+HOST+":"+PORT
	hostname, _ := os.Hostname()
	username := ""
	password := ""


	opts := MQTT.NewClientOptions().AddBroker(server).
		SetClientID(hostname+strconv.Itoa(time.Now().Second())).SetCleanSession(true)
	if username != "" {
		opts.SetUsername(username)
		if password != "" {
			opts.SetPassword(password)
		}
	}


	opts.OnConnect = func(client MQTT.Client) {
		if token := client.SubscribeMultiple(map[string]byte{
			HW_SUB_TOPIC: 0,
			BS_SUB_TOPIC: 1,
			ZF_SUB_TOPIC: 2,
		}, subHandler);
			token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
	}

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}else {
		Log.Printf("Connected to %s\n", server)
	}

	<-c
}
