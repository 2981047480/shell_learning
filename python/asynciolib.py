import os
import sys
import asyncio

async def hello_world():
    # await asyncio.sleep(1)
    print("hello world")

async def yeehoo():
    print("yeehoo")

async def main():
    await asyncio.gather(hello_world(), hello_world())
    
# asyncio.run(main())

class AsyncioRun():
    def __init__(self) -> None:
        pass

    def get_class_instance(self, dict_args):
        '''
        @desc: 用来返回某个文件中的某个类实例
        @contributor: zephyrzhao@icloud.com

        @input: dict_args: 字典,字典中的key如下:
                    module_name: 需要导入的函数名称(必须)
                    file_path: 需要导入的函数所在的路径(必须)
                    class_name: 类名
                    is_class: 是否在类中
        @return:  
                module实例(为一个class, 可以直接通过:
                get_class_instance().module_name()使用
        '''
        from importlib.util import spec_from_file_location,module_from_spec
        module_name = dict_args["module_name"]
        file_path = dict_args["file_path"]
        class_name = ''.join(list(map(lambda x: x.capitalize(), os.path.basename(file_path)[:-3].split('_'))))\
                    if not dict_args.get("class_name") else dict_args["class_name"]
        spec = spec_from_file_location(module_name, file_path)
        module = module_from_spec(spec)
        sys.modules[module_name] = module
        spec.loader.exec_module(module)
        if dict_args.get("is_class") != False:
            module = getattr(module, class_name)
            module = getattr(module(), module_name)
        else:
            module = getattr(module, module_name)
        return module

    def get_class_instance(self, module_name, file_path, class_name=None, is_class=True):
        '''
        @desc: 用来返回某个文件中的某个类实例
        @contributor: zephyrzhao@icloud.com

        @input: module_name: 需要导入的函数名称(必须)
                file_path: 需要导入的函数所在的路径(必须)
                class_name: 类名
                is_class: 是否在类中
        @return:  
                module实例(为一个class, 可以直接通过:
                get_class_instance().module_name()使用
        '''
        from importlib.util import spec_from_file_location,module_from_spec
        class_name = ''.join(list(map(lambda x: x.capitalize(), os.path.basename(file_path)[:-3].split('_'))))\
                    if not class_name else class_name
        spec = spec_from_file_location(module_name, file_path)
        module = module_from_spec(spec)
        sys.modules[module_name] = module
        spec.loader.exec_module(module)
        if is_class:
            module = getattr(module, class_name)
            module = getattr(module(), module_name)
        else:
            module = getattr(module, module_name)
        return module
    
    def gen_func(self, func, func_list):
        if func_list.__class__.__name__ != 'list':
            raise Exception("Func_list is not a list")
        if func.__class__.__name__ == 'list':
            for fun in func:
                func_list.append(self.gen_func(fun))
            return func_list
        elif func.__class__.__name__ == 'dict':
            for fun in func.values():
                func_list.append(self.gen_func(fun))
            return func_list
        elif func.__class__.__name__ == 'function':
            func_list.append(func)
            return func_list
        else:
            raise Exception("Func type: {}, choose from \
                            list/dict/function".format(func.__class__.__name__))

    async def mutli_func(self, func_list):
        for func in func_list:
            await func()
    
    async def gather(self, *args):
        await asyncio.gather(*args)
    
    def run(self, func):
        asyncio.run(func)

if __name__ == "__main__":
    ar = AsyncioRun()
    func_list = [hello_world, hello_world, yeehoo]
    fun = yeehoo
    # asyncio.run(ar.gen_func(factorial, func_list))
    # ar.run(ar.gather(hello_world(), yeehoo()))
    ar.run(ar.mutli_func(ar.gen_func(fun, func_list)))