package video 

import (
  "log"
  "context"
  "errors"
  "github.com/chromedp/chromedp"
  "github.com/devvildaz/ScrapSubtitlesYT/internal/utils"
  "github.com/chromedp/cdproto/cdp"
  "strings"
)

type VideoSubtitle struct{
  TimeString        string
  TimeInSeconds     int
  Text              string
}

const (
  YOUTUBE_WATCH_URL = "https://www.youtube.com/watch?v="
)

func ExtractSubtitleByVideoID (id string) (map[string]VideoSubtitle,error){
  return nil, nil
}

func NavigateToVideo(ctx context.Context, id string) (error) {
  log.Println("Navigate process...")
 	if err := chromedp.Run(ctx,chromedp.Navigate(YOUTUBE_WATCH_URL + id)); err != nil {
    log.Fatal("Could not navigate to youtube: ", err)
    return err
  }
  return nil
}

func ExtractVideoDescription(ctx context.Context) (string, error) {
  var res string
  sel := `//tp-yt-paper-button[@id="expand" and contains(@class, "style-scope") and contains(@class, "ytd-text-inline-expander")]`

  if err := chromedp.Run(ctx, 
    utils.RunWithTimeOut(&ctx, 2, chromedp.Tasks{ chromedp.WaitVisible(sel), chromedp.Click(sel) }),
  ); err != nil {
    log.Fatal("Could not get description snippet section: ", err)
    return "" ,err
  }
  sel = `//*[@id="description"]//yt-attributed-string[contains(@class, "style-scope") and contains(@class, "ytd-text-inline-expander") and not(@id)]`
  
  if err := chromedp.Run(
      ctx,chromedp.Text(sel, &res, chromedp.NodeVisible),
  ); err != nil {
    log.Fatal("Could not get description text:", err)
    return "", err
  }
  
  return strings.TrimSpace(res), nil
}

func OpenVideoTranscript(ctx context.Context) (error) {
  selMoreOpts := `document.querySelector("#button-shape > button")`
  log.Println("Waiting for the element ", selMoreOpts)
  if err := chromedp.Run(ctx, 
    utils.RunWithTimeOut(&ctx, 2, chromedp.Tasks{ 
      chromedp.WaitVisible(selMoreOpts, chromedp.ByJSPath), 
    })); err != nil {
    log.Fatal("Could not wait the options button: ", err)
    return err 
  }
  selShowTranscript := `//ytd-menu-service-item-renderer[.//yt-formatted-string[contains(., 'Show transcript')]]`
  selTranscriptFrame := `document.querySelector("#segments-container")`
  log.Println("Executing opening of the transcript frame")
  if err := chromedp.Run(ctx, 
    utils.RunWithTimeOut(&ctx, 10, chromedp.Tasks{ 
      chromedp.Click(selMoreOpts, chromedp.ByJSPath), 
      chromedp.WaitVisible(selShowTranscript, chromedp.BySearch),
      chromedp.Click(selShowTranscript, chromedp.BySearch),
      chromedp.WaitVisible(selTranscriptFrame, chromedp.ByJSPath), 
    })); err != nil {
    log.Fatal("Could not open the transcript of the video: ", err)
    return err 
  }
  return nil
}

func ExtractionRawSubtitles(ctx context.Context) ([]*cdp.Node, []*cdp.Node, error) {
  log.Println("Executing process to extract transcript items")

  const ytdTranscriptXpath = `//*[@id="segments-container"]//ytd-transcript-segment-renderer`

  var transcriptTimestamps []*cdp.Node
  var transcriptItems []*cdp.Node

  selTranscriptTextNode := ytdTranscriptXpath + `//yt-formatted-string/text()`
  selTranscriptTimestampNode := ytdTranscriptXpath + `//*[contains(@class,'segment-timestamp')]/text()`

  if err := chromedp.Run(ctx, chromedp.Nodes(selTranscriptTextNode , &transcriptItems, chromedp.BySearch)); err != nil {
    log.Fatal("Could not retrieve the render items from the transcript: ", err)
    return nil, nil, err
  }

  if err := chromedp.Run(ctx, chromedp.Nodes(selTranscriptTimestampNode, &transcriptTimestamps, chromedp.BySearch)); err != nil {
    log.Fatal("Could not retrieve the render stamps from the transcript: ", err)
    return nil, nil, err
  }
  log.Println("Printing extracted transcript items")
  log.Println("Length of Timestamps: ", len(transcriptTimestamps))
  log.Println("Length of Items: ", len(transcriptItems))
  
  if len(transcriptItems) == len(transcriptTimestamps) {
    log.Println("Timestamps and Items have the same length, there won't be issues with the merging")
  } else {
    log.Fatal("Don't have the same length")
    return nil, nil, errors.New("Size issue")
  }
  return transcriptItems, transcriptTimestamps, nil
}

func FormatRawSubtitles(items []*cdp.Node, timestamps []*cdp.Node) ([]*VideoSubtitle, error) {
  log.Println("Formatting subtitles into structured data...")
  response := []*VideoSubtitle{} 
  for i:= 0; i < len(items); i++ {
    var timestamp = strings.TrimSpace(timestamps[i].NodeValue)
    var textItem = strings.TrimSpace(items[i].NodeValue)
    seconds, err := utils.GetTimeInSeconds(timestamp)
    if err != nil {
      log.Fatal("Error with the translation of timestamp")
      return nil, err
    }
    response = append(response, &VideoSubtitle{
      Text: textItem,
      TimeInSeconds: seconds, 
      TimeString: timestamp,
    })
  }
  return response, nil
}
