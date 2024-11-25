# ffmpeg-video-concat

A Go application that downloads video segments from a remote storage location and concatenates them into a single video file using FFmpeg.

## Features

- Reads video segment information from a JSON file
- Downloads video segments from a specified S3 storage URL
- Concatenates downloaded segments into a single video file using FFmpeg
- Handles errors gracefully with detailed error messages

## Prerequisites

- Go 1.x or higher
- FFmpeg installed on your system
- Internet connection to access the remote storage

## Installation

1. Clone this repository or download the source code
2. Ensure FFmpeg is installed on your system
   ```bash
   # Ubuntu/Debian
   sudo apt-get install ffmpeg

   # MacOS
   brew install ffmpeg
   ```

## Usage

1. Create a JSON file named `videos.json` with your video segment information:
   ```json
   [
     {
       "video_piece": "/path/to/video1.mp4"
     },
     {
       "video_piece": "/path/to/video2.mp4"
     }
   ]
   ```

2. Update the `baseURL` constant in the code to point to your storage location:
   ```go
   const baseURL = "https://your-storage-url.com/s3/storage"
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

The application will:
- Create a `video` directory to store downloaded segments
- Download all video segments specified in `videos.json`
- Concatenate the segments into `output_full_video.mp4`

## Project Structure

```
.
├── main.go              # Main application code
├── videos.json          # Input JSON file with video segment information
├── video/              # Directory for downloaded video segments
└── output_full_video.mp4 # Final concatenated video file
```

## Functions

### `main()`
- Entry point of the application
- Orchestrates the video processing workflow

### `downloadVideo(dir, url string) string`
- Downloads a video segment from the specified URL
- Saves the video to the specified directory
- Returns the filename of the downloaded video

### `concatenateVideos(videoFiles []string, dir, outputFile string)`
- Creates a list file for FFmpeg
- Concatenates all video segments using FFmpeg
- Outputs the final video file

### `runCommand(command string) error`
- Helper function to execute shell commands
- Used for running FFmpeg commands

## Error Handling

The application includes comprehensive error handling for:
- JSON file reading and parsing
- Video download issues
- Directory creation
- File operations
- FFmpeg execution

## Limitations

- Requires FFmpeg to be installed on the system
- All video segments must be compatible for concatenation
- Assumes video segments are accessible via HTTP/HTTPS
- Memory usage depends on video file sizes

## License

Licensed with MIT license.
