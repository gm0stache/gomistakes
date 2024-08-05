package main

func main() {
	i := 0
	c := make(chan struct{}, 1)
	go func() {
		i++
		<-c
	}()
	c <- struct{}{}
	println(i)
}
