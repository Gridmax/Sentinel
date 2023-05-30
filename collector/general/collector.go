package general

import (
  "fmt"
  "os"
  "strconv"

  "github.com/Gridmax/Sentinel-utility/sysinfo/cpu"
)


func CpuInfo() string {
  cpu, err := cpu.Get()
  if err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err)
  }
  total := strconv.FormatUint(cpu.Total, 10)
  return total
}
