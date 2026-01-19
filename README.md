# GitHub Action Prototype

A simple Go program that validates JSON files in the config folder with automated CI/CD pipeline.

## Features

- ✅ Validates JSON syntax in configuration files
- ✅ Supports multiple JSON files in the config directory
- ✅ Provides detailed validation results with success/failure indicators
- ✅ Exit code 1 if any JSON file is invalid (useful for CI/CD)
- ✅ Automated CI/CD pipeline with scheduled runs and manual triggers

## Project Structure

```
.
├── .github/
│   ├── workflows/
│   │   └── cicd.yaml        # CI/CD pipeline configuration
│   └── SCHEDULE.md          # Schedule documentation
├── config/                   # JSON configuration files
│   ├── app.json
│   ├── database.json
│   └── service.json
├── build/                    # Build output directory
├── main.go                   # Main application
├── main_test.go              # Unit tests
├── go.mod                    # Go module file
├── Makefile                  # Build automation
└── README.md
```

## Prerequisites

- Go 1.21 or higher

## Installation

```bash
# Clone the repository
git clone <repository-url>
cd github-action-prototype

# Download dependencies
go mod download
```

## Usage

### Build the application

```bash
make build
```

### Run the application

```bash
# Validate JSON files in the default config directory
make run

# Or run directly
./build/json-validator

# Validate JSON files in a custom directory
./build/json-validator /path/to/config
```

### Run tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage
```

### Other make targets

```bash
make fmt            # Format code
make vet            # Run go vet
make tidy           # Tidy go modules
make clean          # Clean build artifacts
make help           # Show all available targets
```

## CI/CD Pipeline

The project includes a GitHub Actions workflow that:
- ✅ Runs tests automatically
- ✅ Builds the application
- ✅ Validates code quality
- ✅ Executes on scheduled intervals
- ✅ Can be triggered manually

### Automated Schedule

The CI/CD pipeline runs automatically at:
- **Monday** at 8:00 AM UTC
- **Tuesday** at 8:00 PM UTC
- **Wednesday** at 10:00 PM UTC

See [.github/SCHEDULE.md](.github/SCHEDULE.md) for timezone conversions and detailed schedule information.

### Manual Trigger

You can manually trigger the workflow:

1. Go to **Actions** tab in GitHub
2. Select **Go Service CI/CD** workflow
3. Click **Run workflow**
4. Optionally provide a reason
5. Click **Run workflow** button

Or use GitHub CLI:
```bash
gh workflow run cicd.yaml -f reason="Testing changes"
```

## Example Output

```
===========================================
JSON Configuration File Validator
===========================================
Validating JSON files in: config
-------------------------------------------
✅ app.json: VALID
✅ database.json: VALID
✅ service.json: VALID
===========================================
Summary: 3 valid, 0 invalid
===========================================
```

## Exit Codes

- `0`: All JSON files are valid
- `1`: One or more JSON files are invalid or an error occurred

## Testing

The project includes comprehensive unit tests:
- JSON validation tests with valid and invalid JSON
- Config folder validation tests
- Edge case handling

Run tests with:
```bash
go test -v ./...
```

## Development

### Adding New Configuration Files

1. Create a new `.json` file in the `config/` directory
2. Run the validator to test: `make run`
3. Tests will automatically validate the new file

### Modifying the Schedule

Edit [.github/workflows/cicd.yaml](.github/workflows/cicd.yaml) and update the `schedule` section with new cron expressions.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests: `make test`
5. Create a pull request

## License

MIT
