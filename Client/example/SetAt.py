# !/usr/bin/env python

import pygame
pygame.init()

'''
创建一个窗口，并设置窗口标题
元组，代表分辨率
色深
'''
screen = pygame.display.set_mode((640,480),0,32)
all_colors = pygame.Surface((4096,4096),depth=24)

for r in range(256):
    print (r + 1,"out of 256")
    x = (r&15)*256
    y = (r>>4)*256
    for g in range(256):
        for b in range(256):
            all_colors.set_at((x+g,y+b),(r,g,b))

pygame.image.save(all_colors,r"res\allcolors.bmp")