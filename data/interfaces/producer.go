package interfaces

type IProducer interface {
	Publish(topic, key string, objects ...any) error
}
