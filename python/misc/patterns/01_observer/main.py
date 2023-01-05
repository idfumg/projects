from slack import setupSlackEventHandlers
from log import setupLogEventHandlers
from user import register_user

setupSlackEventHandlers()
setupLogEventHandlers()

register_user("Bob", "1234", "bob@mail.com")
