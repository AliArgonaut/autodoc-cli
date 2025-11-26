from google.adk.agents import SequentialAgent, ParallelAgent
from google.adk.agents.llm_agent import LlmAgent

model = "gemini-2.5-flash-lite"

introduction_agent = LlmAgent(
    model=model,
    name="introduction_agent",
    description="""given information about a project, generate, in markdown language ONLY, an introduction 
    to a professional production README.md file 
    the intro should cover the project name and project description. Answer questions like "what is this project, what does it do, who made it, why?" 
    your output is put in the {intro_draft}. Be sure to create appropriate markdown elements that fit a production README. Use only the fields from the original payload:
    n (program name), d (program description), t (filetree)
    """,
    output_key="intro_draft"
)

utility_agent = LlmAgent(
    model=model,
    name="utility_agent",
    description="""given information about a project, generate, in markdown language, a guide to installing and using the application. 
        have code examples or terminal arguments as markdown elements in the result. Use only the fields from the original payload:
         t (filetree), a (function signatures)
""",
    output_key="usage_draft"
)

judger_agent = LlmAgent(
    model=model,
    name="judger_agent",
    description="""You are the reviewer of documentation drafts.
    Given the project payload and the outputs from subagents:
    - intro_draft
    - usage_draft
    - funcdoc_draft
    Check for:
      1. Accuracy: do the drafts match the project name, description, and functions?
      2. Completeness: are all important sections present?
      3. Markdown formatting: headings, code blocks, lists, etc.
    Return an object with validated drafts after making improvements if needed.
    """,
    output_key="validated_docs"
)

funcdoc_agent = LlmAgent(
    model=model,
    name="funcdoc_agent",
    description="""given information about a project, generate, 
    in markdown language, a reference for the functions in the project. focus on important
    functions and services only, and give a more technical guide to understanding the codebase itself. 
    Use only the fields from the original payload:
         t (filetree), a (function signatures)
    """,
    output_key="funcdoc_draft"

)

parallel_documentation_team = ParallelAgent(
    name="parallel_doc_team",
    sub_agents=[introduction_agent, utility_agent, funcdoc_agent]
)

root_agent = SequentialAgent(
    name='root_agent',
    description="the first reciever of a piece of sotfware's metadata. Call the sub agents in order and then feed their collective outputs to a judger agent. Then, return the judger agent's validated_docs text ",
    sub_agents=[parallel_documentation_team, judger_agent]
)
