package main
import "fmt"

// IN THIS CODE, I CONSIDERED THAT THE DIFFERENT LENGTHS ARE OF TYPE INT.

type Shape interface{
	String() string
}

type Rectangle struct {
	X int
	Y int
}

type Circle struct {
	R int
}

func(rectangle Rectangle) String() string{
	return fmt.Sprintf("Square of width %d and length %d",rectangle.X,rectangle.Y)
}

func(circle Circle) String() string{
	return fmt.Sprintf("Circle of radius %d",circle.R)
}

func main() {
	r := &Rectangle{2, 3}
    c := &Circle{5}

    shapes := []Shape{r, c}

    for _, s := range shapes {
        fmt.Println(s)
	}
}
