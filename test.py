#coding=utf-8
import urllib

def getHtml(url):
    page = urllib.urlopen(url)
    html = page.read()
    return html

html = getHtml("http://mp.weixin.qq.com/s?__biz=MjM5MjAxNDM4MA==&mid=2666146678&idx=4&sn=cea40dd45ce7a2733204d5d1faa4906a&scene=0#wechat_redirect")


f=file("test.html","w+")

# li=["hello world\n","hello china\n"]

f.writelines(html)

f.close()

print html