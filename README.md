# Docker 微服務打包

這是一個簡單的API Gateway與兩個backend services(FastAPI + Go)的微服務系統測試專案。

### 專案目的

本專案的主要目的為了解 SWE 是如何運用 Docker 進行環境隔離與打包的**正式部署工作流程 (Deployment Workflow)**。



## [ Development 階段 ]
我們在這個階段會先打造一個簡單的微服務系統(不會很難啦放心)，並且簡單測試看看可以過的話就進入到打包階段。

同時，為了對齊 Google 等大廠的開發標準，本專案採用 **Monorepo（單一儲存庫）** 結構。我們以「一個資料夾作為一個服務的 Scope (作用域)」，確保每個服務在 Docker 打包時，都能擁有完全獨立的建置環境與依賴，互不干擾。


### 專案架構
```text
docker-practice/
├── frontend/          # 前端: index.html
├── gateway/           # Gateway: gateway.go (Port 8080)
├── python_service/    # Python 服務: python_service.py (Port 8001)
└── go_service/        # Go 服務: go_service.go (Port 8002)
```

``` mermaid
graph LR
    Client[前端 index.html] -->|Port 8080| Gateway[Go Gateway gateway.go]
    Gateway -->|/api/fastapi| FastAPI[Python 服務 python_service.py:8001]
    Gateway -->|/api/go| Go[Go 服務 go_service.go:8002]
```

### 步驟一：初始化專案與 Git

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
   gateway/gateway
   go_service/go_service
   .vscode/
   .DS_Store
   ```

### 步驟二：建立服務資料夾
這邊一次建立四個服務的資料夾

```bash
mkdir frontend gateway python_service go_service
```

### 步驟三：設定各服務環境與依賴
依序執行以下commands來建立我們的環境

1. Python 服務環境：
   ```bash
   cd python_service
   python3 -m venv .venv
   source .venv/bin/activate
   echo -e "fastapi\nuvicorn" > requirements.txt
   pip install -r requirements.txt
   cd ..
   ```

2. Go 服務環境：
   ```bash
   cd gateway && go mod init gateway && go mod tidy && cd ..
   cd go_service && go mod init go_service && go mod tidy && cd ..
   ```

### 步驟四：準備程式碼檔案(前端頁面 + Gateway服務 + 兩個後端服務)
請建立以下四個檔案，並分別貼上我們的程式碼：

1. 前端: 建立 `frontend/index.html`
2. Gateway: 建立 `gateway/gateway.go`
3. Python 服務: 建立 `python_service/python_service.py`
4. Go 服務: 建立 `go_service/go_service.go`

### 步驟五：本地端測試 (Local Run)

請開啟三個不同的終端機（運行 Python 的需先執行 `source .venv/bin/activate`），分別執行：

1. 啟動 Python 服務：
   ```bash
   cd python_service
   source .venv/bin/activate
   python python_service.py
   ```

2. 啟動 Go 服務：
   ```bash
   cd go_service
   go run go_service.go
   ```

3. 啟動 Gateway：
   ```bash
   cd gateway
   go run gateway.go
   ```

4. 測試：用瀏覽器開啟 `frontend/index.html` 點擊測試按鈕。

### 步驟六：提交至 Git

確認測試無誤後，將目前的進度提交：
```bash
git add .
git commit -m "Initial commit: Setup basic microservices architecture"
```


## [ Docker打包階段 ]
這個階段就是我們的重頭戲了，Docker要來打包了。
實務上，大型微服務會將每個服務拆成獨立資料夾，像是Google之類的大廠會**依照服務拆分資料夾**，確保每個服務只打包自己需要的code和dependencies，不會互相干擾。

目前為止我們的檔案結構應該長這樣
```bash
docker-practice
    ├── frontend
    │   └── index.html
    ├── gateway
    │   ├── gateway.go
    │   └── go.mod
    ├── go_service
    │   ├── go.mod
    │   └── go_service.go
    ├── python_service
    │   ├── python_service.py
    │   └── requirements.txt
    └── README.md
```

我們要在每一個服務資料夾底下各自加入一個叫`Dockerfile`的檔案，不用副檔名，D要大寫。因為Docker只會抓取檔名完全符合檔案。
現在應該要變成這樣
```bash
docker-practice
    ├── frontend
    │   ├── Dockerfile
    │   └── index.html
    ├── gateway
    │   ├── Dockerfile
    │   ├── gateway.go
    │   └── go.mod
    ├── go_service
    │   ├── Dockerfile
    │   ├── go.mod
    │   └── go_service.go
    ├── python_service
    │   ├── Dockerfile
    │   ├── python_service.py
    │   └── requirements.txt
    └── README.md
```