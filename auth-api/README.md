# Ghosecorp Auth API

Go service for first-party authentication in `ghosecorp-auth`.

The current implementation supports:

- Basic user signup
- Basic email/password login
- HttpOnly cookie-based sessions
- `GET /v1/me` for the current authenticated user

This service is intentionally layered so OAuth2, MFA, email verification, password resets, JWT/JWE access tokens, and service-to-service authentication can be added later without rewriting the basic auth flow.

## Architecture

```text
handler/httpHandler -> usecase -> repository -> PostgreSQL
                    -> security
                    -> domain
```

## Project Structure

```text
auth-api/
  cmd/
    server/
      main.go                  # App bootstrap and dependency wiring
  internal/
    config/
      config.go                # Environment/config loading
    domain/
      user.go                  # Core user model
      session.go               # Session domain types
    handler/
      httpHandler/
        auth.go                # Signup and login handlers
        me.go                  # Current-user handler
        middleware.go          # Session cookie authentication
        routes.go              # Route registration
    repository/
      credential.go            # credentials table access
      session.go               # sessions table access
      user.go                  # users table access
    security/
      cookies.go               # Session cookie helpers
      password.go              # Password hashing and verification
      token.go                 # Session token generation and hashing
    usecase/
      auth.go                  # Signup/login business flow
      session.go               # Session validation flow
  go.mod
  go.sum
  README.md
```

## Database

The service uses PostgreSQL.

It expects the schema from the root `setup` folder to already exist, especially:

- `users`
- `credentials`
- `sessions`

The local database used during development is:

```text
ghose_cloud_auth_db
```

## Environment

Set `DATABASE_URL` before running the service:

```bash
export DATABASE_URL="postgres://<username>:<password>@<host>:<port>/<db_name>?sslmode=disable"
```

Optional values:

```bash
export PORT=8001
export APP_ENV=dev
```

When `APP_ENV=prod`, session cookies are marked `Secure`.

## Install Dependencies

```bash
go mod tidy
```

## Run

```bash
go run cmd/server/main.go
```

The service starts on:

```text
http://localhost:8001
```

## API

### Signup

```bash
curl -i -X POST http://localhost:8001/v1/auth/signup \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"Password@1234"}'
```

### Login

```bash
curl -i -c cookies.txt -X POST http://localhost:8001/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"Password@1234"}'
```

On success, the API sets an HttpOnly session cookie named:

```text
gc_session
```

### Current User

```bash
curl -i -b cookies.txt http://localhost:8001/v1/me
```

## Security Notes

- Passwords are stored as hashes, never plain text.
- Session tokens are sent to the client as HttpOnly cookies.
- Only a hash of the session token should be stored in the database.
- Internal database IDs should not be exposed in JSON responses.
- Use HTTPS in production so `Secure` cookies work correctly.

## Future Additions

Planned extensions can build on the current structure:

- Email verification using `auth-mailer`
- Password reset flow
- Logout and session revocation
- MFA/TOTP
- OAuth login with Google/Microsoft
- OAuth2 provider endpoints with Authorization Code + PKCE
- JWT/JWE support for service and API access tokens
