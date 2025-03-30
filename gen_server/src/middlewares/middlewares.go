package middlewares

import (
	"bytes"
	"context"
	"gen_server/src/utils"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	Body       *bytes.Buffer
}

func (rec *responseRecorder) WriteHeader(statusCode int) {
	rec.statusCode = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}

func (rec *responseRecorder) Write(b []byte) (int, error) {
	rec.Body.Write(b)
	return rec.ResponseWriter.Write(b)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var requestBody bytes.Buffer
		tee := io.TeeReader(r.Body, &requestBody)
		bodyBytes, err := io.ReadAll(tee)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		log.Printf("Request: %s %s, Body: %s", r.Method, r.URL.Path, requestBody.String())
		rec := &responseRecorder{ResponseWriter: w, Body: &bytes.Buffer{}}
		next.ServeHTTP(rec, r)
		log.Printf("Response Status: %d, Body: %s", rec.statusCode, rec.Body.String())
	})
}

func skipAuthPath(path string) bool {
	return path == "/api/auth/sign-in" || path == "/api/auth/sign-up" || path == "/api" || path == "/api/local/gym/assign"
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ipAddr, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), "ipAddr", ipAddr)
		r = r.WithContext(ctx)
		if skipAuthPath(r.URL.Path) {
			log.Printf("Auth skipped for path %s", r.URL.Path)
			next.ServeHTTP(w, r)
			return
		}
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		if !strings.Contains(authHeader, "Bearer ") {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		accessToken := strings.Replace(authHeader, "Bearer ", "", 1)
		claims, err := utils.ParseAccessToken(accessToken)
		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		userIdInt, err := strconv.ParseInt(claims.Subject, 10, 64)
		ctx = context.WithValue(r.Context(), "userId", userIdInt)
		r = r.WithContext(ctx)
		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
