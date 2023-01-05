from utils import create_user
from event import postEvent

def register_user(name: str, password: str, email: str):
    user = create_user(name, password, email)
    postEvent("user_registered", user)