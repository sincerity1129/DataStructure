class LinkedList:
    def __init__(self, datas):      
        self.datas = datas
        self.default_list = dict(pointer = None, val = datas[0])
        
    def _append_linked_list(self, next_node, now_node):
        next_node["pointer"] = now_node
        return next_node["pointer"]
    
    def set_linked_list(self):
        datas = self.datas
        basic_list = self.default_list
        for data in datas:
            if data==datas[0]:
                continue
            basic_list = self._append_linked_list(
                                basic_list, dict(pointer = None, val = data))
        return self.default_list
        
    def update_linked_list(self, update_node, data):
        update_node = self._append_linked_list(
            update_node, dict(pointer = update_node[""], val = data))
        return self.default_list

if __name__ == "__main__":
    users = ["민수", "철수", "영희"]
    linked_tool = LinkedList(users)
    linked_tool.set_linked_list()
    result = linked_tool.update_linked_list(linked_tool.default_list["pointer"], "만수")
    print(result)
