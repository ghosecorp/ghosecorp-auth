import psycopg2
from dotenv import load_dotenv
import os

load_dotenv()

DB_CONFIG = {
    "dbname": os.getenv("auth_postgres_sql_db", "auth"),
    "user": os.getenv("auth_postgres_sql_username", "postgres"),
    "password": os.getenv("auth_postgres_sql_password", "password123"),
    "host": os.getenv("auth_postgres_sql_host", "127.0.0.1"),
    "port": int(os.getenv("auth_postgres_sql_port", 5432))
}


SCHEMA_SQL = """
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- =========================
-- USERS
-- =========================
CREATE TABLE IF NOT EXISTS users (
    user_id BIGSERIAL PRIMARY KEY,
    public_id UUID DEFAULT gen_random_uuid() UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    phone_number VARCHAR(20) UNIQUE,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- =========================
-- USER DETAILS
-- =========================
CREATE TABLE IF NOT EXISTS user_details (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL UNIQUE REFERENCES users(user_id) ON DELETE CASCADE,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    tax_address TEXT,
    permanent_address TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- =========================
-- CREDENTIALS
-- =========================
CREATE TABLE IF NOT EXISTS credentials (
    user_id BIGINT PRIMARY KEY REFERENCES users(user_id) ON DELETE CASCADE,
    password_hash TEXT NOT NULL,
    password_salt TEXT,
    last_changed TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- =========================
-- SESSIONS
-- =========================
CREATE TABLE IF NOT EXISTS sessions (
    session_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    refresh_token_hash TEXT NOT NULL,
    user_agent TEXT,
    ip_address TEXT,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- =========================
-- EMAIL VERIFICATIONS
-- =========================
CREATE TABLE IF NOT EXISTS email_verifications (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    token TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- =========================
-- PASSWORD RESETS
-- =========================
CREATE TABLE IF NOT EXISTS password_resets (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    token TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    used BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- =========================
-- ROLES
-- =========================
CREATE TABLE IF NOT EXISTS roles (
    role_id BIGSERIAL PRIMARY KEY,
    role_name VARCHAR(50) UNIQUE NOT NULL
);

-- =========================
-- USER ROLES
-- =========================
CREATE TABLE IF NOT EXISTS user_roles (
    user_id BIGINT REFERENCES users(user_id) ON DELETE CASCADE,
    role_id BIGINT REFERENCES roles(role_id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, role_id)
);

-- =========================
-- OAUTH ACCOUNTS
-- =========================
CREATE TABLE IF NOT EXISTS oauth_accounts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    provider VARCHAR(50) NOT NULL,
    provider_user_id TEXT NOT NULL,
    access_token TEXT,
    refresh_token TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(provider, provider_user_id)
);

-- =========================
-- MFA
-- =========================
CREATE TABLE mfa (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,

    provider VARCHAR(50) NOT NULL, 
    -- 'totp', 'sms', 'email'

    secret TEXT,       -- for TOTP (Google/Microsoft Authenticator)
    phone_number TEXT, -- for SMS
    email TEXT,        -- for email OTP

    enabled BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    UNIQUE(user_id, provider)
);

-- =========================
-- OAUTH PROVIDER CLIENTS
-- =========================
CREATE TABLE IF NOT EXISTS oauth_provider_clients (
    client_id TEXT PRIMARY KEY,
    client_secret TEXT NOT NULL,
    client_name VARCHAR(150),
    is_confidential BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- =========================
-- CLIENT REDIRECT URIS
-- =========================
CREATE TABLE IF NOT EXISTS oauth_provider_client_redirects (
    id BIGSERIAL PRIMARY KEY,
    client_id TEXT NOT NULL REFERENCES oauth_provider_clients(client_id) ON DELETE CASCADE,
    redirect_uri TEXT NOT NULL,
    UNIQUE(client_id, redirect_uri)
);

-- =========================
-- AUTHORIZATION CODES
-- =========================
CREATE TABLE IF NOT EXISTS oauth_provider_authorization_codes (
    code TEXT PRIMARY KEY,
    client_id TEXT NOT NULL REFERENCES oauth_provider_clients(client_id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    redirect_uri TEXT NOT NULL,
    code_challenge TEXT,
    code_challenge_method VARCHAR(10),
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- =========================
-- ACCESS TOKENS
-- =========================
CREATE TABLE IF NOT EXISTS oauth_provider_access_tokens (
    access_token_hash TEXT PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    client_id TEXT NOT NULL REFERENCES oauth_provider_clients(client_id) ON DELETE CASCADE,
    expires_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- =========================
-- REFRESH TOKENS
-- =========================
CREATE TABLE IF NOT EXISTS oauth_provider_refresh_tokens (
    refresh_token_hash TEXT PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    client_id TEXT NOT NULL REFERENCES oauth_provider_clients(client_id) ON DELETE CASCADE,
    expires_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- =========================
-- CONSENTS
-- =========================
CREATE TABLE IF NOT EXISTS oauth_provider_consents (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    client_id TEXT NOT NULL REFERENCES oauth_provider_clients(client_id) ON DELETE CASCADE,
    scopes TEXT,
    granted BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, client_id)
);

-- =========================
-- INDEXES
-- =========================
CREATE INDEX IF NOT EXISTS idx_sessions_user_id ON sessions(user_id);
CREATE INDEX IF NOT EXISTS idx_sessions_expires_at ON sessions(expires_at);

CREATE INDEX IF NOT EXISTS idx_email_verifications_user_id ON email_verifications(user_id);
CREATE INDEX IF NOT EXISTS idx_password_resets_user_id ON password_resets(user_id);

CREATE INDEX IF NOT EXISTS idx_oauth_user_id ON oauth_accounts(user_id);

CREATE INDEX idx_mfa_user_id ON mfa(user_id);

CREATE INDEX IF NOT EXISTS idx_oauth_redirects_client ON oauth_provider_client_redirects(client_id);
CREATE INDEX IF NOT EXISTS idx_oauth_auth_codes_client ON oauth_provider_authorization_codes(client_id);
CREATE INDEX IF NOT EXISTS idx_oauth_access_tokens_user ON oauth_provider_access_tokens(user_id);
CREATE INDEX IF NOT EXISTS idx_oauth_access_tokens_client ON oauth_provider_access_tokens(client_id);
"""


def setup_database():
    try:
        conn = psycopg2.connect(**DB_CONFIG)
        conn.autocommit = True
        cur = conn.cursor()

        print("Creating tables and indexes...")

        cur.execute(SCHEMA_SQL)

        print("Database setup completed successfully!")

        cur.close()
        conn.close()

    except Exception as e:
        print("Error:", e)


if __name__ == "__main__":
    setup_database()
    