package jwtconfig

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faridlan/auth-go/helper"
	"github.com/faridlan/auth-go/model/web"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func LoadPrivateKey(filePath string) (*ecdsa.PrivateKey, error) {
	keyBytes, err := os.ReadFile(filePath) // Read private key file
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %v", err)
	}

	block, _ := pem.Decode(keyBytes) // Decode PEM block (ignoring errors for brevity)
	if block == nil {
		return nil, errors.New("failed to decode private key PEM block")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes) // Parse ECDSA key
	if err != nil {
		return nil, fmt.Errorf("failed to parse ECDSA private key: %v", err)
	}

	return privateKey, nil
}

func GenerateJWT(claim *web.Claim) (string, error) {

	claim.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 10))

	config, err := helper.GetEnv()
	if err != nil {
		panic(err)
	}
	path := config.GetString("PRIVATE_KEY")
	privateKey, err := LoadPrivateKey(path)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	return token.SignedString(privateKey)

}

func VerifyToken(tokenString string, publicKey *ecdsa.PublicKey) (*web.Claim, *jwt.Token, error) {

	claim := &web.Claim{}

	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method is ES256.
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the public key for verification.
		return publicKey, nil
	})

	if err != nil {
		return nil, nil, err
	}

	if claims, ok := token.Claims.(*web.Claim); ok && token.Valid {
		return claims, token, nil
	}

	return nil, nil, errors.New("invalid token")

}

func GenerateAndStorePrivateKey(filePath string) (*ecdsa.PrivateKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	// Marshal the private key to PEM format (using DER encoding for ECDSA)
	derStream, err := x509.MarshalECPrivateKey(privateKey) // Import from crypto/x509
	if err != nil {
		return nil, fmt.Errorf("failed to marshal EC private key to DER: %v", err)
	}

	// Create a pem.Block with the DER encoded data
	pemBlock := &pem.Block{
		Type:  "EC PRIVATE KEY", // Adjusted for ECDSA (more accurate)
		Bytes: derStream,
	}

	pemBytes := pem.EncodeToMemory(pemBlock)
	if pemBytes == nil {
		return nil, errors.New("failed to encode private key to PEM format")
	}

	// Write the PEM encoded key to a file with appropriate permissions
	err = os.WriteFile(filePath, pemBytes, 0600) // Restrict access to owner only
	if err != nil {
		return nil, fmt.Errorf("failed to write private key to file: %v", err)
	}

	return privateKey, nil
}

func ParseJwtAuth(ctx *fiber.Ctx) (string, error) {

	auth := ctx.Get("Authorization")
	authString := auth[7:]
	// Parse the token without validating the signature
	token, _, err := new(jwt.Parser).ParseUnverified(authString, jwt.MapClaims{})
	if err != nil {
		log.Fatalf("Error parsing token: %v", err)
	}

	// Extract the claims
	claims := token.Claims.(jwt.MapClaims)
	if user, ok := claims["user"].(map[string]interface{}); ok {
		if whitelist, ok := user["whitelist"].(string); ok {
			return whitelist, nil
		} else {
			log.Fatalf("whitelist not found or not a string in user map")
		}
	} else {
		log.Fatalf("user not found or not a map in claims")
	}

	return "", nil

}
