package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/adammck/dynamixel/network"
	"github.com/adammck/dynamixel/servo/ax"

	"github.com/jacobsa/go-serial/serial"

	"github.com/ArT-Programming/dynamixelBridge/config"
	"github.com/hypebeast/go-osc/osc"
)

func main() {

	fmt.Println("### Welcome to go-osc receiver demo")
	fmt.Println("Press \"q\" to exit")
	//Load conf
	c := &config.Config{}

	_, err := c.GetConfig("default.yaml")

	if err != nil {
		fmt.Printf("Configuration error: %s\n", err)
		os.Exit(1)
	}

	options := serial.OpenOptions{
		PortName:              c.Serial,
		BaudRate:              1000000,
		DataBits:              8,
		StopBits:              1,
		MinimumReadSize:       0,
		InterCharacterTimeout: 100,
	}

	serial, err := serial.Open(options)
	if err != nil {
		fmt.Printf("Serial open error: %s\n", err)
		os.Exit(1)
	}

	network := network.New(serial)

	network.Flush()

	servo, err := ax.New(network, 1)

	if err != nil {
		fmt.Printf("servo init error: %s\n", err)
		os.Exit(1)
	}

	err = servo.Ping()
	if err != nil {
		fmt.Printf("ping error: %s\n", err)
		os.Exit(1)
	}

	addr := "127.0.0.1:8765"
	server := &osc.Server{}
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		fmt.Println("Couldn't listen: ", err)
	}
	defer conn.Close()

	go func() {
		fmt.Println("Start listening on", addr)

		for {
			packet, err := server.ReceivePacket(conn)
			if err != nil {
				fmt.Println("Server error: " + err.Error())
				os.Exit(1)
			}

			if packet != nil {
				switch packet.(type) {
				default:
					fmt.Println("Unknown packet type!")

				case *osc.Message:
					s := strings.Split(fmt.Sprint(packet.(*osc.Message)), ",")
					message := strings.Split(s[0], "/")
					fmt.Print(message[2])

					data := strings.Split(s[1], " ")

					for i := 1; i < len(data); i++ {
						if data[0][i-1] == 'f' {
							f, _ := strconv.ParseFloat(data[i], 32)
							fmt.Println(int(f * 1024))
							err = servo.SetGoalPosition(int(f * 1023))
							if err != nil {
								fmt.Printf("move error: %s\n", err)
								os.Exit(1)
							}

						}
						if data[0][i-1] == 'i' {
							i, _ := strconv.ParseInt(data[i], 10, 32)
							fmt.Println(i)

						}
						if data[0][i-1] == 'T' {
							fmt.Println("True")
						}
						if data[0][i-1] == 'F' {
							fmt.Println("False")
						}
					}

				}
			}
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		c, err := reader.ReadByte()
		if err != nil {
			os.Exit(0)
		}

		if c == 'q' {
			os.Exit(0)
		}
	}
}
