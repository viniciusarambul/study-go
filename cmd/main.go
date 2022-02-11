package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/viniciusarambul/study-go/infra/kafka"
	repository2 "github.com/viniciusarambul/study-go/infra/repository"
	usecase2 "github.com/viniciusarambul/study-go/usecase"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(docker.for.mac.localhost:3306)/fullcycle")
	if err != nil {
		log.Fatalln(err)
	}

	repository := repository2.CourseMySQLRepository{Db: db}
	usecase := usecase2.CreateCourse{Repository: repository}

	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "appgo",
	}
	topics := []string{"courses"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)

	go consumer.Consume(msgChan)

	for msg := range msgChan {
		var input usecase2.CreateCourseInputDto
		json.Unmarshal(msg.Value, &input)
		output, err := usecase.Execute(input)
		if err != nil {
			fmt.Println("error: ", err)
		} else {
			fmt.Println(output)
		}
	}
	//{"name":"Curso Full Cycle","description":"Full cycle 3.0","status":"Pending"}
}
