from fastapi import APIRouter
from ..services import generateservice
from ..models.AgentContext import AgentContext
router = APIRouter()


@router.post("/generate")
async def generate_route(req: AgentContext):
    result = await generateservice.generate(req)
    return result
