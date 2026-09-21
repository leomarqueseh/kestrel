package auth

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const claimsContextKey contextKey = "auth_claims"

// RequireAuth validates the Bearer access token and injects its claims
// into the request context for downstream handlers/middleware to use.
func RequireAuth(tokens *TokenIssuer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			rawToken, found := strings.CutPrefix(header, "Bearer ")
			if !found || rawToken == "" {
				writeError(w, http.StatusUnauthorized, "missing bearer token")
				return
			}

			claims, err := tokens.Parse(rawToken)
			if err != nil || claims.TokenType != "access" {
				writeError(w, http.StatusUnauthorized, "invalid or expired token")
				return
			}

			ctx := context.WithValue(r.Context(), claimsContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// RequireRole enforces least privilege: only the listed roles may proceed.
// Must run after RequireAuth, which populates the claims in context.
func RequireRole(roles ...Role) func(http.Handler) http.Handler {
	allowed := make(map[Role]bool, len(roles))
	for _, r := range roles {
		allowed[r] = true
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value(claimsContextKey).(*Claims)
			if !ok {
				writeError(w, http.StatusUnauthorized, "authentication required")
				return
			}
			if !allowed[claims.Role] {
				writeError(w, http.StatusForbidden, "insufficient permissions")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// ClaimsFromContext retrieves the authenticated user's claims injected by
// RequireAuth. Handlers use this to know who is making the request.
func ClaimsFromContext(ctx context.Context) (*Claims, bool) {
	claims, ok := ctx.Value(claimsContextKey).(*Claims)
	return claims, ok
}
