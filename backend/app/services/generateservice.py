from ..models.AgentContext import AgentContext
from google.adk.runners import Runner
from ..agents.first_agent import root_agent


def generate(params: AgentContext):

    prompt = f"""
    programName : {params.n},
    programDescription: {params.d},
    programFileTree: {params.t},
    programFunctionSignatures: {params.a}
    """
    runner = Runner(agent=root_agent)
    response = runner.run(prompt)
    print(response)
