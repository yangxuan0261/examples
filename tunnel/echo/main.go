package main

import (
	"os"
	"bufio"
	"fmt"
	"flag"

	"github.com/micro/go-micro/tunnel"
	"github.com/micro/go-micro/transport"
)

var (
	address = flag.String("address", ":10001", "tunnel address")
	nodes = flag.String("nodes", "", "tunnel nodes")
	channel = flag.String("channel", "default", "the channel")
)

func readLoop(c transport.Socket) {
	for {
		m := new(transport.Message)
		if err := c.Recv(m); err != nil {
			return
		}
		fmt.Println(string(m.Body))
	}
}

func writeLoop(c transport.Socket, ch chan []byte) {
	for {
		b := <-ch
		if err := c.Send(&transport.Message{
			Body: b,
		}); err != nil {
			return
		}
	}
}

func accept(l tunnel.Listener, ch chan []byte) {
	connCh := make(chan chan []byte)

	go func() {
		var conns []chan []byte

		for {
			select {
			case b := <-ch:
				// send message to all conns
				for _, c := range conns {
					select {
					case c <- b:
					default:
					}
				}
			case c := <-connCh:
				conns = append(conns, c)
			}
		}
	}()


	for {
		c, err := l.Accept()
		if err != nil {
			return
		}

		fmt.Println("Accepting new connection")

		// pass to reader
		rch := make(chan []byte)
		connCh <- rch

		// print out what we read
		go readLoop(c)
		go writeLoop(c, rch)
	}
}

func main() {
	flag.Parse()

	// create a tunnel
	tun := tunnel.NewTunnel(
		tunnel.Address(*address),
		tunnel.Nodes(*nodes),
	)
	if err := tun.Connect(); err != nil {
		fmt.Println(err)
		return
	}
	defer tun.Close()

	l, err := tun.Listen(*channel)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	ch := make(chan []byte)

	go accept(l, ch)

	c, err := tun.Dial(*channel)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	// read and print what we get back
	go readLoop(c)

	// read input and send over the tunnel
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		buf := scanner.Bytes()

		// send over the dialled conn
		if err := c.Send(&transport.Message{
			Body: buf,
		}); err != nil {
			fmt.Println(err)
		}

		// send to all listeners
		ch <- buf
	}
}
