package database

import (
	"fmt"
	"io"
	"net/http"
)

func GetGameList() io.ReadCloser {
	steamDbUrl := "https://api.steampowered.com/ISteamApps/GetAppList/v2/"
	resp, _ := http.Get(steamDbUrl)
	return resp.Body
}

func ParseGameList(gameList io.ReadCloser) {
	result, _ := io.ReadAll(gameList)
	fmt.Printf("%+v", string(result))
}

func SaveGameList(gameStringList []string) {

}
