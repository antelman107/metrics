package nosql

type CounterInterface interface {
	GetValue() (int64, error)
}
