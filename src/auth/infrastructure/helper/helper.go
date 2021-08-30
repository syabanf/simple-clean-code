package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"reflect"
	"strings"
	"time"

	"sagara-test/src/common/constant"

	jwt "github.com/dgrijalva/jwt-go"
)

func PhoneNumberPrefix(phone string) (validPhone string, err error) {
	validPhone = phone

	if !(strings.HasPrefix(phone, "0") ||
		strings.HasPrefix(phone, "62") ||
		strings.HasPrefix(phone, "+62")) {
		err = fmt.Errorf("invalid prefix")
	}

	zeroPrefix := strings.HasPrefix(phone, "0")
	if zeroPrefix == true {
		validPhone = strings.Replace(phone, "0", "+62", 1)
	}
	noPlusPrefix := strings.HasPrefix(phone, "62")
	if noPlusPrefix == true {
		validPhone = strings.Replace(phone, "62", "+62", 1)
	}
	return
}

func GetNamedStruct(data interface{}) []string {
	var value []string
	val := reflect.ValueOf(data)
	for i := 0; i < val.Type().NumField(); i++ {
		if val.Type().Field(i).Tag.Get("db") == "" {
			continue
		}
		value = append(value, val.Type().Field(i).Tag.Get("db"))
	}
	return value
}

func CreateHash(key string) (hash string) {
	hasher := md5.New()
	hasher.Write([]byte(key))
	hash = hex.EncodeToString(hasher.Sum(nil))
	return
}

func Encrypt(data []byte, passPhrase string) (chipered []byte, err error) {
	block, _ := aes.NewCipher([]byte(CreateHash(passPhrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println("Error chiper: ", err)
		return
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Println("Error gcm nonce: ", err)
		return
	}

	chipered = gcm.Seal(nonce, nonce, data, nil)
	return
}

func Decrypt(data []byte, passPhrase string) (plain []byte, err error) {
	key := []byte(CreateHash(passPhrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println("Error decrypt: ", err)
		return
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println("Error decrypt gcm: ", err)
		return
	}

	nonceSize := gcm.NonceSize()
	nonce, chiperTxt := data[:nonceSize], data[nonceSize:]
	plain, err = gcm.Open(nil, nonce, chiperTxt, nil)
	if err != nil {
		log.Println("Error decrypt gcm open: ", err)
		return
	}
	return
}

func GenerateToken(id string, roleId, isVerified int64) (tokens map[string]string, expireAt int64, err error) {
	// Create new token
	expireAt = time.Now().Add(time.Minute * 120).Unix()
	claims := AuthClaims{
		ID:         id,
		RoleID:     roleId,
		IsVerified: isVerified,
		StandardClaims: jwt.StandardClaims{
			Issuer:    constant.AppName,
			ExpiresAt: expireAt,
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString([]byte(constant.JWTKey))
	if err != nil {
		return
	}

	// Create refresh token
	refreshClaims := AuthClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			Issuer:    constant.AppName,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	fresh := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		refreshClaims,
	)

	refreshToken, err := fresh.SignedString([]byte(constant.JWTKey))
	if err != nil {
		return
	}

	return map[string]string{
		"access_token":  signedToken,
		"refresh_token": refreshToken,
	}, expireAt, nil
}
