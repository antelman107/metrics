package nosql

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCounterMock(t *testing.T) {
	testsData := []struct {
		name  string
		key   string
		count int64
	}{
		{
			"Simple",
			"test",
			120,
		},
	}

	for _, tt := range testsData {
		t.Run(tt.name, func(t *testing.T) {
			obj := NewCounterMock(tt.key)
			for i := int64(1); i <= tt.count; i++ {

				val, err := obj.GetValue()
				if err != nil {
					t.Error(err)
				}
				assert.Equal(t, i, val)
			}
		})
	}
}
