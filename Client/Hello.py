# !/usr/bin/env python

# 指定图片文件名称
background_image_filename = r'res\sushiplate.jpg'
mouse_image_filename = r'res\fugu.png'


import pygame
from pygame.locals import *
from sys import exit

# 初始化pygame，为使用硬件做准备
pygame.init()

'''
创建一个窗口，并设置窗口标题
元组，代表分辨率
色深
'''
screen = pygame.display.set_mode((640,480),0,32)
pygame.display.set_caption("Hello,World!")

# 加载并转换图像
background = pygame.image.load(background_image_filename).convert()
mouse_cursor = pygame.image.load(mouse_image_filename).convert_alpha()

# 游戏主循环
while True:
    for event in pygame.event.get():
        # 接收到退出事件，退出程序
        if event.type == QUIT:
            exit()
    # 将背景图画上去
    screen.blit(background,(0,0))
    
    # 获得鼠标位置
    x,y = pygame.mouse.get_pos()
    # 计算光标的左上角位置
    x-= mouse_cursor.get_width()/2
    y-= mouse_cursor.get_height()/2
    '''
    绘制光标
    绘制后调用update更新，否则画面一片漆黑
    '''
    screen.blit(mouse_cursor,(x,y))

    # 刷新画面
    pygame.display.update()