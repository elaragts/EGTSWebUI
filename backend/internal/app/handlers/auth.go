package handlers

import (
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/internal/database"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type AuthHandler struct {
}

func (a AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	accessCode := r.Form.Get("accessCode")
	if username == "" || password == "" || accessCode == "" {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		log.Printf("Invalid form: username=%s, password=[REDACTED], accessCode=%s", username, accessCode)
		return
	}
	baid, doesExist, err := database.GetBaidFromAccessCode(accessCode)
	if err != nil {
		http.Error(w, "Error getting access code", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if !doesExist {
		http.Error(w, "Access code not found", http.StatusBadRequest)
		return
	}

	if unique, foundType, err := database.IsAuthUserUnique(username, baid); err != nil || !unique {
		if err != nil {
			http.Error(w, "Error checking if user is unique", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		if foundType == database.USERNAMEFOUND {
			http.Error(w, "Username already exists", http.StatusBadRequest)
			return
		}
		http.Error(w, "Access code already linked", http.StatusBadRequest)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	err = database.InsertAuthUser(baid, username, string(hash))
	if err != nil {
		http.Error(w, "Error inserting user", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)

}

//func (a AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
//	err := r.ParseForm()
//	if err != nil {
//		http.Error(w, "Invalid form", http.StatusBadRequest)
//		return
//	}
//	username := r.Form.Get("username")
//	password := r.Form.Get("password")
//	if username == "" || password == "" {
//		http.Error(w, "Invalid form", http.StatusBadRequest)
//		log.Printf("Invalid form: username=%s, password=[REDACTED]", username)
//		return
//	}
//}
