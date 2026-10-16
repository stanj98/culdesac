# Culdesac - GoJSON Database Engine

A lightweight, modern JSON document database built from scratch in Go (inspired by MongoDB and CockroachDB architectures). This project implements a local file-based NoSQL document store with concurrent write protection using mutex locking.

---

<img width="3499" height="1805" alt="image" src="https://github.com/user-attachments/assets/fe634491-4b89-4a2f-8a88-6e32bcffc566" />

## Features

- **Document-Based Storage:** Organizes collections into directories where individual records are stored as JSON files.
- **Concurrency Control:** Thread-safe operations utilizing fine-grained `sync.Mutex` maps to prevent data corruption during simultaneous read/write cycles.
- **Custom Struct Mappings:** Go-native data typing with support for nested structs (e.g., User profiles with nested Address attributes).
- **Embedded Lifecycle Management:** Handles directory creation, file-path sanitization, and atomic writes via temporary file staging and renaming.

---

## Project Structure

```text
golang-database/
├── main.go
├── go.mod
└── go.sum
