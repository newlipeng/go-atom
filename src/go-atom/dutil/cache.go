package dutil

import (
	"strconv"
)

func GetMultiKey(key interface{}, prefix string) (ret []string, ok bool) {
	ok = true
	switch key.(type) {
	case []string:
		keys, ok := key.([]string)
		if ok {
			for _, v := range keys {
				ret = append(ret, prefix + v)
			}
		}
	case []int:
		keys, ok := key.([]int)
		if ok {
			for _, v := range keys {
				ret = append(ret, prefix + strconv.Itoa(v))
			}
		}
	case []int32:
		keys, ok := key.([]int32)
		if ok {
			for _, v := range keys {
				ret = append(ret, prefix+strconv.FormatInt(int64(v), 10))
			}
		}
	case []int64:
		keys, ok := key.([]int64)
		if ok {
			for _, v := range keys {
				ret = append(ret, prefix+strconv.FormatInt(v, 10))
			}
		}
	case []interface{}:
		keys, ok := key.([]interface{})
		if ok {
			for _, v := range keys {
				tmpV, tmpOk := Interface2String(v)
				if tmpOk {
					ret = append(ret, prefix + tmpV)
				}
			}
		}
	}
	if len(ret) == 0 {
		ok = false
	}
	return
}
