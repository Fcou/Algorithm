package main

import "fmt"

func removeDuplicatewes(a []int) int {
	left, right, size := 0, 1, len(a)
	for ; right < size; right++ {
		if a[left] == a[right] {
			continue
		}
		left++
		a[left], a[right] = a[right], a[left] //此调换语法要记住
	}

	return left + 1
}
func main() {
	a := []int{1, 1, 2, 2, 2, 4, 4, 5, 6, 7, 7, 7, 7, 9}

	lens := removeDuplicatewes(a)

	fmt.Println("The len of new arrays is:", lens)

}
