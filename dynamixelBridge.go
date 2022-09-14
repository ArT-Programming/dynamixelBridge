package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/adammck/dynamixel/network"
	"github.com/adammck/dynamixel/servo"
	"github.com/adammck/dynamixel/servo/ax"

	"github.com/jacobsa/go-serial/serial"

	"github.com/ArT-Programming/dynamixelBridge/config"
	"github.com/hypebeast/go-osc/osc"
)

type servoax interface {
	SetGoalPosition(int) error
}

func main() {

	fmt.Println("Starting dynamixelBridge")
	fmt.Println("Press \"q\" [Enter] to exit")
	//Load conf
	c := &config.Config{}

	_, err := c.GetConfig("default.yaml")

	if err != nil {
		fmt.Printf("Configuration error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Setting Status return Level: %b\n", c.StatusReturnLevel)
	fmt.Printf("Using serial port: %s\n", c.Serial)

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

	servos := make(map[string]*servo.Servo)

	for _, s := range c.Servos {
		fmt.Printf("Adding %s\n", s.Path)
		newServo, err := ax.New(network, s.ServoID)
		newServo.SetReturnLevel(c.StatusReturnLevel)
		if err != nil {
			fmt.Printf("Servo init error: %s\n", err)
			os.Exit(1)
		}
		err = newServo.Ping()
		if err != nil {
			fmt.Printf("Ping error: %s\n", err)
		}
		path := strings.Replace(s.Path, "/", "", 1)
		servos[path] = newServo
	}

	addr := "0.0.0.0:8765"
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

					data := strings.Split(s[1], " ")

					for i := 1; i < len(data); i++ {
						if data[0][i-1] == 'f' {
							f, _ := strconv.ParseFloat(data[i], 32)
							se := servos[strings.ReplaceAll(message[2], " ", "")]
							if se != nil {

								err := se.SetGoalPosition(int(f * 1023))
								if err != nil {
									fmt.Printf("move error: %s\n", err)
									os.Exit(1)

								}
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
			fmt.Printf("Setting Status return Level: 2")
			for _, s := range servos {
				s.SetReturnLevel(2)
			}
			os.Exit(0)
		}
	}
}
