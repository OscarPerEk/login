# VoiceLine Task 2

This repo creates a small app with the following tools and frameworks:

- Gin: web framework
- Auth0: authentication provider
- html/template: to handle structs in html
- sqlite3: database
- GORM: go library for interacting with database

A lot of the code come from [Auth0 Quickstart](https://auth0.com/docs/quickstart/webapp/golang) example app.
On their website they provide a small go app that has a login/logout page.
I used this app as the backbone and saved the user information to a sqlite3 database.
I then created a new page which simply lists the content of the database.

## Running the App

To run the app, make sure you have **go** installed.

Sign up to [Auth0](https://auth0.com) and provide your Auth0 credentials in `.env`.

```bash
# .env

AUTH0_CLIENT_ID={yourClientId}
AUTH0_DOMAIN={yourDomain}
AUTH0_CLIENT_SECRET={yourClientSecret}
AUTH0_CALLBACK_URL=http://localhost:3000/callback
```

Once you've set your Auth0 credentials in the `.env` file, run `go mod vendor` to download the Go dependencies.

Run `go run main.go` to start the app and navigate to [http://localhost:3000/](http://localhost:3000/).
