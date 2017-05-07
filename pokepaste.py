#!/usr/bin/env python3

__version__ = '0.9'
__author__ = 'Felix Friedlander <felixphew0@gmail.com>'

from string import Template
from wsgiref.simple_server import make_server
from Crypto.Cipher import Blowfish

import re
import sqlite3
import os
import json
import cgi
import html

try:
    import crypto_secrets
except ModuleNotFoundError:
    print('''It looks like you're missing crypto_secrets.py.

crypto_secrets.py contains all the secrets necessary for the various
URL obfuscation schemes that PokePaste has used. This includes the mod
and two keys used for the original obfuscation scheme, and the blowfish
key for the new scheme.

If you're setting up your own instance of PokePaste, the most important
parameter is crypto_secrets.key, which should be a bytestring of 448
bits or shorter.
''')
    exit()

cipher = Blowfish.new(crypto_secrets.key)

# We pass image requests to open() basically unmodified,
# so this regex is needed to filter out any funny business.
img_re = re.compile(r'img/(pokemon/\d+-\d+|items/\d+).png')

conn = sqlite3.connect('pokepaste.db')

pokemon_data = json.load(open('data/pokemon.json', encoding='utf-8'))
item_data = json.load(open('data/items.json', encoding='utf-8'))
move_data = json.load(open('data/moves.json', encoding='utf-8'))

html_template = {}
for html_file in ('paste', 'paste-mon'):
    fd = open('html/{}.html'.format(html_file))
    html_template[html_file] = Template(fd.read())
    fd.close()

html_static = {}
for html_file in ('404', 'create', 'howto'):
    fd = open('html/{}.html'.format(html_file))
    html_static[html_file] = fd.read()
    fd.close()

imgcss = ''
for pokemon in pokemon_data.values(): imgcss += '''
\t\t.pokemon-{id}-{no} {{
\t\t\tbackground-image: url('../img/pokemon/{id}-{no}.png');
\t\t}}
'''.format(id=pokemon['id'], no=pokemon['no'])
for item in item_data.values(): imgcss += '''
\t\t.item-{id} {{
\t\t\tbackground-image: url('../img/items/{id}.png');
\t\t}}
'''.format(id=item['id'])
html_template['paste'] = Template(
        html_template['paste'].safe_substitute(imgcss=imgcss))

def encrypt_id_v1(id):
    return (id * crypto_secrets.key1) % crypto_secrets.mod

def decrypt_id_v1(id):
    return (id * crypto_secrets.key2) % crypto_secrets.mod

def encrypt_id_v2(id):
    return cipher.encrypt(id.to_bytes(8, 'big')).hex()

def decrypt_id_v2(id):
    return int.from_bytes(cipher.decrypt(bytes.fromhex(id)), 'big')

def format_paste(paste, title, author, notes):

    paste = html.escape(paste)

    html_mons = ''

    for mon in paste.split('\r\n\r\n'):

        if mon.strip() == '': continue

        mon_lines = mon.splitlines()
        line = mon_lines.pop(0)

        pokemonid = '0-0'
        itemid = 0

        try:
            index = len(line)
    
            if '(M)' in line:
                index = line.rindex('(M)')
                lindex = index + 1
                rindex = index + 2
                if '(' not in line[rindex + 1:]:
                    line = (line[:lindex]
                          + '<span class="gender-m">'
                          + line[lindex:rindex]
                          + '</span>'
                          + line[rindex:])
    
            if '(F)' in line:
                index = line.rindex('(F)')
                lindex = index + 1
                rindex = index + 2
                if '(' not in line[rindex + 1:]:
                    line = (line[:lindex]
                          + '<span class="gender-f">'
                          + line[lindex:rindex]
                          + '</span>'
                          + line[rindex:])
    
            if ')' in line[:index] and '(' in line[line[:index].rindex('('):]:
                rindex = line[:index].rindex(')')
                lindex = line[:rindex].rindex('(') + 1
                if line[lindex:rindex] in pokemon_data:
                    pokemon = pokemon_data[line[lindex:rindex]]
                    pokemonid = '{}-{}'.format(pokemon['id'], pokemon['no'])
                    line = (line[:lindex]
                          + '<span class="type-{}">'.format(pokemon['type'])
                          + line[lindex:rindex]
                          + '</span>'
                          + line[rindex:])
            else:
                lineparts = line[:index].partition('@')
                if lineparts[0].strip() in pokemon_data:
                    pokemon = pokemon_data[lineparts[0].strip()]
                    pokemonid = '{}-{}'.format(pokemon['id'], pokemon['no'])
                    line = ('<span class="type-{}">'.format(pokemon['type'])
                          + lineparts[0]
                          + '</span>'
                          + lineparts[1]
                          + lineparts[2]
                          + line[index:])
    
            lineparts = line.rpartition('@')
            if lineparts[2].strip() in item_data:
                item = item_data[lineparts[2].strip()]
                itemid = item['id']
                try:
                    line = (lineparts[0]
                          + lineparts[1]
                          + '<span class="type-{}">'.format(item['type'])
                          + lineparts[2]
                          + '</span>')
                except KeyError:
                    pass
        
        finally:
            mon_formatted = line + '\n'

        for line in mon_lines:
            try:
                if line[0] == '-' and line[1:].strip() in move_data:
                    move_type = move_data[line[1:].strip()]['type']
                    mon_formatted += '<span class="type-{}">'.format(move_type)
                    mon_formatted += line[0]
                    mon_formatted += '</span>'
                    mon_formatted += line[1:]
                elif ':' in line:
                    index = line.index(':') + 1
                    mon_formatted += '<span class="heading">'
                    mon_formatted += line[:index]
                    mon_formatted += '</span>'
                    mon_formatted += line[index:]
                else:
                    mon_formatted += line
            except:
                mon_formatted += line
            finally:
                mon_formatted += '\n'

        mon_formatted += '\n'

        html_mons += html_template['paste-mon'].substitute(pokemonid=pokemonid,
                                                           itemid=itemid,
                                                           paste=mon_formatted)

    if title:
        title = html.escape(title)
    else:
        title = 'Untitled'
        
    if author:
        author = html.escape(author)
    else:
        author = 'Anonymous'
        
    if notes:
        notes = html.escape(notes).replace('\r\n', '</p><p>')
    else:
        notes = ''

    return html_template['paste'].substitute(mons=html_mons,
                                             title=title,
                                             author=author,
                                             notes=notes)

def retrieve_paste(id, start_response):
    c = conn.cursor()
    c.execute('SELECT paste,title,author,notes FROM pastes WHERE id=?;', (id,))
    paste = c.fetchone()
    if paste:
        response = format_paste(*paste).encode('utf-8')
        status = '200 OK'
        headers = [
            ('Content-Type', 'text/html; charset=utf-8'),
            ('Content-Length', str(len(response)))
        ]
        start_response(status, headers)
        return [response]
    else:
        return generic_404(start_response)

def parse_pokemon(mon):
    # Example input:
    # ffffffffff (Gyarados) (M) @ Leftovers
    # Ability: Intimidate
    # Level: 57
    # Shiny: Yes
    # Happiness: 64
    # EVs: 252 Atk / 4 SpD / 252 Spe
    # Adamant Nature
    # IVs: 21 HP / 7 Atk / 19 Def / 30 SpA / 0 SpD
    # - Dragon Dance
    # - Substitute
    # - Waterfall
    # - Crunch

    # Example output (curSet dictionary)
    # {'item': {'name': 'Leftovers', 'id': 234}, 'gender': 'M', 'species': 'Gyarados', 'nick': 'ffffffffff', 'ability': 'Intimidate', 'level': 57, 'shiny': True, 'happiness': 64, 'evs': {'hp': 0, 'atk': 252, 'def': 0, 'spa': 0, 'spd': 4, 'spe': 252}, 'ivs': {'hp': 21, 'atk': 7, 'def': 19, 'spa': 30, 'spd': 0, 'spe': 31}, 'moves': [{'name': 'Dragon Dance', 'moveid': 349, 'type': 'dragon', 'classification': 0}, {'name': 'Substitute', 'moveid': 164, 'type': 'normal', 'classification': 0}, {'name': 'Waterfall', 'moveid': 127, 'type': 'water', 'classification': 1}, {'name': 'Crunch', 'moveid': 242, 'type': 'dark', 'classification': 1}]}

    text = mon.split("\n")
    curSet = None

    for x in range(0, len(text)):
        line = text[x].strip()
        if (line == '' or line == '---'):
            curSet = None

        #Initialization - first line of import
        elif (curSet is None):
            curSet = {}

            #Parse item
            atIndex = line.rfind(' @ ')
            if (atIndex != -1):
                item = line[atIndex+3:].strip()
                if (item in item_data):
                    curSet['item'] = {'name': item, 'id': item_data[item]['id']}
                else:
                    curSet['item'] = None
                line = line[:atIndex]

            #Parse gender
            if (line[-4:] == ' (M)'):
                curSet['gender'] = 'M'
                line = line[0:-4]
            if (line[-4:] == ' (F)'):
                curSet['gender'] = 'F'
                line = line[0:-4]

            #Parse species name and nick (if any)
            parenIndex = line.rfind(' (')
            if (line[-1:] == ')' and parenIndex != -1):
                line = line[0:-1]
                curSet['species'] = line[parenIndex+2:]
                line = line[:parenIndex]
                curSet['nick'] = line
            else:
                curSet['species'] = line
                curSet['nick'] = ""

        #Parse other aspects of Pokemon, each with its own line:
        elif (line[:9] == 'Ability: '):
            line = line[9:]
            curSet['ability'] = line
        elif (line == 'Shiny: Yes'):
            curSet['shiny'] = True
        elif (line[:7] == 'Level: '):
            line = line[7:]
            curSet['level'] = int(line)
        elif (line[:11] == 'Happiness: '):
            line = line[11:]
            curSet['happiness'] = int(line)
        elif (line[:5] == 'EVs: '):
            line = line[5:]
            evLines = line.split('/')
            curSet['evs'] = {'hp': 0, 'atk': 0, 'def': 0, 'spa': 0, 'spd': 0, 'spe': 0}
            for y in range(len(evLines)):
                evLine = evLines[y].strip()
                spaceIndex = evLine.find(' ')
                if (spaceIndex == -1):
                    continue
                statId = evLine[spaceIndex+1:].lower()
                statVal = int(evLine[:spaceIndex])
                if (statId in curSet['evs']):
                    curSet['evs'][statId] = statVal
        elif (line[:5] == 'IVs: '):
            line = line[5:]
            ivLines = line.split('/')
            curSet['ivs'] = {'hp': 31, 'atk': 31, 'def': 31, 'spa': 31, 'spd': 31, 'spe': 31}
            for y in range(len(ivLines)):
                ivLine = ivLines[y].strip()
                spaceIndex = ivLine.find(' ')
                if (spaceIndex == -1):
                    continue
                statId = ivLine[spaceIndex + 1:].lower()
                statVal = int(ivLine[:spaceIndex])
                if (statId in curSet['ivs']):
                    curSet['ivs'][statId] = statVal
        elif (line[:-6].lower() == 'nature'):
            test = 0
            #nature.json not created yet...do this later
        #Parse move
        elif (line[:1] == '-' or line[:1] == '~'):
            line = line[1:].strip()
            if ('moves' not in curSet):
                curSet['moves'] = []
            #Check for Hidden Power case. Also (possibly) adjust for IVs
            if (line in move_data):
                curSet['moves'].append({'name': line, 'moveid': int(move_data[line]['id']), 'type': move_data[line]['type'], 'classification': int(move_data[line]['classification'])})

    # print(curSet)
    return curSet

def generic_404(start_response, status='404 Not Found'):
    response = html_static['404'].encode('utf-8')
    headers = [
        ('Content-Type', 'text/html; charset=utf-8'),
        ('Content-Length', str(len(response)))
    ]
    start_response(status, headers)
    return [response]

def application(environ, start_response):
    method = environ['REQUEST_METHOD']
    path = environ['PATH_INFO'].strip('/')

    if method == 'GET':

        if len(path) == 16:
            # Requesting a paste - new encrypted ID
            try:
                id = decrypt_id_v2(path)
            except:
                return generic_404(start_response)
            else:
                return retrieve_paste(id, start_response)

        elif path.isdigit():
            # Requesting a paste - unencrypted or old encrypted ID?
            id = int(path)
            if id >= 256:
                # Old encrypted ID
                id = decrypt_id_v1(id)
                if id >= 1000:
                    # Prevent using old hash format on new pastes
                    return generic_404(start_response)
            return retrieve_paste(id, start_response)


        elif not path:
            # Requesting /, return the (static) submit page
            response = html_static['create'].encode('utf-8')
            status = '200 OK'
            headers = [
                ('Content-Type', 'text/html; charset=utf-8'),
                ('Content-Length', str(len(html_static['create'])))
            ]
            start_response(status, headers)
            return [response]

        elif path == 'howto':
            # Requesting the (static) PokePaste HOWTO
            response = html_static['howto'].encode('utf-8')
            status = '200 OK'
            headers = [
                ('Content-Type', 'text/html; charset=utf-8'),
                ('Content-Length', str(len(html_static['howto'])))
            ]
            start_response(status, headers)
            return [response]


        elif img_re.fullmatch(path):
            # Requesting an image
            try:
                rfile = open(path, mode='rb')
                fsize = os.stat(path).st_size
                bsize = os.statvfs(path).f_bsize
            except:
                return generic_404(start_response)
            else:
                status = '200 OK'
                headers = [
                    ('Content-Type', 'image/png'),
                    ('Content-Length', str(fsize))
                ]
                start_response(status, headers)
                if 'wsgi.file_wrapper' in environ:
                    return environ['wsgi.file_wrapper'](rfile, bsize)
                else:
                    return iter(lambda: rfile.read(bsize), b'')

        else:
            return generic_404(start_response)

    elif method == 'POST':

        if path == 'create':
            # Submit a new paste
            form = cgi.FieldStorage(fp=environ['wsgi.input'], environ=environ)
            if form.getvalue('paste'):
                metadata = {}
                c = conn.cursor()
                c.execute('INSERT INTO pastes(paste,title,author,notes)'
                          'VALUES(?,?,?,?)',
                          [form.getvalue(key) for key in ('paste', 'title',
                                                          'author', 'notes')])
                conn.commit()
                status = '302 Found'
                headers = [
                    ('Location', '/{}'.format(encrypt_id_v2(c.lastrowid)))
                ]
                start_response(status, headers)
                return [b'']
            else:
                return generic_404(start_response, '400 Bad Request')
            response = form['paste'].value.encode('utf-8')
            status = '200 OK'
            headers = [
                ('Content-Type', 'text/plain; charset=utf-8'),
                ('Content-Length', str(len(form['paste'].value)))
            ]
            start_response(status, headers)
            return [response]
        else:
            return generic_404(start_response, '405 Method Not Allowed')

if __name__ == '__main__':
    with make_server('', 8000, application) as httpd:
        try:
            httpd.serve_forever()
        except KeyboardInterrupt:
            quit()
