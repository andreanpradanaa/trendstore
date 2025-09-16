Layered + DD

TRENDSTORE/
├── cmd/
│   └── main.go                      # Entry point aplikasi
├── internal/
│   ├── domain/                      # CORE BUSINESS LOGIC
│   │   ├── product/                 # Domain product
│   │   │   ├── entity.go            # Struct dan methods bisnis
│   │   │   ├── value_object.go      # Price, Stock, dll
│   │   │   └── repository.go        # Interface repository
│   │   ├── order/                   # Domain order
│   │   │   ├── entity.go
│   │   │   ├── value_object.go
│   │   │   └── repository.go
│   │   └── user/                    # Domain user
│   │       ├── entity.go
│   │       ├── value_object.go
│   │       └── repository.go
│   │
│   ├── application/                 # USE CASES & COORDINATION
│   │   ├── product/
│   │   │   ├── service.go           # ProductService (organisir workflow)
│   │   │   ├── dto/                 # Request/Response objects
│   │   │   └── mapper/              # Mapping DTO <> Entity
│   │   ├── order/
│   │   │   ├── service.go           # OrderService
│   │   │   ├── dto/
│   │   │   └── mapper/
│   │   └── user/
│   │       ├── service.go           # UserService
│   │       ├── dto/
│   │       └── mapper/
│   │
│   └── infrastructure/              # TECHNICAL IMPLEMENTATIONS
│       ├── http/                    # Delivery layer
│       │   ├── product/
│       │   │   └── handler.go       # HTTP handlers (Gin/Echo)
│       │   ├── order/
│       │   │   └── handler.go
│       │   ├── user/
│       │   │   └── handler.go
│       │   └── middleware/          # Auth, logging, etc
│       │
│       ├── persistence/             # Data access implementations
│       │   ├── mysql/               # MySQL implementations
│       │   │   ├── product_repository.go
│       │   │   ├── order_repository.go
│       │   │   └── user_repository.go
│       │   └── redis/               # Cache implementations
│       │
│       ├── payment/                 # External services
│       │   ├── gateway_interface.go
│       │   ├── midtrans_gateway.go
│       │   └── stripe_gateway.go
│       │
│       └── messaging/               # Events & notifications
│           ├── eventbus/
│           └── notifier/
│
├── pkg/                             # SHARED UTILITIES
│   ├── response/                    # Standard response format
│   ├── logger/                      # Logging utilities
│   ├── validator/                   # Custom validation
│   ├── database/                    # DB connection pool (Singleton)
│   └── util/                        # General utilities
│
├── config/                          # CONFIGURATION
│   ├── app.go                       # App config struct
│   └── load.go                      # Config loading logic
│
├── db/
│   └── migration/                   # Database migrations
│       ├── 000001_init_schema.up.sql
│       └── 000001_init_schema.down.sql
│
├── scripts/                         # Deployment & maintenance scripts
├── deploy/                          # Docker & Kubernetes configs
├── go.mod
├── go.sum
├── Makefile
└── README.md