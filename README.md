# toolr

`toolr` is a universal, cross-platform CLI tool launcher and registry built to make your workflow smoother by centralizing access to frequently used console-based helper tools, scripts, and micro-programs — no matter the language.

With `toolr`, you can register, manage, and execute console tools from a single command like:

```bash
toolr calc add 5 2
toolr myparser file.txt
toolr rusttool --opt value
```
## ✨ Features

* 🔧 Register any tool (Python, C#, Rust, Go, Node.js, etc.)
* ⚡ One-liner execution from anywhere
* 🧩 Project or function-level dispatch
* 🖥️ Terminal UI for tool management using [Bubble Tea](https://github.com/charmbracelet/bubbletea)
* 🌐 Cross-platform: Works on **Windows**, **Linux**, and **macOS**
* 🗂️ Configurable folder structure and metadata
* 🚀 Fast tool execution with parameter passing
* 💡 Language-agnostic — run any compiled or interpreted helper

---

## 📁 Folder Structure

```plaintext
toolr/
├── cmd/                # CLI command handlers
├── internal/
│   ├── config/         # Tool registry and config management
│   ├── exec/           # Tool execution logic
│   └── tui/            # Terminal UI using Bubble Tea
├── tools/              # Optional internal helper tools
├── data/               # Registered tool metadata (JSON/YAML)
├── go.mod
└── main.go             # Entry point
```

---

## 🧪 Example Tool: Python Calculator

```python
# calc.py
import sys

def main():
    if len(sys.argv) < 4:
        print("Usage: python calc.py [add|sub|mul|div] num1 num2")
        return

    op = sys.argv[1]
    a, b = float(sys.argv[2]), float(sys.argv[3])

    match op:
        case "add": print(f"Result: {a + b}")
        case "sub": print(f"Result: {a - b}")
        case "mul": print(f"Result: {a * b}")
        case "div": print(f"Result: {a / b if b != 0 else 'Error: Division by zero'}")
        case _: print("Unknown operation")
```

To register:

```bash
toolr register --name calc --path /path/to/calc.py --exec python3
```

To run:

```bash
toolr calc add 4 5
```

---

## 📦 Installation

### Pre-built Binary

Coming soon...

### From Source

```bash
git clone https://github.com/your-username/toolr.git
cd toolr
go build -o toolr .
```

---

## 🧠 How It Works

1. **Register Tools**
   Add any script or executable via CLI or terminal UI. The metadata is stored locally (JSON/YAML).

2. **Dispatch Commands**
   `toolr` parses your command, matches the tool, and executes it using its configured interpreter or binary.

3. **TUI Control Panel**
   Navigate, register, or delete tools from the terminal UI built using Bubble Tea.
