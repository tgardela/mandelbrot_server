package mandel

import (
    "testing"
    "image"
    "reflect"
    "os"
)

func TestIfMandelbrotReturnsImg (t *testing.T){
    mb := Mandelbrot(image.Rect(0, 0, 200, 200), -2-2i, 2+2i)
    if reflect.TypeOf(mb) != reflect.TypeOf(image.NewNRGBA(image.Rect(0, 0, 200, 200))){
        t.Error("FAIL: TestIfMandelbrotReturnsImg: types dont match")
    }
}


func TestIsImageSavedOnDisc (t *testing.T) {
    mb := Mandelbrot(image.Rect(0, 0, 200, 200), -2-2i, 2+2i)
    SaveImgM(mb)
    if _, err := os.Stat("mandelbrot.png"); err != nil {
        t.Error("FAIL: TestIsImageSavedOnDisc: File was not created")
    }
}


func TestGetWidthAndHeightFromPoint_returnsProperValues (t *testing.T) {
    width, height := getWidthAndHeightFromPoint(image.Rect(0, 0, 200, 200))
    if width != 200 || height != 200 {
         t.Error("FAIL: TestGetWidthAndHeightFromPoint: Dimensions dont match")
    }
}
