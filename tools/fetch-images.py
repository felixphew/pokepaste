#!/usr/bin/env python3

from os import makedirs
from urllib.request import urlretrieve
from urllib.error import URLError

makedirs('img/pokemon', exist_ok=True)

for n in range(1, 802):
    for i in range(32): 
        try:
            r = hex(0x1000000|0x159a55e5*(n+i*0x10000)&0xFFFFFF)[3:]
            urlretrieve('http://n-3ds-pgl-contents.pokemon-gl.com/share/images/pokemon/{}.png'.format(r), filename='img/pokemon/{}-{}.png'.format(n, i))
        except URLError:
            continue

makedirs('img/items', exist_ok=True)

for itemId in range(921):
    try:
        r = hex(0x1000000|0x159a55e5*(itemId)&0xFFFFFF)[3:]
        urlretrieve('http://n-3ds-pgl-contents.pokemon-gl.com/share/images/item/{}.png'.format(r), filename='img/items/{}.png'.format(itemId))
    except URLError:
        pass
