package main

/*
Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
*/

import "fmt"

func main() {
	var (
		next     int64 = 1
		temp     int64 = 0
		previous int64 = 1
		sum      int64 = 0
	)

	for i := 0; i < 50; i++ {
		next = previous + next
		if next%2 == 0 && next < 4000000 {
			sum = sum + next
			fmt.Println(next)
		}
		temp = previous
		previous = next
		next = temp
	}
	fmt.Println("Sum is: ", sum)
}