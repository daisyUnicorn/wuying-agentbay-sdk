# Wuying AgentBay SDK

Wuying AgentBay SDK provides APIs for Python, TypeScript, and Golang to interact with the Wuying AgentBay cloud runtime environment. This environment enables running commands, executing code, and manipulating files.

## Features

- **Session Management**: Create, retrieve, list, and delete sessions
- **File Management**: Read files in the cloud environment
- **Command Execution**: Run commands
- **ADB Operations**: Execute ADB shell commands in mobile environments (Android)

## Installation

(Note: The following installation methods will be available in the future. Please refer to the current project documentation for setup instructions.)

### Python

```bash
pip install wuying-agentbay-sdk
```

## Usage

### Python

```python
from wuying_agentbay import AgentBay

# Initialize with API key
agent_bay = AgentBay(api_key="your_api_key")

# Create a session
session = agent_bay.create()

# Execute a command
result = session.command.execute_command("ls -la")
print(result)

# Read a file
content = session.filesystem.read_file("/path/to/file.txt")
print(content)

# Execute an ADB shell command (for mobile environments)
adb_result = session.adb.shell("ls /sdcard")
print(adb_result)
```

## Authentication

Authentication is done using an API key, which can be provided in several ways:

1. As a parameter when initializing the SDK
2. Through environment variables (`AGENTBAY_API_KEY`)

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.