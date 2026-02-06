# Go Command Cheatsheet

## Build & Run

| Command | Description |
|---------|-------------|
| `go run .` | Compile and run (doesn't create binary) |
| `go run main.go` | Run a specific file |
| `go build` | Compile binary (output: `<module>.exe`) |
| `go build -o bin/app.exe` | Compile to specific path |
| `go install` | Compile and install to `$GOBIN` |

## Module Management

| Command | Description |
|---------|-------------|
| `go mod init <name>` | Create new module |
| `go mod tidy` | Add missing / remove unused dependencies |
| `go mod download` | Download dependencies to cache |
| `go mod verify` | Verify dependencies haven't been modified |
| `go mod graph` | Print dependency graph |
| `go mod vendor` | Copy dependencies into `vendor/` folder |

## Dependencies

| Command | Description |
|---------|-------------|
| `go get <pkg>` | Add/update dependency |
| `go get <pkg>@latest` | Get latest version |
| `go get <pkg>@v1.2.3` | Get specific version |
| `go get -u <pkg>` | Update to latest minor/patch |
| `go get -u ./...` | Update all dependencies |
| `go list -m all` | List all dependencies |
| `go list -m -u all` | List available updates |

## Testing

| Command | Description |
|---------|-------------|
| `go test` | Run tests in current package |
| `go test ./...` | Run all tests recursively |
| `go test -v` | Verbose test output |
| `go test -run TestName` | Run specific test |
| `go test -cover` | Show coverage percentage |
| `go test -coverprofile=coverage.out` | Generate coverage file |
| `go tool cover -html=coverage.out` | View coverage in browser |
| `go test -bench .` | Run benchmarks |
| `go test -race` | Detect race conditions |

## Code Quality

| Command | Description |
|---------|-------------|
| `go fmt ./...` | Format all code |
| `go vet ./...` | Report suspicious code |
| `go fix ./...` | Update code to use newer APIs |
| `gofmt -s -w .` | Simplify and format code |

## Information

| Command | Description |
|---------|-------------|
| `go version` | Show Go version |
| `go env` | Show all environment variables |
| `go env GOPATH` | Show specific variable |
| `go doc <pkg>` | Show package documentation |
| `go doc <pkg>.<Symbol>` | Show specific symbol docs |
| `go list -m` | Show current module name |

## Build Flags

| Flag | Description |
|------|-------------|
| `-o <path>` | Output binary path |
| `-v` | Verbose (print package names) |
| `-race` | Enable race detector |
| `-ldflags="-s -w"` | Strip debug info (smaller binary) |
| `-tags <tag>` | Build with specific tags |

## Cross-Compilation

```powershell
# Linux
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o app

# macOS
$env:GOOS="darwin"; $env:GOARCH="amd64"; go build -o app

# Windows (reset)
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o app.exe
```

## Common Workflows

```powershell
# Start new project
go mod init myproject
go get github.com/some/dependency

# Daily development
go run .                    # quick run
go build && .\app.exe       # build and run

# Before committing
go fmt ./...                # format
go vet ./...                # check for issues
go test ./...               # run tests
go mod tidy                 # clean dependencies

# Update dependencies
go get -u ./...             # update all
go mod tidy                 # clean up
```
