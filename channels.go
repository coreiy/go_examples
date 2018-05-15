package main

import "time"

func factorialRec(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorialRec(n-1)
}

func factorialRecConcurrency(n int, cIn chan int) chan int {
	cOut := make(chan int)
	if n <= 1 {
		go func() {
			cOut <- 1
		}()
		return cOut
	}
	go func() {
		cOut <- (n * <-factorialRecConcurrency(n-1, cIn))
	}()
	return cOut
}

func main() {
	factorial1 := 20
	cIn := make(chan int)
	currentTime := time.Now()
	cOut := factorialRecConcurrency(factorial1, cIn)
	println(factorial1, "! = ", <-cOut)
	println("elapsed: ", time.Since(currentTime))
	currentTime = time.Now()
	println(factorial1, "! = ", factorialRec(factorial1))
	println("elapsed: ", time.Since(currentTime))
}
