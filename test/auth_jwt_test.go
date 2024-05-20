package test

// import (
// 	"crypto/ecdsa"
// 	"crypto/elliptic"
// 	"crypto/rand"
// 	"fmt"
// 	"testing"

// 	jwtconfig "github.com/faridlan/auth-go/jwt_config"
// )

// func TestCreatetUserHash(t *testing.T) {

// 	pass := "secret123456"

// 	password, err := hashPassword(pass)
// 	if err != nil {
// 		panic(err)
// 	}

// 	user := &UserHash{
// 		Username: "john_doe",
// 		Password: password,
// 	}

// 	userResponse, err := CreateUserHash(user)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(userResponse)

// }

// func TestLoginJWT(t *testing.T) {
// 	username := "john_doe"
// 	pass := "zzzz"

// 	user, err := Login(username, pass)
// 	if err != nil {
// 		panic(err)
// 	}

// 	token, err := jwtconfig.GenerateJWT(username)
// 	if err != nil {
// 		panic(err)
// 	}

// 	UserResponse := UserLoginResponse{
// 		User: &UserResponse{
// 			Id:       user.ID,
// 			Username: user.Username,
// 		},
// 		Token: token,
// 	}

// 	fmt.Println("Id: " + UserResponse.User.Id)
// 	fmt.Println("Username: " + UserResponse.User.Username)
// 	fmt.Println("Token: " + UserResponse.Token)

// }

// func TestLoginJwt(t *testing.T) {

// 	Username := "nullhakim"
// 	Password := "secret1234567"

// 	user, err := FindUser(Username, Password)
// 	if err != nil {
// 		panic(err)
// 	}

// 	token, err := jwtconfig.GenerateJWT(user.Username)
// 	if err != nil {
// 		panic(err)
// 	}

// 	privateKey, err := jwtconfig.LoadPrivateKey("private.pem")
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Verify the generated token.
// 	claims, err := jwtconfig.VerifyToken(token, &privateKey.PublicKey)
// 	if err != nil {
// 		fmt.Println("Error verifying token:", err)
// 		return
// 	}

// 	fmt.Println("Verified Token Claims:", claims)

// }

// func TestPrivateKey(t *testing.T) {
// 	// Specify the desired file path for the private key
// 	filePath := "private.pem" // Replace with your preferred location

// 	_, err := jwtconfig.GenerateAndStorePrivateKey(filePath)
// 	if err != nil {
// 		fmt.Println("Error generating and storing private key:", err)
// 		return
// 	}

// 	fmt.Println("Private key successfully generated and stored in:", filePath)
// }

// func TestGenerateKey(t *testing.T) {

// 	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(privateKey)
// 	fmt.Println(privateKey.PublicKey)
// }

// func TestGeneratedJWT(t *testing.T) {

// 	// Generate a JWT token.
// 	token, err := jwtconfig.GenerateJWT("jhon_doe")
// 	if err != nil {
// 		fmt.Println("Error generating token:", err)
// 		return
// 	}

// 	fmt.Println("Generated Token:", token)

// 	// token := "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG4gZG9lIiwiZXhwIjoxNzE1MTg5ODM1fQ.zsNxzbjQ1nebFG1i9REsDzfFAPM0Gtud6ukHQHFzBMWVg-0MCfC3O-VgVnadnTPXbPlF0g91YPsEIIv3oANupg"
// 	privateKey, err := jwtconfig.LoadPrivateKey("private.pem")
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Verify the generated token.
// 	claims, err := jwtconfig.VerifyToken(token, &privateKey.PublicKey)
// 	if err != nil {
// 		fmt.Println("Error verifying token:", err)
// 		return
// 	}

// 	fmt.Println("Verified Token Claims:", claims)

// }
