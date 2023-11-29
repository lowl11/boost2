package configurator

import (
	"github.com/IBM/sarama"
	"github.com/lowl11/boost2/data/enums/mechanisms"
)

func (configurator *Configurator) Config() *sarama.Config {
	if configurator.config.Net.SASL.Mechanism == "" {
		configurator.config.Net.SASL.Mechanism = mechanisms.Plain
	}

	return configurator.config
}

func (configurator *Configurator) SetConsumer() *Configurator {
	configurator.config.Consumer.Return.Errors = true
	return configurator
}

func (configurator *Configurator) SetProducer() *Configurator {
	configurator.config.Producer.Return.Successes = true
	return configurator
}

func (configurator *Configurator) SetHosts(hosts []string) *Configurator {
	configurator.hosts = hosts
	return configurator
}

func (configurator *Configurator) Hosts() []string {
	return configurator.hosts
}

func (configurator *Configurator) Group() string {
	return configurator.groupName
}

func (configurator *Configurator) SetGroup(groupName string) *Configurator {
	configurator.groupName = groupName
	return configurator
}

func (configurator *Configurator) SetMechanism(mechanism string) *Configurator {
	configurator.config.Net.SASL.Mechanism = sarama.SASLMechanism(mechanism)
	return configurator
}

func (configurator *Configurator) SetAuth(username, password string) *Configurator {
	configurator.config.Net.SASL.User = username
	configurator.config.Net.SASL.Password = password
	return configurator
}
