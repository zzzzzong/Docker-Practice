# 這是用fast api框架實作
from fastapi import FastAPI
import uvicorn

app = FastAPI()

@app.get("/api/fastapi")
def handle_fastapi_request():
    return {"message": "FastAPI 收到了"}

if __name__ == "__main__":
    uvicorn.run(app, host="127.0.0.1", port=8001)
