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

	image_file, err := os.Open("Marceline.jpg")
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