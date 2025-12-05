from google.adk.agents import SequentialAgent
from ..CodeParsingAgent.agent import Code_Parsing_Agent
from ..DocsWriterAgent.agent import Docs_Writer_Agent


root_agent = SequentialAgent(
    name="root_agent",
    sub_agents=[Code_Parsing_Agent, Docs_Writer_Agent],
    description="a pipeline for documenting code automatically"
)
