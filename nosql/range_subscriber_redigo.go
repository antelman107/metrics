package nosql

import (
	"github.com/gomodule/redigo/redis"
)

const (
	ZpopScript = `
if redis.call("EXISTS", KEYS[1]) == 1 then
    local keyvalues = redis.call("ZRANGEBYSCORE", KEYS[1], ARGV[1], ARGV[2], "LIMIT", ARGV[3], ARGV[4])
    for k,v in pairs(keyvalues) do
		redis.call("ZREM", KEYS[1], k, v)
	end	
    return keyvalues
else
    return {}
end
`
)

type RangeSubscriberRedigo struct {
	key       string
	redisPool *redis.Pool
}

func (sh *RangeSubscriberRedigo) Get(scoreFrom, scoreTo int64, limit, offset int) (data [][]byte, err error) {
	var conn redis.Conn
	if conn = sh.redisPool.Get(); conn.Err() != nil {
		return nil, conn.Err()
	}
	defer conn.Close()

	data, err = redis.ByteSlices(
		redis.NewScript(
			1,
			ZpopScript,
		).Do(
			conn,
			sh.key,
			scoreFrom,
			scoreTo,
			offset,
			limit,
		))

	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewRangeSubscriberRedigo(
	key string,
	redisPool *redis.Pool,
) RangeSubscriberInterface {
	return &RangeSubscriberRedigo{
		key:       key,
		redisPool: redisPool,
	}
}
