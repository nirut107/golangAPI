package middleware

import (
	"net/http"
	"context"
)

func contextWithUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, "userID", userID)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenStr := cookie.Value
		claims, err := jwtValidate(tokenStr)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := r.Context()
		ctx = contextWithUserID(ctx, claims.ID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}