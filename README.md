# Go Secret Leak Detector

A lightweight DevSecOps tool that scans files for leaked secrets such as API keys

## Features
- Fast regex secret detection
- Simple CLI usage
- Minimal dependencies

## Usage
```bash
go build -o leakhound .
./leakhound <file>

or

go run . <file>