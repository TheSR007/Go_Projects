package main

import (
	"log"
	"github.com/codecat/go-enet"
)

func SetupEnetServer() {
	// Initialize enet
	enet.Initialize()
	defer enet.Deinitialize() // Ensures that enet is deinitialized when the program exits

	// Create a host listening on 0.0.0.0:1234
	host, err := enet.NewHost(enet.NewListenAddress(1234), 32, 1, 0, 0)
	if err != nil {
		log.Fatalf("Couldn't create host: %s", err.Error())
	}
	defer host.Destroy()

	log.Println("ENet server started on port 1234")

	// The event loop
	for {
		// Wait until the next event
		ev := host.Service(1000)

		// Handle event types
		switch ev.GetType() {
		case enet.EventConnect: // A new peer has connected
			log.Printf("New peer connected: %s", ev.GetPeer().GetAddress())

		case enet.EventDisconnect: // A connected peer has disconnected
			log.Printf("Peer disconnected: %s", ev.GetPeer().GetAddress())

		case enet.EventReceive: // A peer sent us some data
			// Get the packet
			packet := ev.GetPacket()

			// Process packet data
			packetBytes := packet.GetData()
			packet.Destroy() // Explicitly destroy the packet

			// Respond "pong" to "ping"
			if string(packetBytes) == "ping" {
				log.Println("Received ping")
				ev.GetPeer().SendString("pong", ev.GetChannelID(), enet.PacketFlagReliable)
				continue
			}

			// Disconnect the peer if they say "bye"
			if string(packetBytes) == "bye" {
				log.Println("Peer said bye, disconnecting...")
				ev.GetPeer().Disconnect(0)
				continue
			}
		}
	}
}
