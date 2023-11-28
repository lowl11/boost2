package configurator

func (configurator *Configurator) loadBasic() *Configurator {
	configurator.config.Net.SASL.Enable = true
	return configurator
}
