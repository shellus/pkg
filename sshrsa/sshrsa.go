package sshrsa

import (
	"io/ioutil"
	"github.com/shellus/pkg/sutil"
	"encoding/pem"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
)
// 加密
func Encrypt(origData []byte) ([]byte, error) {
	buf, err := ioutil.ReadFile(sutil.HomeDir()+`/.ssh/id_rsa`)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(buf)

	r, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.EncryptPKCS1v15(rand.Reader, &r.PublicKey, origData)
}

// 解密
func Decrypt(ciphertext []byte) ([]byte, error) {
	buf, err := ioutil.ReadFile(sutil.HomeDir()+`/.ssh/id_rsa`)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(buf)

	r, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, r, ciphertext)
}