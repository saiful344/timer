package helper 

import(
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func Encrypt(key, data []byte)([]byte, error){
	blockcip ,  err := aes.NewCipher(key)
	if err != nil{
		return nil, err
	}

	gcm, err := cipher.NewGCM(blockcip)

	if err != nil{
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = rand.Read(nonce); err != nil{
		retunr nil, err
	}
	cipherText := gcm.Seal(nonce, nonce, data, nil)

	return cipherText, nil
}	