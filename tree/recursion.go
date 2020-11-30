package tree

//Factorial N! eg:5! = 1*2*3*4*5
func Factorial(number, count int) int {

	if count == 1 {
		return number
	}
	return Factorial(number*(count-1), count-1)

}
