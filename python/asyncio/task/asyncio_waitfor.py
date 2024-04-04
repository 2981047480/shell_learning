import time
import asyncio


async def wait_hello_world(seconds):
    print("Hello")
    await asyncio.sleep(seconds)
    print("World!")


async def main():
    # 我们手动创建一个事件循环
    func_list = [
        wait_hello_world(1),
        wait_hello_world(2),
        wait_hello_world(10)
    ]
    await_list = []
    for func in func_list:
        await_list.append(asyncio.create_task(func))
    for await_func in await_list:
        try:
            # await asyncio.wait_for(await_func, timeout=5)
            await asyncio.wait_for(asyncio.shield(await_func), timeout=5)
            # 有了shield，函数就不会被取消了
            print("任务已完成")
        except asyncio.TimeoutError:
            await await_func
            print(await_func.cancelled())
            print("任务超时")


if __name__ == "__main__":
    start_time = time.time()
    asyncio.run(main())
    end_time = time.time()
    print(f"程序运行时间：{end_time - start_time}秒")
