package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Test")
}

func TestHelloHandler(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/test", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}

func BilangHalo(w http.ResponseWriter, r *http.Request) {
	nama := r.URL.Query().Get("nama")

	if nama == "" {
		fmt.Fprint(w, "Halo banh")
	} else {
		fmt.Fprintf(w, "Halo bang %s", nama)
	}
}

func TestBilangHalo(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hallo?nama=Nama", nil)
	recorder := httptest.NewRecorder()

	BilangHalo(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
