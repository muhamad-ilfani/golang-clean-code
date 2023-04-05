package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func callBackRegistration(ctx context.Context) handler {
	return func(consumer *kafka.Consumer, msg *kafka.Message) (err error) {
		go func() {
			type data struct {
				Username string
				Email    string
			}

			var payload data

			err = json.Unmarshal(msg.Value, &payload)
			if err != nil {
				fmt.Println(err)

			}

			fmt.Println("RESPONSE KAFKA = ", payload)
		}()

		return nil
	}
}
