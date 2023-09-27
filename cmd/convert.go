package cmd

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/spf13/cobra"
)

var inputFile string
var pngOutputFile string

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert a JPG image to PNG",
	Run: func(cmd *cobra.Command, args []string) {
		err := convertJPGtoPNG(inputFile, pngOutputFile)
		if err != nil {
			fmt.Println("Error converting file:", err)
			return
		}
		fmt.Println("File converted successfully")
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&inputFile, "input", "i", "input.jpg", "Input JPG file")
	convertCmd.Flags().StringVarP(&pngOutputFile, "output", "o", "output.png", "Output PNG file")
}

func convertJPGtoPNG(inputPath, outputPath string) error {
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
