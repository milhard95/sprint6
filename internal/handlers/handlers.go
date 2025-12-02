package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

var index []byte

func HandleRoot(res http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodGet {
		http.Error(res, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data, err := getIndex()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "text/html")
	res.WriteHeader(http.StatusOK)
	res.Write(data)

}

func HandleUpload(res http.ResponseWriter, req *http.Request) {

	err := req.ParseMultipartForm(1024)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	file, _, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	bufer, err := io.ReadAll(file)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := service.ConverText(string(bufer))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	fileName := newFileName()
	err = os.WriteFile(fileName, []byte(result), 0755)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	msg := struct {
		Original string `json:"original"`
		Result   string `json:"result"`
	}{
		Original: string(bufer),
		Result:   result,
	}

	response, err := json.MarshalIndent(msg, "", "    ")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	res.Write(response)

}

func newFileName() string {
	timestamp := time.Now().UTC().Format("2006-01-02_15-04-05")
	return timestamp + ".txt"
}

func getIndex() ([]byte, error) {

	if len(index) != 0 {
		return index, nil
	}

	nameArr := []string{"../", "./", "/", ""}

	var nameFile string
	for _, v := range nameArr {
		nameFile = v + "index.html"
		_, err := os.Stat(nameFile)
		if !errors.Is(err, os.ErrNotExist) {
			data, err := os.ReadFile(nameFile)
			if err != nil {
				return []byte{}, err
			}
			index = data
			return index, nil
		}
	}

	return []byte{}, errors.New("file 'index.html' does not exist")

}
