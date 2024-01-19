package structs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lwabish/go/pkg/util"
)

// IsValidCPUCores Parse string/int types to ints.
// Example:
// []interface{}{"1", "2", "3"} => []int{1, 2, 3}, nil
// []interface{}{"1-3"} => []int{1, 2, 3}, nil
// []interface{}{"1-2", 4, 5} => []int{1, 2, 4, 5}, nil
func IsValidCPUCores(interfaces []interface{}) ([]int, error) {
	var result []int
	for _, inter := range interfaces {
		ret, err := ParseInts(inter)
		if err != nil {
			return nil, err
		}
		result = append(result, ret...)
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("empty slice")
	}
	if HasDuplicates(result) {
		return nil, fmt.Errorf("duplicate core not allowed: %v", result)
	}
	return result, nil
}

// ParseInts Transform string/int/float64 to ints.
// Example:
//
//				"1-3" -> [1, 2, 3]
//				3     -> [3]
//	 			3.0   -> [3]
func ParseInts(v interface{}) ([]int, error) {
	var result []int
	switch val := v.(type) {
	// encoding/json treat ints as float64,so there is no int case
	case string:
		ret, err := StringToInts(val)
		if err != nil {
			return nil, err
		}
		result = append(result, ret...)
	case float64:
		if !EqualToInt(val) {
			return nil, fmt.Errorf("%v not int", val)
		}
		result = append(result, int(val))
	default:
		return nil, fmt.Errorf("unexpected type %v", val)
	}
	return result, nil
}

// EqualToInt Check if a float type number is equal to int type.
// Example: 3.12 -> false, 3 -> true
func EqualToInt[T float32 | float64](f T) bool {
	intStr := fmt.Sprintf("%v", int(f))
	floatStr := fmt.Sprintf("%v", f)
	return strings.Compare(intStr, floatStr) == 0
}

// StringToInts Transform string to int list.
// Example: "1-3" -> [1, 2, 3]
func StringToInts(str string) ([]int, error) {
	strSplits := strings.Split(str, "-")
	if len(strSplits) > 2 {
		return nil, fmt.Errorf("%s invalid", str)
	} else if len(strSplits) == 1 {
		ret, err := strconv.Atoi(strSplits[0])
		if err != nil {
			return nil, fmt.Errorf("string %s not invalid int", strSplits[0])
		}
		return []int{ret}, nil
	} else if len(strSplits) == 0 {
		return nil, fmt.Errorf("string must be not empty")
	}

	// two parts
	left, err := strconv.Atoi(strSplits[0])
	if err != nil {
		return nil, fmt.Errorf("%s left - strings not invalid int", strSplits[0])
	}
	right, err := strconv.Atoi(strSplits[1])
	if err != nil {
		return nil, fmt.Errorf("%s right - strings not invalid int", strSplits[1])
	}

	if left > right {
		return nil, fmt.Errorf("left %d > right %d not invalid", left, right)
	}
	var result = make([]int, right-left+1)
	var index int
	for i := left; i <= right; i++ {
		result[index] = i
		index++
	}
	return result, nil
}

// HasDuplicates Checks if the given slice has duplicates.
func HasDuplicates[T util.Hashable](data []T) bool {
	var m = make(map[T]struct{})
	for _, d := range data {
		if _, ok := m[d]; ok {
			return true
		}
		m[d] = struct{}{}
	}
	return false
}
