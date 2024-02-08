from collections import deque

class TreeNode:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None

def dfs(node):
    if node is not None:
        print(node.value, end=" ")  # 현재 노드의 값을 출력합니다.
        dfs(node.left)              # 왼쪽 서브트리를 재귀적으로 탐색합니다.
        dfs(node.right)             # 오른쪽 서브트리를 재귀적으로 탐색합니다.

def bfs(root):
    if root is None:
        return

    # 큐 생성 및 루트 노드 추가
    queue = deque()
    queue.append(root)

    # BFS 알고리즘을 사용하여 트리의 노드를 방문하고 출력
    while queue:
        node = queue.popleft()
        print(node.value, end=" ")  # 현재 노드의 값을 출력합니다.

        # 현재 노드의 자식 노드를 큐에 추가합니다.
        if node.left:
            queue.append(node.left)
        if node.right:
            queue.append(node.right)

# 이진 트리 생성
root = TreeNode(1)
root.left = TreeNode(2)
root.right = TreeNode(3)
root.left.left = TreeNode(4)
root.left.right = TreeNode(5)
root.right.left = TreeNode(6)
root.right.right = TreeNode(7)

# 전위 순회 실행
print("DFS:")
dfs(root)
print("\nBFS:")
bfs(root)