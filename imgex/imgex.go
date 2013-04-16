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

	for i := bounds.Min.Y; i <= bounds.Max.Y; i++ {
		for j := bounds.Min.X; j <= bounds.Max.X; j++ {
			r, g, b, _ := picture.At(j,i).RGBA()

			gray := (r+g+b)/3

			fmt.Printf("%v ", gray)
		}
		fmt.Print("\n")
	}



}



//normalize takes an image, and returns the same image normalized such that the average value is 0x0FFF (in Go, colors go from 0x0000 to 0xFFFF).
//I'm not sure what the best way to do this is, yet. I'm liking the algorithm described here: http://en.wikipedia.org/wiki/Normalization_(image_processing)
//but I'm not sure if it will do what I want it to.
//#TODO: Implement this method.
// func normalize(picture image) image {
// 	return picture
// }
