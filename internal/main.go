package main

import (
  "context"
  "log"

  "github.com/devvildaz/ScrapSubtitlesYT/internal/video"
  "github.com/devvildaz/ScrapSubtitlesYT/internal/utils"
  "github.com/chromedp/chromedp"
)

const (
  YOUTUBE_ID="B_yitbh-XVk"
)

func main() {
  ctx, cancel := chromedp.NewContext(context.Background())
  defer cancel() 
  
  err := video.NavigateToVideo(ctx, YOUTUBE_ID)

  if err != nil {
    return 
  }

  description, err := video.ExtractVideoDescription(ctx)

  if err != nil {
    return 
  }

  log.Println(YOUTUBE_ID + " Description")
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
  
  err = utils.StoreMapAsJSONFile(response, YOUTUBE_ID)
  if err != nil {
    return 
  }
}


