package osde

import (
  "fmt"
  "runtime"
)

func DetectOS() string {
  os := runtime.GOOS
  os += ":"
  fmt.Println(os)
  return os
}
