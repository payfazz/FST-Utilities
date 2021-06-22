package utilities

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func ConvertStringToInt(data string) (int, error) {
	var (
		err error
		tmp int64
	)
	// res, err = strconv.Atoi(data)
	tmp, err = ConvertStringToInt64(data)
	if err != nil {
		return 0, err
	}
	return int(tmp), nil
}
func ConvertStringToInt64(data string) (int64, error) {
	var (
		err error
		tmp float64
	)
	if strings.Contains(data, ",") {
		data = strings.Replace(data, ",", "", -1)
	}
	tmp, err = strconv.ParseFloat(data, 64)
	if err != nil {
		return 0, err
	}
	tmp = math.Round(tmp)
	return int64(tmp), nil
}
func ConvertStringtoIntArray(data string) ([]int, error) {
	strArr := strings.Split(data, ",")
	res := []int{}
	for _, val := range strArr {
		num, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		res = append(res, num)
	}
	return res, nil
}
func ConvertToInt(data interface{}) (int, error) {
	if data == nil {
		return 0, nil
	}
	if res, ok := data.(int64); ok {
		return int(res), nil
	} else {
		return 0, fmt.Errorf("parsing to int error")
	}
}
func ConvertToInt64(data interface{}) (int64, error) {
	if res, ok := data.(int64); ok {
		return res, nil
	} else {
		return 0, fmt.Errorf("parsing to int64 error")
	}
}
func ConvertToString(data interface{}) (string, error) {
	if data == nil {
		return "", nil
	}
	if res, ok := data.(string); ok {
		return res, nil
	} else {
		return "", fmt.Errorf("parsing to string error")
	}
}

// Only Accept Array of Pointer Struct
func ConvertStructToInterface(data interface{}) ([][]interface{}, error) {
	result := [][]interface{}{}
	header := []interface{}{}
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Slice {
		return nil, fmt.Errorf("required Slice")
	}
	for i := 0; i < val.Len(); i++ {
		var row reflect.Value
		tmp := []interface{}{}
		switch val.Index(i).Kind() {
		case reflect.Ptr:
			row = val.Index(i).Elem()
		case reflect.Struct:
			row = val.Index(i)
		default:
			return nil, fmt.Errorf("the data is not struct or pointer of struct")
		}

		for x := 0; x < row.NumField(); x++ {
			field := row.Type().Field(x)
			if len(header) < row.NumField() {
				tag, ok := field.Tag.Lookup("alias")
				if ok {
					header = append(header, tag)
				} else {
					header = append(header, field.Name)
				}
			}
			tmp = append(tmp, fmt.Sprintf("%v", row.Field(x)))
		}
		if len(result) == 0 {
			result = append(result, header)
		}
		result = append(result, tmp)
	}

	return result, nil
}
