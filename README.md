import-lyrics
===

Simple tool to import a batch of lyrics into the [baelor-api](https://github.com/baelorswift/api).

#### Getting Started

- Create a file called `config.json` in the root of the repo in the following format
``` json
{
  "api_address": "http://<the address of the api>/v1",
  "api_key": "<your baelor api key>",
  "song_id": "<the id of the song you want to add lyrics to>"
}
```
- Update the `lyrics` const inside `main.go` with the lyrics list (remember to escape the `"` chars).
- Run `go run main.go`

![;)](https://chandeww.files.wordpress.com/2015/02/whatgif.gif)
