import asyncio


def callback(future):
    print("Function callback", future.result())


async def main():
    future = asyncio.Future()
    future.add_done_callback(callback)
    future.set_result("111")

if __name__ == "__main__":
    asyncio.run(main())
