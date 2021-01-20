package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	log.SetFlags(log.Llongfile)
	if len(os.Args) >= 3 {

		serverAddress := fmt.Sprintf("%s:%s", os.Args[1], os.Args[2])

		l, err := net.Listen("tcp", serverAddress)
		if err != nil {
			log.Fatal(err)
		}
		defer l.Close()

		fmt.Println("Listening on " + serverAddress)
		fmt.Println()

		for {
			fmt.Println("esperando una conexion...")

			conn, err := l.Accept()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("conexion recibida/aceptada")

			fmt.Println("leyendo...")

			reader := bufio.NewReader(conn)

			for {

				content, err := reader.ReadString('\n')

				if err != nil {
					if strings.Contains(err.Error(), "host") {
						break
					} else {
						log.Fatal(err)
					}
				}
				fmt.Println("se a leido!")
				fmt.Println()

				fmt.Println("Message received:", string(content))
			}
		}

	}
}
