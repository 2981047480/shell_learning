from system_info import get_system_info
# from system_info import get_system_info


# class SystemInfo(object):
#     def __init__(self):
#         pass

#     def run(self):
#         memoryview=system_info().get_system_info().get_memory_g()
#         for i in memoryview.keys():
#             print(memoryview[i])

def run():
    memoryview=get_system_info()
    memorydict=memoryview.get_memory_g()
    for i in memorydict.keys():
        print(i,memorydict[i])

run()