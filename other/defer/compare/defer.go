package compare

func Test() {
	ch1 := make(chan  int)
	ch2 := make(chan  int)
	defer close(ch1)
	defer close(ch2)
	defer close(ch2)
	defer close(ch2)
	defer close(ch2)
	defer close(ch2)
	defer close(ch2)
	defer close(ch2)
}
