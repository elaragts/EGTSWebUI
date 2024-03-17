package handlers

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/internal/database"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/internal/model"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/pkg"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math"
	"net/http"
	"time"
)

type AuthHandler struct {
}

func (a AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
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
		//log.Printf("Invalid form: username=%s, password=[REDACTED], accessCode=%s", username, accessCode)
		return
	}
	if len(username) > 20 {
		http.Error(w, "Username must be less than or equal to 20 characters long", http.StatusBadRequest)
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
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}
	simpleUser := model.SimpleAuthUser{
		Username: user.Username,
		Baid:     user.Baid,
	}

	http.SetCookie(w, &cookie)
	json.NewEncoder(w).Encode(simpleUser)
}

func (a AuthHandler) Session(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Authorization")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(pkg.ConfigVars.SessionSecret), nil
	})
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var simpleUser model.SimpleAuthUser
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
		baid := uint(math.Round(claims["sub"].(float64)))
		//log baid
		log.Printf("baid: %d", baid)
		username, found, _ := database.GetUsernameByBaid(baid)
		if !found {
			//print user
			http.Error(w, "Unauthorized", http.StatusInternalServerError)
			return
		}
		simpleUser = model.SimpleAuthUser{
			Username: username,
			Baid:     baid,
		}
	}
	json.NewEncoder(w).Encode(simpleUser)
}

func (a AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "Authorization",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}
