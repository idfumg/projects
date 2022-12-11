# import asyncio
# from websockets import connect

# async def hello(uri):
#     async with connect(uri) as websocket:
#         await websocket.send("Hello world!")
#         msg = await websocket.recv()
#         print(f"Received: {msg}")

# asyncio.run(hello("ws://localhost:8765"))

import asyncio
import pathlib
import ssl
import websockets

ssl_context = ssl.SSLContext(ssl.PROTOCOL_TLS_CLIENT)
localhost_pem = pathlib.Path(__file__).with_name("localhost.pem")
ssl_context.load_verify_locations(localhost_pem)

async def hello():
    uri = "wss://localhost:8765"
    async with websockets.connect(uri, ssl=ssl_context) as websocket:
        name = input("What's your name? ")

        await websocket.send(name)
        print(f">>> {name}")

        greeting = await websocket.recv()
        print(f"<<< {greeting}")

if __name__ == "__main__":
    asyncio.run(hello())