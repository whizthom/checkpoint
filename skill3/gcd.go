package main


// QUESTION:
// Compute the GCD of two integers using Euclid's algorithm.

func GCD(a, b int) int{

	if a < 0 {
		a = -a
	}

	if b < 0{
		b = -b
	}

	for b != 0 {
		a , b = b , a%b
	}

	return a
}

// QUESTION
// Write a function that takes two uint representing two strictly positive integers and returns their greatest common divisor.
// If any of the input numbers is 0, the function should return 0.

func Gcd(a, b uint) uint{

	if a < 0 {
		a = -a
	}

	if b < 0 {
		b = -b
	}

	if a == 0 || b == 0 {
		return 0
	}

	for b != 0{
		a , b = b , a%b
	}
	return a
}