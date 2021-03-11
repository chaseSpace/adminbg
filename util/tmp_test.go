package util

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestX(t *testing.T) {
	var s = struct {
		Name string
		S    *struct {
			Age int
		}
	}{
		Name: "xxx",
	}

	b, _ := json.Marshal(&s)
	fmt.Println(string(b))
}
