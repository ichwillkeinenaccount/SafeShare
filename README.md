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
- Use [golang-migrate](https://github.com/golang-migrate/migrate) for database migrations
- Use go tool
- Use OpenTelemetry for tracing
- Use sqlc for SQL generation?
- Use a salt for additional security in database?
- CLI version for sharing files and text directly from the terminal ([Cobra](https://github.com/spf13/cobra))?
- Use Cobra for arguments?

## Development
- Clone the repository
- Run `docker-compose -f .development/docker-compose.yml up` to start the development environment
- Open [http://localhost:8081/#/](http://localhost:8081/#/) in your browser to see the Swagger UI
- The API is available at [http://localhost:8080](http://localhost:8080)

### Development Tools 🧰
- [Docker](https://www.docker.com/)
- [Bruno](https://github.com/usebruno/bruno)
- [Swagger-UI](https://swagger.io/tools/swagger-ui/)

## Build with 🛠️

### Server 🖥️
- [Go](https://go.dev/)
- [Swag](https://github.com/swaggo/swag)
- [Viper](https://github.com/spf13/viper)

### Frontend 🌐
- [TypeScript](https://www.typescriptlang.org/)
- [TailwindCSS](https://tailwindcss.com/)

### Database 🗄️
- [PostgreSQL](https://www.postgresql.org/)
- [MinIO](https://github.com/minio/minio)
