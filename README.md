# Autodoc Cli

This project provides a command-line interface (CLI) tool for automatically generating documentation for your projects. It leverages AI agents to analyze your codebase and create comprehensive README files.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Automatic README Generation**: Analyzes your project's structure and code to generate a well-formatted README file.
- **Code Semantic Analysis**: Utilizes AI agents to understand the meaning and relationships within your code.
- **Customizable Configuration**: Allows you to configure aspects like ignored files and project descriptions.
- **FastAPI Backend**: A robust backend service to process documentation generation requests.
- **CLI Interface**: A user-friendly command-line tool for interacting with the documentation generation service.

## Installation

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/AliArgonaut/Auto-Doc_Release.git
    cd Auto-Doc_Release
    ```

2.  **Set up the Backend**:
    *   Ensure you have Python and pip installed.
    *   Install the backend dependencies:
        ```bash
        cd backend
        pip install -r requirements.txt
        ```
    *   Set up your Google API key by creating a `.env` file in the `backend` directory with the following content:
        ```
        GOOGLE_API_KEY=YOUR_API_KEY
        ```
    *   Run the FastAPI backend server:
        ```bash
        uvicorn app.main:app --reload
        ```

3.  **Set up the CLI**:
    *   Ensure you have Go installed.
    *   Navigate to the `cli` directory:
        ```bash
        cd ../cli
        ```
    *   Build the CLI executable:
        ```bash
        go build -o autodoc .
        ```

## Usage

### Initialize a Project

Before generating documentation, initialize your project with the `init` command. This will create a configuration file (`autodoc_config.json`) for your project.

```bash
# Navigate to your project's root directory
cd /path/to/your/project

# Run the init command from the cloned Autodoc_Cli/cli directory after building
./autodoc init
```

This command will communicate with the backend to generate a default `autodoc_config.json` file.

### Configure Your Project

Open the generated `autodoc_config.json` file in your project's root directory and customize it as needed:

```json
{
  "name": "Your Project Name",
  "developer": "Your Name",
  "description": "A brief description of your project.",
  "ignore": [
    ".env",
    "venv",
    ".gitignore",
    "node_modules"
  ]
}
```

-   `name`: The name of your application.
-   `developer`: The name of the developer or team.
-   `description`: A short description of the project.
-   `ignore`: A list of file or directory patterns to exclude from documentation generation.

### Generate Documentation

Once your project is initialized and configured, run the `generate` command from your project's root directory:

```bash
# Ensure you are in your project's root directory
./autodoc generate
```

This command will:
1.  Gather metadata about your project (configuration, file structure).
2.  Send the code and metadata to the backend API.
3.  Receive the generated documentation from the backend.
4.  Create a `README.md` file in your project's root directory with the generated documentation.

### Health Check

You can check the health and connectivity of the backend API using the `health` command:

```bash
./autodoc health
```

This command will ping the backend server to ensure it's running and responsive.

## Example Usage

Assuming you have a simple Python project with the following structure:

```
my_python_project/
├── main.py
├── utils.py
└── autodoc_config.json
```

And your `autodoc_config.json` is set up as follows:

```json
{
  "name": "My Python Project",
  "developer": "Dev Team",
  "description": "A sample project to demonstrate Autodoc CLI.",
  "ignore": []
}
```

After running `./autodoc generate` from the `my_python_project` directory, you would find a `README.md` file generated with content similar to this:

```markdown
# My Python Project

A sample project to demonstrate Autodoc CLI.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)

## Features

- **Description of Functionality**: This section would detail the core features of your project.
- **Key Components**: Outline the main parts or modules of your application.

## Installation

```bash
# Example installation instructions
pip install -r requirements.txt
```

## Usage

```python
# Example usage of your project's code
from utils import helper_function

result = helper_function(10)
print(result)
```

## Codebase Overview

### `main.py`

This file serves as the main entry point for the application.

- **`main_function(input_data)`**:
    -   **Description**: Processes the input data and orchestrates other operations.
    -   **Input**: `input_data` (type: any) - The data to be processed.
    -   **Output**: (type: any) - The result of the processing.
    -   **References**: Calls `helper_function` from `utils.py`.

### `utils.py`

This file contains utility functions for the project.

-   **`helper_function(number)`**:
    -   **Description**: A helper function that performs a specific calculation.
    -   **Input**: `number` (type: int) - The input number.
    -   **Output**: (type: int) - The result of the calculation.

## Contributing

Contributions are welcome! Please follow these steps:

1.  Fork the repository.
2.  Create a new branch for your feature (`git checkout -b feature/your-feature`).
3.  Make your changes and commit them (`git commit -m 'Add some feature'`).
4.  Push to the branch (`git push origin feature/your-feature`).
5.  Open a Pull Request.

Please ensure your code adheres to the project's coding standards and includes relevant tests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
