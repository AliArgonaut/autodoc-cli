from pydantic import BaseModel


class ConfigTemplate(BaseModel):
    name: str
    developer: str
    description: str
    ignore: list[str]

    def toJsonTemplate(self):
        return self.model_dump()
