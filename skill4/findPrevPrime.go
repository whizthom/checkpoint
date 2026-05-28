package main

// Write a function that returns the first prime number that is equal
//  or inferior to the int passed as parameter.
// If there are no primes inferior to the int passed as parameter the
//  function should return 0.

func FindPrevPrime(nb int)int{
	if nb < 1{
		return 0
	}

	for i := nb; i >= 2; i--{
		if isPrime(i){
			return i
		}
	}
	return 0
}

func isPrime(n int)bool{
	if n < 2 { return false}
	if n==2 || n == 3{ return true}
	if n%2 == 0 || n %3 == 0 {return false}
return true 
}