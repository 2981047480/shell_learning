import subprocess


class run_command():
    def __init__(self):
        pass

    def run_cmd(self, command, if_cmd_output=False, timeout=300):
        command_info = ''
        error_msg = ''
        res = ''
        if if_cmd_output == True:
            try:
                command_exec_info = subprocess.Popen(command, shell=True, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
                # command_info = command_exec_info.communicate(timeout=timeout)
                while command_exec_info.poll() is None:
                    line = command_exec_info.stdout.readline().decode('utf-8')
                    print(line)
                    res += line
            except Exception as e:
                error_msg = "Exception: {}".format(e)
                res += error_msg

        else:
            command_exec_info = subprocess.Popen(command, shell=True, stdout=subprocess.PIPE, stderr=subprocess.STDOUT,\
                                                    universal_newlines=True)
            try:
                res = ''
                stdout = ''
                stderr = ''
                (stdout,stderr) = command_exec_info.communicate(timeout=timeout)
                res += stderr if stderr else stdout
            except Exception as e:
                error_msg = "Timeout error: {}".format(e)
                command_exec_info.kill()
                res += error_msg


        return  res
    
run_command().run_cmd('ping -c 10 baidu.com', if_cmd_output=False, timeout=5)