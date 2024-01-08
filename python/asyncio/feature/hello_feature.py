import asyncio


def main():
    future = asyncio.Future()
    print(future.done())
    future.set_result("hello feature")
    print(future.done())
    return future


if __name__ == "__main__":
    future = main()
    print(future.result())
