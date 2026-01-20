package middleware

import (
	"context"
	"net/http"
	"strings"

	jwtpkg "github.com/rseigha/goecomapi/pkg/jwt"
	"go.uber.org/zap"
)

func JWTAuth(jwt *jwtpkg.JWT, logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				http.Error(w, "missing authorization header", http.StatusUnauthorized)
				return
			}
			parts := strings.SplitN(auth, " ", 2)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				http.Error(w, "invalid authorization header", http.StatusUnauthorized)
				return
			}
			token := parts[1]
			claims, err := jwt.ParseToken(token)
			if err != nil {
				logger.Debug("token parse error", zap.Error(err))
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}
			// attach claims to context
			ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
			ctx = context.WithValue(ctx, "user_email", claims.Email)
			ctx = context.WithValue(ctx, "user_role", claims.Role)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}