package main
import (
	"fmt"
	"strconv"
)

// struct
type Vector3 struct {
	x int
	y float32
	z int
}

// main function
func main() {
  
  // array of structs
	coordinates := []Vector3{}
	index := 0

	for xc := 0; xc < 10; xc++ {
		for zc := 0; zc < 10; zc++ {
      
			vec := Vector3{x: xc, y: 0, z:zc}
			coordinates = append(coordinates,vec)
			fmt.Printf(strconv.Itoa(coordinates[index].x) + " X " + strconv.Itoa(coordinates[index].z) + " Z \n")
      
			index++
		}
	}
}
