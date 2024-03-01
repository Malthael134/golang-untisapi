package v1_test

import (
	"fmt"
	"testing"

	v1 "github.com/malthael134/golang-untisapi/v1"
)

func TestMain(t *testing.T) {
	u := v1.Untis{
		Url: "Hi",
	}

	fmt.Println(u)

}
