# go-web-tmpl

A minimal Go REST API template using the standard library.

## Structure

```
.
├── cmd/server/       # entry point
├── internal/
│   ├── config/       # env-based config with cleanenv
│   ├── handler/      # HTTP handlers
│   └── middleware/   # logging & recovery middleware
```

## Getting started

```sh
cp .env.example .env
make run
```

## Routes

| Method | Path      | Description   |
|--------|-----------|---------------|
| GET    | /health   | Health check  |
