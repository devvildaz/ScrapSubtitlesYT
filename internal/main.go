package main

import (
  "context"
  "log"

  "flag"

  "github.com/devvildaz/ScrapSubtitlesYT/internal/video"
  "github.com/devvildaz/ScrapSubtitlesYT/internal/utils"
  "github.com/chromedp/chromedp"
)


func main() {
  var youtube_id string

  flag.StringVar(&youtube_id, "youtube_id", "", "Specify a youtube video's ID")
  flag.Parse()
  if youtube_id == "" {
    log.Fatal("Please enter a ID of a youtube video that you wanna analyze")
  }

  ctx, cancel := chromedp.NewContext(context.Background())
  defer cancel() 
  
  err := video.NavigateToVideo(ctx, youtube_id)

  if err != nil {
    return 
  }

  description, err := video.ExtractVideoDescription(ctx)

  if err != nil {
    return 
  }

  log.Println(youtube_id + " Description")
  log.Println(description)

  err = video.OpenVideoTranscript(ctx)
  if err != nil {
    return
  }

  items, timestamps, err := video.ExtractionRawSubtitles(ctx)
  if err != nil {
    return 
  }

  response, err := video.FormatRawSubtitles(items, timestamps)
  if err != nil {
    return 
  }
  
  err = utils.StoreMapAsJSONFile(response, youtube_id)
  if err != nil {
    return 
  }
}


