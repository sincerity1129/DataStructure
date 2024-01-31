class Node:
    def __init__(self, data):
        self.data = data
        self.pointer = None
        
class LinkedList:
    @classmethod
    def _append(cls, now_node, next_node, final=True):
        '''
        now_node: 현재 노드 가져옴
        next_node: now_node.pointer를 통해 다음 포인터로 이동
        filnal: 마지막에 노드 추가를 위한 로직 수행(중간 삽입의 경우 로직 제외)
        '''
        # 마지막에 추가 해주는 로직
        if final and now_node.pointer != None:
            while now_node.pointer:
                if now_node.pointer==None:
                    break
                now_node=now_node.pointer
        now_node.pointer = next_node
        return next_node
    
    @classmethod
    def set_linked_list(cls, basic_node, datas, index=None, exist=False, final=True):
        '''
        basic_node: 기본적으로 세팅된 노드 정보
        datas: 넣고자 하는 데이터 정보
        index: 중간에 노드 삽입 필요 시 적용 옵션(기본 마지막에 노드 추가)
        exist: 중간에 노드 추가 시 적용 옵션(default:False-제일 뒤에 노드 추가)
        '''
        now_node = basic_node
        # 중간 삽입을 위한 로직
        if index!=None:
            try:
                for _ in range(index): 
                    now_node = now_node.pointer
                exist_node = now_node.pointer 
            except:
                print(f"{index}번째 노드가 존재하지 않습니다.")
                return

        for data in datas:
            now_node = cls._append(now_node, Node(data), final)
        if exist:
            now_node.pointer=exist_node
        return basic_node
    
    @classmethod
    def _delete(cls, delete_node, exist_node):
        delete_node.pointer = exist_node
    
    @classmethod
    def delete_linked_list(cls, basic_node, index):
        '''
        basic_node: 전체 LinkedList node
        index: 삭제하고자 하는 값
        '''
        try:
            node_page = 0
            now_node = basic_node
            while now_node.pointer:
                if node_page==(index-1):
                    delete_node = now_node.pointer
                elif node_page==(index+1):
                    exist_node = now_node.pointer
                    break
                now_node = now_node.pointer
                node_page+=1
            cls._delete(delete_node, exist_node)
        except:
            print(f"{index}번째 노드가 존재하지 않습니다.")
            return
        return basic_node
        
if __name__ == "__main__":
    data_list = ["민수", "철수", "영희"]
    # 초기 값 세팅 필수
    basic_node = Node(data_list[0])
    # 리스트 넣기
    LinkedList.set_linked_list(basic_node, data_list[1:])
    # 가장 마지막에 넣기
    LinkedList.set_linked_list(basic_node, "b")
    # 중간에 값 넣기
    result = LinkedList.set_linked_list(basic_node, ["1","2"], index=2, exist=True,final=False)
    # 삭제
    result = LinkedList.delete_linked_list(basic_node, index=2)
    print(result)
    