package nosql

type ListPublisherInterface interface {
	Publish(data []byte) error
}
