package nosql

type RangePublisherInterface interface {
	Publish(score int64, data []byte) error
}
