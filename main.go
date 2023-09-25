package main

import (
	"fmt"
)

func main() {
	imageURL := "https://example.com/image.jpg"
	inputFilename := "input.jpg"
	outputFilename := "output.png"

	err := DownloadFile(imageURL, inputFilename)
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return
	}

	err = ConvertJPGtoPNG(inputFilename, outputFilename)
	if err != nil {
		fmt.Println("Error converting file:", err)
		return
	}

	fmt.Println("Image successfully downloaded and converted to PNG")
}
