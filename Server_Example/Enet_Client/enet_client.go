package main

import (
	"log"
	"github.com/codecat/go-enet"
)

func main() {
	// Initialize enet
	enet.Initialize()
	defer enet.Deinitialize() // Ensure ENet is cleaned up on exit

	// Create a client host
	client, err := enet.NewHost(nil, 1, 1, 0, 0)
	if err != nil {
		log.Fatalf("Couldn't create host: %s", err.Error())
	}
	defer client.Destroy() // Ensure the client host is destroyed on exit

	// Connect the client host to the server
	peer, err := client.Connect(enet.NewAddress("127.0.0.1", 1234), 1, 0)
	if err != nil {
		log.Fatalf("Couldn't connect: %s", err.Error())
	}

	// The event loop
	for {
		// Wait until the next event
		ev := client.Service(1000)

		// Send a ping if no event occurs
		if ev.GetType() == enet.EventNone {
			err := peer.SendString("ping", 0, enet.PacketFlagReliable)
			if err != nil {
				log.Printf("Failed to send ping: %s", err.Error())
			}
			continue
		}

		switch ev.GetType() {
		case enet.EventConnect: // We connected to the server
			log.Println("Connected to the server!")

		case enet.EventDisconnect: // We disconnected from the server
			log.Println("Lost connection to the server!")

		case enet.EventReceive: // The server sent us data
			packet := ev.GetPacket()
			data := string(packet.GetData()) // Convert byte slice to string
			
			if data == "pong" {
				log.Println("Received Pong")
			}else{
				log.Printf("Received %d bytes from server: %s", len(packet.GetData()), data)
			}
			
			packet.Destroy() // Clean up the packet
		}
	}
}
