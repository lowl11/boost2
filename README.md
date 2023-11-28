# Boost 2.0

# Kafka Example
```go
func main() {
    app := boost.New()

    kafkaConfig := configurator.
		New().
		SetConsumer().
		SetProducer().
		SetHosts(config.Get("KAFKA_HOSTS").Strings()).
		SetAuth(config.Get("KAFKA_USER").String(), config.Get("KAFKA_PASS").String()).
		SetMechanism(mechanisms.Plain)

	// consumer example
    err := consumer.
        New("c360_valid_attrs", kafkaConfig).
        StartConsume(func(message *sarama.ConsumerMessage) error {
            fmt.Printf("Message value: %s\n", message.Value)
            return nil
        })
    if err != nil {
        log.Fatal(err)
    }

    // producer example
    prod, err := producer.NewSync(kafkaConfig)
    if err != nil {
        log.Fatal("Create producer error: ", err)
    }
    
    prod.Publish("c360_valid_attrs", "123", "")
	
    app.Run()
}
```
