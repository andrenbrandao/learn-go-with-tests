package main

func Sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	result := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		result[i] = Sum(numbers)
	}

	return result
}

func SumAllTails(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	result := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		if len(numbers) == 0 {
			result[i] = 0
		} else {
			result[i] = Sum(numbers[1:])
		}
	}

	return result
}
