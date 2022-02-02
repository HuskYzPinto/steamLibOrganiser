package main

import (
	"github.com/HuskYzPinto/steamLibOrganiser/database"
)

func main() {
	rawList := database.GetGameList()
	database.ParseGameList(rawList)

	//database.SaveGameList(parsedList)
}
