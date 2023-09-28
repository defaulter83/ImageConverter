package cmd

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var inputFile string
var outputExtension string

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert an image file to another format",
	Run: func(cmd *cobra.Command, args []string) {
		err := convertImage(inputFile, outputExtension)
		if err != nil {
			fmt.Println("Error converting file:", err)
			return
		}
		fmt.Println("File converted successfully")
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input image file")
	convertCmd.Flags().StringVarP(&outputExtension, "output", "o", "png", "Output file extension (e.g., jpg, png)")
	convertCmd.MarkFlagRequired("input")
}

func convertImage(inputPath, outputExt string) error {
	// Open the input file
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// Decode the input image
	img, format, err := image.Decode(inputFile)
	if err != nil {
		return err
	}

	// Prepare the output file name
	outputPath := strings.TrimSuffix(inputPath, filepath.Ext(inputPath)) + "." + outputExt

	// Create the output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Encode and save the image in the new format
	switch strings.ToLower(outputExt) {
	case "jpg", "jpeg":
		err = jpeg.Encode(outputFile, img, nil)
	case "png":
		err = png.Encode(outputFile, img)
	default:
		return fmt.Errorf("unsupported output format: %s", outputExt)
	}

	if err != nil {
		return err
	}

	fmt.Printf("Image converted from %s to %s\n", format, outputExt)
	return nil
}
