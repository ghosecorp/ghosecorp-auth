auth-api/
│
├── cmd/
│   └── server/
│       └── main.go              # Entry point (bootstraps everything)
│
├── internal/                   # Private application code (core logic)
│   │
│   ├── config/
│   │   └── config.go           # Env loading, config structs
│   │
│   ├── domain/                 # Business models (pure logic)
│   │   ├── user.go
│   │   └── session.go
│   │
│   ├── repository/             # DB layer (Postgres queries)
│   │   ├── user_repository.go
│   │   └── session_repository.go
│   │
│   ├── usecase/                # Business logic (core rules)
│   │   ├── auth_usecase.go     # login/signup logic
│   │   └── session_usecase.go
│   │
│   ├── delivery/               # Transport layer
│   │   └── http/
│   │       ├── handler.go      # route handlers
│   │       ├── routes.go
│   │       └── middleware.go   # JWT, logging, recovery
│   │
│   ├── pkg/                    # Shared internal utilities
│   │   ├── jwt/
│   │   │   └── jwt.go          # token generation/validation
│   │   ├── hash/
│   │   │   └── bcrypt.go       # password hashing
│   │   └── logger/
│   │       └── logger.go
│   │
│   └── infrastructure/         # External systems
│       ├── db/
│       │   └── postgres.go     # DB connection
│       └── cache/
│           └── redis.go        # session/cache (optional)
│
├── api/                        # API contracts (optional but pro-level)
│   ├── openapi.yaml
│   └── proto/                  # if using gRPC later
│
├── migrations/                 # DB migrations
│   ├── 001_create_users.sql
│   └── 002_create_sessions.sql
│
├── scripts/
│   └── seed.go                 # optional test data
│
├── deployments/
│   ├── Dockerfile
│   ├── docker-compose.yml
│   └── k8s/                    
│
├── .env
├── .env.example
├── go.mod
├── go.sum
└── README.md

psql -h 127.0.0.1 -U adminuser -d ghose_cloud_auth_db -c "\dt"
