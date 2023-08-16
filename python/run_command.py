import subprocess
import paramiko
import pexpect

class run_command():
    def __init__(self):
        pass

    def run_cmd(self, command, is_cmd_output=False, timeout=300):
        command_info = ''
        error_msg = ''
        res = ''
        if is_cmd_output == True:
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
    # 
    def run_command(self, command, logfile='exec.log'):
        # python_path = os.path.join('/usr/local/')
        command_info = ''
        try:
            command_exe_info = pexpect.spawn(command=command, encoding='utf-8')
            for i in command_exe_info.readlines():
                command_info += i
            return command_info
        except Exception as e:
            error_msg = "Command error: {}".format(e)
            print(error_msg)
            return Exception
        
    def run_command_sudo(self, command, is_sudo=False):
        sudo = 'sudo'
        if is_sudo:
            command = sudo + command
            command_output = self.run_cmd(command)
            return  command_output
        else:
            self.run_cmd(command)
            command_output = self.run_cmd(command)
            return command_output

print(run_command().run_command_sudo('ls -l /tmp/aaa'))