# MiniLink

A minimal URL shortener written in Go that uses YAML files as a database.

## Features

- Fast HTTP redirects using Fiber web framework
- Simple YAML-based URL mapping
- CLI interface with Cobra
- Docker support
- Minimal dependencies

## Installation

```bash
git clone <repository-url>
cd minilink
go mod download
go build -o minilink
```

## Usage

### Configuration

Edit the `links.yaml` file to add your URL routing rules:

```yaml
routes:
  gh:
    default: "https://github.com"
    rules:
      - query: "x=y"
        url: "https://github.com/Meraj"
        passthrough: false
  gf:
    default: "https://github.com"
    rules:
      - query: "x=y"
        url: "https://github.com/Meraj"
        passthrough: true
  yt:
    default: "https://youtube.com"
    rules:
      - query: "v=x"
        url: "https://google.com"
        passthrough: false
      - query: "y=2"
        url: "https://meraj.com"
        passthrough: false
```

### Running the Server

Start the server with default settings:

```bash
./minilink serve
```

Or specify custom port and config file:

```bash
./minilink serve --port 3000 --config my-links.yaml
```

### CLI Options

- `--port, -p`: Port to run the server on (default: 8080)
- `--config, -c`: Path to the YAML config file (default: links.yaml)

### API Endpoints

- `GET /`: Health check endpoint
- `GET /:shortcode`: Redirects to the configured URL

### Examples

Once the server is running on port 8080:

**Basic redirects:**
- Visit `http://localhost:8080/gh` → redirects to `https://github.com`
- Visit `http://localhost:8080/google` → redirects to `https://google.com`

**Query-based routing:**
- Visit `http://localhost:8080/gh?x=y` → redirects to `https://github.com/Meraj`
- Visit `http://localhost:8080/gf?x=y` → redirects to `https://github.com/Meraj?x=y` (with passthrough)
- Visit `http://localhost:8080/yt?v=x` → redirects to `https://google.com`
- Visit `http://localhost:8080/yt?y=2` → redirects to `https://meraj.com`

**Error handling:**
- Visit `http://localhost:8080/invalid` → returns 404 error

## Docker

### Build Image

```bash
docker build -t minilink .
```

### Run Container

```bash
docker run -p 8080:8080 -v $(pwd)/links.yaml:/app/links.yaml minilink
```

## Development

### Project Structure

```
minilink/
├── main.go              # Entry point
├── cmd/
│   └── serve.go         # Serve command
├── internal/
│   ├── config/
│   │   └── config.go    # YAML config loader
│   └── server/
│       └── server.go    # HTTP server
├── links.yaml           # URL database
├── go.mod              # Go module
├── README.md           # This file
└── Dockerfile          # Container config
```

### Adding New Routes

Edit the `links.yaml` file to add new routing rules:

```yaml
routes:
  mysite:
    default: "https://mywebsite.com"
    rules:
      - query: "page=about"
        url: "https://mywebsite.com/about"
        passthrough: false
  blog:
    default: "https://myblog.com"
    rules:
      - query: "tag=tech"
        url: "https://myblog.com/tech"
        passthrough: true
```

**Route Configuration:**
- `default`: The URL to redirect to when no query rules match
- `rules`: Array of query-based routing rules
- `query`: Query parameter pattern to match (e.g., "x=y")
- `url`: Target URL for this rule
- `passthrough`: If true, append original query parameters to the target URL

## License

MIT License