package handler

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Response struct {
	Data []struct {
		Password string `json:"password"`
		Ip       string `json:"ip"`
		Port     string `json:"port"`
		Title    string `json:"title"`
	} `json:"data"`
}

func decryptCBC(ciphertext, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// Remove padding
	padding := plaintext[len(plaintext)-1]
	return plaintext[:len(plaintext)-int(padding)], nil
}

func fetchAndDecryptData() ([]string, error) {
	url := "http://api.skrapp.net/api/serverlist"
	headers := map[string]string{
		"Accept":          "/",
		"Accept-Language": "zh-Hans-CN;q=1, en-CN;q=0.9",
		"AppVersion":      "1.3.1",
		"User-Agent":      "SkrKK/1.3.1 (iPhone; iOS 13.5; Scale/2.00)",
		"Content-Type":    "application/x-www-form-urlencoded",
		"Cookie":          "PHPSESSID=fnffo1ivhvt0ouo6ebqn86a0d4",
	}

	data := "data=4265a9c353cd8624fd2bc7b5d75d2f18b1b5e66ccd37e2dfa628bcb8f73db2f14ba98bc6a1d8d0d1c7ff1ef0823b11264d0addaba2bd6a30bdefe06f4ba994ed"

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	l, err := hex.DecodeString(string(body))
	if err != nil {
		return nil, err
	}

	key := []byte("65151f8d966bf596")
	iv := []byte("88ca0f0ea1ecf975")
	m, err := decryptCBC(l, key, iv)
	if err != nil {
		return nil, err
	}

	var result Response
	if err := json.Unmarshal(m, &result); err != nil {
		return nil, err
	}

	var outputs []string
	for _, o := range result.Data {
		p := fmt.Sprintf("aes-256-cfb:%s@%s:%s", o.Password, o.Ip, o.Port)
		q := base64.StdEncoding.EncodeToString([]byte(p))
		r := fmt.Sprintf("ss://%s#%s", q, o.Title)
		outputs = append(outputs, r)
	}
	return outputs, nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	outputs, err := fetchAndDecryptData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s := strings.Join(outputs, "|")

	url := fmt.Sprintf("https://url.v1.mk/sub?target=singbox&url=%s%s", url.QueryEscape(s), `&insert=false&config=https%3A%2F%2Fraw.githubusercontent.com%2FACL4SSR%2FACL4SSR%2Fmaster%2FClash%2Fconfig%2FACL4SSR_Online_Mini.ini&emoji=true&list=false&xudp=false&udp=false&tfo=false&expand=true&scv=false&fdn=false`)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(w, "")
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.WriteHeader(resp.StatusCode)

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
