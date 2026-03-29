# ghosecorp-auth
ghosecorp-auth

# GHOSECORP AUTH SERVICE

A scalable, microservices-based authentication system supporting:

* Email and password authentication
* OAuth-based login (Google, Microsoft, etc.)
* Multi-Factor Authentication (TOTP)
* OAuth provider capabilities for third-party integrations

---

## Architecture Overview

This repository contains multiple services:

```
auth/
├── auth-api        # Core authentication service (Go)
├── auth-mailer     # Email service (notifications, OTP, etc.)
├── setup           # Database setup scripts (Python)
```

---

## Tech Stack

* **Backend Services:** Go
* **Database Setup:** Python
* **Database:** PostgreSQL

---

## Database Setup

Database initialization is handled via Python scripts.

---

### 1. Install Python Dependencies

```bash
cd setup
pip install -r requirements.txt
```

---

### 2. Configure Environment Variables

Create a `.env` file in the root directory (`auth/`):

```
auth_postgres_sql_username=postgres
auth_postgres_sql_password=yourpassword
auth_postgres_sql_db=ghose_cloud_auth_db
auth_postgres_sql_host=127.0.0.1
auth_postgres_sql_port=5432
```

### Notes

* Do not use quotes
* Do not add spaces around `=`
* Ensure `.env` is included in `.gitignore`

---

### 3. Run Database Setup

```bash
cd setup
python db_setup.py
```

---

### Alternative: Bash Script

```bash
cd setup
chmod +x setup_db.sh
./setup_db.sh
```

---

## Running Services

### Auth API (Go)

```bash
cd auth-api
go mod tidy
go run main.go
```

---

### Mailer Service

```bash
cd auth-mailer
go mod tidy
go run main.go
```

---

## Features

### Authentication

* Email and password login
* Session management

### OAuth Login

* Google
* Microsoft
* Extensible to other providers

### Multi-Factor Authentication

* TOTP-based authentication
* Compatible with Google Authenticator and Microsoft Authenticator

### OAuth Provider

* Supports third-party applications
* Authorization Code flow with PKCE
* Token issuance and validation

---

## Security Considerations

* Do not commit `.env` files
* Always hash passwords and tokens
* Use HTTPS in production environments

---

## Future Enhancements

* API gateway integration
* Rate limiting and abuse protection
* Distributed tracing and monitoring
* Service-to-service authentication

---

## Contributors

* **Anubroto Ghose** — [https://github.com/anubrotoGhose](https://github.com/anubrotoGhose)

---
