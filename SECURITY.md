# Security Policy

We take the security of our services seriously. If you believe you have found a security vulnerability in the **ghosecorp-auth** repository, please read this document to learn how to report it.

## Supported Versions

Currently, security updates are actively provided for the following versions:

| Version | Supported          |
| ------- | ------------------ |
| v0.1.x  | :white_check_mark: |
| < v0.1  | :x:                |

## Reporting a Vulnerability

Please **do not** open public GitHub issues for security vulnerabilities. Instead, report them privately using the following process:

1. Send an email to **ghosecorp@gmail.com** (or contact the maintainers directly).
2. Include a detailed description of the vulnerability, steps to reproduce, and a proof of concept if available.
3. We will acknowledge receipt of your vulnerability report within 48 hours and work with you to resolve the issue promptly.

## Security Best Practices for Deployment

When deploying this service in production environments:
* **HTTPS**: Force all traffic over TLS (HTTPS/WSS).
* **Secrets Management**: Do not commit private configuration files (`.env`) to version control. Use secure vault services or environment variable injection.
* **Database Access**: Limit database permissions for the auth database user to only the schemas and tables it requires.
* **Token Expirations**: Set short expiration windows on access tokens and rotate refresh tokens securely.
