import asyncio


def hello_loop():
    print("hello loop, 这个函数会被延迟调用")


async def main():
    loop = asyncio.get_event_loop()
    loop.call_soon(hello_loop)
    await asyncio.sleep(1)


if __name__ == "__main__":
    loop = asyncio.new_event_loop()
    try:
        loop.run_until_complete(main())
    finally:
        loop.close()
