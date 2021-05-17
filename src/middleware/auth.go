package middleware

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Eklow-AI/Gotham/src/models"
	"github.com/dgrijalva/jwt-go"
)

// SecureRequest middleware is used to secure routes by requiring
// a valid JWT with the incoming request
func SecureRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			// if the cookie is not set, return an unathorized status
			if err == http.ErrNoCookie {
				resp := map[string]interface{}{
					"success": false,
					"message": "No cookie found, user is not authorized",
					"code":    401,
				}
				json.NewEncoder(w).Encode(resp)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// for any other type of error, return status 500
			resp := map[string]interface{}{
				"success": false,
				"message": "Internal server error",
				"code":    500,
			}
			json.NewEncoder(w).Encode(resp)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Get the JWT from cookie and init a new 'Claims' instance
		tokenString := cookie.Value
		claims := models.Claims{}
		// Parse the JWT string and store the values in claims struct instance
		// We're passing the key in this method as well. This part will return an error
		// if the JWT has expired or if the signature does not match
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return os.Getenv("jwtKey"), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				resp := map[string]interface{}{
					"success": false,
					"message": "User is not Authorized",
					"code":    401,
				}
				json.NewEncoder(w).Encode(resp)
				return
			}
			resp := map[string]interface{}{
				"success": false,
				"message": "Bad Request",
				"code":    400,
			}
			json.NewEncoder(w).Encode(resp)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !token.Valid {
			resp := map[string]interface{}{
				"success": false,
				"message": "User is not authorized",
				"code":    401,
			}
			json.NewEncoder(w).Encode(resp)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Everything, is well just pass the request to the next handler down the chain
		next.ServeHTTP(w, r)
	})
}
