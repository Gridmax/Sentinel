
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

func Get(timeConfig string) float64{
  get, _ := time.ParseDuration(timeConfig)
  convertSecond := get.Seconds()
  return convertSecond
}

