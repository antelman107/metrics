package nosql

type PubSubPublisherInterface interface {
	Publish(key string, data []byte) error
}
