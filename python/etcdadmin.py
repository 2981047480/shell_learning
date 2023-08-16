import os
import sys
import json
import yaml
import etcd3
import argparse
sys.path.append(os.path.dirname(__file__))

class etcd_get():
    def __init__(self) -> None:
        self.etcdclient = etcd3.client()
        self.etcd_decode = lambda x: x[0].decode('utf-8')
        self.cm = ConfigManager()

    def read_args(self):
        '''
        读参数的模块
        '''
        parser = argparse.ArgumentParser(description='Get etcd conf')

        status = argparse.ArgumentParser(add_help=False)
        config_set.add_argument('-p','--path',help='Path')
        sub_parser = parser.add_subparsers(dest='operation')
        sub_parser.add_parser("status", parents=[status])

        config = argparse.ArgumentParser(add_help=False)
        config_get = argparse.ArgumentParser(add_help=False)
        config_set.add_argument('-p','--path',help='Path')
        config_parser = config.add_subparsers(dest="config_operation")
        config_parser.add_parser('get', parents=[config_get])

        config_set = argparse.ArgumentParser(add_help=False)
        config_set.add_argument('-p','--path',help='Path')
        config_parser.add_parser('set', parents=[config_set])
        sub_parser.add_parser('config', parents=[config])
        
        return parser.parse_args()


    def get_sensors_config(self):
        args = self.read_args()
        path = self.get_path()
        print(path)
        if args.config_operation == 'get':
            status = self.etcd_get(path)
            print(status)
        elif args.config_operation == 'set':
            self.etcd_put(path, args.value)
        # return status


    def get_prefix(self, path):
        res = self.etcdclient.get_prefix_response(path)
        return res 


    def etcd_get(self, path):
        try:
            if path:
                res = self.etcdclient.get(path)
                # etcd_decode = lambda x: x[0].decode('utf-8')
                # return json.loads(res[0].decode('utf-8')) #这里其实是个byte类型的，所以要先给转一下编码，然后再用json读进来
                if res[0]!=None:
                    return self.etcd_decode(res)
                else:
                    return None
            else:
                print('Please enter path you want to get from etcd')
        except Exception as e:
            print('Error: {}'.format(e))


    def etcd_put(self, path, value):
        etcdclient = self.etcdclient
        cm=self.cm
        try:
            if path and self.etcd_get(path):
                print('Key: {} is existing.'.format(path))
                table = cm.create_entire_table(['old_value', 'new_value'], [self.etcd_get(path), value])
                print(table)
                res = input('Please enter Yes/No:\n')
                if res.lower() == 'yes':
                    etcdclient.put(path, value)
                else:
                    print('You entered is not yes')
            elif path and self.etcd_get(path)==None:
                etcdclient.put(path, value)
            else:
                raise ValueError('Path you entered is None.')
        except Exception as e:
            print('Error: {}'.format(e))


if __name__ == '__main__':
    main = etcd_get()
    main.get_sensors_config()
    # res = main.get_prefix('/sensors_analytics/sp/client/')
    # for i in res.kvs:
    #     print(i.key.decode('utf-8'), main.etcd_decode(i.value))
    # config = main.etcd_get('/sensors_analytics/sp/client/redis')
    # print(ConfigManager().create_entire_table(['/sensors_analytics/sp/client/redis'], [yaml.dump(config)]))