import logging
import os

class logger():
    def __init__(self):
        self.ERROR='%(asctime)s \E[1;31m ERROR \E[0m'
        self.INFO='%(asctime)s \E[1;32m INFO \E[0m'
        self.WARN="%(asctime)s \E[1;33m WARNING \E[0m"
        self.INFO_LOG = '%(asctime)s - %(levelname)s - %(message)s'
        logging.basicConfig(filename='/tmp/example.log')

    def test(self):
        logging.log(msg=self.ERROR,level=0)

print_log=logger()
print_log.test()