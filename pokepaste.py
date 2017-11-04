from flask import Flask, Markup, escape, render_template, request, flash, redirect, abort

from cryptography.hazmat.primitives.ciphers import Cipher
from cryptography.hazmat.primitives.ciphers.algorithms import Blowfish
from cryptography.hazmat.primitives.ciphers.modes import ECB
from cryptography.hazmat.backends import default_backend

from mysql.connector import connect as MySQL

import re
import json

app = Flask(__name__)
app.config.from_pyfile('pokepaste.cfg')

cipher = Cipher(
        Blowfish(app.config['CRYPTO_KEY']),
        ECB(),
        default_backend()
)

conn = MySQL(
    user='pokepaste',
    database='pokepaste',
    charset='utf8mb4',
    collation='utf8mb4_unicode_ci',
    autocommit=True
)

res = {
    'name1': r'^([A-Z][a-z0-9:\']+(?:[- ][A-Z][a-z0-9:\']+)*)(?=$| \(| @)',
    'name2': r'(?<=\()([A-Z][a-z0-9:\']+(?:[- ][A-Z][a-z0-9:\']*)*)(?=\))',
    'gender': r'(?<=\()([MF])(?=\))',
    'item': r'(?<=@ )([A-Z][a-z0-9:\']+(?:[- ][A-Z][a-z0-9:\']*)*)',
    'move': r'^(- )(?=([A-Z][a-z\']+(?:[- ][A-Z][a-z\']*)*))',
    'hp': r'(?<=\[)([A-Z][a-z]+)(?=\])',
    'attr': r'^([A-Za-z]+:)',
    'stats': r'([0-9]+ (HP|Atk|Def|SpA|SpD|Spe))'
}
res = {n: re.compile(r) for n, r in res.items()}

data = ['pokemon', 'items', 'moves']
data = {n: json.load(open('data/{}.json'.format(n))) for n in data}

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/create', methods=['POST'])
def create():
    # Check that we have a paste (everything else is optional)
    if 'paste' not in request.form:
        flash('Something went wrong there - did you try to submit an empty paste?')
        return redirect('/')

    # Create the paste
    c = conn.cursor()
    c.execute('INSERT INTO pastes (paste, title, author, notes) VALUES (%s, %s, %s, %s)',
              [request.form.get(k) for k in ('paste', 'title', 'author', 'notes')])

    # Encrypt ID and return
    encryptor = cipher.encryptor()
    url = encryptor.update(c.lastrowid.to_bytes(8, 'big')).hex()
    return redirect('/' + url)


@app.route('/<cryptid>')
def paste(cryptid):
    if len(cryptid) == 16:
        # Looks like an encrypted ID, check validity
        try:
            cryptid = bytes.fromhex(cryptid)
        except ValueError:
            abort(404)

        # Read was successful, decrypt
        decryptor = cipher.decryptor()
        pasteid = int.from_bytes(decryptor.update(cryptid), 'big')
        decryptor.finalize()

    elif cryptid.isdigit():
        # Looks like an old ID of some kind
        pasteid = int(cryptid)

        # Greater than 256 is pseudo-encrypted.
        # This scheme was so stupidly designed that I'm just hardcoding
        # the keys, we're not using it anymore anyway.
        if pasteid >= 256:
            pasteid = (pasteid * (2 ** 31 - 1)) % (2 ** 32)

        # Prevent grabbing new pastes with old URL schemes
        if pasteid >= 1000:
            abort(403)

    else:
        abort(404)

    c = conn.cursor()
    c.execute('SELECT paste, title, author, notes FROM pastes WHERE id = %s',
              (pasteid,))
    paste = c.fetchone()

    if not paste:
        abort(404)

    return render_paste(*paste)

def render_paste(paste, title, author, notes):
    mons = []

    ismon = False
    for line in paste.splitlines():

        line = line.rstrip()

        if not line:
            # Blank line means end of a set
            if ismon:
                mon['text'] += '\n'
                mons.append(mon)
                ismon = False

        elif not ismon:
            # First line in a set
            mon = { 'monid': 0, 'monno': 0, 'itemid': 0, 'text': Markup() }
            # Try with nickname first
            match = res['name2'].split(line, maxsplit=1)
            if len(match) != 3:
                # Now try without nickname
                match = res['name1'].split(line, maxsplit=1)
            # If we matched, lookup and highlight the first part of the string
            # If not, the rest of the tests use the whole string
            if len(match) == 3 and match[1] in data['pokemon']:
                moninfo = data['pokemon'][match[1]]
                mon['monid'] = moninfo['id']
                mon['monno'] = moninfo['no']
                mon['text'] += (escape(match[0])
                                + Markup('<span class="type-{}">'.format(moninfo['type']))
                                + escape(match[1])
                                + Markup('</span>'))
                line = match[2]
            # Try gender
            match = res['gender'].split(line, maxsplit=1)
            if len(match) == 3:
                mon['text'] += (escape(match[0])
                                + Markup('<span class="gender-{}">'.format(match[1].lower()))
                                + escape(match[1])
                                + Markup('</span>'))
                line = match[2]
            # Try item
            match = res['item'].split(line, maxsplit=1)
            if len(match) == 3 and match[1] in data['items']:
                iteminfo = data['items'][match[1]]
                mon['itemid'] = iteminfo['id']
                if 'type' in iteminfo:
                    mon['text'] += (escape(match[0])
                                    + Markup('<span class="type-{}">'.format(iteminfo['type']))
                                    + escape(match[1])
                                    + Markup('</span>'))
                else:
                    mon['text'] += escape(match[0] + match[1])
                line = match[2]
            mon['text'] += escape(line)
            mon['text'] += '\n'
            ismon = True

        else:
            # Normal line in a set
            # Try move
            match = res['move'].split(line, maxsplit=1)
            if len(match) == 4 and match[2] in data['moves']:
                mon['text'] += (escape(match[0])
                                + Markup('<span class="type-{}">'.format(data['moves'][match[2]]['type']))
                                + escape(match[1])
                                + Markup('</span>'))
                                #+ escape(match[3]))
                line = match[3]
                # Successful move, try hidden power types
                match = res['hp'].split(line)
                if len(match) > 1 and len(match) % 2 == 1:
                    mon['text'] += escape(match[0])
                    for i in range(1, len(match), 2):
                        mon['text'] += (Markup('<span class="type-{}">'.format(match[i].lower()))
                                        + escape(match[i])
                                        + Markup('</span>')
                                        + escape(match[i+1]))
                else:
                    mon['text'] += line
            else:
                # Try attribute
                match = res['attr'].split(line, maxsplit=1)
                if len(match) == 3:
                    mon['text'] += (escape(match[0])
                                    + Markup('<span class="attr">')
                                    + escape(match[1])
                                    + Markup('</span>'))
                    line = match[2]
                    # Successful attribute, try EVs/IVs
                    match = res['stats'].split(line)
                    if len(match) > 1 and len(match) % 3 == 1:
                        mon['text'] += escape(match[0])
                        for i in range(1, len(match), 3):
                            mon['text'] += (Markup('<span class="stat-{}">'.format(match[i+1].lower()))
                                            + escape(match[i])
                                            + Markup('</span>')
                                            + escape(match[i+2]))
                    else:
                        mon['text'] += escape(line)
                else:
                    mon['text'] += escape(line)
            mon['text'] += '\n'


    # Add last mon (necessary if no trailing newline)
    if ismon:
        mon['text'] += '\n'
        mons.append(mon)
        ismon = False

    return render_template('paste.html', mons=mons, title=title,
                           author=author, notes=notes)

@app.errorhandler(403)
@app.errorhandler(404)
@app.errorhandler(500)
def not_found(error):
    return render_template('error.html', error=error), error.code
