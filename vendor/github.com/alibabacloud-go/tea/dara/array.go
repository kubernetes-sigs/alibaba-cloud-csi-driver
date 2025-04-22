package dara

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// ArrContains checks if an element is in the array
func ArrContains(arr interface{}, value interface{}) bool {
	arrValue := reflect.ValueOf(arr)
	valueValue := reflect.ValueOf(value)

	// Ensure arr is a slice
	if arrValue.Kind() != reflect.Slice {
		return false
	}

	for i := 0; i < arrValue.Len(); i++ {
		elem := arrValue.Index(i)

		// Ensure the array element is a pointer
		if elem.Kind() == reflect.Ptr {
			if valueValue.Kind() == reflect.Ptr && elem.Pointer() == valueValue.Pointer() {
				return true
			}
			if elem.Elem().Interface() == valueValue.Interface() {
				return true
			}
		} else if elem.Kind() == reflect.Interface {
			elem = elem.Elem()
			if valueValue.Kind() == reflect.Ptr && elem.Interface() == valueValue.Pointer() {
				return true
			}
			if elem.Interface() == valueValue.Interface() {
				return true // Return the index if found
			}
		}
	}

	return false
}

// ArrIndex returns the index of the element in the array
func ArrIndex(arr interface{}, value interface{}) int {
	arrValue := reflect.ValueOf(arr)
	valueValue := reflect.ValueOf(value)

	// Ensure arr is a slice
	if arrValue.Kind() != reflect.Slice {
		return -1
	}

	for i := 0; i < arrValue.Len(); i++ {
		elem := arrValue.Index(i)

		// Ensure the array element is a pointer
		if elem.Kind() == reflect.Ptr {
			// Dereference the pointer to get the underlying value
			if valueValue.Kind() == reflect.Ptr && elem.Pointer() == valueValue.Pointer() {
				return i
			}
			if elem.Elem().Interface() == valueValue.Interface() {
				return i // Return the index if found
			}
		} else if elem.Kind() == reflect.Interface {
			elem = elem.Elem()
			if valueValue.Kind() == reflect.Ptr && elem.Interface() == valueValue.Pointer() {
				return i
			}
			if elem.Interface() == valueValue.Interface() {
				return i // Return the index if found
			}
		}
	}

	return -1 // Return -1 if not found
}

func handlePointer(elem reflect.Value) string {
	if elem.IsNil() {
		return "" // Skip nil pointers
	}

	// Dereference the pointer
	elem = elem.Elem()
	return handleValue(elem)
}

func handleValue(elem reflect.Value) string {
	switch elem.Kind() {
	case reflect.String:
		return elem.String()
	case reflect.Int:
		return fmt.Sprintf("%d", elem.Interface())
	case reflect.Float64:
		return fmt.Sprintf("%f", elem.Interface())
	case reflect.Bool:
		return fmt.Sprintf("%t", elem.Interface())
	default:
		return "" // Skip unsupported types
	}
}

func ArrJoin(arr interface{}, sep string) string {
	var strSlice []string
	var str string

	arrValue := reflect.ValueOf(arr)

	// Ensure arr is a slice
	if arrValue.Kind() != reflect.Slice {
		return ""
	}

	for i := 0; i < arrValue.Len(); i++ {
		elem := arrValue.Index(i)

		if elem.Kind() == reflect.Ptr {
			str = handlePointer(elem)
		} else if elem.Kind() == reflect.Interface {
			str = handleValue(elem.Elem())
		} else {
			str = handleValue(elem)
		}

		if str != "" {
			strSlice = append(strSlice, str)
		}
	}

	return strings.Join(strSlice, sep)
}

// ArrShift removes the first element from the array
func ArrShift(arr interface{}) interface{} {
	arrValue := reflect.ValueOf(arr)

	// Ensure arr is a pointer to a slice
	if arrValue.Kind() != reflect.Ptr || arrValue.Elem().Kind() != reflect.Slice {
		return nil
	}

	// Get the slice from the pointer
	sliceValue := arrValue.Elem()

	// Ensure the slice is not empty
	if sliceValue.Len() == 0 {
		return nil
	}

	// Get the first element
	firstElem := sliceValue.Index(0)

	// Create a new slice with one less element
	newArrValue := reflect.MakeSlice(sliceValue.Type(), sliceValue.Len()-1, sliceValue.Cap())

	// Copy the elements after the first one to the new slice
	reflect.Copy(newArrValue, sliceValue.Slice(1, sliceValue.Len()))

	// Set the original slice to the new slice
	sliceValue.Set(newArrValue)

	// Return the removed first element
	return firstElem.Interface()
}

// ArrPop removes the last element from the array
func ArrPop(arr interface{}) interface{} {
	arrValue := reflect.ValueOf(arr)

	// Ensure arr is a pointer to a slice
	if arrValue.Kind() != reflect.Ptr || arrValue.Elem().Kind() != reflect.Slice {
		return nil
	}

	// Get the slice from the pointer
	sliceValue := arrValue.Elem()

	// Ensure the slice is not empty
	if sliceValue.Len() == 0 {
		return nil
	}

	// Get the last element
	lastIndex := sliceValue.Len() - 1
	lastElem := sliceValue.Index(lastIndex)

	// Create a new slice with one less element
	newArrValue := reflect.MakeSlice(sliceValue.Type(), sliceValue.Len()-1, sliceValue.Cap()-1)

	// Copy the elements before the last one to the new slice
	reflect.Copy(newArrValue, sliceValue.Slice(0, lastIndex))

	// Set the original slice to the new slice
	sliceValue.Set(newArrValue)

	// Return the removed last element
	return lastElem.Interface()
}

// ArrUnshift adds an element to the beginning of the array
func ArrUnshift(arr interface{}, value interface{}) int {
	arrValue := reflect.ValueOf(arr)

	// Ensure arr is a pointer to a slice
	if arrValue.Kind() != reflect.Ptr || arrValue.Elem().Kind() != reflect.Slice {
		return 0
	}

	// Get the slice from the pointer
	sliceValue := arrValue.Elem()

	// Create a new slice with one additional element
	newArrValue := reflect.MakeSlice(sliceValue.Type(), sliceValue.Len()+1, sliceValue.Cap()+1)

	// Set the new element as the first element
	newArrValue.Index(0).Set(reflect.ValueOf(value))

	// Copy the old elements to the new slice, starting at index 1
	reflect.Copy(newArrValue.Slice(1, newArrValue.Len()), sliceValue)

	// Set the original slice to the new slice
	sliceValue.Set(newArrValue)

	// Return the new length of the slice
	return newArrValue.Len()
}

// ArrPush adds an element to the end of the array
func ArrPush(arr interface{}, value interface{}) int {
	arrValue := reflect.ValueOf(arr)

	// Ensure arr is a pointer to a slice
	if arrValue.Kind() != reflect.Ptr || arrValue.Elem().Kind() != reflect.Slice {
		return 0
	}

	// Get the slice from the pointer
	sliceValue := arrValue.Elem()

	// Create a new slice with one additional element
	newArrValue := reflect.MakeSlice(sliceValue.Type(), sliceValue.Len()+1, sliceValue.Cap()+1)

	// Copy the old elements to the new slice
	reflect.Copy(newArrValue, sliceValue)

	// Set the new element as the last element
	newArrValue.Index(sliceValue.Len()).Set(reflect.ValueOf(value))

	// Set the original slice to the new slice
	sliceValue.Set(newArrValue)

	// Return the new length of the slice
	return newArrValue.Len()
}

// ConcatArr concatenates two arrays
func ConcatArr(arr1 interface{}, arr2 interface{}) interface{} {
	var result []interface{}
	value1 := reflect.ValueOf(arr1)
	value2 := reflect.ValueOf(arr2)

	// 检查 arr1 和 arr2 是否为切片
	if value1.Kind() != reflect.Slice || value2.Kind() != reflect.Slice {
		panic("ConcatArr: both inputs must be slices")
	}

	// 如果两个切片的类型相同
	if value1.Type() == value2.Type() {
		// 创建一个新的切片，类型与输入切片相同
		result := reflect.MakeSlice(value1.Type(), 0, value1.Len()+value2.Len())

		// 复制第一个切片的元素
		for i := 0; i < value1.Len(); i++ {
			result = reflect.Append(result, value1.Index(i))
		}
		// 复制第二个切片的元素
		for i := 0; i < value2.Len(); i++ {
			result = reflect.Append(result, value2.Index(i))
		}
		return result.Interface() // 返回类型相同的切片
	}

	// 否则返回 []interface{}
	for i := 0; i < value1.Len(); i++ {
		result = append(result, value1.Index(i).Interface())
	}
	for i := 0; i < value2.Len(); i++ {
		result = append(result, value2.Index(i).Interface())
	}
	return result
}

// ArrAppend inserts a new pointer at a specified index in a pointer array.
func ArrAppend(arr interface{}, value interface{}, index int) {
	arrV := reflect.ValueOf(arr)
	if arrV.Kind() != reflect.Ptr || arrV.Elem().Kind() != reflect.Slice {
		return
	}

	sliceV := arrV.Elem()

	if index < 0 || index > sliceV.Len() {
		return
	}

	valueV := reflect.ValueOf(value)

	// 创建一个容纳新值的切片
	newSlice := reflect.Append(sliceV, reflect.Zero(sliceV.Type().Elem()))
	reflect.Copy(newSlice.Slice(index+1, newSlice.Len()), newSlice.Slice(index, newSlice.Len()-1))
	newSlice.Index(index).Set(valueV)

	// 更新原始切片
	sliceV.Set(newSlice)
	return
}

// ArrRemove removes an element from the array
func ArrRemove(arr interface{}, value interface{}) {
	arrValue := reflect.ValueOf(arr)

	// Ensure arr is a pointer to a slice
	if arrValue.Kind() != reflect.Ptr || arrValue.Elem().Kind() != reflect.Slice {
		return
	}

	// Get the slice from the pointer
	slice := arrValue.Elem()
	index := ArrIndex(slice.Interface(), value)
	// If index is found, remove the element at that index
	if index != -1 {
		// Remove the element at the specified index
		newSlice := reflect.MakeSlice(slice.Type(), 0, slice.Len()-1)

		// Copy elements before the index
		newSlice = reflect.AppendSlice(slice.Slice(0, index), slice.Slice(index+1, slice.Len()))
		// Set the new slice back to the original reference
		slice.Set(newSlice)
	}
}

func SortArr(arr interface{}, order string) interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("SortArr: input must be a slice")
	}

	// 创建一个新的切片来存储排序结果
	newSlice := reflect.MakeSlice(v.Type(), v.Len(), v.Cap())
	for i := 0; i < v.Len(); i++ {
		newSlice.Index(i).Set(v.Index(i))
	}

	order = strings.ToLower(order)

	sort.SliceStable(newSlice.Interface(), func(i, j int) bool {
		return compare(newSlice.Index(i), newSlice.Index(j), order)
	})

	return newSlice.Interface()
}

func compare(elemI, elemJ reflect.Value, order string) bool {
	valI := reflect.Indirect(elemI)
	valJ := reflect.Indirect(elemJ)

	// 对interface{}类型处理实际类型
	if elemI.Kind() == reflect.Interface {
		valI = reflect.Indirect(elemI.Elem())
	}
	if elemJ.Kind() == reflect.Interface {
		valJ = reflect.Indirect(elemJ.Elem())
	}

	if valI.Kind() != valJ.Kind() {

		if order == "asc" {
			return valI.Kind() < valJ.Kind()
		}
		return valI.Kind() > valJ.Kind()

	}

	switch kind := valI.Kind(); kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compareNumbers(valI.Int(), valJ.Int(), order)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return compareNumbers(int64(valI.Uint()), int64(valJ.Uint()), order)
	case reflect.Float32, reflect.Float64:
		return compareNumbers(valI.Float(), valJ.Float(), order)
	case reflect.String:
		return compareStrings(valI.String(), valJ.String(), order)
	case reflect.Struct:
		return compareStructs(valI, valJ, order)
	default:
		panic("SortArr: unsupported element type")
	}
}

func compareNumbers(a, b interface{}, order string) bool {
	switch order {
	case "asc":
		return a.(int64) < b.(int64)
	case "desc":
		return a.(int64) > b.(int64)
	default:
		return a.(int64) > b.(int64)
	}
}

func compareStrings(a, b string, order string) bool {
	switch order {
	case "asc":
		return a < b
	case "desc":
		return a > b
	default:
		return a > b
	}
}

func compareStructs(valI, valJ reflect.Value, order string) bool {
	if valI.NumField() > 0 && valJ.NumField() > 0 {
		fieldI := reflect.Indirect(valI.Field(0))
		fieldJ := reflect.Indirect(valJ.Field(0))
		if fieldI.Kind() == fieldJ.Kind() {
			switch fieldI.Kind() {
			case reflect.String:
				return compareStrings(fieldI.String(), fieldJ.String(), order)
			case reflect.Int:
				return compareNumbers(fieldI.Int(), fieldJ.Int(), order)
			}
		}
	}
	return false
}
