package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

func PostMessage(topic, message string) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"kafka:29092"}, config)
	if err != nil {
		log.Fatalf("Erro ao criar o produtor: %v", err)
	}

	defer producer.Close()

	mensagemKafka := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err = producer.SendMessage(mensagemKafka)
	if err != nil {
		return err
	}
	return nil
}
