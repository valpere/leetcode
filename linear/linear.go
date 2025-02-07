package linear

// 287. Find the Duplicate Number
// Two pinters
func FindDuplicate2Pointers(nums []int) int {
	ln := len(nums)
	if ln < 2 {
		return 0
	}

	// Find the intersection point of the two pointers
	slow := nums[0]
	fast := nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	// Find the entrance of the cycle
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

// 287. Find the Duplicate Number
// Binary search
func FindDuplicateBinarySearch(nums []int) int {
	ln := len(nums)
	if ln < 2 {
		return 0
	}
	left := 1
	right := ln - 1

	for left < right {
		mid := left + (right-left)/2
		count := 0

		// Count the numbers less than or equal to mid
		for _, num := range nums {
			if num <= mid {
				count += 1
			}
		}

		// If count is greater than mid, the duplicate lies in the left half
		if count > mid {
			right = mid
		} else { // Otherwise, it lies in the right half
			left = mid + 1
		}
	}

	return left
}

// 169. Majority Element
// Boyer-Moore Voting Algorithm
func MajorityElement(nums []int) int {
	count := 0
	var candidate int

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count += 1
		} else {
			count -= 1
		}
	}

	return candidate
}

// 121. Best Time to Buy and Sell Stock
// One Pass
func MaxProfit(prices []int) int {
	minprice := int(^uint(0) >> 1)
	maxprofit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minprice {
			minprice = prices[i]
		} else if prices[i]-minprice > maxprofit {
			maxprofit = prices[i] - minprice
		}
	}
	return maxprofit
}

// 461. Hamming Distance
// Brian Kernighan's Algorithm
func HammingWeight(n int) int {
	res := 0
	for n != 0 {
		res++
		n &= (n - 1)
	}

	return res
}

// 461. Hamming Distance
func HammingDistance(x int, y int) int {
	return HammingWeight(x ^ y)
}
