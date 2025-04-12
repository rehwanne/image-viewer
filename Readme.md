# Single Image Display Service

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)
![Gin](https://img.shields.io/badge/Gin-1.9.1-00ADD8?logo=go)
![Docker](https://img.shields.io/badge/Docker-✓-2496ED?logo=docker)
![AI-Generated](https://img.shields.io/badge/AI-Generated-FFD700)

> **Note**: This project was generated with AI assistance (100% vibe coded) and has been verified/tested for functionality.

A minimalist web service that only displays the most recently uploaded image and automatically discards previous versions.

## Key Features

- 🖼️ **Single Image Policy**: Maintains only one image (`current.jpg`) at a time
- ⚡ **Auto-Refresh**: Displays updates without page reloads (2s interval)
- 📤 **Dual Upload**: Supports both curl and web form uploads
- 🗑️ **Automatic Cleanup**: Immediately discards older uploads
- 🐳 **Docker-Ready**: Production-ready containerization

## Technology Stack

- **Backend**: Go 1.21+ with Gin framework
- **Frontend**: Vanilla JS with auto-refresh
- **Container**: Alpine-based Docker image (~5MB runtime)
- **AI-Assisted**: Code generated with AI and manually verified

## Quick Start

### Prerequisites
- Go 1.21+ or Docker

### Run Natively
```bash
git clone https://github.com/your-repo/image-display-service.git
cd image-display-service
go mod download
mkdir images
go run .
