package main

func main() {
	Sum(5, 4, 3, 2)

}

func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
