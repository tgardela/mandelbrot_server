package mandelServer
//package main

import (
    "bytes"
    "image"
    "image/png"
    "log"
    "net/http"
    "strconv"
    "scaleworks/mandel"
)

func MandelHandlerSubarea(w http.ResponseWriter, r *http.Request) {
    width, height := getWidthAndheighFromUrl(r)
    re0, im0, re1, im1 := getBoundariesFromUrl(r)

    m := mandel.Mandelbrot(image.Rect(0, 0, width, height), complex(re0, im0), complex(re1, im1))
    writeImage(w, &m)
}


func MandelHandler(w http.ResponseWriter, r *http.Request) {
    width, height := getWidthAndheighFromUrl(r)
    lowerBoundary :=  -2-2i
    upperBoundary := 2+2i
    m := mandel.Mandelbrot(image.Rect(0, 0, width, height), lowerBoundary, upperBoundary)
    writeImage(w, &m)
}


func RedirectToMandelbrot(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/mandelbrot", 301)
}


func getWidthAndheighFromUrl(r *http.Request) (int, int){
    width, werr := strconv.Atoi(r.URL.Query().Get("width"))
    if werr != nil { width = 500 }
    
    height, herr := strconv.Atoi(r.URL.Query().Get("height")) 
    if herr != nil { height = 500 }
    
    return width, height
}


func getBoundariesFromUrl(r *http.Request) (float64, float64, float64, float64){
    //nie czyta dobrze wartosci ale dobrze rusyje dla calkowitych, float nie dziala
    re0, errre0:= strconv.ParseFloat(r.URL.Query().Get("re"), 64)
    if errre0 != nil { re0 = -2. }
    
    im0, errim0:= strconv.ParseFloat(r.URL.Query().Get("im0"), 64)
    if errim0 != nil { im0 = -2. }
    
    re1, errre1:= strconv.ParseFloat(r.URL.Query().Get("re1"), 64)
    if errre1 != nil { re1 = 2. }
    
    im1, errim1:= strconv.ParseFloat(r.URL.Query().Get("im1"), 64)
    if errim1 != nil { im1 = 2. }
    
    return re0, im0, re1, im1
}


// writeImage encodes an image 'img' in png format and writes it into ResponseWriter.
func writeImage(w http.ResponseWriter, img *image.Image) {
    buffer := new(bytes.Buffer)
    if err := png.Encode(buffer, *img); err != nil {
        log.Println("unable to encode image.")
    }

    w.Header().Set("Content-Type", "image/png")
    w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
    if _, err := w.Write(buffer.Bytes()); err != nil {
        log.Println("unable to write image.")
    }
}


func main() {
    mux := http.NewServeMux()
    
    mansub := http.HandlerFunc(MandelHandlerSubarea)
    man := http.HandlerFunc(MandelHandler)
    red := http.HandlerFunc(RedirectToMandelbrot)
    
    mux.Handle("/mandelbrot/subarea", mansub)
    mux.Handle("/mandelbrot", man)
    mux.Handle("/", red)
    
    log.Println("Listening on 8080")
    err := http.ListenAndServe(":8080", mux)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}