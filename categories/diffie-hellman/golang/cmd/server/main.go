package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

const (
	p = 97
	g = 5
)

var clientKeyMap = map[string]int{}

func getSharedKeys(w http.ResponseWriter, r *http.Request) {
	resp := map[string]any{
		"p": p,
		"g": g,
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func decodeAndGetSharedSecret(priv, clientSharedKey int) int {
	res := math.Mod(math.Pow(float64(clientSharedKey), float64(priv)), float64(p))
	return int(res)
}

func handshake(w http.ResponseWriter, r *http.Request) {
	// key should be in header
	clientId := r.Header.Get("x-client-id")
	clientSharedKey := r.Header.Get("x-client-shared-key")
	if clientId == "" || clientSharedKey == "" {
		err := fmt.Errorf("No client ID or shared key found.")
		fmt.Printf("handshake error=%v\n", err)
		return
	}
	clientShared, err := strconv.Atoi(clientSharedKey)
	if err != nil {
		fmt.Printf("handshake error=%v\n", err)
		return
	}
	privKey := rand.Intn(p)
	fmt.Printf("client public key (client id: %v): %v\n", clientId, clientShared)
	fmt.Printf("[server] symmetric key generate parameter: client public key=%v, server priv=%v, p=%v\n", clientShared, privKey, p)
	sharedSecret := decodeAndGetSharedSecret(privKey, clientShared)
	fmt.Printf("symmetric key for client id %v is: %v\n", clientId, sharedSecret)
	clientKeyMap[clientId] = sharedSecret
	fmt.Printf("generate shared parameters: g=%v, priv=%v, p=%v\n", g, privKey, p)
	serverSharedKey := int(math.Mod(math.Pow(float64(g), float64(privKey)), p))
	fmt.Printf("server public key=%v\n", serverSharedKey)
	w.Header().Set("x-server-shared-key", strconv.Itoa(int(serverSharedKey)))
	w.WriteHeader(http.StatusOK)
	fmt.Println("endhandshake===============")
}

func main() {
	fmt.Println("Diffe Hellman server simulation")
	fmt.Printf("p=%v, g=%v\n", p, g)
	http.HandleFunc("/key", getSharedKeys)
	http.HandleFunc("/handshake", handshake)
	err := http.ListenAndServe(":8091", nil)
	if err != nil {
		fmt.Printf("err=%v", err)
	}
}
