package mock_producer

import "github.com/lowl11/boost2/log"

func (producer *Producer) Publish(topic, key string, objects ...any) error {
	log.Info("Producer message to ", topic, " by key ", key, ". Objects count: ", len(objects))
	for _, obj := range objects {
		log.Info("Produce object: ", obj)
	}

	return nil
}
