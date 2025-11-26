from ..models.AgentContext import AgentContext
from google.adk.runners import InMemoryRunner
from app.agents.first_agent.agent import root_agent


async def generate(params: AgentContext):
    prompt = f"""
    programName : {params.n},
    programDescription: {params.d},
    programFileTree: {params.t},
    programFunctionSignatures: {params.a}
    """
    runner = InMemoryRunner(agent=root_agent)
    response = await runner.run_debug(prompt)
    print("========================================================================================================")
    print(response["validated_docs"])
    print("========================================================================================================")
    return response["validated_docs"]
