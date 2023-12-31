package main

import (
	"fmt"
	"os"

	"github.com/song940/fileinfo-go/fileinfo"
)

func main() {
	filename := os.Args[1]
	out := fileinfo.Parse(filename)
	fmt.Println("Title:", out.Title)
	fmt.Println("Year:", out.Year)
	fmt.Println("Source:", out.Source)
	fmt.Println("Resolution:", out.Resolution)
	fmt.Println("VideoCodec:", out.VideoCodec)
	fmt.Println("AudioCodec:", out.AudioCodec)
	fmt.Println("Studio:", out.Studio)
	fmt.Println("Channel:", out.Channel)
	fmt.Println("Type:", out.Type)
}
