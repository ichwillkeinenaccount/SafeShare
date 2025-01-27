# SafeShare 🔗

> **Note**: At the moment this is only a hobby project in active development! Everything is still experimental and under
> construction! 🚧

Safe, simple and easy to use Text- and File-sharing service.

SafeShare aims to be a simple service to quickly and securely share text, passwords or files. Everything is fully
end-2-end
encrypted so even the server can't access the data.
SafeShare is written in Go and vanilla JavaScript/TypeScript.

## Roadmap 🗺️

| Feature          | Version | Status      |
|------------------|:-------:|-------------|
| e2e Text Sharing |   1.0   | In Progress |
| SSO              |   1.0   | Planned     |
| e2e File Sharing |   1.1   | Not Started |
| Internal Auth    |   1.0   | Not Started |
| Public Mode      |   1.0   | Not Started |
| Theming          |   2.0   | Not Started |
| Dark Mode        |   2.0   | Not Started |
| i18n             |   2.0   | Not Started |
| Feature Flags    |   2.0   | Not Started |

### Technical ideas 💡
- Use three word urls for sharing
- Use structured logging (slog?)
- Use [golang-migrate](https://github.com/golang-migrate/migrate) for database migrations
- Use MinIO as a storage backend
- Use [Viper](https://github.com/spf13/viper) and [Cobra](https://github.com/spf13/cobra) for configuration and CLI
- Use xyz for openapi generation
  - Use Bruno for API testing
- Use go tool
- Use OpenTelemetry for tracing
- Use sqlc for SQL generation?
- Use a salt for additional security in database?
- CLI version for sharing files and text directly from the terminal?

## Build with 🛠️

- [Go](https://go.dev/)
- [TypeScript](https://www.typescriptlang.org/)
- [TailwindCSS](https://tailwindcss.com/)
- [PostgreSQL](https://www.postgresql.org/)
- [MinIO](https://github.com/minio/minio)



.
├── cmd/
│   └── <app-name>/
│       └── main.go
├── pkg/
│   └── <your-packages>/
├── internal/
│   └── <internal-packages>/
├── api/
│   └── <api-definitions>/
├── web/
│   └── <frontend-files>/
├── configs/
│   └── <config-files>/
├── scripts/
│   └── <scripts>/
├── build/
│   └── <build-output>/
├── deployments/
│   └── <deployment-configs>/
├── test/
│   └── <test-data>/
├── go.mod
├── go.sum
└── README.md