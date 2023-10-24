import asyncio
import time


async def request1():
    print("task 1")
    time.sleep(5)


async def request2():
    print("task 2")
    time.sleep(3)


async def main(loop: asyncio.AbstractEventLoop):
    tasks = [loop.create_task(fn()) for fn in [request1, request2]]
    return await asyncio.gather(*tasks)


if __name__ == '__main__':
    start = time.time()
    asyncio.run(main(asyncio.get_event_loop()))
    print(f"{(time.time() - start):.3f} seconds")
