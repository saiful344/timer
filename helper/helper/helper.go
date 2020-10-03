package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Message struct {
	Msg string `json:'msg'`
}

func Hashing(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(Hashing(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func Decrypt(data []byte, passphrase string) []byte {
	key := []byte(Hashing(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

// func main() {

// 	for{
// 		pwd := getPwd()
// 		hash := hashAndSalt(pwd)

// 		pwd2 := getPwd();
// 		pwdMatch:=  comparePasswords(hash, pwd2)

// 		fmt.Println("password mastch?",pwdMatch)
// 	}

// testing one
// fmt.Println("Starting the application...")
// ciphertext := Encrypt([]byte("Hello World"), "password")
// fmt.Printf("Encrypted: %x\n", ciphertext)
// plaintext := Decrypt(ciphertext, "password")
// fmt.Printf("Decrypted: %s\n", plaintext)
// }
func getPwd() []byte {
	fmt.Println("isi oy")

	var pwd string
	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Println(err)
	}
	return []byte(pwd)
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func ComparePasswords(hasher string, plainpwd []byte) bool {
	byteHash := []byte(hasher)
	err := bcrypt.CompareHashAndPassword(byteHash, plainpwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func Message_Json(m string) string {
	var msg Message
	msg.Msg = m
	data, _ := json.Marshal(msg)
	// return
	return string(data)
}
