package youtube

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var (
	query     = flag.String("query", "pubg live streaming", "Search string")
	maxResult = flag.Int64("max-results", 6, "Max Youtube results")
)

type YouTube struct {
}

type result struct {
	VideoTitle string `json:"title"`
	VideoId    string `json:"video_link"`
	Thumbnail  string `json:"thumbnail_link"`
}

var results []result

func apiKey() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env file loading error - ", err)
		return ""
	}
	ytAPI := os.Getenv("YOUTUBE_API_KEY")
	return ytAPI
}

func (yt YouTube) GetYtData() interface{} {

	flag.Parse()
	ctx := context.Background()
	developerKey := apiKey()
	service, err := youtube.NewService(ctx, option.WithAPIKey(developerKey))

	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	var temp []string
	var thumbnail, title []string
	temp = append(temp, "id", "snippet")
	//Make the API to YouTube
	call := service.Search.List(temp).Q(*query).MaxResults(*maxResult)
	response, err := call.Do()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	videos := make(map[string]string)

	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
			title = append(title, item.Snippet.Title)
			thumbnail = append(thumbnail, item.Snippet.Thumbnails.Default.Url)
		default:
			continue
		}
	}
	var res result
	var count int = -1
	for i := range videos {
		count = count + 1
		res.VideoTitle = title[count]
		res.VideoId = fmt.Sprint("https://www.youtube.com/watch?v=" + i)
		res.Thumbnail = thumbnail[count]
		results = append(results, res)
	}
	return results
}
