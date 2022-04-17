package main

func maxProduct(nums []int) int {
	max1 := -1
	max2 := -1

	for _, n := range nums {
		if n > max1 {
			max2 = max1
			max1 = n
		} else if n > max2 {
			max2 = n
		}
	}

	return (max1 - 1) * (max2 - 1)
}
