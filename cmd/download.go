package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var url string
var outputFile string

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a JPG image from a URL",
	Run: func(cmd *cobra.Command, args []string) {
		err := downloadFile(url, outputFile)
		if err != nil {
			fmt.Println("Error downloading file:", err)
			return
		}
		fmt.Println("File downloaded successfully")
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringVarP(&url, "url", "u", "", "URL of the image to download")
	downloadCmd.Flags().StringVarP(&outputFile, "output", "o", "input.jpg", "Output file name")
	downloadCmd.MarkFlagRequired("url")
}

func downloadFile(url, filename string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	return err
}
