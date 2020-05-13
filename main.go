package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"crypto/rand"
	"github.com/iotaledger/goshimmer/plugins/analysis/packet"

)

var additionRate = flag.Float64("nps", 10.0, "Amount of nodes to add each second")
var pattern = flag.String("pattern", "distribute", "Pattern for spamming")

// A heartbeat packet
type Packet struct {
	OwnID       []byte
	OutboundIDs [][]byte
	InboundIDs  [][]byte
}

var methods = map[string]func([]string, map[string][]string){
	"flood":         flood,
	"flood-reverse": floodReverse,
	"distribute":    distribute,
}

// Scenario ideas:
// 1. Distribute reporting times over 5 seconds evenly
// 2. Start sending heartbeats in reverse order (nodes)
// 3. Start sending in random order

func main() {
	flag.Parse()
	nodes := readNodes()
	links := readLinks()

	nodeCounter := 0
	for _, _ = range nodes {
		nodeCounter++
	}
	fmt.Println("Nodecount is: ", nodeCounter)

	linkCounter := 0
	for _, element := range links {
		for _, _ = range element {
			linkCounter++
		}
	}
	fmt.Println("Linkcount is: ", linkCounter)
	actualPattern := methods[*pattern]
	actualPattern(nodes, links)
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("I'm alive.")
		}
	}
}

// flood() floodes the server with heartbeats starting from the first node.
// Heartbeat messages are sent immediately after one another.
func flood(nodes []string, links map[string][]string) {
	for _, node := range nodes {
		own := []byte(node)

		out := make([][]byte, len(links[node]))
		for i, neighbor := range links[node] {
			out[i] = []byte(neighbor)
		}

		inSize := 0

		for _, link := range links {
			for _, neighbor := range link {
				if neighbor == node {
					inSize++
				}
			}
		}

		in := make([][]byte, inSize)
		inSize = 0
		for key, value := range links {
			for _, neighbor := range value {
				if neighbor == node {
					in[inSize] = []byte(key)
					inSize++
				}
			}
		}
		packet := &packet.Heartbeat{OwnID: own, OutboundIDs: out, InboundIDs: in}

		go sendPacket(packet)
	}
}

// floodReverse() floodes the server with heartbeats starting from the last node.
// Heartbeat messages are sent immediately after one another.
func floodReverse(nodes []string, links map[string][]string) {
	for i := len(nodes) - 1; i > 0; i-- {
		node := nodes[i]
		own := []byte(node)

		out := make([][]byte, len(links[node]))
		for i, neighbor := range links[node] {
			out[i] = []byte(neighbor)
		}

		inSize := 0

		for _, link := range links {
			for _, neighbor := range link {
				if neighbor == node {
					inSize++
				}
			}
		}

		in := make([][]byte, inSize)
		inSize = 0
		for key, value := range links {
			for _, neighbor := range value {
				if neighbor == node {
					in[inSize] = []byte(key)
					inSize++
				}
			}
		}
		packet := &packet.Heartbeat{OwnID: own, OutboundIDs: out, InboundIDs: in}

		go sendPacket(packet)
	}
}

func distribute(nodes []string, links map[string][]string) {
	delay := time.Duration(float64(time.Second) / (*additionRate))
	fmt.Println("Delay is: ", delay)
	fmt.Println("Node addition rate: ", *additionRate, " node(s)/sec")
	for i, node := range nodes {
		own := []byte(node)

		out := make([][]byte, len(links[node]))
		for i, neighbor := range links[node] {
			out[i] = []byte(neighbor)
		}

		inSize := 0

		for _, link := range links {
			for _, neighbor := range link {
				if neighbor == node {
					inSize++
				}
			}
		}

		in := make([][]byte, inSize)
		inSize = 0
		for key, value := range links {
			for _, neighbor := range value {
				if neighbor == node {
					in[inSize] = []byte(key)
					inSize++
				}
			}
		}
		packet := &packet.Heartbeat{OwnID: own, OutboundIDs: out, InboundIDs: in}

		go sendPacket(packet)
		if i > len(nodes)/5 {
			time.Sleep(time.Duration(delay))
		}
	}
}

func sendPacket(p *packet.Heartbeat) {
	ticker := time.NewTicker(5 * time.Second)
	conn, err := net.Dial("tcp", "0.0.0.0:16178")
	//i := 0
	if err != nil {
		fmt.Println(err.Error())
	}
	for {
		select {
		case <-ticker.C:
			data, err := packet.NewHeartbeatMessage(p)
			if err != nil {
				fmt.Println(err)
			}
			_, _ = conn.Write(data)
		}
	}
	// Deliberately keep the connection open until timeout on server (10s)
	conn.Close()

}

func GenRandomBytes(size int) (blk []byte, err error) {
	blk = make([]byte, size)
	_, err = rand.Read(blk)
	return
}

var nodes = []string{}
var links = make(map[string][]string)

func readNodes() []string {
	file, err := os.Open("generated_nodes.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nodes = append(nodes, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return nodes
}

func readLinks() map[string][]string {
	file, err := os.Open("generated_links.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		source := line[:32]
		target := line[32:]
		if len(links[source]) < 4 {
			links[source] = append(links[source], target)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}