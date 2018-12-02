package nosql

type ListSubscriberInterface interface {
	Get() ([]byte, error)
}
