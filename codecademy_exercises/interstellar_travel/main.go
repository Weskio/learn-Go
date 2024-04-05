package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int){
  fmt.Printf("The fuel left in the tank is %d liters", fuel)
}


// Create the function calculateFuel() here
func calculateFuel(planet string) int{
 var fuel int

 switch planet{
  case "Venus":
  fuel = 300000
  case "Mercury":
  fuel = 500000
  case "Mars":
  fuel = 700000
  default:
  fuel = 0
 }

 return fuel
}


// Create the function greetPlanet() here
func greetPlanet(planet string){
  fmt.Println("Hello passengers, were headed to", planet, "! Enjoy the trip")
}


// Create the function cantFly() here
func cantFly(){
  fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int{
 var fuelRemaining, fuelCost int

 fuelRemaining = fuel
 fuelCost = calculateFuel("Mars")

 if fuelRemaining >= fuelCost{
  greetPlanet("Mars")
  fuelCost -= fuelRemaining
 } else {
  cantFly()
 }

 return fuelRemaining
}


func main() {
  // Test your functions!
  fuel := 1000000 
  planetChoice:= "Mars"

  flyToPlanet(planetChoice, fuel)

  fuelGauge(calculateFuel(planetChoice))

  // fuelGauge(95)
  
  // fmt.Println()

  // fmt.Printf("The fuel needed for this planet is %d liters",calculateFuel("Mars"))


  
  // Create `planetChoice` and `fuel`
  
  // And then liftoff!
  
}