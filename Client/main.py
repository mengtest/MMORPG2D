import wx
from socket import *
import views.LoginFrame as lf

class App(wx.App):
    def OnInit(self):
        # frame = wx.Frame(parent = None,title ='第一个窗口程序')
        # frame.Show()
        return True

class MyFrame(wx.Frame):
    def __init__(self,parent,id):
        wx.Frame.__init__(self,parent,id,title="测试",pos=(100,100),size=(300,300))
        panel = wx.Panel(self)
        wx.StaticText(panel,label='李维民',pos=(50,50))


def main():
    print("main is call!")
    # 加载数据

def startClient():
    BUFSIZE = 1024
    ADDR = ('127.0.0.1',8565)

    while True:
        data = input('>')
        if not data:
            break
        tcpClisock = socket(AF_INET,SOCK_STREAM)
        tcpClisock.connect(ADDR)
        tcpClisock.send(data.encode())
        data = tcpClisock.recv(BUFSIZE).decode()
        print(data)
        tcpClisock.close()

if __name__== '__main__':
    # main()
    # app = App()
    # # frame = MyFrame(parent=None,id=1)
    # loginframe = lf.LoginFrame(parent=None,id = 1)
    # loginframe.Show()
    # app.MainLoop()
    root = startClient()
    root.MainLoop()