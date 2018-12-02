package nosql

type RangeSubscriberInterface interface {
	Get(scoreFrom, scoreTo int64, limit, offset int) ([][]byte, error)
}
