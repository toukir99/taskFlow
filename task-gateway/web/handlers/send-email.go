package handlers

import (
	"context"
	"fmt"
	"math/rand"
	"task-gateway/db"
	"time"
)

func GenerateOTP() string {
	otp := rand.Intn(900000) + 100000
	return fmt.Sprintf("%06d", otp)
}
func SendEmail(userEmail string) error {
	otp := GenerateOTP()

	const (
		otpPrefix     = "otp:"
		otpExpiration = 5 * time.Minute
	)

	otpKey := otpPrefix + userEmail

	err := db.GetRedisClient().Set(context.Background(), otpKey, otp, otpExpiration).Err()
	if err != nil {
		return fmt.Errorf("failed to store OTP: %v".err)
	}
}
