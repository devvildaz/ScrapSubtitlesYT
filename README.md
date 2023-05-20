# ScrapSubtitlesYT

A Golang project that extracts the subtitles of a YouTube video (for now) based on its ID.

## Description

ScrapSubtitlesYT is a command-line tool built in Golang that allows you to fetch and extract the subtitles of a YouTube video. It utilizes scraping techinques to retrieve the video metadata and extract the available subtitles. Currently, the tool supports extracting subtitles based on the video ID provided.

## Features

- Fetches and extracts subtitles from YouTube videos
- Supports extracting subtitles by video ID
- Provides command-line interface (CLI) for easy usage

## Usage
1. Clone the repository
```
git clone https://github.com/devvildaz/ScrapSubtitlesYT
cd ScrapSubtitlesYT
```
2. Install dependencies
```
go mod download
```
3. Build the executable
```
go build -o <output-file> ./internal/
```
4. Run the executable
```
./<output-file> -youtube_id <id of your youtube video>
```
5. The application will generate a JSON file that includes the subtitles and corresponding timestamps in the specified format.
Note: To improve the formatting of JSON files, you can employ tools such as jq.
```
cat <youtube-id>.json | jq '.'

[
...
  {
    "TimeString": "32:22",
    "TimeInSeconds": 1942,
    "Text": "of thousands of german soldiers occupied at a time when manpower was critically low"
  },
  {
    "TimeString": "32:28",
    "TimeInSeconds": 1948,
    "Text": "and the experience gained from the amphibious landings at salerno proved invaluable to the planners of"
  },
  {
    "TimeString": "32:34",
    "TimeInSeconds": 1954,
    "Text": "d-day still historians remained divided on the campaign with many arguing that it was an"
  },
...
]
```
Another note: If you want to obtain the entire subtitles as a single large text, you can use jq to map all the Text Attributes.
```
cat <youtube-id>.json | jq -r 'map(.Text) | join(" ")'
```
## TODOs

- [ ] Implement language selection for subtitles extraction (for now it's getting the default selection by your location or the video)
- [ ] Enhance error handling and user feedback
- [ ] Improve documentation and code comments
- [ ] Improve the structure of the result
- [ ] Integration with ChatGPT?

## Author

- [devvildaz](https://github.com/devvildaz)

Feel free to contribute, report issues, and submit pull requests to improve this project.
