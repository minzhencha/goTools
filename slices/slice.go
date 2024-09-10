package slices

import (
	"log"
	"reflect"
)

// IsSlice 判断是否为切片类型
func IsSlice(slice interface{}) bool {
	return reflect.TypeOf(slice).Kind() == reflect.Slice || reflect.TypeOf(slice).Kind() == reflect.Array
}

// InSlice 元素是否在切片中
func InSlice(elem interface{}, slice interface{}) bool {
	if IsSlice(slice) {
		sliceValue := reflect.ValueOf(slice)
		for i := 0; i < sliceValue.Len(); i++ {
			if sliceValue.Index(i).Interface() == elem {
				return true
			}
		}
	}

	return false
}

// TypesEqual 判断两个切片类型是否相等
func TypesEqual(slice1, slice2 interface{}) bool {
	return reflect.TypeOf(slice1) == reflect.TypeOf(slice2)
}

// ValuesEqual 判断两个切片是否相等(按顺序)
/*
	eg: []int{1,2,3,4,5} == []int{1,2,3,4,5}
        []int{1,2,3,4,5} != []int{5,4,3,2,1}
*/
func ValuesEqual(slice1, slice2 interface{}) bool {
	return reflect.DeepEqual(slice1, slice2)
}

// ValuesSame 判断两个切片的值是否相同(不按顺序)
/*
	排序后也相等
	eg: []int{1,2,3,4,5} == []int{5,4,3,2,1}
*/
func ValuesSame(slice1, slice2 interface{}) bool {
	if TypesEqual(slice1, slice2) {
		// 切片值
		slice1Value := reflect.ValueOf(slice1)
		slice2Value := reflect.ValueOf(slice2)

		// 创建一个 map 用于存储切片中的元素
		elementsMap := make(map[interface{}]bool)

		// 遍历第二个切片，并将元素存储在 map 中
		for i := 0; i < slice2Value.Len(); i++ {
			elementsMap[slice2Value.Index(i).Interface()] = true
		}

		// 检查第一个切片中的元素是否存在于 map 中，如果不存在则返回 false
		for i := 0; i < slice1Value.Len(); i++ {
			elem := slice1Value.Index(i).Interface()
			if ok := elementsMap[elem]; !ok {
				return false
			}
		}

		return true
	} else {
		log.Println("slice1 和 slice2 类型不相等！")
		return false
	}
}

// Contains slice1 是否包含 slice2
func Contains(slice1, slice2 interface{}) bool {
	if TypesEqual(slice1, slice2) {
		// 切片值
		//slice1Value := reflect.ValueOf(slice1)
		slice2Value := reflect.ValueOf(slice2)

		// 遍历第二个切片，检查元素是否存在于第一个切片中
		for i := 0; i < slice2Value.Len(); i++ {
			elem := slice2Value.Index(i).Interface()
			if !InSlice(elem, slice1) {
				return false
			}
		}

		return true
	} else {
		log.Println("slice1 和 slice2 类型不相等！")
		return false
	}
}

// Merges 两个切片合并
func Merges(slice1, slice2 interface{}) interface{} {
	if TypesEqual(slice1, slice2) {
		return reflect.AppendSlice(reflect.ValueOf(slice1), reflect.ValueOf(slice2)).Interface()
	} else {
		log.Println("slice1 和 slice2 类型不相等！")
		return nil
	}
}

// Sub 切片相减，查找第一个切片中不存在于第二个切片的元素，返回一个切片
func Sub(slice1, slice2 interface{}) interface{} {
	if TypesEqual(slice1, slice2) {
		// 切片值
		slice1Value := reflect.ValueOf(slice1)
		slice2Value := reflect.ValueOf(slice2)

		// 创建一个 map 用于存储切片中的元素
		elementsMap := make(map[interface{}]bool)

		// 遍历第二个切片，并将元素存储在 map 中
		for i := 0; i < slice2Value.Len(); i++ {
			elementsMap[slice2Value.Index(i).Interface()] = true
		}

		// 检查第一个切片中的元素是否存在于 map 中，如果不存在则加入结果切片
		resultSlice := reflect.MakeSlice(slice1Value.Type(), 0, 0)
		for i := 0; i < slice1Value.Len(); i++ {
			elem := slice1Value.Index(i).Interface()
			if ok := elementsMap[elem]; !ok {
				resultSlice = reflect.Append(resultSlice, reflect.ValueOf(elem))
			}
		}

		return resultSlice.Interface()
	} else {
		log.Println("slice1 和 slice2 类型不相等！")
		return nil
	}
}

// Delete 删除切片中的元素
func Delete(slice interface{}, elem interface{}) interface{} {
	if !IsSlice(slice) {
		log.Println("slice 不是切片类型！")
		return slice
	} else if !InSlice(elem, slice) {
		log.Println("elem 不在切片中！")
		return slice
	}

	sliceValue := reflect.ValueOf(slice)
	// 创建一个新的切片来存储结果
	newSlice := reflect.MakeSlice(sliceValue.Type(), 0, sliceValue.Len()-1)
	for i := 0; i < sliceValue.Len(); i++ {
		if sliceValue.Index(i).Interface() != elem {
			newSlice = reflect.Append(newSlice, sliceValue.Index(i))
		}
	}

	return newSlice
}

// Deduplicate 切片去重
func Deduplicate(input interface{}) interface{} {
	if IsSlice(input) {
		inputSlice := reflect.ValueOf(input)
		outputMap := make(map[interface{}]bool)
		outputSlice := reflect.MakeSlice(inputSlice.Type(), 0, inputSlice.Len())
		for i := 0; i < inputSlice.Len(); i++ {
			elem := inputSlice.Index(i).Interface()
			if ok := outputMap[elem]; !ok {
				outputSlice = reflect.Append(outputSlice, reflect.ValueOf(elem))
				outputMap[elem] = true
			}
		}
		return outputSlice.Interface()
	} else {
		log.Println("传入的值不是切片类型！")
		return nil
	}
}

// DifferentSet 切片差集，找出两个切片相互不存在的元素，组合成一个切片去重返回
func DifferentSet(slice1, slice2 interface{}) interface{} {
	if TypesEqual(slice1, slice2) {
		// 切片值
		slice1Value := reflect.ValueOf(slice1)
		slice2Value := reflect.ValueOf(slice2)

		// 创建一个 map 用于存储切片中的元素
		elementsMap := make(map[interface{}]bool)

		// 遍历第一个切片，并将元素存储在 map 中
		for i := 0; i < slice1Value.Len(); i++ {
			elementsMap[slice1Value.Index(i).Interface()] = true
		}

		// 检查第二个切片中的元素是否存在于 map 中，如果不存在则加入结果切片，如果存在则从 map 中删除该元素
		differentElements := reflect.MakeSlice(slice1Value.Type(), 0, 0)
		for i := 0; i < slice2Value.Len(); i++ {
			elem := slice2Value.Index(i).Interface()
			if ok := elementsMap[elem]; !ok {
				differentElements = reflect.Append(differentElements, reflect.ValueOf(elem))
			} else {
				delete(elementsMap, elem)
			}
		}

		// 将剩余在 map 中的元素也加入结果切片
		for elem := range elementsMap {
			differentElements = reflect.Append(differentElements, reflect.ValueOf(elem))
		}

		// 去重返回
		return Deduplicate(differentElements.Interface())
	} else {
		log.Println("slice1 和 slice2 类型不相等！")
		return nil
	}
}
