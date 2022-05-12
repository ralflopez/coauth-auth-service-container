package middlewares

import (
	"coauth/pkg/config/di"
	"coauth/pkg/middlewares/keys"
	"context"
	"log"
	"net/http"
)

func WithUserMiddleware(next http.Handler, di *di.DIContainer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := di.SessionService.GetLoggedInUser(w, r)
		if err == nil {
			// Do stuff here
			ctx := context.WithValue(context.Background(), keys.WithUserCtxKey, user)
			log.Println(r.RequestURI)
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
