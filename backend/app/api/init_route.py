
from fastapi import APIRouter
from ..services import initservice

router = APIRouter()


@router.get("/init")
def init():
    return initservice.config_file()
