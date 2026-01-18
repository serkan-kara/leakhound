# Go Secret Leak Detector

![LeakHound Secret Scan](../../actions/workflows/leakhound.yml/badge.svg)

LeakHound is a lightweight DevSecOps CLI tool that scans files and directories for leaked secrets such as API keys

Its designed to be fast, simple and easy to integrate into security pipelines

LeakHound never prints full secret values.


## Features
- Scan single file or directories (recursive)
- File and line number reporting
- Multiple secret detectors (AWS, JWT, Private Keys)
- Detector-specific masking strategies
- Regex based secret detection
- Safe output (prevents secret reexposure)
- Ignoring common folders like .git, node_modules, vendor
- Skips common binary file types
- Simple CLI usage
- Minimal dependencies

## Currently Supported Detectors
- AWS Access Key ID (AKIA....)
- JWT Keys
- Private Keys

## CI / GitHub Actions
Leakhound returns non zero exit code when findings are detected.
This makes it suitable as a CI gate to prevent secret leaks from being merged.

## Output Format
file_path:line_number  [TYPE]  MASKED_VALUE

## Example
.env:5  [AWS_ACCESS_KEY_ID]  AKIA************CDEF
tokens.txt:1  [JWT]  eyJ***************c2P0
keys.pem:3  [PRIVATE_KEY]  REDACTED

## Usage
```bash
go build -o leakhound .
./leakhound <path or file>

or

go run . <path or file>