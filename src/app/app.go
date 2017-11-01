package app

import (
	"os"
	"time"

	"github.com/fragmenta/assets"
	"github.com/fragmenta/server/config"
	"github.com/fragmenta/server/log"

	"github.com/bitcubate/cryptojournal/src/lib/mail"
	"github.com/bitcubate/cryptojournal/src/lib/mail/adapters/sendgrid"
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

	finalString := "http://www.gravatar.com/avatar/" + hex.EncodeToString(finalBytes) + "?d=identicon"

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

// appAssets holds a reference to our assets for use in asset setup
var appAssets *assets.Collection

// Setup sets up our application
func Setup() {

	// Setup log
	err := SetupLog()
	if err != nil {
		println("failed to set up logs %s", err)
		os.Exit(1)
	}

	// Log server startup
	msg := "Starting server"
	if config.Production() {
		msg = msg + " in production"
	}

	log.Info(log.Values{"msg": msg, "port": config.Get("port")})
	defer log.Time(time.Now(), log.Values{"msg": "Finished loading server"})

	// Set up external service interfaces (twitter)
	SetupServices()

	// Set up our assets
	SetupAssets()

	// Setup our view templates
	SetupView()

	// Setup our database
	SetupDatabase()

	// Setup our authentication and authorisation
	SetupAuth()

	// Setup our router and handlers
	SetupRoutes()

	// Setup mail from config
	SetupMail()

}

// SetupLog sets up logging
func SetupLog() error {

	// Set up a stderr logger with time prefix
	logger, err := log.NewStdErr(log.PrefixDateTime)
	if err != nil {
		return err
	}
	log.Add(logger)

	// Set up a file logger pointing at the right location for this config.
	fileLog, err := log.NewFile(config.Get("log"))
	if err != nil {
		return err
	}
	log.Add(fileLog)

	return nil
}

// SetupMail sets us up to send mail via sendgrid (requires key).
func SetupMail() {
	mail.Production = config.Production()
	mail.Service = sendgrid.New(config.Get("mail_from"), config.Get("mail_secret"))
}
