package main

import (
	"testing"
	"github.com/hashicorp/go-immutable-radix"
	"fmt"
	"encoding/json"
)

func ToMap(t *iradix.Tree) map[string]interface{} {
	out := make(map[string]interface{}, t.Len())
	t.Root().Walk(func(k []byte, v interface{}) bool {
		out[string(k)] = v
		return false
	})
	return out
}

func Test_Radix2(t *testing.T) {
	r := iradix.New()
	r, _, _ = r.Insert([]byte("/foo"), "value-foo")
	r, _, _ = r.Insert([]byte("/bar"), "value-bar")
	r, _, _ = r.Insert([]byte("/foobar"), "value-foobar")
	b, i, ok := r.Root().LongestPrefix([]byte("/bar*"))
	if ok {
		fmt.Println(string(b))
		fmt.Println(i)
	}
	fmt.Println(r.Len())

	m := ToMap(r)
	j, err := json.Marshal(m)
	if err == nil {
		fmt.Println(string(j))
	}
}
