
# 🟦 Go-Basics — Small hands-on Go examples

This repository contains small, focused Go examples and exercises that demonstrate core Go concepts: custom types, slices, structs, interfaces, channels, concurrency, file I/O, HTTP clients, maps, and basic testing. Each folder contains a compact example you can run to learn by experimenting.

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

**Quick commands**
- Run a single example from a folder (replace `path/to/folder`):
  - `go run path/to/folder/*.go`
- Run all tests in the repository:
  - `go test ./...`

---

**Project structure & short descriptions**

- `cards/` : Deck of cards example
  - **What:** Defines a custom `deck` type with helpers to create, deal, shuffle, print, save to file, and load from file.
  - **Files:** `deck.go`, `deck_test.go`, `main.go`
  - **Run:** `go run cards/*.go`
  - **Notes:** `main` demonstrates dealing a hand, saving to `my_cards`, loading from file, and shuffling.

- `channels/` : Concurrency and channels
  - **What:** Demonstrates launching goroutines to check HTTP links concurrently and reporting results back over a channel.
  - **Files:** `main.go`
  - **Run:** `go run channels/main.go`

- `http/` : HTTP client and custom writer
  - **What:** Performs a simple `http.Get` and shows how to implement `io.Writer` to process streaming response bodies.
  - **Files:** `main.go`
  - **Run:** `go run http/main.go`

- `interface/` : Interfaces with examples
  - **What:** Two small examples showing how interfaces work:
    - `main.go` — greeting bots implementing a `bot` interface
    - `shapes/shapemain.go` — `shape` interface with `triangle` and `square` implementations
    - `printfile/printfilemain.go` — basic file-reading CLI using `io.Copy`
  - **Run:**
    - `go run interface/main.go`
    - `go run interface/shapes/shapemain.go`
    - `go run interface/printfile/printfilemain.go <path-to-file>`

- `map/` : Working with maps
  - **What:** Shows creating, updating, iterating and deleting entries in a `map[string]string`.
  - **Files:** `main.go`
  - **Run:** `go run map/main.go`

- `Structs/` : Structs and methods, pointers
  - **What:** Define `person` and `contactInfo` structs, demonstrate value vs pointer receivers, and updating struct fields.
  - **Files:** `main.go`
  - **Run:** `go run Structs/main.go`

- `truck_manager/` : Simple manager with interface + unit tests
  - **What:** Demonstrates defining an interface `FleetManager`, a `Truck` struct, and a concrete `truckManager` implementation. Includes unit tests in `truck_manager_test.go`.
  - **Files:** `truck_manager.go`, `truck_manager_test.go`
  - **Run tests:** `go test ./truck_manager`

---

**Testing guidance**
- Run all tests in repository: `go test ./...`
- Run tests for a specific package: `go test ./truck_manager`

---

If you'd like, I can:
- add a short `Makefile` or task list to run each example easily,
- add example output for each program in this `README`, or
- run the tests here and report results.

Please tell me which additions you prefer and I will update the `README.md` accordingly.
