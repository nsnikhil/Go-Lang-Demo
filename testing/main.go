package main

func main() {

}

func Ints(v ...int) int {
	return ints(v)

}
func ints(v []int) int {
	if len(v) == 0 {
		return 0
	}
	return ints(v[1:]) + v[0]
}
