package main

func main() {
	one(3)
}

func one(a int) (int, int) {
	return a, a + 5
}
