package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeUntil(t *testing.T) {

	// Assume the token expires in 30 minutes from now
	tokenExpiry := time.Now().Add(11 * time.Minute)

	// Calculate the duration until the token expires
	timeUntilExpiry := time.Until(tokenExpiry)

	fmt.Printf("Time until token expires: %v\n", timeUntilExpiry)

	// Check if the token expires within the next 10 minutes
	if timeUntilExpiry < 10*time.Minute {
		fmt.Println("Token is close to expiring, refresh it.")
	} else {
		fmt.Println("Token is still valid.")
	}

}
