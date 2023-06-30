

class base_tools():
    '''
    一些常用的方法
    '''
    def __init__(self):
        self.byte_to_mb = lambda x: round(x/1024/1024)
        self.get_list_0 = lambda list1: map(lambda x: x.split(':')[0], list1)