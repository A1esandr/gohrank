package main

func Merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {

	if n == 0 {
		return
	}

	buf1 := make(chan int, n)
	buf2 := make(chan int, n)
	closebuf1 := make(chan int, n)
	closebuf2 := make(chan int, n)

	go load(in1, buf1, n)
	go load(in2, buf2, n)

	out1 := make(chan int, n)
	out2 := make(chan int, n)
	closeout := make(chan int, n)

	for i := 0; i < n; i++ {
		go exec(f, buf1, out1, closebuf1)
		go exec(f, buf2, out2, closebuf2)
		go sum(out1, out2, out, closeout)
	}

	go closechan(closebuf1, buf1, n)
	go closechan(closebuf2, buf2, n)
	go closechanout(closeout, out1, out2, n)
}

func load(in1 <-chan int, buf chan int, n int) {
	for i := 0; i < n; i++ {
		select {
		case v := <-in1:
			buf <- v
		}
	}
}

func exec(f func(int) int, in <-chan int, out chan int, closebuf chan int) {
	n0 := <-in
	select {
	case out <- f(n0):
		closebuf <- 1
		return
	}
}

func sum(in1 <-chan int, in2 <-chan int, out chan<- int, closebuf chan int) {
	n1 := <-in1
	n2 := <-in2
	select {
	case out <- n1 + n2:
		closebuf <- 1
		return
	}
}

func closechan(closebuf chan int, buf chan int, n int) {
	for i := 0; i < n; i++ {
		select {
		case <-closebuf:
		}
	}
	close(closebuf)
	close(buf)
}

func closechanout(closeout chan int, out1 chan int, out2 chan int, n int) {
	for i := 0; i < n; i++ {
		select {
		case <-closeout:
		}
	}
	close(closeout)
	close(out1)
	close(out2)
}
