from google.adk.agents.llm_agent import SequentialAgent, ParallelAgent, LLMAgent

model = "gemini-2.5-flash-lite"

introduction_agent = LLMAgent(
    model=model,
    name="introduction_agent",
    description="""given information about a project, generate, in markdown language ONLY, an introduction to a professional production README.md file 
        the intro should cover the project name and project description. Answer questions like "what is this project, what does it do, who made it, why?" 
        your output is put in the {intro_draft}. Be sure to create appropriate markdown elements that fit a production README. 
        """,
    output_key="intro_draft"
)

utility_agent = LLMAgent(
    model=model,
    name="utility_agent",
    description="given information about a project, generate, in markdown language, "
)

parallel_documentation_team = ParallelAgent(
    name="parallel_doc_team",
    sub_agents=[introduction_agent, utility_agent, funcdoc_agent]
)

root_agent = SequentialAgent(
    model=model,
    name='root_agent',
    description="the first reciever of a data object of class AgentContext",
    instruction="""you will receive  payload with info about a software module.
    The payload has four main parts:
        1. n – The name of the module or application.
        Example: "Autodoc CLI"
        2. d – A short description of the module.
        Example: "the CLI layer for a CLI application that auto-documents code"
        3. t – The file tree for the module. Each node has:
            n: the name of the file or directory
            t: the type (d for directory, f for file)
            c: a list of children if the node is a directory
            Example:
            {
                "n": "cmd",
                "t": "d",
                "c": [
                    {"n": "generate.go", "t": "f"},
                    {"n": "health.go", "t": "f"}
                ]
            }
        4. a – Function metadata for each file. Each entry includes:
            f: the file name
            s: a list of functions in that file. Each function has:
            n: function name
            p: parameters (may be null)
            r: return values (may be null)
            Example:
            {
                "f": "generate.go",
                "s": [
                    {"n": "init", "p": null, "r": null}
                ]
            }
        Your goal is to run each subagent with appropriate context
            """,
    sub_agents=[parallel_documentation_team, judger_agent]
)
