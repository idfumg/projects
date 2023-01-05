def create_user(name: str, password: str, email: str) -> dict:
    print('Creating a new user...')
    return {
        'name': name,
        'password': password,
        'email': email,
    }
    
def log(msg: str):
    print('Logging a new msg...', msg)