package loops

func ForLoop(start int, stop int) (sum int) {
	sum = 0

	for x := start; x <= stop; x++ {
		sum += x
	}

	return
}

func WhileLoop(start int, stop int) (sum int) {
	sum = 0
	x := start

	for x <= stop {
		sum += x
		x += 1
	}

	return
}
