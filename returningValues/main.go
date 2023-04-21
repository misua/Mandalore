package main

import (
	"errors"
	"fmt"
)

func main() {

	// sum := calculateSumofDigits(12345)
	// fmt.Println(sum)
	// fmt.Println(calculateSumofDigits(12345))
	// fmt.Println(calculateSumofDigits(-1))

	sum1, err := calculateSumofDigits(12345)
	if err != nil {
		panic(err)
	}

	sum2, err := calculateSumofDigits(-1000)
	if err != nil {
		panic(err)
	}

	fmt.Println(sum1, sum2)

}

func calculateSumofDigits(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("error nega values not allowed")
	}
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum, nil
}
