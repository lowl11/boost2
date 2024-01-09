package configurator

import (
	"github.com/IBM/sarama"
	"github.com/lowl11/boost2/data/enums/mechanisms"
	"time"
)

func (configurator *Configurator) Copy() *Configurator {
	c := *configurator
	return &c
}

func (configurator *Configurator) Config() *sarama.Config {
	if configurator.config.Net.SASL.Mechanism == "" {
		configurator.config.Net.SASL.Mechanism = mechanisms.Plain
	}

	return configurator.config
}

func (configurator *Configurator) SetConsumer(offset int64) *Configurator {
	configurator.config.Consumer.Return.Errors = true
	configurator.config.Consumer.Offsets.Initial = offset
	configurator.config.Consumer.Offsets.AutoCommit.Enable = false
	return configurator
}

func (configurator *Configurator) SetAutocommit() *Configurator {
	configurator.config.Consumer.Offsets.AutoCommit.Enable = true
	return configurator
}

func (configurator *Configurator) SetBufferSize(size int) *Configurator {
	configurator.config.ChannelBufferSize = size
	return configurator
}

func (configurator *Configurator) SetMaxProcessTime(timeout time.Duration) *Configurator {
	configurator.config.Consumer.MaxProcessingTime = timeout
	return configurator
}

func (configurator *Configurator) SetMaxWaitTime(timeout time.Duration) *Configurator {
	configurator.config.Consumer.MaxWaitTime = timeout
	return configurator
}

func (configurator *Configurator) SetProducer() *Configurator {
	configurator.config.Producer.Return.Successes = true
	configurator.config.Producer.RequiredAcks = sarama.WaitForAll
	configurator.config.Producer.Retry.Max = 5
	configurator.config.Producer.Timeout = time.Second * 60
	return configurator
}

func (configurator *Configurator) SetProducerBatchSize(size int) *Configurator {
	configurator.config.Producer.Flush.Bytes = size
	return configurator
}

func (configurator *Configurator) SetFetchSize(size int) *Configurator {
	configurator.config.Consumer.Fetch.Default = int32(size)
	return configurator
}

func (configurator *Configurator) SetHosts(hosts []string) *Configurator {
	configurator.hosts = hosts
	return configurator
}

func (configurator *Configurator) Hosts() []string {
	return configurator.hosts
}

func (configurator *Configurator) SetMechanism(mechanism string) *Configurator {
	configurator.config.Net.SASL.Mechanism = sarama.SASLMechanism(mechanism)
	return configurator
}

func (configurator *Configurator) SetAuth(username, password string) *Configurator {
	configurator.config.Net.SASL.Enable = true
	configurator.config.Net.SASL.Handshake = true
	configurator.config.Net.SASL.User = username
	configurator.config.Net.SASL.Password = password
	return configurator
}

func (configurator *Configurator) SetCustom(apply func(config *sarama.Config)) {
	apply(configurator.config)
}
