package main

import (
  "fmt"
)

func askOrder() (string) {
  var input string
  fmt.Print("What would you like to eat: ")
  // Get the input from the user
  fmt.Scan(&input)
  return input
}

// WRITE CONTAINS FUNCTION BELOW
func contains(menu []string, order string) bool{
 for _,value:= range menu{
  if value == order{
    return true
  }
 }
 return false
}

func main() {
  fastfoodMenu := []string{"Burgers,", "Nuggets,", "Fries"}
  fmt.Println("The fast food menu has these items:", fastfoodMenu)
 
  var total int
  var order string
  // WRITE INDEFINITE LOOP ASKING FOR ORDERS BELOW
  for{
    order = askOrder()
    if order == "quit"{
      break
    }

	if contains(fastfoodMenu, order){
		total +=4;
	} else{
		fmt.Println("Your order is not on the menu")
	}
  }

  // WRITE DEFINITE LOOP COUNTING $2 BILLS BELOW
  for amount :=total; amount > 0; amount -=2 {
    fmt.Println("That's $2 bill, your remaining amount is", amount)
  }


  fmt.Println("The total for the order is", total)
}
