package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

// TODO Should move to his own repo/project that should submit only final hash
type Faktur struct {
	privateKey *ecdsa.PrivateKey
	Address    common.Address
}

// TODO
func NewFaktur(p *ecdsa.PrivateKey) *Faktur {
	addr := crypto.PubkeyToAddress(p.PublicKey)
	return &Faktur{privateKey: p,
		Address: addr}
}

// TODO
func (f *Faktur) Serialize() []byte {
	return nil
}

func generateToken(w http.ResponseWriter, r *http.Request) {
	nToken, ok := r.URL.Query()["token_number"]
	if !ok || len(nToken) == 0 {
		http.Error(w, "token_number not found in parameters", 422)
		return
	}
	n, err := strconv.ParseInt(nToken[0], 10, 0)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	var fakturs []*Faktur
	for i := 0; i < int(n); i++ {
		privateKey, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fakturs = append(fakturs, NewFaktur(privateKey))
	}
	var hashList string
	for k, v := range fakturs {
		log.Printf("%d: %+v", k, v)
		tokenHash := common.Bytes2Hex(crypto.Keccak256(v.Address.Bytes()))
		hashList += "hash=" + tokenHash + "&"
	}
	log.Println(hashList)
	req, err := http.NewRequest("GET", "http://127.0.0.1:8090/save?"+hashList, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	fmt.Fprintf(w, string(body))
	/*
		go func() {
			//print
			//Targethash
			//Proofs
			//MerkleRoot
			//TransactionHash
			//privatekey
		}()
		// for each generate token
	*/
}

func main() {
	http.HandleFunc("/generate", generateToken) // set router
	err := http.ListenAndServe(":9090", nil)    // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
