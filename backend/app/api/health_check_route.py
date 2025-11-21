from fastapi import APIRouter
from ..services import healthcheckservice

router = APIRouter()


@router.get("/health")
def check_health():
    return healthcheckservice.check_health()
