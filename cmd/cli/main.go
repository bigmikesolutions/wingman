package main

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"golang.org/x/oauth2"
	"io"
	"log"
	"net/http"
)

var (
	oauth2Config oauth2.Config
	oauth2State  = "randomstate"
	redirectURL  = "http://localhost:8080/oauth2/callback"
	clientID     = "wingman"
	issuerURL    = "http://localhost:8082/realms/wingman"
)

func main() {
	oauth2Config = oauth2.Config{
		ClientID:    clientID,
		RedirectURL: redirectURL,
		Scopes:      []string{"openid", "profile", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  fmt.Sprintf("%s/protocol/openid-connect/auth", issuerURL),
			TokenURL: fmt.Sprintf("%s/protocol/openid-connect/token", issuerURL),
		},
	}

	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/oauth2/callback", handleCallback)
	http.HandleFunc("/protected", handleProtected)

	log.Printf("Server is starting at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var codeVerifier string

func handleLogin(w http.ResponseWriter, r *http.Request) {
	codeChallenge := ""
	codeVerifier, codeChallenge = generatePKCE()

	authURL := oauth2Config.AuthCodeURL(
		oauth2State,
		oauth2.AccessTypeOffline,
		oauth2.SetAuthURLParam("code_challenge", codeChallenge),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
	)

	log.Printf("Redirecting to Authorization URL: %s", authURL)

	http.Redirect(w, r, authURL, http.StatusFound)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	log.Println("###### handle callback")
	for k, v := range r.URL.Query() {
		log.Printf("%s: %s\n", k, v)
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing authorization code", http.StatusBadRequest)
		return
	}

	token, err := oauth2Config.Exchange(
		context.Background(),
		code,
		oauth2.SetAuthURLParam("code_verifier", codeVerifier),
	)
	if err != nil {
		http.Error(w, "Failed to exchange code for token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token.AccessToken,
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	})
	http.Redirect(w, r, "/protected", http.StatusFound)
}

func handleProtected(w http.ResponseWriter, r *http.Request) {
	//for _, cookie := range r.Cookies() {
	//	log.Printf("Cookie: %s = %s", cookie.Name, cookie.Value)
	//}

	tokenCookie, err := r.Cookie("access_token")
	if err != nil || tokenCookie == nil {
		http.Error(w, "No token found", http.StatusUnauthorized)
		return
	}

	log.Printf("Access token: %s\n", tokenCookie.Value)

	client := oauth2Config.Client(context.Background(), &oauth2.Token{AccessToken: tokenCookie.Value})
	resp, err := client.Get("http://localhost:4180/probes/health")
	if err != nil {
		http.Error(w, "Failed to access protected resource: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	log.Printf("Protected resource status: %s\n", resp.Status)

	w.WriteHeader(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
	}

	_, _ = w.Write(body)
}

func generatePKCE() (string, string) {
	codeVerifier := generateCodeVerifier()
	hash := sha256.Sum256([]byte(codeVerifier))
	codeChallenge := base64.RawURLEncoding.EncodeToString(hash[:]) // No padding
	return codeVerifier, codeChallenge
}

// Generate a secure random string (code_verifier)
func generateCodeVerifier() string {
	verifier := make([]byte, 32) // 32 bytes = 256-bit security
	_, err := rand.Read(verifier)
	if err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(verifier) // No padding
}
