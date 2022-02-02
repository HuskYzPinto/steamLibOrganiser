package database

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetGameList() io.ReadCloser {
	steamDbUrl := "https://api.steampowered.com/ISteamApps/GetAppList/v2/"
	resp, _ := http.Get(steamDbUrl)
	return resp.Body
}

func ParseGameList(gameList io.ReadCloser) string {
	result, _ := io.ReadAll(gameList)
	var out bytes.Buffer
	err := json.Indent(&out, result, "", "  ")
	if err != nil {
		return ""
	}
	fmt.Printf("%+v", out.String())
	return out.String()
}

//
//func SaveGameList(gameList string) {
//
//
//}
