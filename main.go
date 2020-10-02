package main

import (
	"fmt"
	"log"
	"os"
)

const imageWidth = 256
const imageHeight = 256

func main() {

	file := openFile()
	fmt.Fprintln(file, "P3")
	fmt.Fprintf(file, "%d %d\n255\n", imageWidth, imageHeight)

	for j := imageHeight - 1; j >= 0; j-- {
		fmt.Printf("Scanlines remaining: %d\n", j)
		for i := 0; i < imageWidth; i++ {
			r := float32(i) / (imageWidth - 1)
			g := float32(j) / (imageHeight - 1)
			b := 0.25
			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)
			fmt.Fprintf(file, "%d %d %d\n", ir, ig, ib)
		}
	}
	fmt.Println("Done.")

}

func openFile() *os.File {
	var filename string
	switch len(os.Args) {
	case 2:
		filename = os.Args[1]
	default:
		filename = "image.ppm"
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln("Unable to open file", err)
	}
	file.Seek(0, os.SEEK_SET)
	return file
}
