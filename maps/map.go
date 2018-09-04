package maps

import (
	"strconv"
	"github.com/web-like/devlib/structs"
)

// 获取两个相同结构的map的值的异同  以第一个参数为基准
func Diff(old interface{}, new interface{}) (map[string]string, []string) {

	data := make(map[string]string)
	var keys []string

	fields := structs.Names(old)
	em := structs.Map(new)
	nm := structs.Map(new)
	for _, v := range fields{
		if v == "Id" || em[v] == nm[v] {
			continue
		}

		switch em[v].(type) {
		case int:
			data[v] = strconv.Itoa(em[v].(int)) + "###" + strconv.Itoa(nm[v].(int))
			keys = append(keys, v)
			break
		case int64:
			data[v] = strconv.Itoa(int(em[v].(int))) + "###" + strconv.Itoa(int(nm[v].(int)))
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








