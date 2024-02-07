import asyncio


async def wait_hello_world(seconds):
    print("Hello")
    await asyncio.sleep(seconds)
    print("World!")


def main():
    loop = asyncio.get_event_loop()
    loop.create_task(wait_hello_world(3))
    loop.create_task(wait_hello_world(2))
    loop.run_until_complete(wait_hello_world(1))
    undone_tasks = asyncio.all_tasks(loop)
    print(undone_tasks)
    for task in undone_tasks:
        loop.run_until_complete(task)


if __name__ == "__main__":
    main()
