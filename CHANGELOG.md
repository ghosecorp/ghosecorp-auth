# Changelog

All notable changes to the **ghosecorp-auth** project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2026-05-31

### Added
- **Mailer Microservice (`auth-mailer`)**: Created a standalone Go service to handle transactional emails integrated with the Brevo SMTP API.
- **Database Setup Script**: Added Python setup scripts and SQL schemas supporting standard users, credentials, multi-factor authentication (MFA), and PKCE-based OAuth providers.
- **Verification Scripts**: Added `./test-mailer.sh` script to verify email API requests.
- **Project Structure**: Set up core domain models, configurations, and clean architecture directories inside the `auth-api` service.
