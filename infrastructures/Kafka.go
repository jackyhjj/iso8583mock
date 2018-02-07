// Package infrastructures kafka engine
// @author Valentino <daud.darianus@kudo.co.id>
package infrastructures

import (
	"sync"
	
	log "github.com/sirupsen/logrus"
	config "github.com/spf13/viper"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaProducer is the structure kafka producer adapter structure
type KafkaProducer struct {
	BrokerList string
	producer   *kafka.Producer
	Partition  int32
	mu         sync.Mutex
}

// NewKafkaProducer initialize new instance
func NewKafkaProducer() *KafkaProducer {
	p := &KafkaProducer{
		BrokerList: config.GetString("kafka.broker_list"),
	}
	return p
}

// connect to broker
func (k *KafkaProducer) connect() *kafka.Producer {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  k.BrokerList,
		"request.timeout.ms": config.GetInt("kafka.request_timeout"),
		"acks":               config.GetString("kafka.producer.acks"),
	})

	if err != nil {
		log.Error(err)
		return nil
	}
	k.producer = p

	return p
}

// Publish the message topic
func (k *KafkaProducer) Publish(topic string, payload []byte, deliveryOk chan<- bool) {

	k.mu.Lock()
	defer k.mu.Unlock()
	if k.producer == nil {
		k.connect()
	}

	doneChan := make(chan bool)
	go func() {
		//defer close(doneChan)
		for e := range k.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				m := ev
				if m.TopicPartition.Error != nil {
					log.Errorf("Delivery failed: %v\n", m.TopicPartition.Error)
					doneChan <- false
					deliveryOk <- false
					return
				}

				log.Infof("Delivered message to topic: %s  partition [%d] at offset %v\n",
					*m.TopicPartition.Topic,
					m.TopicPartition.Partition,
					m.TopicPartition.Offset,
				)

				doneChan <- true
				deliveryOk <- true
				return
			default:
				log.Warnf("Ignored event: %s\n", ev)
			}
		}
	}()

	k.producer.ProduceChannel() <- &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: payload,
	}
	// wait for delivery report goroutine to finish
	<-doneChan
	//k.producer.Close()
}
