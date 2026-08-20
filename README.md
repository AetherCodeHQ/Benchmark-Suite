# Benchmark Suite

![CI](https://github.com/Qyroxen/Benchmark-Suite/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Benchmark-Suite/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Benchmark-Suite?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Benchmark-Suite)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Benchmark-Suite)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Benchmark-Suite?style=social)](https://github.com/Qyroxen/Benchmark-Suite/stargazers)

## What is it?

Benchmark Suite is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Benchmark-Suite.git
cd Benchmark-Suite
go build -o benchmarksuite .

# Run
./benchmarksuite --help
```

## CLI Usage

```bash
# Basic usage
./benchmarksuite

# With flags
./benchmarksuite --verbose --output json

# Get help
./benchmarksuite --help
```

## Examples

```bash
# Example 1
./benchmarksuite example1

# Example 2
./benchmarksuite example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o benchmarksuite .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Benchmark-Suite/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Benchmark-Suite?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Benchmark-Suite/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Benchmark-Suite?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Benchmark-Suite/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Benchmark-Suite" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Benchmark-Suite/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Benchmark-Suite" alt="Pull Requests">
  </a>
</p>
