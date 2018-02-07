// Package infrastructures kafka consumer
// @author Tri Wicaksono <tri.wicaksono@kudo.co.id>
package infrastructures

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
	config "github.com/spf13/viper"
)

// KafkaConsumer is the kafka consumer service
type KafkaConsumer struct {
	Consumer *kafka.Consumer
	groupID  string
}

// NewKafkaConsumer instance initialize of consumer
func NewKafkaConsumer() *KafkaConsumer {
	return &KafkaConsumer{}
}

// GetChannelSize consume message
func (kc *KafkaConsumer) GetChannelSize() int {
	size := config.GetInt("kafka.consumer.channel_size")
	if size < 1 {
		return 100
	}
	return size
}

// GetSessionTimeout timeout sessions
func (kc *KafkaConsumer) GetSessionTimeout() int {
	timeout := config.GetInt("kafka.consumer.session_timeout")

	if timeout < 1 {
		return 3000
	}
	return timeout
}

// SetGroupID the consumer
func (kc *KafkaConsumer) SetGroupID(groupID string) *KafkaConsumer {
	kc.groupID = groupID

	return kc
}

// OpenBroker connection with channel
func (kc *KafkaConsumer) OpenBrokerChannel(groupID string) *kafka.Consumer {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               config.GetString("kafka.broker_list"),
		"group.id":                        fmt.Sprintf("cgr.%s", groupID),
		"session.timeout.ms":              kc.GetSessionTimeout(),
		"go.events.channel.enable":        config.GetBool("kafka.consumer.event_channel_enable"),
		"go.application.rebalance.enable": config.GetBool("kafka.consumer.rebalance_enable"),
		"go.events.channel.size":          kc.GetChannelSize(),
		"enable.auto.commit":              config.GetBool("kafka.consumer.auto_commit"),
		"default.topic.config":            kafka.ConfigMap{"auto.offset.reset": "earliest"},
	},
	)

	if err != nil {
		log.Errorf("Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	return c
}

// OpenBroker connection without channel
func (kc *KafkaConsumer) OpenBroker(groupID string) *kafka.Consumer {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    config.GetString("kafka.broker_list"),
		"group.id":             fmt.Sprintf("cgr.%s", groupID),
		"session.timeout.ms":   kc.GetSessionTimeout(),
		"enable.auto.commit":   config.GetBool("kafka.consumer.auto_commit"),
		"default.topic.config": kafka.ConfigMap{"auto.offset.reset": "earliest"},
	},
	)

	if err != nil {
		log.Errorf("Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	return c
}

// Subscribe topic without channel
func (kc *KafkaConsumer) Subscribe(topic string, signalChannel chan os.Signal, fn func(m *kafka.Message, c *kafka.Consumer)) {

	c := kc.OpenBroker(topic)

	err := c.Subscribe(topic, nil)

	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	run := true

	for run == true {
		select {
		case sig := <-signalChannel:
			log.Warnf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				log.WithField("data", string(e.Value)).Debugf("Consume : %v ", e.TopicPartition)
				fn(e, c)
			case kafka.PartitionEOF:
				log.Infof("%% Reached %v\n", e)
			case kafka.Error:
				log.Errorf("%% Error: %v\n", e)
				run = false
			default:
				log.Infof("Ignored %v\n", e)
			}
		}
	}

	log.Warnf("Closing consumer\n")
	c.Close()
}

// SubscribeChannel topic with channel
func (kc *KafkaConsumer) SubscribeChannel(topic string, signalChannel chan os.Signal, fn func(m *kafka.Message, c *kafka.Consumer)) {

	c := kc.OpenBrokerChannel(topic)

	err := c.Subscribe(topic, nil)

	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	run := true

	for run == true {
		select {
		case sig := <-signalChannel:
			log.Warnf("Caught signal %v: terminating consumer topup", sig)
			run = false

		case ev := <-c.Events():
			switch e := ev.(type) {
			case kafka.AssignedPartitions:
				log.Warnf("Assign partition : %v", e)
				c.Assign(e.Partitions)

			case kafka.RevokedPartitions:
				log.Warnf("%v", e)
				c.Unassign()

			case *kafka.Message:
				log.WithField("data", string(e.Value)).Debugf("Consume : %v ", e.TopicPartition)
				fn(e, c)
				_, err := c.CommitMessage(e)

				if err != nil {
					log.Errorf("error commit offset @%v  : %v", e.TopicPartition.Offset, err)
				}

			case kafka.PartitionEOF:
				log.Infof("Reached %v", e)

			case kafka.Error:
				log.Errorf("Error: %v", e)
				run = false
			}
		}
	}

	log.Warnf("Closing consumer\n")
	c.Close()

}
