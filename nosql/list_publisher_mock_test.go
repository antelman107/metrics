package nosql

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewListPublisherMock(t *testing.T) {

	testsData := []struct {
		name string
		data []struct {
			key   string
			value []byte
		}
		storage map[string][][]byte
	}{
		{
			"Simple",
			[]struct {
				key   string
				value []byte
			}{
				{
					key:   "test",
					value: []byte("123"),
				},
			},
			map[string][][]byte{
				"test": {
					[]byte("123"),
				},
			},
		},
	}

	for _, tt := range testsData {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < len(tt.data); i++ {
				obj := NewListPublisherMock(tt.data[i].key)
				err := obj.Publish(tt.data[i].value)
				if err != nil {
					t.Error(err)
				}
				assert.Equal(t, tt.storage, obj.GetStorage())
			}
		})
	}

}
