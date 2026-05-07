func productExceptSelf(nums []int) []int {
	front := make([]int, len(nums))
	for i, _ := range nums {
		if i==0 {front[i]=nums[i];continue} 
		front[i]=front[i-1]*nums[i]
	}
	fmt.Println(front)
	back := make([]int, len(nums))
	for i:=len(nums)-1;i>=0;i--{
		if i==len(nums)-1{back[i]=nums[i];continue}
		back[i]=back[i+1]*nums[i]	
	}
	fmt.Println(back)
	res := make([]int, len(nums))
	for i, _ := range nums {
		if i==0{res[i]=back[i+1];continue}
		if i==len(nums)-1{res[i]=front[i-1];continue}
		res[i] = front[i-1]*back[i+1]
	}

	return res
}
