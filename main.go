package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"encoding/json"

	"github.com/jinzhu/configor"
)

type lyricLine struct {
	Index          int    `json:"index"`
	Content        string `json:"content"`
	IsStructureGap bool   `json:"is_structure_gap"`
	SongID         string `json:"song_id"`
}

var config = struct {
	APIAddress string `json:"api_address"`
	APIKey     string `json:"api_key"`
	SongID     string `json:"song_id"`
}{}

const lyrics = `He said the way my blue eyes shined
Put those Georgia stars to shame that night
I said: \"That's a lie.\"
Just a boy in a Chevy truck
That had a tendency of gettin' stuck
On backroads at night
And I was right there beside him all summer long
And then the time we woke up to find that summer gone

But when you think Tim McGraw
I hope you think my favorite song
The one we danced to all night long
The moon like a spotlight on the lake
When you think happiness
I hope you think that little black dress
Think of my head on your chest
And my old faded blue jeans
When you think Tim McGraw

I hope you think of me

September saw a month of tears
And thankin' God that you weren't here
To see me like that
But in a box beneath my bed
Is a letter that you never read
From three summers back

It's hard not to find it all a little bittersweet
And lookin' back on all of that, it's nice to believe

When you think Tim McGraw
I hope you think my favorite song
The one we danced to all night long
The moon like a spotlight on the lake
When you think happiness
I hope you think that little black dress
Think of my head on your chest

And my old faded blue jeans
When you think Tim McGraw
I hope you think of me

And I'm back for the first time since then
I'm standin' on your street
And there's a letter left on your doorstep
And the first thing that you'll read is:

When you think Tim McGraw
I hope you think my favorite song
Someday you'll turn your radio on
I hope it takes you back to that place
When you think happiness

I hope you think that little black dress
Think of my head on your chest

And my old faded blue jeans
When you think Tim McGraw
I hope you think of me
Oh, think of me
Mmmm...
He said the way my blue eyes shine
Put those Georgia stars to shame that night
I said: \"That's a lie\"`

func main() {
	configor.Load(&config, "config.json")

	lyricLinesStr := strings.Split(lyrics, "\n")
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

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(fmt.Printf("http_%d :: %s", resp.StatusCode, string(body)))
	}
}
