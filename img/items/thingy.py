from urllib.request import urlretrieve

for itemId in range(96, 9999):
    try:
        r = (0x1000000|0x159a55e5*(itemId)&0xFFFFFF)
        urlretrieve('http://n-3ds-pgl-contents.pokemon-gl.com/share/images/item/{}.png'.format(hex(r)[3:]), filename='{}.png'.format(itemId))
    except:
        pass
