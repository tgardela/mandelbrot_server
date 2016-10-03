package mandelServer

import (
    "testing"
    "net/http"
	"net/http/httptest"
)

func TestIfMandelHandlerSubareaReturnsOKStatus (t *testing.T){
    req, err := http.NewRequest("GET", "/health-check", nil)
    if err != nil {
        t.Fatal(err)
    }
    // create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(MandelHandlerSubarea)

    // handlers satisfy http.Handler, its possible to call ServeHTTP method 
    // directly and pass in Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
}


func TestIfMandelHandlerReturnsOKStatus (t *testing.T){
    req, err := http.NewRequest("GET", "/health-check", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(MandelHandler)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
}