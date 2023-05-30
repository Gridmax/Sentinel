package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
  "strconv"

  "github.com/Gridmax/Sentinel/utility/timeconvert"
  "github.com/Gridmax/Sentinel/utility/configload"
)

//type Conf struct {
//  ServerAddress string  `yaml:"server_address"`
//  ServerPort int  `yaml:"server_port"`
//  AgentMode string `yaml:"agent_mode"`
//  AgentPort int `yaml:"agent_port"`
//  HostName string `yaml:"host_name"`
//  HostGroup string `yaml:"host_group"`
//  AgentInterval string `yaml:"agent_interval"`
//}


func main() {
	// Start the client
	startClient()
}


func startClient() {

  config, err := configload.loadConfig("config.yaml")
  if err != nil {
    fmt.Println("Failed to load config: ", err)
    return
  }
	for {

    remoteServer := config.ServerAddress + ":" + strconv.Itoa(config.ServerPort)
    
    conn, err := net.Dial("tcp", remoteServer)
	  
    if err != nil {
			fmt.Println("Failed to connect to server:", err)
			return
		}
		defer conn.Close()

		// Create the payload header and message
		header := "HEADER2"
		message := "Payload message"

		// Encode the header and message into binary format
		headerBytes := []byte(header)
		messageBytes := []byte(message)

		// Calculate the size of the header and message
		headerSize := uint16(len(headerBytes))
		messageSize := uint16(len(messageBytes))

		// Create a byte slice to hold the binary data
		buffer := make([]byte, 4+len(headerBytes)+len(messageBytes))

		// Encode the header size and write it to the buffer
		binary.BigEndian.PutUint16(buffer[:2], headerSize)

		// Encode the message size and write it to the buffer
		binary.BigEndian.PutUint16(buffer[2:4], messageSize)

		// Write the header to the buffer
		copy(buffer[4:4+len(headerBytes)], headerBytes)

		// Write the message to the buffer
		copy(buffer[4+len(headerBytes):], messageBytes)

		// Send the binary data to the server
		_, err = conn.Write(buffer)
		if err != nil {
			fmt.Println("Failed to send data:", err)
			return
		}

		fmt.Println("Message sent to server")


    interval := timeconvert.GetInterval(config.AgentInterval)

    //interval := timeconvert.GetInterval(config.AgentInterval) * time.Second
		// Wait for the specified interval before sending the next message
    time.Sleep(time.Duration(interval) * time.Second)
    
  }
}

