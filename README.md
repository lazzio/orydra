# Orydra - Ory Hydra Client Manager

This is a simple web application to manage Ory Hydra clients, written in Go.

## Features

- Manage Ory Hydra clients
- View clients
- Add clients
- Update clients
- Delete clients

## Usage

Environnement variables are stored in the `.env` file. You can use the `.env.default` file as a template.

```
PORT=8080
POSTGRES_PORT=5432
POSTGRES_USER=root
POSTGRES_PASSWORD=root
POSTGRES_DB=hydra_dev
POSTGRES_HOST=localhost
POSTGRES_SSLMODE=disable
POSTGRES_CLIENT_TABLE=hydra_client
```

```bash
# Set environment variables
cp .env.default .env

# Install dependencies
go mod tidy

# Run the application
go run main.go
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.