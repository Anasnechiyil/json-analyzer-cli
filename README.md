
# JSON Analyzer CLI

[![Go Report Card](https://goreportcard.com/badge/github.com/Anasnechiyil/json-analyzer-cli)](https://goreportcard.com/report/github.com/Anasnechiyil/json-analyzer-cli)

JSON Analyzer CLI is a command-line tool to analyze large JSON files efficiently. It provides commands for counting keys, filtering, querying, and handling other JSON data operations. This tool is particularly useful when dealing with massive JSON files that traditional tools struggle to handle.

## Features

- **Count** the number of keys in a JSON file
- **Filter** JSON data based on specific criteria
- **Query** specific fields and values
- Efficiently handles large files by streaming JSON data
- Customizable and extensible command structure

## Installation

### Prerequisites
- Go 1.18 or later

### Install
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/json-analyzer-cli.git
   cd json-analyzer-cli

2. Build the tool:
    ```bash
    go build -o json-analyzer-cli

3. (Optional) Move the executable to your $PATH:
    ```bash
    mv json-analyzer-cli /usr/local/bin/

### Usage

## Basic Command Structure
Run the tool by entering:
```bash
    json-analyzer-cli <command> [flags] <file>
```

### Commands
Count Keys in JSON

Count the total number of keys in a JSON file. This is helpful for quickly understanding the structure or size of the JSON data.

Usage:
```bash
    json-analyzer-cli count yourfile.json
```

Example:

Given a JSON file sample.json:
```json
    {
        "name": "Example",
        "details": {
            "age": 25,
            "location": "Unknown"
        }
    }
```
Running:

```bash
    json-analyzer-cli count sample.json
```

Outputs:

```plaintext
    Total keys: 2
```

## Future Commands (Coming Soon)

- Filter: Filter JSON by keys or values.

- Query: Query specific JSON fields and extract matching data.

## Development
Project Structure

- cmd/ - Contains command implementations for different CLI operations.

- main.go - The entry point for the CLI tool.

- README.md - Documentation for the project.

## Testing
Run tests to ensure the functionality of the tool:
```bash
    go test ./...
```
Add tests in the cmd/ directory for new commands and features.

## Contributing

We welcome contributions! Please follow the steps below to contribute to the project:

- Fork the repository.
- Create a new branch for your feature (git checkout -b feature-branch).
- Make your changes and commit them (git commit -m "Add a new feature").
- Push to the branch (git push origin feature-branch).
- Open a pull request.
Feel free to open issues for feature requests, bug reports, or general feedback.

## License

This project is licensed under the MIT License. See the <a href="LICENSE">LICENSE</a> file for details.

## Acknowledgments

- **jq** for inspiration on command-line JSON manipulation.
- **Cobra** for the CLI framework.
