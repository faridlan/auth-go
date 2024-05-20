package helper

import (
	"golang.org/x/crypto/bcrypt"
)

// func generateSalt() []byte {
// 	salt := make([]byte, 16)
// 	rand.Read(salt)
// 	return salt
// }

func HashPassword(password string) ([]byte, error) {
	// Append the password to the salt to create a combined byte slice

	// Use bcrypt to generate a hashed password from the combined byte slice
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Return the hashed password
	return hashedPassword, nil
}

func ComparePasswords(hashedPassword []byte, password string) bool {
	// Use bcrypt to compare the hashed password with the provided password
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)) == nil
}

// func Login(username, password string) (*UserHash, error) {
// 	// Find the user by username
// 	user, err := FindUserHash(username)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to find user '%s': %w", username, err)
// 	}

// 	// Hash the provided password with the retrieved salt
// 	providedHash, err := hashPassword(password)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to hash password: %w", err)
// 	}

// 	// Debugging output
// 	fmt.Println("Stored Username:", user.Username)
// 	fmt.Println("Stored Password:", string(user.Password))
// 	fmt.Println("Provided Password:", string(providedHash))

// 	// Compare the hashed passwords
// 	if !comparePasswords(user.Password, password) {
// 		return nil, errors.New("password mismatch")
// 	}

// 	fmt.Println("Authentication successful for user:", user.Username)
// 	return user, nil
// }
