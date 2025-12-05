from google.adk.agents.llm_agent import Agent
instructions = """

You receive a single JSON payload containing:
  - app_name
  - description
  - filetree
  - code (filename → content)

Follow this exact process:

1. for each codefile (that is, for each file in the code part of the payload) generate a semantic description of that file and the functions wihtin it. 
2. include function signatures, what variables are taken in and returned, and if any functions in the file reference other functions.
3. your goal is to create a semantic understanding of the code in a way that is useful for a documentation writer and return it to root agent



"""


Code_Parsing_Agent = Agent(
    model="gemini-2.5-flash-lite",
    name='Code_Parsing_Agent',
    description='an assistant that analyzes the semantic meaning of code',
    instruction=instructions,
    output_key="semantic_understanding"
)
