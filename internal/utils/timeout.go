package utils

import (
  "context"
  "time"
  "github.com/chromedp/chromedp"
  "strings"
  "strconv"
  "errors"
  "math"
//  "github.com/chromedp/cdproto/cdp"
)

func RunWithTimeOut(ctx *context.Context, timeout time.Duration, tasks chromedp.Tasks) chromedp.ActionFunc{
  return func (ctx context.Context) error {
    timeoutContext, cancel := context.WithTimeout(ctx, timeout * time.Second)
    defer cancel()
    return tasks.Do(timeoutContext)
  }
}

func GetTimeInSeconds(timeStr string) (int, error) {
  components := strings.Split(timeStr, ":")
  rate := 60.0
  var seconds int = 0
  var i int

  for i=len(components); i > 0; i-- {
    value, err := strconv.Atoi(components[i-1])
    if err != nil {
      return 0, errors.New("invalid time format")
    }
    seconds += value * int(math.Pow(rate, float64(len(components)-i)))
  }

	return seconds, nil
}
