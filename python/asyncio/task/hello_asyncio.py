import asyncio


async def wait_hello_world(seconds):
    print("Hello")
    await asyncio.sleep(seconds)
    print("World!")


async def main():
    # 我们手动创建一个事件循环
    func_list = [
        wait_hello_world(2),
        wait_hello_world(3)
    ]
    for func in func_list:
        await func

if __name__ == "__main__":
    asyncio.run(main())
