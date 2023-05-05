package main

import (
	"net/http";
	"os";
	"fmt";
	"encoding/json";
)

type StorageInfo struct {
	File		string `json:"filename"`
	Content		string `json:"content"`
}

func ShowMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")	
}

func SaveValue(w http.ResponseWriter, r *http.Request) {
	var info StorageInfo

	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	f, err := os.Create(info.File)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	if _, err := f.WriteString(info.Content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
