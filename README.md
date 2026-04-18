# Flow

Flow is a multi-language workflow management tool used to define, compile, and execute workflows. It adopts a modern multi-language technology stack, including Go, OCaml, Rust, and React, to build a complete workflow management ecosystem.

## Project Structure

```plaintext
├── flow/         # Command-line tool (Go)
├── flowc/        # Compiler (OCaml)
├── flowe/        # Executor (Rust)
├── flowd-back/   # Backend service (Go + Gin)
├── flowd-front/  # Frontend interface (React + TypeScript)
├── project/      # Project-related files
└── .flow/        # Workflow definition files
```

## Core Features

- **Workflow Definition**: Define workflows through `.flow` files, including multiple steps and dependencies between steps
- **Compilation Process**: Use `flowc` compiler to compile `.flow` files into JSON format workflow definitions
- **Dependency Analysis**: Use DFS algorithm to detect cyclic dependencies between steps
- **Execution Order**: Determine the execution order of steps based on dependencies
- **Step Execution**: Use `flowe` executor to execute each step
- **Status Recording**: Record execution results (success/failure) to SQLite database
- **Web Interface**: Provide intuitive workflow execution status display through `flowd-front`
- **API Interface**: `flowd-back` provides RESTful API interfaces for querying workflow execution status

## Technology Stack

| Module | Technology | Purpose |
| -------- | ------------ | --------- |
| Command-line tool | Go + Cobra | Provide user interaction interface, execute workflows |
| Compiler | OCaml | Compile .flow files into JSON format |
| Executor | Rust | Execute workflow steps |
| Backend service | Go + Gin | Provide API interfaces, access database |
| Frontend interface | React + TypeScript + Vite | Display workflow execution status |
| Data storage | SQLite | Record workflow execution status |

## Installation

### Prerequisites

- Go 1.20+
- OCaml 4.14+
- Rust 1.70+
- Node.js 18+
- pnpm

### Build and Install

1. **Compile flow command-line tool**

```bash
cd flow
go build -o flow main.go
```

1. **Compile flowc compiler**

```bash
cd flowc
dune build
cp _build/default/bin/main.exe flowc
```

1. **Compile flowe executor**

```bash
cd flowe
cargo build --release
cp target/release/flowe flowe
```

1. **Build frontend**

```bash
cd flowd-front
pnpm install
pnpm run build
```

1. **Compile backend service**

```bash
cd flowd-back
go build -o flowd-back main.go
```

## Usage

### 1. Define Workflow

Create a `.flow` file, for example `test.flow`:

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

### 3. View Execution Status

Start backend service:

```bash
./flowd-back
```

Start frontend development server:

```bash
cd flowd-front
pnpm run dev
```

Then visit `http://localhost:5173` to view workflow execution status.

## Workflow Definition Syntax

### Basic Syntax

```flow
workflow workflow_name doing
  step step_name is
    command
  end with
    deps  [dependency1, dependency2],
  finish
done
```

### Example

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

## API Interfaces

### Get Service Status

```plaintext
GET /api/status
```

### Get Workflow Execution Status

```plaintext
GET /api/flow
```

## Architecture Description

### Workflow Process

1. **Define Workflow**: User creates `.flow` file, defines workflow name, steps and dependencies
2. **Compile Workflow**: Use `flowc` compiler to compile `.flow` file into JSON format
3. **Execute Workflow**: Use `flow run` command to execute workflow, process dependencies and execute steps in order
4. **Record Status**: Execution results are recorded to SQLite database
5. **View Status**: View workflow execution status through web interface or API interface

### Core Component Interaction

```plaintext
+-------------+      +-------------+      +-------------+
|   flow      | ---> |   flowc     | ---> |   JSON      |
+-------------+      +-------------+      +-------------+
      |                                        |
      v                                        v
+-------------+      +-------------+      +-------------+
|   flowe     | <--- |   flow run  | <--- | Dependency  |
+-------------+      +-------------+      | Analysis    |
      |                                   +-------------+ 
      v
+-------------+
|  Execute    |
|  Steps      |
+-------------+
      |
      v
+-------------+      +-------------+      +-------------+
|  Update     | ---> |  SQLite     | <--- |  flowd-back |
|  Database   |      +-------------+      +-------------+
+-------------+                                 |
                                                v
                                          +-------------+
                                          |  flowd-front|
                                          +-------------+
```

## Application Scenarios

- **CI/CD Pipeline**: Define and execute continuous integration/continuous deployment processes
- **Data Processing Pipeline**: Build and execute data processing workflows
- **Automated Tasks**: Define and execute various automated tasks
- **DevOps Toolchain**: As part of the DevOps toolchain, manage various automated processes

## License

MIT License

## Contribution

Welcome to submit Issues and Pull Requests!

## Author

 hunter-hongg
