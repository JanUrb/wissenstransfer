// main package declaration
package main

/*
Idea: The task of the program is to collect reports from different hard- or software
components and send them to some kind of output. The output might be the internet, ipc or else.
To simplify the program, the output will be simply printed using the log package.
*/

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

//ReportDelay defines the interval between sending reports to the output.
// const ReportDelay time.Duration = 5 * time.Second
const ReportDelay = 5 * time.Second

//Event is a generic type that holds event information
type Event string

//Component is some kind of software/hardware component
type Component struct {
	Name      string
	Attribute string
}

//report is used by components to simulate reporting an event to the report channel.
//The function report has an object of type Component as receiver.
func (c Component) report(reportCh chan<- Event) {
	for { //for loop with no condition -> runs forever

		// delay := .... uses the short declaration of var delay time.Duration = ....

		//sleep for a random amount of time (1-5 seconds)
		delay := time.Duration(rand.Intn(5)+1) * time.Second
		time.Sleep(delay)

		//Create the event
		attrValue := rand.Intn(100)                                             //rand.Intn(100) returns a random number between 0 and 100
		eventStr := fmt.Sprintf("%s -> %s: %d", c.Name, c.Attribute, attrValue) //use the Name and the Attribute of the component to create the event
		ev := Event(eventStr)                                                   //cast string to Event

		reportCh <- ev //send event to channel
	}
}

//mergeReports reads the reports and sends them to the output. The output might be some log, web or some
//form of IPC.
func mergeReports(reports <-chan Event, output chan<- Event) {
	//a ticker sends a signal on the internal channel ticker.C. The interval between the signals is defined by the argument ReportDelay.
	ticker := time.NewTicker(ReportDelay)

	//when the ticker ticks, send all reports to the output
	for {
		<-ticker.C               //blocks until a value is send to the channel of the ticker
		for v := range reports { //iterate over the values stored in reports. Blocks, if reports is empty
			output <- v
		}
	}
}

//sendReports simulates sending reports to some other component. To simplify the program,  the output is just printed.
func sendReports(output <-chan Event) {
	for ev := range output {
		log.Println(ev)
	}
}

//entry point of the application, only one main function per program!
func main() {
	reportChannel := make(chan Event, 15) //buffered with space for 15 elements
	outputChannel := make(chan Event)     //unbuffered

	//Here, three components are created.

	//initialize a struct by using the names of the fields. Using the name of the fields has the advantage that you can leave out
	//some fields or change the order of initialization.
	cpu := Component{
		Name:      "CPU",
		Attribute: "Usage",
	}

	//this way of initialization does not use the names. The values are mapped according to the order of the fields in the struct declaration.
	connections := Component{"Connections", "Number of conn"}

	//you can also initialize a struct and leave the fields empty and initialize them later.
	workerRoutines := Component{}
	workerRoutines.Name = "Worker"
	workerRoutines.Attribute = "Status"

	//Start the functions in the background

	//the keyword go in front of a function call starts the function as a goroutine.
	go cpu.report(reportChannel)
	go connections.report(reportChannel)
	go workerRoutines.report(reportChannel)

	go mergeReports(reportChannel, outputChannel)

	go sendReports(outputChannel)

	//example of an anonymous function started as goroutine
	go func(delay time.Duration) {
		time.Sleep(delay)
		log.Println("Hello after", delay)
	}(2 * time.Second) // <- function call!

	//If you dont sleep, the program will exit here.
	time.Sleep(1 * time.Minute)
	log.Println("Time is over!")
}

//the init function is similar to pythons __init__ . It is executed before the main function. Image a go program with multiple packages
//and every package has a init function, all init functions will be executed before the main function. It can be used to initialize some configuration. Generally,
//it is not recommended to do too much in the init function because it makes you program harder to read.
func init() {
	log.SetFlags(0) //disable time etc from logging
}
