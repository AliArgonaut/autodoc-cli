from pydantic import BaseModel


class CodeFile(BaseModel):
    filename: str
    code: str
