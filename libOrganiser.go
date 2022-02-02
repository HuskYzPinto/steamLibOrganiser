package main

import (
	"github.com/HuskYzPinto/steamLibOrganiser/database"
)

func main() {
	rawList := database.GetGameList()
	database.ParseGameList(rawList)
	//
	//fmt.Printf("%+v", parsedList)
	//fmt.Printf("%T", parsedList)
}
