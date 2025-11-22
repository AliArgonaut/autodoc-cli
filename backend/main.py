from fastapi import FastAPI
from app.api.health_check_route import router as health_check_router
from app.api.init_route import router as init_router


app = FastAPI()

app.include_router(health_check_router, prefix="/api")
app.include_router(init_router, prefix="/api")
