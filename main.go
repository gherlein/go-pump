package main

import (
	"fmt"
	"log"
	"strings"

	"go.bug.st/serial"
)

var port serial.Port

func main() {
	openPort()
	var buffer []byte = []byte{0x44, 0x45, 0x46, 0x47}
	sendBuffer(buffer)
	readBuffer()
}

// This example prints the list of serial ports and use the first one
// to send a string "10,20,30" and prints the response on the screen.
func openPort() {

	// Open the first serial port detected at 9600bps N81
	mode := &serial.Mode{

		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	var err error
	port, err = serial.Open("/dev/ttyUSB0", mode)
	if err != nil {
		log.Fatal(err)
	}
}

func sendBuffer(buffer []byte) {
	n, err := port.Write(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sent %v bytes\n", n)
}

func readBuffer() {
	buff := make([]byte, 100)
	for {
		// Reads up to 100 bytes
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		}
		for i := 0; i < n; i++ {
			fmt.Printf("[%02X] ", buff[i])
		}
		fmt.Println()
	}
}

func sendReceive() {
	// Send th.e string "10,20,30\n\r" to the serial port
	n, err := port.Write([]byte("10,20,30\n\r"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sent %v bytes\n", n)

	// Read and print the response

	buff := make([]byte, 100)
	for {
		// Reads up to 100 bytes
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		}

		fmt.Printf("%s", string(buff[:n]))

		// If we receive a newline stop reading
		if strings.Contains(string(buff[:n]), "\n") {
			break
		}
	}
}
