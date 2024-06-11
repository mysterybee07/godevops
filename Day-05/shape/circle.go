package shape

const PI = 3.14

func Area(radius float64) float64 {
	return PI * radius * radius
}

func Perimeter(radius float64) float64 {
	return 2 * PI * radius
}
