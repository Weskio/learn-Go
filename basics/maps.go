package main

import "fmt"  

func main(){
	//fmt.Println("Hello world")

    score := make( map[string] int)
	score["Cho"] = 45
	score["Badassou"] = 3
	score["Togbe"]= 23

	//delete(score, "Badassou")
	//fmt.Println(score)

	for k,v := range score{
		if v >= 45{
			fmt.Printf("%v scored %v\n", k, v)
		}
		
	}

	//fmt.Println(score["Cho"])
}
