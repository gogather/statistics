package statistics

import (
	"fmt"
	"testing"
)

func Test_Lang(t *testing.T) {
	json, err := GetRepos("duguying", "token")

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(json)
	}

}
