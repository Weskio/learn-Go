package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true
  eludedGuards := rand.Intn(100)

  if eludedGuards >= 50{
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else{
    isHeistOn = false
    fmt.Println("Plan a better disguise next time!")
  }

  openedVault := rand.Intn(100)
  fmt.Println( isHeistOn )

  if isHeistOn && openedVault >=70{
    fmt.Println("Grab and GO!")
  } else if isHeistOn{
    isHeistOn = false
    fmt.Println("The vault can't be opened might as well call the heist off")
  }

  leftSafely := rand.Intn(5)

  if isHeistOn{
    switch leftSafely {
      case 0:
      isHeistOn = false
      fmt.Println("You failed to steal")
      case 1:
      isHeistOn = false
      fmt.Println("Turns out vault doors don't open from the inside")
       case 2:
       isHeistOn = false
      fmt.Println("Fight broke out between the squad")
       case 3:
       isHeistOn = false
      fmt.Println("Retreat! Kwame killed a civilian")
      default:
      fmt.Println("Start the getaway car now!")
    }
   
   if isHeistOn{
    amtStolen  := rand.Intn(1000000) + 10000
    fmt.Println (amtStolen, "was the amount stolen")
   }
    
  }
}
