package examples

import "fmt"

func Loops() {
	//FOR LOOP

	for i := 0; i < 10; i++ {
		fmt.Println("FOR LOOP WITH A COUNTER", i)
	}

	//Golang doesn't have "while" or "do while" loops but we can acheive the
	//same with "for" loop by passing only boolean condition

	i := 0

	for i < 10 {
		fmt.Println("FOR LOOP AS WHILE", i)
		i++
	}

	//FOR loop without any condition will create an infinite loop.

	i = 0

	for {
		if i >= 10 {
			break //ending the loop on 10th iteration.
		}

		fmt.Println("INFINITE FOR LOOP", i)
		i++
	}

	//We can also skip the condition part when writing a counter for loop. (also creates an infinite loop)

	for j := 0; ; j++ {
		if j >= 10 {
			break //ending the loop on 10th iteration.
		}

		fmt.Println("COUNTER FOR LOOP WITHOUT CONDITION", j)
	}
}
