package handler

import (
	"encoding/json"
	"net/http"
	"portfolio/backend/service"
	_struct "portfolio/backend/struct"
	"strings"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "" && !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		http.Error(w, "Content-Type debe ser application/json", http.StatusUnsupportedMediaType)
		return
	}

	var form _struct.ContactForm
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&form)
	if err != nil {
		http.Error(w, "Error trying to decode the file", http.StatusBadRequest)
		return
	}

	if form.Name == "" || form.Email == "" || form.Message == "" || form.Subject == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	emailError := service.SendEmail(form)

	if emailError != nil {
		http.Error(w, "Error sending the email", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Message sent successfully",
	})
}
