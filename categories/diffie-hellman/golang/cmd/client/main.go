package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/oklog/ulid/v2"
)

var (
	serverEndpoint string
	cliKey         string
)

func getEnv(key, defaultVal string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultVal
	}

	return v
}

func init() {
	serverEndpoint = getEnv("SERVER_ENDPOINT", "localhost:8091")
}

type KeyResponse struct {
	P int `json:"p"`
	G int `json:"g"`
}

func generateCliKey() string {
	return strings.ToLower(ulid.Make().String())
}

func getKey() (*KeyResponse, error) {
	keyResp, err := http.Get(fmt.Sprintf("http://%s/key", serverEndpoint))
	if err != nil {
		log.Fatalf("error while getting key: %v", err)
		return nil, err
	}
	defer keyResp.Body.Close()
	keyRespBody, _ := io.ReadAll(keyResp.Body)
	var keyObj KeyResponse
	json.Unmarshal(keyRespBody, &keyObj)
	return &keyObj, nil
}

func handshake(sharedKey int) (int, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("http://%s/handshake", serverEndpoint), nil)
	req.Header.Set("x-client-id", cliKey)
	req.Header.Set("x-client-shared-key", strconv.Itoa(sharedKey))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	serverSharedKey, err := strconv.Atoi(resp.Header.Get("x-server-shared-key"))
	if err != nil {
		return -1, err
	}
	return serverSharedKey, nil
}

func init() {
	cliKey = generateCliKey()
}

func main() {
	fmt.Println("Start Diffe Hellman simulation")
	key, err := getKey()
	if err != nil {
		fmt.Printf("key fetch error=%v\n", err)
		return
	}

	fmt.Printf("fetched key, p=%v, g=%v\n", key.P, key.G)
	priv := rand.Intn(key.P)
	fmt.Printf("generate shared parameters: g=%v, priv=%v, p=%v\n", key.G, priv, key.P)
	shared := math.Mod(math.Pow(float64(key.G), float64(priv)), float64(key.P))
	fmt.Printf("client public key=%v\n", shared)
	serverSharedKey, err := handshake(int(shared))
	if err != nil {
		fmt.Printf("handshake err=%v\n", err)
	}
	fmt.Printf("server public key=%v\n", serverSharedKey)
	fmt.Printf("[client] symmetric key generate parameter: server public key=%v, client priv=%v, p=%v\n", serverSharedKey, priv, key.P)
	symmetricKey := math.Mod(math.Pow(float64(serverSharedKey), float64(priv)), float64(key.P))
	fmt.Printf("symmetric key=%v\n", symmetricKey)
}
