package main

import (
	"fmt"
	"math/rand"
	"time"
)

// A synchronized queue consists of elements that need to be processed in a particular
// sequence. Passenger queue and ticket processing queues are types of synchronized queues.

const (
	messagePassStart = iota
	messageTicketStart
	messagePassEnd
	messageTicketEnd
)

type Queue struct {
	waitPass    int
	waitTicket  int
	playPass    bool
	playTicket  bool
	queuePass   chan int
	queueTicket chan int
	message     chan int
}

// The New method on Queue initializes message, queuePass, and queueTicket with nil values.
func (queue *Queue) New() {
	queue.message = make(chan int)
	queue.queuePass = make(chan int)
	queue.queueTicket = make(chan int)

	go func() {
		var message int
		for {
			select {
			case message = <-queue.message:
				switch message {
				case messagePassStart:
					queue.waitPass++
				case messagePassEnd:
					queue.playPass = false
				case messageTicketStart:
					queue.waitTicket++
				case messageTicketEnd:
					queue.playTicket = false
				}
				if queue.waitPass > 0 && queue.waitTicket > 0 && !queue.playPass &&
					!queue.playTicket {
					queue.playPass = true
					queue.playTicket = true
					queue.waitTicket--
					queue.waitPass--
					queue.queuePass <- 1
					queue.queueTicket <- 1
				}
			}
		}
	}()
}

// The StartTicketIssue method starts the issuing of a ticket for passengers standing in a
// queue. The StartTicketIssue method on Queue sends messageTicketStart to the
// message queue and queueTicket receives the message.
func (Queue *Queue) StartTicketIssue() {
	Queue.message <- messageTicketStart
	<-Queue.queueTicket
}

// The EndTicketIssue method finishes the issuing of a ticket to a passenger standing in the
// queue.
func (Queue *Queue) EndTicketIssue() {
	Queue.message <- messageTicketEnd
}

// The ticketIssue method starts and finishes the issuing of a ticket to the passenger.
// The ticketIssue method invokes the StartTicketIssue and
// EndTicketIssue methods after Sleep calls for 10 seconds and two seconds. The ticket is
// issued after the ticket is processed.
func ticketIssue(Queue *Queue) {
	for {
		// Sleep up to 10 seconds.
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		Queue.StartTicketIssue()
		fmt.Println("Ticket Issue starts")
		// Sleep up to 2 seconds.
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println("Ticket Issue ends")
		Queue.EndTicketIssue()
	}
}

// The StartPass method starts the passenger queue moving toward the ticket counter. The
// StartPass method on Queue sends messagePassStart to the message queue and
// queuePass receives the message.
func (Queue *Queue) StartPass() {
	Queue.message <- messagePassStart
	<-Queue.queuePass
}

// The EndPass method stops the passenger queue moving toward the ticket counter. The
// EndPass method on Queue sends messagePassEnd to the message queue in the following
// code. The passenger is moved to the counter for ticket processing, and the passenger is then
// out of the queue.
func (Queue *Queue) EndPass() {
	Queue.message <- messagePassEnd
}

// The passenger methods starts and ends passenger movement to the queue. The
// passenger method invokes the StartPass method, and the EndPass method ends after
// sleep calls for 10 seconds and two seconds. The passenger moves into the queue and
// reaches the ticket counter.
func passenger(Queue *Queue) {
	// fmt.Println("starting the passenger Queue")
	for {
		// fmt.Println("starting the processing")
		// Sleep up to 10 seconds.
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		Queue.StartPass()
		fmt.Println(" Passenger starts")
		// Sleep up to 2 seconds.
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println(" Passenger ends")
		Queue.EndPass()
	}
}

func main() {
	var Queue *Queue = &Queue{}
	// fmt.Println(Queue)
	Queue.New()
	fmt.Println(Queue)
	var i int
	for i = 0; i < 10; i++ {
		// fmt.Println(i, "passenger in the Queue")
		go passenger(Queue)
	}
	// close(Queue.queuePass)
	var j int
	for j = 0; j < 5; j++ {
		// fmt.Println(i, "ticket issued in the Queue")
		go ticketIssue(Queue)
	}
	select {}
}
