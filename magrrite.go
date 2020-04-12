package main

import (
	"fmt"
	"image"
	_ "image/jpeg" // decode jpeg image
	_ "image/png"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

func rgbToGrayScale(R, G, B uint32) uint32 {
	grayVal := (19595*R + 38470*G + 7471*B + 1<<15) >> 24
	return grayVal
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: magrrite [FILE]\n\nConvert picture into an ASCII art\n\nExample:\nmagrrite /home/prdpx7/Pictures/birthday.jpg")
		return
	}
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Invalid image file path")
		return
	}
	defer f.Close()

	ASCIIMap := [...]string{".", ",", ":", ";", "+", "*", "?", "%", "S", "#", "@"}
	// offset = 256/len(ASCII_MAP)
	pixelOffset := uint32(25)
	imgWidth := 70

	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Println("Unable to decode image")
		return
	}

	resizedImg := resize.Resize(uint(imgWidth), 0, img, resize.Lanczos3)

	pixelArray := []string{}

	for y := resizedImg.Bounds().Min.Y; y < resizedImg.Bounds().Max.Y; y++ {
		// first iteration:  01 02 03 04 05 06
		// second iteration: 11 12 13 14 15 16
		// ....
		for x := resizedImg.Bounds().Min.X; x < resizedImg.Bounds().Max.X; x++ {

			R, G, B, _ := resizedImg.At(x, y).RGBA()

			grayVal := rgbToGrayScale(R, G, B)
			// get respective ascii char for given y,x based on grayscale value
			// higher the value, higher the dense ascii char in ASCIIMap
			pixelChar := ASCIIMap[(grayVal / pixelOffset)]
			pixelArray = append(pixelArray, pixelChar)
		}
	}

	pixelLength := len(pixelArray)
	for i := 0; i < pixelLength; i += imgWidth {
		// print each row of pixel matrix
		fmt.Println("\t\t\t", strings.Join(pixelArray[i:i+imgWidth], ""))
	}

}
