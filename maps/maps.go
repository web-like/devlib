package maps

import (
	"strconv"
	"github.com/web-like/devlib/structs"
		)

type maps struct {
}

// 获取两个相同结构的map的值的异同  以第一个参数为基准
func Diff(old interface{}, new interface{}, ignore... string) (map[string]string, []string) {

	ignore = append(ignore, "Id")
	data := make(map[string]string)
	var keys []string

	fields := structs.Names(old)
	em := structs.Map(old)
	nm := structs.Map(new)
	for _, v := range fields{
		index := SliceExist(ignore, v)

		if index != len(ignore) || em[v] == nm[v] {
			continue
		}

		switch em[v].(type) {
		case int:
			data[v] = strconv.Itoa(em[v].(int)) + "###" + strconv.Itoa(nm[v].(int))
			keys = append(keys, v)
			break
		case int64:
			data[v] = strconv.Itoa(int(em[v].(int64))) + "###" + strconv.Itoa(int(nm[v].(int64)))
			keys = append(keys, v)
			break
		default:
			data[v] = em[v].(string) + "###" + nm[v].(string)
			keys = append(keys, v)
			break
		}

	}
	return data, keys
}


func SliceExist(s []string, key string) (i int) {

	num := len(s)

	if num == 0 {
		return 0
	}

	for k, v := range s{
		if key == v {
			return k
		}
	}
	return  num
}







