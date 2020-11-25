package main

import "fmt"

func Quicksort(ar []int){
	if len(ar) <= 1 {
		return
	}
	split := devision(ar)

	Quicksort(ar[:split])
	Quicksort(ar[split:])
}

func devision(ar []int) int{
	pivot := ar[len(ar)/2]

	left := 0
	right := len(ar) - 1


	for {
		for ; ar[left] < pivot; left++ {
		}

		for ; ar[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}
		ar[left], ar[right] = ar[right], ar[left]
	}
}



func main() {
	ar := []int{0,2,17,87,-23,56,134,200,4,5,6,8,10,20,13}
	Quicksort(ar)
	fmt.Println(ar)
}
