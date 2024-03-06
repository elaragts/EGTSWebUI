package middleware

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/pkg"
	"net/http"
	"time"
)

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		var ctx context.Context
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}
			ctx = context.WithValue(r.Context(), "baid", claims["sub"])
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
