# 🂡 GoCardsLab — A Hands-on Go Learning Project

A beginner-friendly Go project designed to explore **core Go concepts** through a practical example: simulating a deck of cards 🎴.  
This project also includes examples and exercises on **Go fundamentals** like channels, interfaces, and structs — making it an excellent foundation for anyone learning Go.


---

## 🛠️ Prerequisites

Before running the project, make sure you have:

- **Go installed on your system**

  👉 [Install Go from the official site](https://go.dev/doc/install)

- *(Optional)* **VS Code** with the **Go extension** for an enhanced development experience.

---

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

---

## 💡 Learning Goals
- By exploring this project, you’ll learn how to:
- Create and use custom types
- Work efficiently with slices and functions
- Define and implement structs and interfaces
- Use channels and goroutines for concurrency
- Write and execute unit tests in Go
- Perform file I/O operations safely and efficiently
