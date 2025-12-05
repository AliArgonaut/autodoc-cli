# Auto-Doc CLI

Auto-Doc is a command-line interface (CLI) tool designed to automate the generation of documentation for your codebase. It works by sending a snapshot of your workspace to an agent backend, which then processes the code to create a comprehensive `README.md` file and other relevant documentation.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [Initializing a Project](#initializing-a-project)
  - [Generating Documentation](#generating-documentation)
  - [Checking Backend Health](#checking-backend-health)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Automated Documentation Generation**: Creates a `README.md` file from your source code.
- **Workspace Snapshotting**: Sends your project's file structure and code content to a backend agent for processing.
- **Project Initialization**: Sets up a new documentation project with a default configuration file.
- **Backend Health Check**: Allows you to verify the connectivity and responsiveness of the auto-doc backend services.

## Installation

This project is a CLI tool, and its installation would typically involve building it from source or downloading a pre-compiled binary. Assuming you have Go installed, you can build the project using the following commands:

```bash
# Clone the repository (if you haven't already)
git clone <repository-url>
cd <repository-directory>

# Build the CLI tool
go build -o auto-doc .
```

This will create an executable file named `auto-doc` in your current directory.

## Usage

The `auto-doc` CLI tool supports several commands to manage your documentation workflow.

### Initializing a Project

To start a new documentation project, use the `init` command. This command will set up the necessary folder structure and generate a configuration file named `autodoc_config.json` with sensible defaults.

```bash
./auto-doc init
```

This command interacts with the auto-doc backend to fetch a configuration template, which is then saved as `autodoc_config.json` in your current directory.

### Generating Documentation

Once your project is initialized and configured, you can generate the documentation using the `generate` command. This command will send your workspace snapshot to the backend agent for processing and will create/update your `README.md` file with the generated documentation.

```bash
./auto-doc generate
```

This command performs the following actions:
1. Gathers your project's configuration from `autodoc_config.json`.
2. Builds a file tree representation of your project.
3. Identifies and reads the content of supported language files.
4. Sends this information to the backend API (`http://localhost:8000/api/generate`).
5. Receives the generated documentation from the backend.
6. Writes the documentation to a `README.md` file in your current working directory.

### Checking Backend Health

To ensure the auto-doc backend services are running and accessible, you can use the `health` command.

```bash
./auto-doc health
```

This command pings the backend API (`http://localhost:8000/api/health`) and prints the response, which is useful for troubleshooting connectivity issues.

## Configuration

The `autodoc_config.json` file is used to configure the auto-doc tool. It should be placed in the root of your project. The following fields are supported:

| Field       | Type    | Description                                       | Default Value |
| :---------- | :------ | :------------------------------------------------ | :------------ |
| `name`      | string  | The name of the application or project.           | ""            |
| `developer` | string  | The name of the developer.                        | ""            |
| `description` | string  | A brief description of the project.               | ""            |
| `ignore`    | []string | A list of file or directory patterns to ignore. | `[]`          |

**Example `autodoc_config.json`:**

```json
{
  "name": "My Awesome Project",
  "developer": "Jane Doe",
  "description": "A project that does amazing things.",
  "ignore": [
    "node_modules",
    "*.test.js",
    ".git"
  ]
}
```

When you run `auto-doc init`, a default configuration file will be generated. You can then modify this file to suit your project's specific needs.

## Contribution

Contributions are welcome! Please refer to the contributing guidelines in the project's repository for more information on how to contribute.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.