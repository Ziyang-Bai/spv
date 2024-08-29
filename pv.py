import time
import threading
from queue import Queue
import requests

def worker(q, url, delay):
    while not q.empty():
        start_time = time.time()
        response = requests.get(url)
        response_time = time.time() - start_time
        print(f"线程 {threading.current_thread().name} 响应时间: {response_time} 秒")
        time.sleep(delay)

def main():
    url = input("请输入要访问的URL: ")
    count = int(input("请输入要访问的次数: "))
    interval = int(input("请输入每次访问之间的间隔时间（秒）: "))
    threads = int(input("请输入线程数: "))

    if threads <= 0:
        print("线程数必须为正整数！")
        return

    q = Queue()
    for _ in range(count):
        q.put(1)

    threads_list = []
    for i in range(threads):
        t = threading.Thread(target=worker, args=(q, url, interval / 1000))
        t.start()
        threads_list.append(t)

    for t in threads_list:
        t.join()

if __name__ == "__main__":
    main()
