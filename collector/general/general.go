package general

import (
  "fmt"
  "os"
  "strconv"
  "strings"
  "time"

  "github.com/Gridmax/Sentinel-utility/sysinfo/cpu"
  "github.com/Gridmax/Sentinel-utility/sysinfo/ram"
  "github.com/Gridmax/Sentinel-utility/sysinfo/uptime"
)


func CpuInfo() string {
  cpu, err := cpu.Get()
  if err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err)
  }
//  fmt.Println(cpu)
  cpu_total := strconv.FormatUint(cpu.Total, 10)
  cpu_idle := strconv.FormatUint(cpu.Idle, 10)
  cpu_usage := strconv.FormatUint(cpu.System + cpu.User, 10)
//  fmt.Println(cpu_idel)
  cpu_info := []string{"cpu_total", cpu_total, "cpu_idel", cpu_idle, 
    "cpu_usage", cpu_usage}
  cpu_infos := strings.Join(cpu_info, ":")
  cpu_infos += ":"
//  fmt.Printf(strings.Join(cpu_info, ":"))
  return cpu_infos
}

func RamInfo() string {
  ram, err := ram.Get()
  if err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err)
  }
  ram_total := strconv.FormatUint(ram.Total, 10)
  ram_idle := strconv.FormatUint(ram.Free, 10)
  ram_usage := strconv.FormatUint(ram.Used, 10)

  ram_info := []string{"ram_total", ram_total, "ram_idle", 
  ram_idle, "ram_usage", ram_usage}
  ram_infos := strings.Join(ram_info,":")
  ram_infos += ":"

  return ram_infos
}

func UpInfo() string {
  uptime, err := uptime.Get()
  if err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err)
  }
//  fmt.Println(uptime)
  uptimeSecond := uptime.Seconds()
//  fmt.Println(uptimeSecond)

//  up := time.ParseDuration(uptime)
//  fmt.Println(up.Seconds())
  up := "uptime:"
  up += fmt.Sprintf("%f", uptimeSecond)
  up += ":"
  return up
}

func LogTime() string {
  now := time.Now()
  timestamp := now.Unix()
  timeNow := "timestamp:"
  timeNow += strconv.FormatInt(timestamp, 10)
  timeNow += ":"
  return timeNow

}

func GeneralInfo(name string, group string) string{
  var generalInfo string
  generalInfo += "host_name:" + name + ":host_group:" + group + ":" 
  generalInfo += LogTime()
  generalInfo += UpInfo()
  generalInfo += CpuInfo()
  generalInfo += RamInfo()
  return generalInfo
}
//func GeneralInfo() map[string] {
//  cpus, err := cpu.Get()
//  if err != nil{
//    fmt.Fprintf(os.Stderr, "%s\n", err)
//  }
//  cpu_total := strconv.FormatUint(cpus.Total, 10)
//  cpu_idel := strconv.FormatUint(cpus.User, 10)
 
//  info := map[string]string{
//    "cpu_total":cpu_total,
//    "cpu_idel":cpu_idel,
//  }
//  return info
//}
