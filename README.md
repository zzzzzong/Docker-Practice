# Docker Microservices Practice Project

這是一個簡單的API Gateway與兩個backend services(FastAPI + Go)的微服務系統測試專案。

## 專案目的

本專案的主要目的為**系統設計入門**，並透過實際操作，了解如何運用 Docker 進行環境隔離與打包的**正式部署工作流程 (Deployment Workflow)**。


## 專案架構
- **前端**: `index.html` (純 HTML/JS)
- **Gateway**: `gateway.go` (Go 實作的反向代理，Port 8080)
- **Python 服務**: `python_service.py` (FastAPI，Port 8001)
- **Go 服務**: `go_service.go` (Go 內建 HTTP，Port 8002)
``` mermaid
graph LR
    Client[前端 index.html] -->|Port 8080| Gateway[Go Gateway gateway.go]
    Gateway -->|/api/fastapi| FastAPI[Python 服務 python_service.py:8001]
    Gateway -->|/api/go| Go[Go 服務 go_service.go:8002]
```

## 步驟一：初始化專案與 Git

1. 建立專案資料夾並進入：
   ```bash
   mkdir docker-practice
   cd docker-practice
   git init
   ```

2. 建立 `.gitignore` 檔案：
   ```text
   .venv/
   __pycache__/
   *.pyc
   .env
   gateway
   backend
   .vscode/
   .DS_Store
   ```

## 步驟二：設定 Python 環境 (FastAPI)

1. 建立並啟動虛擬環境：
   ```bash
   python3 -m venv .venv
   source .venv/bin/activate
   ```

2. 建立 `requirements.txt`：
   ```text
   fastapi
   uvicorn
   ```

3. 安裝套件：
   ```bash
   pip install -r requirements.txt
   ```

## 步驟三：設定 Go 環境

初始化 Go 模組：
   ```bash
   go mod init docker-practice
   go mod tidy
   ```

## 步驟四：準備前端頁面 + Gateway服務 + 兩個後端服務
請建立以下四個檔案，並分別貼上我們的程式碼：

1. 前端: 建立 `index.html`
2. Gateway: 建立 `gateway.go`
3. Python 服務: 建立 `python_service.py`
4. Go 服務: 建立 `go_service.go`

## 步驟五：本地端測試 (Local Run)

請開啟三個不同的終端機（運行 Python 的需先執行 `source .venv/bin/activate`），分別執行：

1. 啟動 Python 服務：
   ```bash
   python python_service.py
   ```

2. 啟動 Go 服務：
   ```bash
   go run go_service.go
   ```

3. 啟動 Gateway：
   ```bash
   go run gateway.go
   ```

4. 測試：用瀏覽器開啟 `index.html` 點擊測試按鈕。

## 步驟六：提交至 Git

確認測試無誤後，將目前的進度提交：
```bash
git add .
git commit -m "Initial commit: Setup basic microservices architecture"
```