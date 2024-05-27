package error_test

import (
	"errors"
	"fmt"
	"testing"
)

func GetFibonacci(n int) ([]int, error) {
	if n < 2 || n > 100 {
		return nil, errors.New("n should be in [2, 100]")
	} else {
		fibList := []int{1, 1}
		for i := 2; i < n; i++ {
			fibList = append(fibList, fibList[i-2]+fibList[i-1])
		}
		return fibList, nil
	}
}

var LessThanTwoError = errors.New("n should not be less than 2")
var LargerThanOneHundredError = errors.New("n should not be larger than 100")

func GetFibonacci2(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	} else if n > 100 {
		return nil, LargerThanOneHundredError
	} else {
		fibList := []int{1, 1}
		for i := 2; i < n; i++ {
			fibList = append(fibList, fibList[i-2]+fibList[i-1])
		}
		return fibList, nil
	}
}

func TestGetFibonacciSuccess(t *testing.T) {
	if v, err := GetFibonacci(5); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestGetFibonacciFail(t *testing.T) {
	if v, err := GetFibonacci(-10); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestGetFibonacci2(t *testing.T) {
	if v, err := GetFibonacci2(-5); err != nil {
		if errors.Is(err, LessThanTwoError) {
			fmt.Println("It is less.")
		} else if errors.Is(err, LargerThanOneHundredError) {
			fmt.Println("It is larger.")
		}
	} else {
		t.Log(v)
	}
}

// 最佳实践：避免嵌套，及早失败
