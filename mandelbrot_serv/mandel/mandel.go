package mandel
//package main

import (
    "flag"
    "fmt"
    "image"
    "image/color"
    "image/png"
    "image/draw"
    "os"
    "strings"
    "strconv"
    "math/cmplx"
)
 
 
func Mandelbrot(size image.Rectangle, lower, upper complex128) image.Image{
    rMin   := real(lower)
    rMax   := real(upper)
    iMin   := imag(lower)
    red    := 230.
    green  := 235.
    blue   := 255.

    width, height := getWidthAndHeightFromPoint(size)
    scale := float64(width) / (rMax - rMin)
    bounds := image.Rect(0, 0, width, height)
    mandel := image.NewNRGBA(bounds)
    draw.Draw(mandel, bounds, image.NewUniform(color.Black), image.ZP, draw.Src)
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            fEsc := getMandelbrot(complex(float64(x)/scale+rMin, float64(y)/scale+iMin))
            mandel.Set(x, y, color.NRGBA{uint8(red * fEsc), uint8(green * fEsc), uint8(blue * fEsc), 255})
        }
    }
    return mandel
}


func getWidthAndHeightFromPoint(size image.Rectangle) (int, int){
    dimensions := strings.Split(size.Size().String(), ",")
    width, _ := strconv.Atoi(dimensions[0][1:])
    height, _ := strconv.Atoi(strings.Trim(dimensions[1], ")"))
    
    return width, height
}


func getMandelbrot(a complex128) float64 {
    maxEsc := 32
    i := 0
    for z := a; cmplx.Abs(z) < 2 && i < maxEsc; i++ {
        z = z*z + a
    }
    return float64(maxEsc-i) / float64(maxEsc)
}


func SaveImgM(img image.Image) {
    f, err := os.Create("mandelbrot.png")
    if err != nil {
        fmt.Println(err)
        return
    }
    if err = png.Encode(f, img); err != nil {
        fmt.Println(err)
    }
    if err = f.Close(); err != nil {
        fmt.Println(err)
    }
}


func main() {
    // no check for error input ie: 6xv instead of 600
    width := flag.Int("width", 200, "an int")
    height := flag.Int("height", 200, "an int")
    flag.Parse()

    mb := Mandelbrot(image.Rect(0, 0, *width, *height), -2-2i, 2+2i)
    
    SaveImgM(mb)
}