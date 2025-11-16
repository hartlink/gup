# GUP - Ubuntu Server Maintenance CLI Tool

GUP (Go Update) is a command-line tool developed in Go that uses Bubble Tea and Cobra to automate common maintenance tasks on Ubuntu servers.

## 🚀 Features

- **Interactive Interface**: Uses Bubble Tea for a modern and attractive user experience
- **Structured Commands**: Implemented with Cobra for robust command and argument handling
- **Minimal Dependencies**: Designed to use as few dependencies as possible
- **Standalone Binary**: Generates an independent executable binary
- **Multilingual (i18n)**: Full support for Spanish and English, with automatic system language detection

## 📦 Installation

### Prerequisites
- Go 1.21 or higher
- Ubuntu system (for maintenance functions)
- Administrator permissions (sudo) for system commands

### Compile from source

```bash
# Clone the repository
git clone https://github.com/hartlink/gup.git
cd gup

# Download dependencies
go mod tidy

# Compile the binary
make build

# Or compile directly with Go
go build -o gup main.go
```

### Deploy to Server

After compilation, you can deploy the binary to your Ubuntu server:

```bash
# Copy the binary to your server
scp build/gup user@your-server:/tmp/

# On the server, move it to a system path
ssh user@your-server
sudo mv /tmp/gup /usr/local/bin/
sudo chmod +x /usr/local/bin/gup

# Verify installation
gup version
```

Alternatively, compile directly on the server:

```bash
# On the server
git clone https://github.com/hartlink/gup.git
cd gup
make build
sudo cp build/gup /usr/local/bin/
```

## 🛠️ Usage

### Available Commands

#### `gup update`
Updates the system package list by running `apt update`.

```bash
gup update
```

> **Note**: The command will use `sudo` automatically if needed.

**Options:**
- `-v, --verbose`: Shows detailed information during execution
- `-l, --lang`: Select language (es/en)

### Languages

GUP automatically detects your system language, but you can change it manually:

```bash
# Spanish
gup --lang es

# English  
gup --lang en

# Applicable to any command
gup version --lang es
gup update --lang en
```

For more information about the translation system, see [docs/i18n.md](docs/i18n.md).

### Permanent Configuration

You can set your preferred language permanently:

```bash
# Create configuration file
mkdir -p ~/.gup
echo '{"language":"en"}' > ~/.gup/config.json
```

See more options in [docs/config.md](docs/config.md).

### Examples

```bash
# Update package list
gup update

# Update with detailed output
gup update --verbose

# View help
gup --help
gup update --help
```

## 🔧 Development

### Project Structure

```
cli_go/
├── main.go              # Main entry point
├── cmd/                 # Cobra commands
│   ├── root.go          # Root command
│   ├── update.go        # Update command
│   ├── demo.go          # Demo command
│   └── version.go       # Version command
├── internal/            # Internal logic
│   └── ui.go           # Bubble Tea interface
├── pkg/                 # Public packages
│   ├── i18n/           # Internationalization system
│   │   └── i18n.go     # ES/EN translations
│   └── config/         # Configuration system
│       └── config.go   # Config management
├── docs/               # Documentation
│   ├── i18n.md        # Internationalization guide
│   └── config.md      # Configuration guide
├── build/              # Compiled binary
│   └── gup            # Executable
├── go.mod              # Go module
├── go.sum              # Dependencies checksums
├── Makefile            # Build commands
├── .gitignore         # Ignored files
├── README.md           # This file (English)
└── README_ES.md        # Spanish version
```

### Development Commands

```bash
# Compile
make build

# Clean binaries
make clean

# Run in development mode
make dev

# Install dependencies
make deps
```

### Adding New Commands

1. Create a new file in `cmd/` (e.g., `cmd/upgrade.go`)
2. Implement the command using Cobra
3. Use the Bubble Tea interface from `internal/ui.go`
4. Register the command in `cmd/root.go`

## 📝 Roadmap

- [ ] `gup upgrade` - Upgrade system packages
- [ ] `gup cleanup` - Clean unnecessary packages
- [ ] `gup status` - Show system status
- [ ] `gup logs` - View system logs
- [ ] `gup services` - Manage systemd services
- [ ] Customizable configuration
- [ ] Advanced logging
- [ ] Unit tests

## 🌍 Language Support

This project is available in:
- 🇺🇸 [English](README.md) (this file)
- 🇲🇽 [Español](README_ES.md)

## 🤝 Contributing

Contributions are welcome. Please:

1. Fork the project
2. Create a branch for your feature (`git checkout -b feature/new-feature`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature/new-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the GPLv3. See the `LICENSE` file for more details.

## ⚠️ Warnings

- **Root Permissions**: Many commands require administrator permissions
- **Compatibility**: Designed specifically for Ubuntu/Debian systems
- **Responsible Use**: Always review commands before running them in production
