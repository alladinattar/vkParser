package rabbitmq

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"os"
)

func GetTask() error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"vkTasks", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {

			log.Printf("Received a message: %s", d.Body)
			err := scrapeUser(string(d.Body))
			if err != nil {
				log.Println(err)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
	return nil
}

func scrapeUser(id string) error {
	friends, err := getUsers(id, "friends.get")
	if err != nil {
		return err
	}
	followers, err := getUsers(id, "users.getFollowers")
	if err!=nil{
		return err
	}

	for _, friend := range friends{

	}

	for _, follower := range followers{

	}
	return nil
}

func getUsers(id string, vkMethod string) ([]int, error) {
	url := "https://api.vk.com/method/" + vkMethod + "?user_id=" + id + "&v=5.52&access_token=" + os.Getenv("TOKEN")
	friends, err := http.Get(url)
	if err != nil {
		return []int{}, err
	}
	var usersResponse struct {
		Response struct {
			Items []int `json:"items"`
		} `json:"response"`
	}
	err = json.NewDecoder(friends.Body).Decode(&usersResponse)
	if err != nil {
		return []int{}, err
	}
	return usersResponse.Response.Items, nil
}
