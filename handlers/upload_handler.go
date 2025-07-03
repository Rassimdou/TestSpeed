package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//read and discard all upload data
	n, err := io.Copy(io.Discard, r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	fmt.Println("Received upload of size:", n, "bytes")

	// Respond with a success message
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "success", "message": "File uploaded successfully"}`))

}
