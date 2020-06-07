package main

import (
	"fmt"
	"gofullstack/SocialMedia"
)

func main() {
	myPost := SocialMedia.NewPost("Zack", SocialMedia.Moods["sad"], "Life in Saudi", "I regret not staying the in US", "Guru.org", "", "", []string{"Saudi", "is", "a", "prison"})
	fmt.Printf("MyPost: %+v\n", myPost)
}
