package sender

import (
	"github.com/Shopify/sarama"
	"github.com/hablof/logistic-package-api/internal/config"
	"github.com/hablof/logistic-package-api/internal/model"
	"google.golang.org/protobuf/proto"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
	topic    string
}

func NewKafkaProducer(cfg config.Kafka) (*KafkaProducer, error) {

	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	sp, err := sarama.NewSyncProducer(cfg.Brokers, config)
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{
		producer: sp,
		topic:    cfg.Topic,
	}, nil
}

func (kp *KafkaProducer) Send(subdomain *model.PackageEvent) error {
	proto.Marshal()
	msg := sarama.ProducerMessage{
		Topic: kp.topic,
		Value: sarama.ByteEncoder(),
	}
	kp.producer.SendMessage()
}
