# Flow

A modern multi-language workflow management tool for defining, compiling, and executing workflows with a focus on dependency management and real-time status tracking.

## 🚀 Features

- **Multi-language Support**: Go, OCaml, Rust, React stack
- **Workflow Definition**: Define complex workflows with `.flow` files
- **Dependency Management**: Automatic cyclic dependency detection using DFS
- **Real-time Status**: Track execution progress via web interface
- **RESTful API**: Query workflow status programmatically
- **SQLite Storage**: Persistent execution history

## 📁 Project Structure

```plaintext
├── flow/         # CLI tool (Go)
├── flowc/        # Compiler (OCaml)
├── flowe/        # Executor (Rust)
├── flowd-back/   # Backend API (Go + Gin)
├── flowd-front/  # Web UI (React + TypeScript)
├── project/      # Project resources
└── .flow/        # Workflow definitions
```

## 🛠️ Tech Stack

| Component | Technology | Purpose |
| ----------- | ------------ | --------- |
| CLI Tool | Go + Cobra | User interface & workflow execution |
| Compiler | OCaml | Flow file parsing & compilation |
| Executor | Rust | Step execution & monitoring |
| Backend | Go + Gin | API services & database access |
| Frontend | React + TypeScript + Vite | Status visualization |
| Database | SQLite | Execution history storage |

## 📦 Installation

### Prerequisites

- Go 1.20+
- OCaml 4.14+
- Rust 1.70+
- Node.js 18+
- pnpm

### Build Instructions

```bash
# CLI Tool
cd flow && go build -o flow main.go

# Compiler
cd flowc && dune build && cp _build/default/bin/main.exe flowc

# Executor
cd flowe && cargo build --release && cp target/release/flowe flowe

# Frontend
cd flowd-front && pnpm install && pnpm run build

# Backend
cd flowd-back && go build -o flowd-back main.go
```

## 🚀 Quick Start

### 1. Define Workflow

Create `test.flow`:

```flow
workflow test doing
  step step1 is
    echo 'Step 1'
  end 
  step step2 is
    echo 'Step 2'
  end 
  step step3 is
    echo 'Step 3'
  end with
    deps  [step2]
  finish
done
```

### 2. Execute Workflow

```bash
./flow run test
```

### 3. Monitor Progress

```bash
# Start backend
./flowd-back

# Start frontend (development)
cd flowd-front && pnpm run dev
```

Visit `http://localhost:5173` for real-time status.

## 📝 Workflow Syntax

### Basic Structure

```flow
workflow workflow_name doing
  step step_name is
    command
  end with
    deps  [dependency1, dependency2]
  finish
done
```

### Example: Build Pipeline

```flow
workflow build doing
  step install is
    npm install
  end 
  step test is
    npm run build
  end with
    deps  [install]
  end 
  step deploy is
    npm run deploy
  end with
    deps  [build]
  finish
done
```

## 🔌 API Reference

### Service Status

```bash
GET /api/status
```

### Workflow Status

```bash
GET /api/flow
```

## 🏗️ Architecture

### Workflow Execution Flow

1. **Define**: Create `.flow` file with steps and dependencies
2. **Compile**: `flowc` converts to JSON format
3. **Execute**: `flow run` processes dependencies and executes steps
4. **Record**: Results stored in SQLite database
5. **Monitor**: View status via web interface or API

### Component Interaction

```plaintext
┌─────────┐    ┌─────────┐    ┌─────────┐
│  flow   │───▶│  flowc  │───▶│  JSON   │
└─────────┘    └─────────┘    └─────────┘
       │           │           │
       │           │           │
       ▼           ▼           ▼
┌─────────┐    ┌─────────┐    ┌─────────┐
│  flowe  │◀───│  flow   │◀───│  Dep    │
└─────────┘    └─────────┘    │  Graph  │
       │           │           └─────────┘
       │           │
       ▼           ▼
┌─────────┐    ┌─────────┐    ┌─────────┐
│  Update │───▶│ SQLite  │◀───│  API    │
└─────────┘    └─────────┘    └─────────┘
       │                         │
       │                         │
       ▼                         ▼
┌─────────┐    ┌─────────┐    ┌─────────┐
│  Web UI │    │  Status │    │  Front  │
└─────────┘    └─────────┘    └─────────┘
```

## 🎯 Use Cases

- **CI/CD Pipelines**: Continuous integration/deployment workflows
- **Data Processing**: ETL pipelines and data transformation
- **Automation**: Task scheduling and batch processing
- **DevOps**: Infrastructure provisioning and deployment

## 📄 License

MIT

## 🤝 Contributing

Issues and Pull Requests are welcome!

## 👤 Author

hunter-hongg
