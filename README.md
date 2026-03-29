# LRU Cache [![Go Reference](https://pkg.go.dev/badge/github.com/tarvarrs/lru-cache.svg)](https://pkg.go.dev/github.com/tarvarrs/lru-cache) [![Go Report Card](https://goreportcard.com/badge/github.com/tarvarrs/lru-cache)](https://goreportcard.com/report/github.com/tarvarrs/lru-cache) [![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Thread-safe LRU (Least Recently Used) cache implementation in Go with generics (Go 1.26+).  
Supports `Get`, `Set`, and `Clear` operations with O(1) time complexity.

## Features

- **Generic** – works with any comparable key and any value type.
- **Thread-safe** – uses a mutex for concurrent access.
- **O(1) operations** – based on a hash map + doubly linked list.
- **Eviction** – automatically removes the least recently used entry when capacity is reached.
- **Zero dependencies** – only the standard library.

## Installation

```bash
go get github.com/tarvarrs/lru-cache@latest
```
