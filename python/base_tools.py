

class base_tools():
    '''
    一些常用的方法
    '''
    def __init__(self):
        self.byte_to_mb = lambda x: round(x/1024/1024)
        self.get_list_0 = lambda list1: map(lambda x: x.split(':')[0], list1)
        self.time_transform = lambda x : list(map(lambda x: datetime.fromtimestamp(x/1000).strftime("%Y-%m-%d %H:%M:%S"),x))