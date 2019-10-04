import pygame
from pygame.locals import *

class Player(pygame.sprite.Sprite):
    def __init__(self):        
        super(Player,self).__init__()
        self.surf = pygame.Surface((75,25))
        self.surf.fill((255,255,255))
        self.rect = self.surf.get_rect()

if __name__== '__main__':
    pygame.init()
    sceen = pygame.display.set_mode((800,600))
    player = Player()

    running = True
    while running:
        for event in pygame.event.get():
            if event.type == KEYDOWN:
                if event.key == K_ESCAPE:
                    running = False
                elif event.type==QUIT:
                    running = False

    # surf = pygame.Surface((50,50))
    # surf.fill((255,255,255))
    # rect = surf.get_rect()

    screen.blit(player.surf,(400,300))
    pygame.display.flip()