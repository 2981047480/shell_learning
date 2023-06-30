import psutil
import os
import logger
from IPy import IP

class get_system_info():
    def __init__(self):
        self.to_G=1024*1024*1024
        self.run_user = os.environ.get('USER')


    def get_memory_g(self) -> dict:
        '''
        获取内存信息
        '''
        mem_info = psutil.virtual_memory()
        to_G = self.to_G
        mem={}
        mem['total'] = str(int(mem_info.total/to_G))+'GB'
        mem['used'] = str(mem_info.used/to_G)+'GB'
        mem['free'] = str(mem_info.free/to_G)+'GB'
        mem['percent'] = str(int(mem_info.percent))+'%'
        mem['available'] = str(mem_info.available/to_G)+'GB'
        mem['swap'] = str(psutil.swap_memory().total/to_G)+'GB'
        return mem
    
    def get_cpu_info(self) -> dict:
        '''
        获取cpu信息
        '''
        cpu_info = {}
        cpu_info['count'] = psutil.cpu_count()
        cpu_info['count_logical'] = psutil.cpu_count_logical()
        cpu_info['true_logical'] = psutil.cpu_count(logical=False)
        return cpu_info
    
    def get_disk_info(self) -> dict:
        '''
        获取磁盘信息，返回值为磁盘信息和磁盘使用率
        '''
        disk_info = {}
        disk_usage = {}
        disk_info['partations'] = psutil.disk_partitions()
        for disk_column in disk_info['partations']:
            mount_point = disk_column.mountpoint()
            disk_usage[mount_point] = psutil.disk_usage(disk_name)
        return disk_info,disk_usage
    
    def env_info(self):
        current_login_user = psutil.users()
        current_user = os.environ['USER']

    def netmask_to_address(self,netmask):
        ip=IP(netmask)
        try:
            for i in ip:
                print(i)
        except Exception as e:
            print('Exception: {}').format(e)

    def ip_to_netmask(self,ip):
        try:
            netmask=IP(ip,make_net=True)
            print(netmask)

        except Exception as e:
            print('Exception: {}').format(e)
    
    def 
        

get_info=get_system_info()
ip_info=get_info.ip_to_netmask('192.168.1.0-192.168.1.3')

# memoryview=system_info.get_memory_g()
# for i in memoryview.keys():
#     print(memoryview[i])