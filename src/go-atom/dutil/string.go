package dutil

import (
	"github.com/golang/snappy"
	"strconv"
)

func Interface2String(input interface{}) (output string, ok bool) {
	ok = true
	switch input.(type) {
	case int:
		output = strconv.Itoa(input.(int))
	case int32:
		tmp := input.(int32)
		output = strconv.FormatInt(int64(tmp),10)
	case int64:
		tmp := input.(int64)
		output = strconv.FormatInt(tmp,10)
	case string:
		output = input.(string)
	default :
		ok = false
	}
	return
}

func SnappyDecode(input []byte) (output []byte) {
	output, err := snappy.Decode(nil, input)
	if nil != err {
		output = []byte{}
	}
	return
}

func SnappyEncode(input []byte) (output []byte) {
	return snappy.Encode(nil, input)
}