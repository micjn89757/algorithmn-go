package algorithmn

// BinarySearchFirstImplement 二分查找第一种实现 [left, right]
func BinarySearchFirstImplement(nums []int, target int) int {
	left := 0
	right := len(nums)
	mid := 0

	for left < right {
		mid = (left + right) / 2

		if target > nums[mid] {
			left = mid + 1
		}else if target < nums[mid] {
			right = mid
		}else {
			return mid
		}
	}

	return -1
}


// BinarySearchSecondImplement 二分查找第二种实现 [left, right)
func BinarySearchSecondImplement(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := 0

	for left <= right {
		mid = (left + right) / 2

		if target > nums[mid] {
			left = mid + 1
		}else if target < nums[mid] {
			right = mid - 1
		}else {
			return mid
		}
	}
	
	return -1
}

