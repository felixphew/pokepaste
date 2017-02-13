PokePaste
=========

What is this?
-------------

An answer to "why can't Pastebin syntax-highlight Pokemon sets?".

Pokepaste is a simple pastebin, with a clean interface, that supports
highlighting of the syntax created by [Pokemon Showdown][]. This syntax has
become a standard, and is supported by multiple tools, yet no pastebin (as
far as I'm aware) has support for syntax-highlighting it.

  [Pokemon Showdown]: https://github.com/zarel/Pokemon-Showdown

Features
--------

### Working ###

*   Highlights mon names, moves, and (where appropriate) items with the
    appropriate type
*   Image preview for mons and items
*   Clean and mobile-friendly layout

### To come ###

*   Tooltips for pokemon, items and moves
*   Pokemon Showdown integration (hopefully bi-directional)
*   QR code import (maybe)
*   QR code *export* (unlikely, but would be cool)
*   World domination

Running
-------

PokePaste only needs Python 3+, compiled with SQlite support (if you don't
know, yours probably is). Create `pokepaste.db` from `schema.sql`, run
`pokepaste.py`, and you should be up and running.

If you plan on putting this on anything resembling the public Internet,
running behind a reverse proxy such as [nginx][] is recommended.

  [nginx]: https://nginx.org
