package proceed_test

import (
	"fmt"
	"testing"

	"github.com/ani8570/learngo/010.file/f"
	"github.com/ani8570/learngo/010.file/proceed"
)

func testEnjson(t *testing.T) {
	m := f.Member{Name: "Lee", Age: 10, Active: true}
	c := proceed.Enjson(m)
	s := string(c)
	asd := 123
	fmt.Println(s)
	if asd != 123 {
		t.Errorf("Wrong")
	}
	//t.Errorf("123")
}
