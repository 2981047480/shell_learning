import logging
import logging.handlers
import os

class logger():
    def __init__(self):
        self.ERROR = "\E[1;31m"
        self.INFO = '\E[1;32m'
        self.WARN = "\E[1;33m"
        self.ZERO = '\033[0m'
        self.LOGGER_SETTINGS = {
            'ERROR' : self.ERROR,
            'INFO' : self.INFO,
            'WARN' : self.WARN,
            'ZERO' : self.ZERO
        }
        self.LOG_PATH = '/tmp/log'
        # levelname_color = LoggerColouredFormatter.COLOR_OF_LEVEL[levelname.upper()] + levelname + '\033[0m'
        # recoder.levelname = logging

    def log_info(self, massage):
        loglevel = messages.level
        self.LOG_FORMAT = "%(asctime)s - " + self.LOGGER_SETTINGS[loglevel] + "%(level) - %(message)s" + self.LOGGER_SETTINGS['ZERO']
        logging.basicConfig(filename = self.LOG_PATH, level = logging.WARNING, format = self.LOG_FORMAT)
        logging.Formatter.format(self.LOG_FORMAT)

        console = logging.StreamHandler()
        console.setLevel(loglevel)
        console.setFormatter(self.LOG_FORMAT)
        
        



    def test(self):
        logging.info(self.ERROR)

print_log=logger()
print_log.test()