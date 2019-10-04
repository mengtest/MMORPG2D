import wx
from socket import *
import views.LoginFrame as lf
import pygame
from pygame.locals import *

def main():
    print("main is call!")
    pygame.init()
    screen = pygame.display.set_mode((800,600))

    running = True
    while running:
        for event in pygame.event.get():
            if event.type == KEYDOWN:
                if event.key == K_ESCAPE:
                    running = False
                elif event.type==QUIT:
                    running = False

    surf = pygame.Surface((50,50))
    surf.fill((255,255,255))
    rect = surf.get_rect()

    screen.blit(surf,(400,300))
    pygame.display.flip()


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
    main()
    # root = startClient()
    # root.MainLoop()