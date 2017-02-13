#!/usr/bin/env python3

from wsgiref.simple_server import make_server

import os
import cgi

# We pass image requests to open() basically unmodified,
# so this regex is needed to filter out any funny business.
import re
img_re = re.compile(r'img/(pokemon/\d+-\d+|items/\d+).png')

import sqlite3
conn = sqlite3.connect('pokepaste.db')

import json
pokemon_data = json.load(open('data/pokemon.json'))
item_data = json.load(open('data/items.json'))
move_data = json.load(open('data/moves.json'))

html_data = {}
for html_file in ('404', 'create', 'paste', 'paste-mon'):
    fd = open('html/{}.html'.format(html_file))
    html_data[html_file] = fd.read()
    fd.close()

css_data = {}
for css_file in ('paste',):
    fd = open('css/{}.css'.format(css_file))
    css_data[css_file] = fd.read()
    fd.close()

def format_paste(pasteid, paste):

    html_mons = ''

    for mon in paste.split('\r\n\r\n'):

        if mon.strip() == '': continue

        mon_lines = mon.splitlines()
        line = mon_lines.pop(0)

        pokemonid = '0-0'
        itemid = 0

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
            index = line.rindex('(M)')
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
            lineparts = line.partition('@')
            if lineparts[0].strip() in pokemon_data:
                pokemon = pokemon_data[lineparts[0].strip()]
                pokemonid = '{}-{}'.format(pokemon['id'], pokemon['no'])
                line = ('<span class="type-{}">'.format(pokemon['type'])
                      + lineparts[0]
                      + '</span>'
                      + lineparts[1]
                      + lineparts[2])

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

        mon_formatted = line

        for line in mon_lines:
            mon_formatted += '\n'
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

        html_mons += html_data['paste-mon'].format(pokemonid=pokemonid,
                                                   itemid=itemid,
                                                   paste=mon_formatted)

    return html_data['paste'].format(pasteid=pasteid, mons=html_mons)

def generic_404(start_response):
    response = html_data['404'].encode('utf-8')
    status = '404 Not Found'
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

        if not path:
            # Requesting /, return the (static) submit page
            response = html_data['create'].encode('utf-8')
            status = '200 OK'
            headers = [
                ('Content-Type', 'text/html; charset=utf-8'),
                ('Content-Length', str(len(html_data['create'])))
            ]
            start_response(status, headers)
            return [response]

        elif path.isdigit():
            # Requesting a specific paste, format and return it
            c = conn.cursor()
            c.execute('SELECT * FROM pastes WHERE id=?;', (int(path),))
            paste = c.fetchone()
            if paste:
                response = format_paste(paste[0], paste[1]).encode('utf-8')
                status = '200 OK'
                headers = [
                    ('Content-Type', 'text/html; charset=utf-8'),
                    ('Content-Length', str(len(response)))
                ]
                start_response(status, headers)
                return [response]
            else:
                return generic_404(start_response)

        elif path in ['css/{}.css'.format(css_file) for css_file in css_data]:
            # Requesting a static CSS file
            response = css_data[path[4:-4]].encode('utf-8')
            status = '200 OK'
            headers = [
                ('Content-Type', 'text/css; charset=utf-8'),
                ('Content-Length', str(len(css_data[path[4:-4]])))
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
            if 'paste' in form and form['paste'].value:
                c = conn.cursor()
                c.execute('INSERT INTO pastes(paste) VALUES(?)',
                          (form['paste'].value,))
                conn.commit()
                status = '302 Found'
                headers = [
                    ('Location', '/{}'.format(c.lastrowid))
                ]
                start_response(status, headers)
                return [b'']
            else:
                # TODO Error
                pass
            response = form['paste'].value.encode('utf-8')
            status = '200 OK'
            headers = [
                ('Content-Type', 'text/plain; charset=utf-8'),
                ('Content-Length', str(len(form['paste'].value)))
            ]
            start_response(status, headers)
            return [response]
        else:
            # TODO 404 or 405
            pass

if __name__ == '__main__':
    with make_server('', 8000, application) as httpd:
        httpd.serve_forever()
