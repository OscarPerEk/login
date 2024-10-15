package callback

import (
	"01-Login/platform/authenticator"
	"01-Login/web/app/types"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ProfileType map[string]interface{}

func MarshalProfile(user *types.User, token ProfileType) *types.User {
	if name, ok := token["name"].(string); ok {
		user.Name = name
	}
	if name, ok := token["given_name"].(string); ok {
		user.GivenName = name
	}
	if name, ok := token["family_name"].(string); ok {
		user.FamilyName = name
	}
	if name, ok := token["nickname"].(string); ok {
		user.Nickname = name
	}
	if name, ok := token["picture"].(string); ok {
		user.Picture = name
	}
	return user
}

// Handler for our callback.
func Handler(auth *authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		if ctx.Query("state") != session.Get("state") {
			ctx.String(http.StatusBadRequest, "Invalid state parameter.")
			return
		}

		// Exchange an authorization code for a token.
		token, err := auth.Exchange(ctx.Request.Context(), ctx.Query("code"))
		if err != nil {
			ctx.String(http.StatusUnauthorized, "Failed to convert an authorization code into a token.")
			return
		}

		idToken, err := auth.VerifyIDToken(ctx.Request.Context(), token)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Failed to verify ID Token.")
			return
		}

		var profile ProfileType
		if err := idToken.Claims(&profile); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		session.Set("access_token", token.AccessToken)
		session.Set("profile", profile)
		if err := session.Save(); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Save user information to database
		db, err := gorm.Open(sqlite.Open("app_db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		// Migrate the schema
		db.AutoMigrate(&types.User{})

		// Create
		var user types.User
		db.Create(MarshalProfile(&user, profile))

		// Redirect to logged in page.
		ctx.Redirect(http.StatusTemporaryRedirect, "/user")
	}
}
