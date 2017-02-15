package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	/*empty slices with certain sizes cannot be initiated with 
	var s []int for example, since this will create a nil slice
	has 0 length and 0 capacity. Recommended using make{} if 
	you want to create a dynamic sized array/slice
	*/
	
	//make a 2 dimensional array of vertical length dy, each row  	 //contains a slice of length dx 
	result:=make([][]uint8,dy)
	for i :=0;i<dy;i++{
		result[i] = make([]uint8,dx)
	}
	
	//fill each row with the functions mentioned, iteration of
	//array is very similar to how you would do it in C
	for j := 0; j<dy; j++{
		for k := 0; k<dx; k++{
			result[j][k] = uint8((j+k)/2)
			// note that unlike in C or Python, you would need
			// to reverse row and column in element
			// iteration since Go treats it in reverse
		}
	}
	
	return result
}

func main() {
	pic.Show(Pic)
}
