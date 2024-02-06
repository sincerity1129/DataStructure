from collections import deque

stack = []
queue = deque()

exp_list = [1,2,3,4,5]

for data in exp_list:
    stack.append(data)
    queue.append(data)

for _ in exp_list:
    print("--------------")
    print(f"stack: {stack.pop()}")
    print(f"queue: {queue.popleft()}")
    print("--------------")
    

