package main

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
)

func main() {

	config := &serial.Config{
		Name: "/dev/ttyUSB0",
		Baud: 9600,
	}

	message := ""

	for {
		fmt.Println("Enter message for adruino")
		fmt.Println("ON,OFF or quit: ")
		_, err := fmt.Scanln(&message)
		if err != nil {
			fmt.Println("Error reading line")
		}
		if message == "quit" {
			break
		}

		s, err := serial.OpenPort(config)
		if err != nil {
			log.Fatal(err)
		}

		n, err := s.Write([]byte(message))
		if err != nil {
			log.Fatal(err)
		}

		buf := make([]byte, 128)
		n, err = s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%q", buf[:n])
	}

}
