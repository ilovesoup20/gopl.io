package main

import "fmt"

var months = [...]string{1: "january", 2: "february", 3: "march", 4: "april", 5: "may", 6: "june", 7: "july", 8: "august", 9: "september", 10: "october", 11: "november", 12: "december"}

func main() {
	q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(q2)
	fmt.Println(summer)

	for _, s := range summer {
		for _, q := range q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	// fmt.Println(summer[:20])	// panic: out of range
	endlessSummer := summer[:5] //extends a slice (within capacity)
	fmt.Println(endlessSummer)

	fmt.Println("")

	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("nums %d cap=%d, len=%d\n", nums, cap(nums), len(nums))

	sub1 := nums[1:]
	fmt.Printf("sub1 %d cap=%d, len=%d\n", sub1, cap(sub1), len(sub1))
	nums[2] = 30
	sub1[0] = 20
	fmt.Printf("nums=%d, sub1=%d\n", nums, sub1)

	sub1 = append(sub1, 7)
	fmt.Printf("sub1 %d cap=%d, len=%d\n", sub1, cap(sub1), len(sub1))
	fmt.Printf("nums=%d, sub1=%d\n", nums, sub1)

	nums[2] = 300
	sub1[0] = 200
	fmt.Printf("nums=%d, sub1=%d\n", nums, sub1)

	fmt.Println(nums)
	nums = append(nums, 8)
	fmt.Println(nums, sub1)
}
