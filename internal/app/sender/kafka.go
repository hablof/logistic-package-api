package sender

import (
	"time"

	"github.com/Shopify/sarama"
	"github.com/hablof/logistic-package-api/internal/config"
	"github.com/hablof/logistic-package-api/internal/model"
	kpb "github.com/hablof/logistic-package-api/pkg/kafka-proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type KafkaProducer struct {
	producer    sarama.SyncProducer
	topic       string
	maxAttempts int
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
		producer:    sp,
		topic:       cfg.Topic,
		maxAttempts: cfg.MaxAttempts,
	}, nil
}

func (kp *KafkaProducer) Send(event *model.PackageEvent) error {
	pbEvent := &kpb.PackageEvent{
		ID:        event.ID,
		PackageID: event.PackageID,
		Type:      0,
		Created:   timestamppb.New(event.Created),
		Payload:   event.Payload,
	}

	switch event.Type {
	case model.Created:
		pbEvent.Type = kpb.EventType_Created

	case model.Updated:
		pbEvent.Type = kpb.EventType_Updated

	case model.Removed:
		pbEvent.Type = kpb.EventType_Removed
	}

	bytes, err := proto.Marshal(pbEvent)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: kp.topic,
		Value: sarama.ByteEncoder(bytes),
	}

	var (
		partition int32
		offset    int64
		// err       error
	)
	for i := 0; i < kp.maxAttempts; i++ {
		partition, offset, err = kp.producer.SendMessage(msg)
		if err == nil {
			break
		}

		log.Debug().Err(err).Msgf("KafkaProducer.Send: failed attempt #%d to send message", i+1)
		time.Sleep(200 * time.Millisecond)
	}

	if err != nil {
		log.Debug().Err(err).Msg("KafkaProducer.Send: failed to send message")
		return err
	}

	log.Debug().Msgf("KafkaProducer.Send: message sent, partition %d, offset %d", partition, offset)

	return nil
}
