from pydantic import BaseModel
from typing import List, Dict, Any
from .CodeFile import CodeFile


class AgentRequestParams(BaseModel):
    app_name: str
    description: str
    Filetree: Dict[str, Any]
    Code: List[CodeFile]
