package maps

import (
	"log"
	"reflect"
)

// IsMap 是否是map类型
func IsMap(map1 interface{}) bool {
	return reflect.TypeOf(map1).Kind() == reflect.Map
}

// InMap key 是否存在 map 中
func InMap(key interface{}, input interface{}) bool {
	return IsMap(input) && reflect.ValueOf(input).MapIndex(reflect.ValueOf(key)).IsValid()
}

// TypesEqual 类型是否相等
func TypesEqual(map1, map2 interface{}) bool {
	return reflect.TypeOf(map1) == reflect.TypeOf(map2)
}

// ValuesEqual 值是否相等
func ValuesEqual(map1, map2 interface{}) bool {
	return reflect.DeepEqual(map1, map2)
}

// Contains map1 是否包含 map2，只匹配键
func Contains(map1, map2 interface{}) bool {
	if TypesEqual(map1, map2) {
		// 获取 map 的键值
		map1Value := reflect.ValueOf(map1)
		map2Value := reflect.ValueOf(map2)

		for _, key := range map2Value.MapKeys() {
			map2Elem := map2Value.MapIndex(key)
			map1Elem := map1Value.MapIndex(key)

			if !map1Elem.IsValid() || map1Elem.Interface() != map2Elem.Interface() {
				return false
			}
		}

		return true
	} else {
		log.Println("map1 和 map2 类型不相等！")
		return false
	}
}

// Merges 合并两个 map，返回一个新的 map，注意：相同的键值会被覆盖(map2 覆盖 map1)
func Merges(map1, map2 interface{}) interface{} {
	if TypesEqual(map1, map2) {
		map1Value := reflect.ValueOf(map1)
		map2Value := reflect.ValueOf(map2)

		mergedMap := make(map[interface{}]interface{})

		for _, key := range map1Value.MapKeys() {
			mergedMap[key.Interface()] = map1Value.MapIndex(key).Interface()
		}

		for _, key := range map2Value.MapKeys() {
			mergedMap[key.Interface()] = map2Value.MapIndex(key).Interface()
		}

		return mergedMap
	} else {
		log.Println("map1 和 map2 类型不相等！")
		return nil
	}
}

// Sub map 相减，查找第一个 map 中不存在于第二个 map 的键值，返回一个 map
func Sub(map1, map2 interface{}) interface{} {
	if TypesEqual(map1, map2) {
		map1Value := reflect.ValueOf(map1)
		map2Value := reflect.ValueOf(map2)

		subMap := make(map[interface{}]interface{})

		for _, key := range map1Value.MapKeys() {
			map2Elem := map2Value.MapIndex(key)
			map1Elem := map1Value.MapIndex(key)

			if !map2Elem.IsValid() || map1Elem.Interface() != map2Elem.Interface() {
				subMap[key.Interface()] = map1Elem.Interface()
			}
		}
		return subMap
	} else {
		log.Println("map1 和 map2 类型不相等！")
		return nil
	}
}

// DifferentSet 映射差集，找出两个映射相互不存在的键值，组合成一个映射去重返回
/*
	同键不同值，保留 map1 的键值
*/
func DifferentSet(map1, map2 interface{}) interface{} {
	if TypesEqual(map1, map2) {
		m1 := reflect.ValueOf(map1)
		m2 := reflect.ValueOf(map2)

		different := make(map[interface{}]interface{})

		// 遍历 m1，找出 m2 中不存在的键值，并添加到 different 中，存在则从 m2 中删除这些键值
		for _, key := range m1.MapKeys() {
			if !m2.MapIndex(key).IsValid() {
				different[key.Interface()] = m1.MapIndex(key).Interface()
			} else {
				m2.SetMapIndex(key, reflect.Value{})
			}
		}
		// 将 m2 中剩下的键值添加到 different 中
		for _, key := range m2.MapKeys() {
			if m2.MapIndex(key).IsValid() {
				different[key.Interface()] = m2.MapIndex(key).Interface()
			}
		}
		return different
	} else {
		log.Println("map1 和 map2 类型不相等！")
		return nil
	}
}
