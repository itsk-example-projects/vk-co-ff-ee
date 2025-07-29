package web

import (
	"fmt"
	"net/http"

	"vk-co-ff-ee/internal/crypto"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := pageData{}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		action := r.FormValue("action")
		input := r.FormValue("input")
		custom := r.FormValue("customKey")
		if action == "encrypt" {
			data.PlainInput = input
			data.EncryptResult = crypto.EncryptVKCoffee(input, custom)
		} else if action == "decrypt" {
			data.CipherInput = input
			if res, err := crypto.DecryptVKCoffee(input, custom); err == nil {
				data.DecryptResult = res
			} else {
				data.DecryptResult = fmt.Sprintf("Error: %v", err)
			}
		}
	}
	if err := pageTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
