import wx

class LoginFrame(wx.Frame):
    def __init__(self,parent,id):
        wx.Frame.__init__(self,parent,id,title="登录&注册",size=(400,300))
        # 创建面板
        panel = wx.Panel(self)
               # 创建文本和输入框
        self.title = wx.StaticText(panel, label="请输入用户名和密码", pos=(140, 20))
        self.label_user = wx.StaticText(panel, label="用户名:", pos=(50, 50))
        self.text_user = wx.TextCtrl(panel, pos=(100, 50), size=(235, 25), style=wx.TE_LEFT)
        self.label_pwd = wx.StaticText(panel, pos=(50, 90), label="密   码:")
        self.text_password = wx.TextCtrl(panel, pos=(100, 90), size=(235, 25), style=wx.TE_PASSWORD)

        self.bt_login = wx.Button(panel, label='登录', pos=(105, 130))  # 创建“确定”按钮
        self.bt_regin = wx.Button(panel, label='注册', pos=(195, 130))  # 创建“取消”按钮
