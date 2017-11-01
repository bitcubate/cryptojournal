// Package users represents the user resource
package users

import (
	"time"

	"github.com/bitcubate/cryptojournal/src/lib/resource"
	"github.com/bitcubate/cryptojournal/src/lib/status"
    "crypto/md5"
	"encoding/hex"
	"io"
	"strings"
)

// This function takes in an email address as a parameter and returns a complete url to the image.
// for example: GetImg("example@example.com") will return this: http://www.gravatar.com/avatar/563dd2c62e66d42a3b4454c410a0b3ab?d=identicon
func GetImg(email string) string {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	h := md5.New()
	io.WriteString(h, email)

	finalBytes := h.Sum(nil)

	finalString := "http://www.gravatar.com/avatar/" + hex.EncodeToString(finalBytes) + "?d=retro"

	return finalString
}

// This function takes in an email address as a parameter and returns just the hash, which you can use to your liking.
func GetHash(email string) string {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	h := md5.New()
	io.WriteString(h, email)

	finalBytes := h.Sum(nil)

	finalString := hex.EncodeToString(finalBytes)

	return finalString
}

// User handles saving and retreiving users from the database
type User struct {
	// resource.Base defines behaviour and fields shared between all resources
	resource.Base

	// status.ResourceStatus defines a status field and associated behaviour
	status.ResourceStatus

	Email   string
	Name    string
	Points  int64
	Role    int64
	Summary string
	Text    string
	Title   string

	PasswordHash    string
	PasswordResetAt time.Time
}
