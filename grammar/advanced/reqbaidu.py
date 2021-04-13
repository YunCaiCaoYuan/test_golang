import requests

def touch_baidu():
    res = requests.get("http://www.baidu.com")
    return res
