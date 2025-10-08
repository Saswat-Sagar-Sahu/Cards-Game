# 🃏 Cards Project (Go)

A simple Go project to simulate a deck of cards. It covers concepts like:

- Custom types
- Slices and functions
- File I/O (save/load deck to/from file)
- Error handling
- Unit testing in Go

---

## 🛠️ Prerequisites

Before running the project, make sure you have:

- **Go installed on your system**

  👉 Install Go from the official site:  
  [https://go.dev/doc/install](https://go.dev/doc/install)

- **VS Code (Optional)** for a better development experience with the Go extension.

---

## ▶️ How to Run the Program

Inside the project directory (e.g., `cards`), run:

```bash
go run main.go deck.go
```

## 📦 Module Initialization (Required for Testing)

Before running tests or managing dependencies, initialize a Go module:
```bash
go mod init cards
```

This creates a go.mod file and prepares your project for package management and testing.

## ✅ You only need to run this once per project.

✅ Running Tests
This project includes unit tests. You can run them via terminal or through VS Code.

🔹 Option 1: Run Tests in Terminal
```
go test
```
This runs all tests in files ending with `_test.go`.

🔹 Option 2: Run Tests in VS Code
- Open the test file (e.g., deck_test.go) in VS Code.
- Hover over the test function or use the Testing Sidebar.
- Click Run Test to execute.

  
Make sure you have the official Go extension installed in VS Code.
