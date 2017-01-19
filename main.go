package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"fmt"

	"github.com/jinzhu/configor"
)

type lyricLine struct {
	Index          int    `json:"index"`
	Content        string `json:"content"`
	IsStructureGap bool   `json:"is_structure_gap"`
	SongID         string `json:"song_id"`
}

const lyricsFileName = "./lyrics.txt"

var config = struct {
	APIAddress string `json:"api_address"`
	APIKey     string `json:"api_key"`
	SongID     string `json:"song_id"`
}{}

func main() {
	configor.Load(&config, "config.json")
	file, _ := ioutil.ReadFile(lyricsFileName)
	lyricLinesStr := strings.Split(string(file), "\n")
	lyricLines := make([]*lyricLine, len(lyricLinesStr))
	for i, lyric := range lyricLinesStr {
		lyricLines[i] = &lyricLine{
			Index:          i,
			Content:        lyric,
			IsStructureGap: lyric == "",
			SongID:         config.SongID,
		}
	}

	url := fmt.Sprintf("%s/lyrics", config.APIAddress)
	for _, lyricLine := range lyricLines {
		json, _ := json.Marshal(lyricLine)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.APIKey))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		if resp.StatusCode == 201 {
			fmt.Println(fmt.Printf("http_%d", resp.StatusCode))
			continue
		}

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(fmt.Printf("http_%d :: %s", resp.StatusCode, string(body)))
	}
}
