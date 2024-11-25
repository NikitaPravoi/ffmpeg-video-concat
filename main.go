package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type VideoData struct {
	VideoPiece string `json:"video_piece"`
}

const baseURL = "https://example.com/s3/storage"

func main() {
	// Step 1: Read JSON file
	fileName := "videos.json"
	dir := "video"

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return
	}

	jsonData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Step 2: Parse JSON
	var videos []VideoData
	err = json.Unmarshal(jsonData, &videos)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Step 2: Download each video and save it locally
	var videoFiles []string
	for _, video := range videos {
		url := baseURL + video.VideoPiece
		fmt.Println("Downloading:", url)
		fileName := downloadVideo(dir, url)
		if fileName != "" {
			videoFiles = append(videoFiles, fileName)
		}
	}

	// Step 3: Concatenate videos (this step assumes you have FFMpeg installed)
	outputFile := "output_full_video.mp4"
	concatenateVideos(videoFiles, dir, outputFile)
}

// Helper function to download a video
func downloadVideo(dir, url string) string {
	// Get the file name from the URL
	parts := strings.Split(url, "/")
	fileName := parts[len(parts)-1]

	// Create a new file to save the video
	out, err := os.Create("./" + dir + "/" + fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return ""
	}
	defer out.Close()

	// Get the video from the URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading video:", err)
		return ""
	}
	defer resp.Body.Close()

	// Write the video to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error saving video:", err)
		return ""
	}

	return fileName
}

// Helper function to concatenate videos using FFMpeg
func concatenateVideos(videoFiles []string, dir, outputFile string) {
	// Generate the input list for FFMpeg
	listFile := "input_list.txt"
	list, err := os.Create(listFile)
	if err != nil {
		fmt.Println("Error creating list file:", err)
		return
	}
	defer list.Close()

	for _, file := range videoFiles {
		list.WriteString(fmt.Sprintf("file './%s/%s'\n", dir, file))
	}

	// Run FFMpeg command to concatenate the videos
	cmd := fmt.Sprintf("ffmpeg -f concat -safe 0 -i %s -c copy %s", listFile, outputFile)
	err = runCommand(cmd)
	if err != nil {
		fmt.Println("Error concatenating videos:", err)
	}
}

// Helper function to run shell commands
func runCommand(command string) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
