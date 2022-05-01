package mcode

import (
	"fmt"
	"time"
	"math/rand"
	"os"
	"log"
	"strings"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
)
//
func MakeRandomString() string {
	// Source String used when generating a random identifier.
	const idSource = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	// Save the length in a constant so we don't look it up each time.
	const idSourceLen = byte(len(idSource))
	//
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	// Create an array with the correct capacity
	id := make([]byte, r.Intn(100))
	// Fill our array with random numbers
	rand.Read(id)
	// Replace each random number with an alphanumeric value
	for i, b := range id {
    		id[i] = idSource[b%idSourceLen]
	}
	return string(id)
}
// GenerateID creates a prefixed random identifier.
func GenerateID(prefix string, length int) string {
	// Source String used when generating a random identifier.
	const idSource = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	// Save the length in a constant so we don't look it up each time.
	const idSourceLen = byte(len(idSource))
	// Create an array with the correct capacity
	id := make([]byte, length)
	// Fill our array with random numbers
	rand.Read(id)
	// Replace each random number with an alphanumeric value
	for i, b := range id {
    	    id[i] = idSource[b%idSourceLen]
	}
	// Return the formatted id
	return fmt.Sprintf("%s_%s", prefix, string(id))
}
/*
    fmt.Println(MakeRandomNumber(2))
    v := []string{"email","text","tel",}
    fmt.Println(MakeRandomString(v))
*/
func MakeRandomNumber(kind string) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	//
	var res string = ""
	switch kind {
    	    case "int":   res = fmt.Sprintf("%d,%d",r1.Intn(100),r1.Intn(100))
    	    case "float": res = fmt.Sprintf("%f", r1.Float32())
    	    case "real":
        	fls := fmt.Sprintf("%f",r1.Float32())
        	res = fmt.Sprintf("%d.%s", r1.Intn(100),fls[2:])
	}
	return res
}
//
func GetRandomString(variants []string) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return fmt.Sprintf("%s", variants[r1.Intn(len(variants))])
}
//
func MakeRandomNames(num int) []string {
	var res = []string{}
	//
	count := 0
	for count < num {
    	    res = append(res,fmt.Sprintf("name(%d)",count))
    	    count++
	}
	return res
}
//
func CreateImage(width int, height int, imgPath string, rgba1 uint8, rgba2 uint8,rgba3 uint8, rgba4 uint8) {
	//
	out, err := os.Create(imgPath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	background := color.RGBA{rgba1, rgba2, rgba3, rgba4}
	rect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rect)
	draw.Draw(img, img.Bounds(), &image.Uniform{background}, image.ZP, draw.Src)
	//
	if strings.HasSuffix(strings.ToLower(imgPath), ".jpg") {
		var opt jpeg.Options
		opt.Quality = 80
		err = jpeg.Encode(out, img, &opt)
	} else {
		err = png.Encode(out, img)
	}
	//
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	//
	fmt.Printf("Image saved to %s\n", imgPath)
}
//
func CreateImageArray(imgPath string, width int, height int, limit uint8, offset uint8) {
    //
    var count uint8 = 0
    for count < limit {
        step := fmt.Sprintf("%sart_%d.jpeg",imgPath,count)
        CreateImage(width, height, step, 250, count, count, 250)
        count = count + offset
        //
        CreateImage(width, height, step, 250, count, 250, count)
        count = count + offset*2
        //
        CreateImage(width, height, step, 250, 250, count, count)
        count = count + offset*3
    }
}
