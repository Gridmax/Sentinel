package errck

import (
  "log"
  "strings"
)

func ErrCheck(message string) {
  if strings.Contains(message, "address already in use"){
    log.Println("1001 err,",message)
  }else if strings.Contains(message, "connection refused"){
    log.Println("1002 err,",message)
  }else if strings.Contains(message, "broken pipe"){
    log.Println("1003 err,", message)
  }else {
    log.Println("1099 err, error unknown", message)
  }
}
