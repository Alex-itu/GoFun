package main

import (
	"fmt"
	"math/rand"
	"time"
)

func client(ack, seq chan int) {
	int x := rand.Intn(1000)
	fmt.Println("The CLIENT is now sending a SYN")
	seq <- x

	int y
	if (ack == x + 1) {
		x = x + 1
		fmt.Println("The CLIENT got the right ACK from SERVER")
		y := <- seq
		y = y + 1

		ack <- y + 1
		seq <- x + 1
		fmt.Println("The CLIENT is sending last step in the three way handsake")
	}


}

func server(syn, seq chan int) {
	int x := <- syn
	fmt.Println("The SERVER is now resirving a SYN")
	
	int y := rand.Intn(1000)
	x = x + 1
	seq <- x
	syn <- y
	fmt.Println("The SERVER is now sending a SYN to the CLIENT")

	if (syn == x + 1) {
		fmt.Println("The SERVER got the right ACK from CLIENT")
		fmt.Println("The three way handsake is now complete and SERVER and CLIENT can exchange infomation")
	}
}

func main() {
	var chAck_Syn = make(chan int)
	var chSeq = make(chan int)

	go client(chAck_Syn, chSeq)
	go server(chAck_Syn, chSeq)

	time.Sleep(time.Millisecond * 1000)
	fmt.Println("DONE")
	
}