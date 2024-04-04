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
    count = 0
    while not await_list[-1].done():
        if count > 5:
            await_list[-1].cancel()
        print("检测到任务尚未完成，一秒钟之后继续检测")
        await asyncio.sleep(1)
        count = count + 1
    for await_func in await_list:
        try:
            await await_func
            print("任务已完成")
        except asyncio.CancelledError:
            print("任务被取消")


if __name__ == "__main__":
    start_time = time.time()
    asyncio.run(main())
    end_time = time.time()
    print(f"程序运行时间：{end_time - start_time}秒")
