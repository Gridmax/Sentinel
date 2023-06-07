package main

import (
  "log"

  "github.com/Gridmax/Sentinel/commun/client"	

)

func main() {
  log.Println("- - - - - - - - - - - - - - -")
  client.Start("config.yaml")
}
