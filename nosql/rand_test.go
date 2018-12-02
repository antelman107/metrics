package nosql

import (
	"fmt"
	"testing"
)

func TestGetRandString(t *testing.T) {
	for i := 0; i < 10; i++ {
		str, err := getRandString()
		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("rand: %+v\n", str)
	}
}
