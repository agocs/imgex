package main
import (
			"fmt"
			//"io/ioutil"
			//"image"
			"os"
			"log"
			"image/jpeg"
		)


func main(){

	filename := "Marceline.jpg" //default

	if len(os.Args) > 1 {
		filename = os.Args[1]
	} else {
		fmt.Println("Usage: imgex.exe [filename]. No file name found, defaulting to Marceline.jpg")
	}

	image_file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer image_file.Close()

	picture, err := jpeg.Decode(image_file)
	if err != nil {
		log.Fatal(err)
	}

	bounds := picture.Bounds()

	fmt.Println("bounds: %v %v %v %v" , bounds.Min.X , bounds.Min.Y , bounds.Max.X, bounds.Max.Y)


}