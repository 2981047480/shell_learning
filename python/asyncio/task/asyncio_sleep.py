import time
import asyncio


async def wait_hello_world(seconds):
    print("Hello")
    await asyncio.sleep(seconds)
    print("World!")


async def main():
    # 我们手动创建一个事件循环
    task1 = asyncio.create_task(
        wait_hello_world(2)
        )
    await task1

if __name__ == "__main__":
    start_time = time.time()
    asyncio.run(main())
    end_time = time.time()
    print(f"程序运行时间：{end_time - start_time}秒")
