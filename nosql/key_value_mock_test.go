package nosql

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestKeyValueMock(t *testing.T) {

	testsData := []struct {
		name             string
		key              string
		setValue         []byte
		setError         error
		ttl              time.Duration
		sleep            time.Duration
		afterSleepExists bool
		existsError      error
	}{
		{
			"Simple",
			"test",
			[]byte("hz"),
			nil,
			5 * time.Millisecond,
			7 * time.Millisecond,
			true,
			nil,
		},
	}

	for _, tt := range testsData {
		t.Run(tt.name, func(t *testing.T) {
			obj := NewKeyValueMock()

			isExists, err := obj.IsExist(tt.key)
			assert.Equal(t, tt.existsError, err)
			assert.Equal(t, false, isExists)

			err = obj.Set(tt.key, tt.ttl, tt.setValue)
			assert.Equal(t, tt.setError, err)

			isExists, err = obj.IsExist(tt.key)
			assert.Equal(t, tt.existsError, err)
			assert.Equal(t, tt.afterSleepExists, isExists)
		})
	}

}
