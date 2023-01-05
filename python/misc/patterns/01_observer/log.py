from event import subscribe
from utils import log

def handleUserRegisteredEvent(user):
    log(f"User registered with email address {user['email']}")

def setupLogEventHandlers():
    subscribe("user_registered", handleUserRegisteredEvent)