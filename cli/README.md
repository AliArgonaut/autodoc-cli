# AUTODOC CLI

Welcome to the AUTODOC CLI! This tool empowers you to automatically generate beautiful READMEs and comprehensive documentation for your codebases using AI agents.

## Table of Contents

- [About](#about)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [Initializing a Project](#initializing-a-project)
  - [Generating Documentation](#generating-documentation)
  - [Checking Health](#checking-health)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## About

The AUTODOC CLI is designed to streamline the documentation process for software projects. By leveraging AI agents, it analyzes your codebase, understands its structure, and generates high-quality documentation, including README files, with minimal manual effort.

## Features

*   **Automated Documentation Generation:** Create READMEs and other documentation artifacts automatically.
*   **AI-Powered Analysis:** Utilizes AI agents to understand code context and generate meaningful documentation.
*   **Project Initialization:** Simplifies the setup of new documentation projects with a default configuration.
*   **Health Checks:** Easily verify the connectivity and health of the AI agent backend.
*   **Customizable Configuration:** Define project details and ignore patterns through a configuration file.

## Installation

To install the AUTODOC CLI, you need to have Go installed on your system.

1.  **Clone the repository:**
    ```bash
    git clone <repository-url>
    cd <repository-directory>
    ```

2.  **Build the executable:**
    ```bash
    go build -o autodoc ./cli/
    ```

This will create an executable file named `autodoc` in your current directory. You can then move this executable to your system's PATH for easy access from any directory.

## Usage

The AUTODOC CLI provides several commands to manage your documentation workflow.

### Initializing a Project

Before generating documentation, you need to initialize your project. This command sets up the necessary configuration file.

```bash
./autodoc init
```

This command will:
1.  Fetch a default configuration template from the AI agent backend.
2.  Create an `autodoc_config.json` file in your current directory with sensible defaults.

### Generating Documentation

Once your project is initialized, you can generate documentation for your codebase.

```bash
./autodoc generate
```

This command performs the following steps:
1.  Gathers project metadata, including application name, description, and ignored files, from `autodoc_config.json`.
2.  Builds a file tree representation of your project's workspace.
3.  Identifies and reads the content of supported programming language files.
4.  Constructs a payload with project metadata and code content.
5.  Sends the payload to the AI agent backend for documentation generation.
6.  Receives the generated documentation from the backend.
7.  Creates or overwrites a `README.md` file in your project's root directory with the generated documentation.

### Checking Health

You can check the health and connectivity of the auto-doc AI agent backend servers.

```bash
./autodoc health
```

This command sends a GET request to the backend to verify its status and prints the response.

## Configuration

The AUTODOC CLI uses a configuration file named `autodoc_config.json` to manage project settings. This file is typically created by the `init` command.

The `autodoc_config.json` file has the following structure:

```json
{
  "name": "your-project-name",
  "developer": "Your Name",
  "description": "A brief description of your project.",
  "ignore": [
    "*.log",
    "temp/",
    "node_modules/"
  ]
}
```

-   **`name`**: The name of your project.
-   **`developer`**: The name of the project developer.
-   **`description`**: A short description of your project.
-   **`ignore`**: A list of file or directory patterns to exclude during the documentation generation process.

## Project Structure

The AUTODOC CLI has the following internal structure:

```
cli/
├── cmd/
│   ├── generate.go
│   ├── health.go
│   ├── init.go
│   └── root.go
├── main.go
├── models/
│   ├── AgentRequestParams.go
│   ├── AgentResponseParams.go
│   ├── CodeFile.go
│   ├── config.go
│   └── treenode.go
└── services/
    ├── generateservice.go
    ├── healthservice.go
    ├── initservice.go
    └── utils/
        ├── BuildTree.go
        ├── ConvertTreeToJSON.go (Deprecated)
        ├── CreateReadme.go
        ├── GetAppName.go
        ├── GetDescription.go
        ├── GetIgnoredFiles.go
        ├── GetJSONData.go
        ├── GetLanguageFilesFromAllPaths.go
        ├── GetNameFromPath.go
        ├── GetParentDirectory.go
        ├── GetTextFromFile.go
        ├── ReadLanguagePaths.go
        └── TraverseTreeForPaths.go
```

### Core Components

| Component                 | Description                                                                                                                            |
| :------------------------ | :------------------------------------------------------------------------------------------------------------------------------------- |
| **`cmd` Directory**       | Contains the definitions for the CLI commands (`generate`, `health`, `init`, `root`).                                                  |
| **`main.go`**             | The entry point of the application, responsible for executing the CLI commands.                                                        |
| **`models` Directory**    | Defines the data structures used throughout the application, including request/response parameters for AI agents and configuration.      |
| **`services` Directory**  | Houses the core logic for the CLI's functionalities, such as documentation generation, health checks, and project initialization.        |
| **`services/utils`**      | Contains various utility functions used by the services, including file system traversal, configuration loading, and API communication. |

## Contributing

We welcome contributions to the AUTODOC CLI! Please refer to the `CONTRIBUTING.md` file (if available) for guidelines on how to contribute.

## License

[Specify the license type here, e.g., MIT License, Apache License 2.0]