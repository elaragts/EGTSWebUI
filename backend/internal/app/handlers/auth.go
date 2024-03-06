package handlers

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/internal/database"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/internal/model"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/pkg"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
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
	var user model.AuthUser
	user.Username = username
	user.Baid = baid
	user.PasswordHash = string(hash)
	err = database.InsertAuthUser(user)
	if err != nil {
		http.Error(w, "Error inserting user", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (a AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	//if use is already logged in
	_, err := r.Cookie("Authorization")
	if err == nil {
		http.Error(w, "User already logged in", http.StatusBadRequest)
		return
	}
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	if username == "" || password == "" {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		log.Printf("Invalid form: username=%s, password=[REDACTED]", username)
		return
	}
	user, found, err := database.GetAuthUserByUsername(username)
	if err != nil {
		http.Error(w, "Error getting user", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if !found {
		http.Error(w, "Username or Password is incorrect", http.StatusUnauthorized)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		http.Error(w, "Username or Password is incorrect", http.StatusUnauthorized)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Baid,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(pkg.ConfigVars.SessionSecret))
	if err != nil {
		http.Error(w, "Error creating token", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	expiration := time.Now().Add(30 * 24 * time.Hour)
	cookie := http.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Expires:  expiration,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}
