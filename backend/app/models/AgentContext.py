from pydantic import BaseModel


class AgentContext(BaseModel):
    n: str  # program name
    d: str  # description found in autodoc config
    t: dict  # filetree
    a: list  # function signatures


def toJsonTemplate(self):
    return self.model_dump()
