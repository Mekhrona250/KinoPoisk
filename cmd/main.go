package main

import (
	"fmt"
	"kinopoisk/pkg"
)

func main() {
	pkg.FillCinemas()
	pkg.FillCharacters()

	fmt.Printf("%+v \n", pkg.Cinemas)

	fmt.Printf("%+v \n", pkg.Characters)

	result := pkg.GetGoodCharacters()

	for _, val := range result{
		fmt.Printf("Name %v IsAlive %v Cinema name %v \n",val.Name , val.IsAlive, val.Cinema.Name)

	
	}
	

	// name, cinema.name, is_alive

}
