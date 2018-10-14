//methods.go

package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

// Structs
type car struct {
  gas_pedal uint16
  brake_pedal uint16
  steering_wheel int16
  top_speed_mph float64
}

func (c car) kmh() float64{
  return float64(c.gas_pedal) * (c.top_speed_mph/usixteenbitmax)
}

func main(){
    a_car := car{ gas_pedal: 22341,
                  brake_pedal:0,
                  steering_wheel: 12561,
                  top_speed_mph: 225.0 }

   fmt.Println(a_car.gas_pedal)
   fmt.Println(a_car.kmh())
}
