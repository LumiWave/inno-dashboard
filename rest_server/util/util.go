package util

import (
	"strconv"
	"time"
)

// string -> int64
func ParseInt(data string) int64 {
	value, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return 0
	}
	return value
}

// string -> int로 변환해서 크기 비교
func CompareString(data1, data2 string) int64 {
	value1 := ParseInt(data1)
	value2 := ParseInt(data2)

	if value1 < value2 {
		return -1 // data1이 작으면
	} else if value1 > value2 {
		return 1 // data1이 더 크면
	}
	return 0 //같으면 0 리턴
}

// string -> int로 변환해서 더하기
func SumString(data1, data2 string) string {
	return strconv.FormatInt(ParseInt(data1)+ParseInt(data2), 10)
}

// string -> int로 변환해서 곱하기
func MultiplyString(data1, data2 string) int64 {
	return ParseInt(data1) * ParseInt(data2)
}

// string -> int로 변환해서 빼기
func SubString(data1, data2 string) string {
	return strconv.FormatInt(ParseInt(data1)-ParseInt(data2), 10)
}

// string 배열에서 원소 찾기
func Contains(x string, arr []string) bool {
	for _, i := range arr {
		if x == i {
			return true
		}
	}
	return false
}

func ContainsInt(x int64, arr []int64) bool {
	for _, i := range arr {
		if x == i {
			return true
		}
	}
	return false
}

func IsTimeBetween(now time.Time, startStr, endStr string) (bool, error) {
	// 문자열을 time.Time 형식으로 변환
	startTime, err := time.Parse(time.RFC3339, startStr)
	if err != nil {
		return false, err
	}

	endTime, err := time.Parse(time.RFC3339, endStr)
	if err != nil {
		return false, err
	}

	// 현재 시간이 startTime과 endTime 사이인지 체크
	return now.After(startTime) && now.Before(endTime), nil
}
