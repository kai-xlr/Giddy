# Giddy

A high-performance text editor built in Go, starting with a Rope data structure for efficient text buffer operations.

## Overview

Giddy uses a Rope data structure—a binary tree of text chunks—to achieve O(log N) performance for insertions, deletions, and line lookups instead of O(N) with standard strings.

### Key Features

- **Immutable Rope**: Modifications create new trees with structural sharing for efficient Undo/Redo
- **Line Tracking**: Fast line-based navigation with O(log N) lookups
- **High Performance**: Optimized for large files with minimal GC pressure

## Project Structure

```
/giddy
  /buffer          # Core rope data structure package
    node.go        # Node struct and tree primitives
    rope.go        # Public API for text manipulation
```

## Development

### Build and Test

```bash
# Build the project
go build ./...

# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run benchmarks
go test -bench=. ./buffer
```

### Code Quality

```bash
# Format code
go fmt ./...

# Run static analysis
go vet ./...

# Tidy dependencies
go mod tidy
```

## Implementation Status

Currently in Phase 1: Building the core Rope data structure with basic text operations.
