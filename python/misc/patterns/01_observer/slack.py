from event import subscribe

def handleUserRegisteredEvent(user):
    postSlackMessage("sales", f"{user['name']} has registered with email address {user['email']}")
    
def postSlackMessage(group: str, msg: str):
    print(f"{group}: {msg}")
    
def setupSlackEventHandlers():
    subscribe("user_registered", handleUserRegisteredEvent)