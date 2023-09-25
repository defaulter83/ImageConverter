package main

import (
	"image/jpeg"
	"image/png"
	"os"
)

func ConvertJPGtoPNG(inputPath, outputPath string) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	img, err := jpeg.Decode(inputFile)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	return png.Encode(outputFile, img)
}
