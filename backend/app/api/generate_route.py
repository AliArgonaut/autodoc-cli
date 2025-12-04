from fastapi import APIRouter
from ..services import generateservice
from ..models.AgentRequestParams import AgentRequestParams
router = APIRouter()


@router.post("/generate", response_model=None)
async def generate_route(req: AgentRequestParams):
    result = await generateservice.generate(req)
    return result
