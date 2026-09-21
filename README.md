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

我們要在每一個服務資料夾底下各自加入一個叫`Dockerfile`的檔案，不用副檔名，D要大寫，因為Docker只會抓取檔名完全符合檔案。
當你打完之後vscode的explorer那邊理論上就會有藍色鯨魚icon出現。

然後貼上我們的對應程式碼，四個服務的Dockerfile內容都不一樣喔。

接下來就要在根目錄建立我們的`docker-compose.yml`，會有紅色鯨魚icon。

也是一樣貼上我們的程式碼。


file tree大概會長這樣

```bash
docker-practice
    ├── docker-compose.yml   <---- new
    ├── frontend
    │   ├── Dockerfile       <---- new
    │   └── index.html
    ├── gateway
    │   ├── Dockerfile       <---- new
    │   ├── gateway.go
    │   └── go.mod
    ├── go_service
    │   ├── Dockerfile       <---- new
    │   ├── go.mod
    │   └── go_service.go
    ├── python_service
    │   ├── Dockerfile       <---- new
    │   ├── python_service.py
    │   └── requirements.txt
    └── README.md
```

在 Docker 的世界裡，每個 Container 都是一台完全獨立的虛擬機器。如果在程式碼中寫死連線到 `127.0.0.1`（localhost），它只會連到「Container 自己的內部」，而找不到外面的其他服務。

為了解決這個問題，Docker Compose 會自動建立一個虛擬網路。不同服務之間必須透過 `docker-compose.yml` 裡定義的 **「服務名稱」**（例如 `python_service` 或 `go_service`）當作 DNS 網址來互相溝通。



### 打包服務

我們的原始碼中已經備妥了各服務的 `Dockerfile` 以及總指揮圖 `docker-compose.yml`。只需在專案根目錄執行以下指令：

```bash
# 建立、打包並啟動所有服務 (--build 代表強制重新打包最新映像檔)
sudo docker compose up --build
```
這樣就會開始編譯、打包，並且結束後會自動run起來。


### 測試服務
啟動完成後，打開瀏覽器並前往：
```text
http://localhost
```

並且按按看按鈕 觀察一下是否有正確回應，並且可以觀察terminal看看docker有沒有正確印出我們埋的echo line.


---
[ Anotations ]

這邊是常用一些command

```bash
sudo docker compose up --build
```

- up: 建立並同時啟動所有服務的container
- --build 強制在啟動前重新打包最新的image, 確保所有修改都有被套用


```bash
### 【映像檔 (Image) 管理】
1. `docker images`                # 列出本機所有下載或打包好的 Image
2. `docker pull <image>`          # 從雲端下載 Image (例如 docker pull python:3.9)
3. `docker rmi <image>`           # 刪除指定的 Image
4. `docker build -t <name> .`     # 讀取當前目錄的 Dockerfile 並打包成 Image

### 【容器 (Container) 基礎操作】
5. `docker ps`                    # 列出「正在執行中」的 Container
6. `docker ps -a`                 # 列出「所有」Container (包含已停止的)
7. `docker run <image>`           # 從 Image 建立並啟動一個新的 Container
8. `docker stop <container>`      # 停止執行中的 Container
9. `docker start <container>`     # 重新啟動已經停止的 Container
10. `docker rm <container>`       # 刪除已停止的 Container

### 【進入與檢查容器】
11. `docker logs <container>`     # 查看特定 Container 的輸出日誌
12. `docker logs -f <container>`  # 即時追蹤日誌 (類似 tail -f，隨時監控)
13. `docker exec -it <cont> bash` # 進入運作中 Container 的終端機 (輕量環境可用 sh)
14. `docker inspect <container>`  # 查看 Container 的詳細底層設定與網路資訊

### 【Docker Compose (微服務管理)】
15. `docker compose up`           # 依照 yml 檔建立並同時啟動所有服務 (前台顯示 Log)
16. `docker compose up -d`        # 建立並啟動所有服務 (背景執行，實務部署最常用)
17. `docker compose down`         # 停止並刪除 compose 建立的所有 Container 與網路
18. `docker compose build`        # 不啟動服務，只強制重新打包有改動的 Image

### 【系統清理 (大掃除)】
19. `docker image prune`          # 清除打包過程中產生、沒名字的廢棄 Image (<none>)
20. `docker system prune`         # 清除所有停止的 Container、無用網路與 Image，釋放空間
```