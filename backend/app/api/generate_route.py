from fastapi import APIRouter
from ..services import generateservice
router = APIRouter()


@router.post("/generate")
async def generate_route(req):
    result = await generateservice.generate(req)
    return result
