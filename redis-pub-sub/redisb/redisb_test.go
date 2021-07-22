package redisb

import (
	"testing"
)

func Test( t *testing.T){
	assert := "ABC" 
	ret := ToUpper("abc")

	if assert != ret {
		t.Error("aqui")
	}
}
