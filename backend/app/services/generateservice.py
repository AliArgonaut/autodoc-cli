from google.adk.sessions import InMemorySessionService
from google.adk.runners import Runner
from fastapi import HTTPException
from ..models.AgentRequestParams import AgentRequestParams
from ..agents.root_agent.agent import root_agent
from google.genai import types
from dotenv import load_dotenv
import uuid
import os


load_dotenv()
API_KEY = os.getenv("GOOGLE_API_KEY")


async def generate(params: AgentRequestParams):
    try:
        APP_NAME = "agents"
        USER_ID = "test_user"
        SESSION_ID = str(uuid.uuid4())

        session_service = InMemorySessionService()

        await session_service.create_session(
            app_name=APP_NAME,
            user_id=USER_ID,
            session_id=SESSION_ID
        )

        runner = Runner(
            agent=root_agent,
            app_name=APP_NAME,
            session_service=session_service,
        )

        new_message = types.Content(
            role="user",
            parts=[types.Part(text=params.json())]
        )

        for event in runner.run(
            user_id=USER_ID,
            session_id=SESSION_ID,
            new_message=new_message
        ):
            final_docs_raw = event.actions.state_delta.get('final_docs')
            # print(final_docs_raw)

        if final_docs_raw.startswith('```'):
            final_docs_raw = final_docs_raw.split(
                '\n', 1)[1].rsplit('\n```', 1)[0]

        print(final_docs_raw)
        return {
            "success": True,
            "documentation": final_docs_raw
        }

    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
        print(f"ERROR in generate: {str(e)}")
        import traceback
        traceback.print_exc()
        return {
            "success": False,
            "documentation": f"Error: {str(e)}"
        }
