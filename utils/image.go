package utils

// import (
// 	"image"
// 	"image/jpeg" // or the appropriate image format package
// 	"os"

// 	"github.com/nfnt/resize"
// )

// func resizeImage(inputPath string, outputPath string, width int, height int) error {
// 	file, err := os.Open(inputPath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	img, _, err := image.Decode(file)
// 	if err != nil {
// 		return err
// 	}

// 	resizedImg := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

// 	outFile, err := os.Create(outputPath)
// 	if err != nil {
// 		return err
// 	}
// 	defer outFile.Close()

// 	if err = jpeg.Encode(outFile, resizedImg, nil); err != nil {
// 		return err
// 	}

// 	return nil
// }
