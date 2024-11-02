package arrays

func Sum(nums []int) int {
	var sum = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(numsToSum ...[]int) []int {
	numOfSlices := len(numsToSum)
	sums := make([]int, numOfSlices)

	for _, nums := range numsToSum {
		sums = append(sums, Sum(nums))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}