import logging
import logging.handlers
import os


class Logger():
    def __init__(self):
        self.ERROR = "\E[1;31m"
        self.INFO = '\E[1;32m'
        self.WARN = "\E[1;33m"
        self.ZERO = '\033[0m'
        self.LOGGER_COLOR = {
            'ERROR' : self.ERROR,
            'INFO' : self.INFO,
            'WARN' : self.WARN,
            'ZERO' : self.ZERO
        }
        self.red = f'\033[31m'
        self.green = f'\033[32m'
        self.yellow = f'\033[33m'
        self.reset = f'\033[0m'
        self.level_dict = {
            'INFO': self.green + 'INFO ' + self.reset,
            'WARNING': self.yellow + 'WARNING ' + self.reset,
            'ERROR': self.red + 'ERROR ' + self.reset
        }
        self.LOG_PATH = '/tmp/logger.log'

    def generate_logger(self, level='INFO', filename=None):
        log_level = self.level_dict[level]
        self.LOG_FORMAT = "%(asctime)s {} %(message)s".format(log_level)
        logger = logging.basicConfig(filename=filename, format = self.LOG_FORMAT)
        return logger

    def log_info(self, message, level='INFO'):
        logger = self.generate_logger(level)
        # logger.info(message)
        logging.warning(message)
        
        



#     def test(self):
#         logging.info('Testing')

print_log=Logger()
print_log.log_info(message='error test')