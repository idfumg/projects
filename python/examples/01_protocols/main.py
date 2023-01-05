from typing import Protocol


class EmailServer(Protocol):
    def connect(self, host: str, port: int) -> None:
        ...

    def starttls(self) -> None:
        ...

    def login(self, login: str, password: str) -> None:
        ...

    def quit(self) -> None:
        ...


class EmailClient:
    def __init__(
        self,
        smtp_server: EmailServer,
    ):
        self.server = smtp_server

    def connect(self, host: str, port: int) -> None:
        self.server.connect(host, port)

    def quit(self) -> None:
        self.server.quit()


class EmailServerInstance:
    def connect(self, host: str, port: int) -> None:
        print('connect')

    def starttls(self) -> None:
        print('starttls')

    def login(self, login: str, password: str) -> None:
        print('login')

    def quit(self) -> None:
        print('quit')


if __name__ == '__main__':
    server = EmailServerInstance()
    client = EmailClient(server)
    client.quit()
