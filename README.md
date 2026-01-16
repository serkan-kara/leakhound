# Go Secret Leak Detector

LeakHound is a lightweight DevSecOps CLI tool that scans files and directories for leaked secrets such as API keys

Its designed to be fast, simple and easy to integrate into security pipelines


## Features
- Scan single file or directories (recursive)
- File and line number reporting
- Regex based secret detection
- Ignoring common folders like .git, node_modules, vendor
- Skips common binary file types
- Simple CLI usage
- Minimal dependencies

## Currently Supported Detectors
- AWS Access Key ID (AKIA....)

## Output Format
file_path:line_number  [TYPE]  MATCH

## Example
config.env:5  [AWS_ACCESS_KEY_ID]  AKIA************CDEF

## Usage
```bash
go build -o leakhound .
./leakhound <path or file>

or

go run . <path or file>