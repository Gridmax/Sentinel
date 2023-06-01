package client

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
  "strconv"
  "log"

  "github.com/Gridmax/Sentinel/utility/timeconvert"
  "github.com/Gridmax/Sentinel/utility/configload"
  "github.com/Gridmax/Sentinel/collector/general"
  "github.com/Gridmax/Sentinel/utility/osde"
)

func Agent(configFile string) {

  config, err := configload.LoadConfig(configFile)
  if err != nil {
    fmt.Println("Failed to load config: ", err)
    return
  }
  for {
    remoteServer := config.ServerAddress + ":" + strconv.Itoa(config.ServerPort)
    conn, err := net.Dial("tcp", remoteServer)
	  if err != nil{ 
      log.Println("1001 err,",err)
		  //fmt.Println("Failed to connect to server:", err)
      return
		}
		defer conn.Close()

		// Create the payload header and message
    header := osde.DetectOS()
		//message := "Payload message"

    message := general.GeneralInfo(config.HostName, config.HostGroup)
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

    log.Println("Message sent successfully, with size", headerSize + messageSize)
//		fmt.Println("Message sent to server")


    interval := timeconvert.GetInterval(config.AgentInterval)

    //interval := timeconvert.GetInterval(config.AgentInterval) * time.Second
		// Wait for the specified interval before sending the next message
    time.Sleep(time.Duration(interval) * time.Second)
  }
}

func Start(configFile string) {
  log.Println("Starting Sentinel Agent")
  log.Println("- - - - - - - - - - - - - - -")
  config, err := configload.LoadConfig(configFile)
  if err != nil {
    fmt.Println("Failed to load config: ", err)
    return
  }
  log.Println("Sentinel Agent successfully started")
  interval := timeconvert.GetInterval(config.AgentInterval)
  for i := 0; i < config.AgentRetry; i ++ {
    if i > 0 {
      log.Println("Failed to connect to Hillock server, retrying the ", i, " times")
    }

    //interval := timeconvert.GetInterval(config.AgentInterval)
    time.Sleep(time.Duration(interval) * time.Second)

    Agent(configFile)

  }
}
