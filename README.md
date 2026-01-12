# IntelliLog 🧠

**IntelliLog** is an advanced AI-powered logging and monitoring solution designed for Backend Engineers. It goes beyond traditional logging by incorporating real-time anomaly detection, auto-resolving capabilities, and a premium glassmorphism dashboard.

![IntelliLog Dashboard](https://via.placeholder.com/800x400?text=IntelliLog+Dashboard+Preview)
*(Note: Replace with actual screenshot if available)*

## 🚀 Features

-   **AI-Driven Analysis**: Automatically detects anomalies (e.g., High Latency, OOM) and provides severity scores and fix suggestions.
-   **High-Performance API**: Built with **Go** and **Echo**, ensuring low-latency log ingestion.
-   **Secure**: Implements **JWT Authentication** and **Rate Limiting**.
-   **Real-Time Dashboard**: A modern, dark-mode web interface to visualize logs and system health instantly.
-   **Docker Ready**: Fully containerized for easy deployment and orchestration.

## 🛠️ Tech Stack

-   **Backend**: Go (Golang), Echo Framework
-   **Frontend**: HTML5, CSS3 (Glassmorphism), Vanilla JS
-   **AI/ML**: Custom Rule-Based Analysis Engine (Simulated LLM)
-   **Infrastructure**: Docker, Docker Compose

## 📋 Prerequisites

-   [Docker](https://www.docker.com/) & Docker Compose
-   [Go 1.25+](https://go.dev/) (Optional, for local dev without Docker)

## ⚡ Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/aykutterzi/intellilog.git
cd intellilog
```

### 2. Run with Docker (Recommended)

```bash
docker-compose up --build
```

The application will start on port `8080`.

### 3. Access the Dashboard

-   Open your browser to: `http://localhost:8080`
-   **Username**: `admin`
-   **Password**: `password`

## 🔌 API Usage

IntelliLog exposes a RESTful API for log ingestion and retrieval.

### Authentication (Get Token)
```bash
POST /login
{
  "username": "admin",
  "password": "password"
}
```

### Ingest Log
```bash
POST /api/logs
Header: Authorization: Bearer <TOKEN>
{
  "level": "ERROR",
  "service": "payment-service",
  "message": "Connection timeout"
}
```

### Retrieve Logs
```bash
GET /api/logs
Header: Authorization: Bearer <TOKEN>
```

## 🧪 Testing

You can use the included script to verify the system:

```bash
chmod +x test_api.sh
./test_api.sh
```

## 📄 License

[MIT](LICENSE)
