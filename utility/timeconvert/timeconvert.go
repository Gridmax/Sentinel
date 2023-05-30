
package timeconvert

import (
  "fmt"
	"time"
)


//func toInterval(timeConfig string) float64{
//  get := timeConvert(timeConfig)
//  interval := get * 60
//  return interval 
//}

func GetInterval(timeConfig string) int{
  get, _ := time.ParseDuration(timeConfig)
  convertSecond := get.Seconds()
  fmt.Println(int(convertSecond))
  return int(convertSecond)

}

