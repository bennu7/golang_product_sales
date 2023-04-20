package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bennu7/golang_product_sales/config/dotenv"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

var TokenKey = dotenv.GetKeyString("TOKEN_KEY")

func ValidateToken(tokenString string) (*Payload, error) {
	parse, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(TokenKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parse.Claims.(jwt.MapClaims)
	if !ok || !parse.Valid {
		return nil, errors.New("token invalid")
	}

	getPayload := claims["payload"]
	var payloadToken = Payload{}

	payloadByte, err := json.Marshal(getPayload)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(payloadByte, &payloadToken)
	if err != nil {
		return nil, err
	}
	fmt.Println("final payloadToken => ", payloadToken)

	return &payloadToken, nil
}

func GenerateToken(payload *Payload) (string, error) {
	TokenExpired, err := strconv.Atoi(dotenv.GetKeyString("TOKEN_EXPIRED"))
	if err != nil {
		return "", err
	}
	duration := time.Duration(TokenExpired)

	payloads := Payload{
		UserId: payload.UserId,
		RoleId: payload.RoleId,
		Start:  time.Now().Unix(),
		Exp:    time.Now().Add(time.Minute * duration).Unix(),
	}

	claims := jwt.MapClaims{
		"payload": payloads,
		"exp":     payloads.Exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(TokenKey))
	if err != nil {
		return "", err
	}

	return signedString, nil
}
