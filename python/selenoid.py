#!/usr/bin/python

from selenium import webdriver
import argparse
import time
import webbrowser
import os
import sys
import logging
import time
from datetime import datetime


def main():
    logging.info('Retrieving browser from 127.0.0.1:4444')
    capabilities = {
        'browserName': 'chrome',
        'version': '79.0',
        'enableVideo': True,
        'enableLog': True,
        'screenResolution': '1024x768x24',
        'name': "qaChallageAccepted",
    }
    url='https://duckduckgo.com'

    driver = webdriver.Remote(desired_capabilities=capabilities, command_executor='http://127.0.0.1:4444/wd/hub')
    time.sleep(30)
    try:
        print(driver.session_id)
        driver.get(url)
        FILENAME = datetime.now().strftime('%Y-%m-%d_%H:%M:%S') + '.png'
        driver.get_screenshot_as_file(FILENAME)
    except Exception as e:
        print(e)
    finally:
        driver.quit()
        logging.info('Done')



if __name__ == '__main__':
    main()