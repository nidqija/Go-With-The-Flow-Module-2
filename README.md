# 🐹 GDGoC Golang Workshop: Gopher Watch 📊

Welcome to the **GDGoC (Google Developer Groups on Campus) Golang Workshop!** 

In this workshop, you will build **Gopher Watch**—a real-time system monitoring web application using **Go (Golang)**, `html/template`, `HTMX`, and system metrics libraries (`gopsutil`).

---

## 🛠️ Prerequisites

Before you start, make sure you have the following installed:

1. **Go (Golang)**: Version `1.24` or higher  
   👉 [Download & Install Go](https://golang.org/dl/)
2. **Git**: For cloning the repository  
   👉 [Download Git](https://git-scm.com/)
3. **Code Editor**: VS Code (recommended with the official *Go extension*) or your preferred IDE.

Check your Go version in your terminal:
```bash
go version
```

---

## 🚀 Step-by-Step Workshop Guide

### Step 1: Clone the Repository

Clone this starter repository to your local machine and open the directory:

```bash
# 1. Clone the repository
git clone <repository-url>

# 2. Change into the project directory
cd <your-directory>
```

---

### Step 2: Explore the Starter Files

When you open the project, you will find the following starter files:

- `main.go`: Starter Go file with code template placeholders (`// to be filled by participants`).
- `index.html`: Dashboard template with HTMX interactive components.
- `signin.html`: Login page template.


---

### Step 3: Install Required Dependencies

This application uses the `gopsutil` package to collect hardware stats (CPU, RAM, OS info).

Run the following commands in your terminal to fetch the library and synchronize your modules:

```bash
# 1. Download the gopsutil package
go get github.com/shirou/gopsutil/v4/cpu

# 2. Tidy dependencies
go mod tidy
```

---

### Step 4: Import Packages in `main.go`

Open `main.go` in your code editor and update the `import` block at the top of the file:

```go
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)
```

---

### Step 5: Fill in the Workshop Tasks (`main.go`)

Work through the following sections in `main.go` alongside the instructor:

#### 📝 Task 1: Define Data Structs
Define the `LoginData` struct (to handle sign-in form state and errors) and the `SystemStats` struct (to hold CPU load, RAM usage, and OS info).

#### 📊 Task 2: Implement `fetchStats()`
Use `cpu.Percent()`, `mem.VirtualMemory()`, and `host.Info()` to collect live system statistics and return a populated `SystemStats` struct.

#### 🎨 Task 3: Parse Templates & Page Rendering
Parse `signin.html` and `index.html` using `template.Must(template.ParseFiles(...))` and create the `renderPage` closure to handle HTMX dynamic template swapping.

#### 🌐 Task 4: Implement Route Handlers
Complete the HTTP handlers for:
- `/`: Redirect to `/signin`
- `/signin`: Process GET & POST sign-in requests
- `/dashboard`, `/stats`, `/memory`, `/api/live-stats`, `/system`: Render dashboard views with live system stats

---

### Step 6: Run & Test the Application

Once your code is complete, run the server:

```bash
go run main.go
```

**Terminal Output:**
```text
Starting server on port 8000
```

1. Open your browser and navigate to:
   ```
   http://localhost:8000
   ```
2. You will be redirected to the **Sign-In Page** (`/signin`).
3. Sign in with the credentials you implemented in your code.
4. View your live **Gopher Watch System Monitor Dashboard**! 🎉

---

## 💡 Helpful Go Commands Cheat Sheet

- `go run main.go`: Compile and run your application immediately.
- `go get <package-path>`: Add an external Go package to your project.
- `go mod tidy`: Automatically manage, add, or remove module dependencies.
- `go fmt main.go`: Auto-format your Go source code.
