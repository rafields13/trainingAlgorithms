package main

import "fmt"

type Video struct {
	ID            int
	Duration      int
	Entertainment int
}

func findBestVideo(t int, videos []Video) (Video, error) {
	var bestVideo Video
	greatestEntertainment := 0

	for _, video := range videos {
		if video.Duration+1 < t && video.Entertainment > greatestEntertainment {
			bestVideo = video
			greatestEntertainment = video.Entertainment
		}
	}

	if bestVideo.ID == 0 {
		return Video{}, fmt.Errorf("error")
	}

	return bestVideo, nil
}

func main() {
	t := 10
	videos := []Video{
		{
			ID:            1,
			Duration:      6,
			Entertainment: 8,
		},
		{
			ID:            2,
			Duration:      7,
			Entertainment: 5,
		},
		{
			ID:            3,
			Duration:      9,
			Entertainment: 9,
		},
		{
			ID:            4,
			Duration:      4,
			Entertainment: 6,
		},
	}

	bestVideo, err := findBestVideo(t, videos)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("Video{\n\tID: %d,\n\tDuration: %d,\n\tEntertainment: %d,\n}", bestVideo.ID, bestVideo.Duration, bestVideo.Entertainment)
	}
}
