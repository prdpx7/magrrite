package main
import (
    "fmt"
    "strings"
    "os"
    "image"
    "github.com/nfnt/resize"
    _ "image/jpeg"
    _ "image/png"
)

func rgbToGrayScale(R,G,B uint32) uint32 {
    gray_val := (19595*R + 38470*G + 7471*B + 1<<15) >> 24
    return gray_val
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

    ASCII_MAP := [...]string {".", ",", ":", ";", "+", "*", "?", "%", "S", "#", "@"}
    // offset = 256/len(ASCII_MAP)
    pixel_offset := uint32(25)
    img_width := 60

    img,_,err := image.Decode(f)
    if err != nil {
        fmt.Println("Unable to decode image")
        return
    }

    resized_img := resize.Resize(uint(img_width), 0, img, resize.Lanczos3)
    
    pixel_arr := []string{}

    for y := resized_img.Bounds().Min.Y; y < resized_img.Bounds().Max.Y; y++ {
        // first iteration:  01 02 03 04 05 06
        // second iteration: 11 12 13 14 15 16
        // ....
        for x := resized_img.Bounds().Min.X; x < resized_img.Bounds().Max.X; x ++ {
            
            R,G,B,_ := resized_img.At(x,y).RGBA()

            gray_val := rgbToGrayScale(R,G,B)
            // get respective ascii char for given y,x based on grayscale value
            // higher the value, higher the dense ascii char in ASCII_MAP
            pixel_char := ASCII_MAP[(gray_val/pixel_offset)]
            pixel_arr = append(pixel_arr, pixel_char)
        }
    }

    pixel_length := len(pixel_arr)
    for i := 0; i < pixel_length; i += img_width {
        // print each row of pixel matrix
        fmt.Println(strings.Join(pixel_arr[i:i+img_width], ""))
    }
    
}
