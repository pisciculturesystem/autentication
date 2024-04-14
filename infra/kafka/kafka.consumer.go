package kafka

type ConsumerKafka struct {
}

func (c *ConsumerKafka) Start() error {

	return nil
}

func NewConsumerKafka() *ConsumerKafka {
	return &ConsumerKafka{}
}
