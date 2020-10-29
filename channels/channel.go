package main
import (
	"fmt"
	"time"
)
func worker(done chan bool)  {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func ping(pings chan<- string, msg string)  {
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string)  {
	msg := <- pings
	pongs <- msg
}

func main()  {
	messages := make(chan string,2)
	go func ()  {
		messages <- "buffered"		
		messages <- "ping"	
	}()

	msg := <- messages
	fmt.Println(msg)
	fmt.Println(<- messages)

	done := make(chan bool, 1)
	go worker(done)
	fmt.Println("aaa")
	<-done
	fmt.Println("bbb")
	// fmt.Println(<- done)


	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings,"passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
