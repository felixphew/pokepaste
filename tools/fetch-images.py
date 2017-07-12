#!/usr/bin/env python3

from os import makedirs
from urllib.request import urlretrieve
from urllib.error import URLError

makedirs('img/pokemon', exist_ok=True)

for mon in range(1, 803):
    for form in range(32): 
        try:
            code = hex(0x1000000 | 0x159a55e5 * (mon + form * 0x10000) & 0xFFFFFF)[3:]
            url = 'http://n-3ds-pgl-contents.pokemon-gl.com/share/images/pokemon/300/{}.png'.format(code)
            urlretrieve(url, filename='img/pokemon/{}-{}.png'.format(mon, form))
        except URLError:
            break
        else:
            print('{}.png -> {}-{}.png'.format(code, mon, form))

makedirs('img/items', exist_ok=True)

for item in range(921):
    try:
        code = hex(0x1000000 | 0x159a55e5 * item & 0xFFFFFF)[3:]
        url = 'http://n-3ds-pgl-contents.pokemon-gl.com/share/images/item/{}.png'.format(code)
        urlretrieve(url, filename='img/items/{}.png'.format(item))
    except URLError:
        pass
    else:
        print('{}.png -> {}.png'.format(code, item))
