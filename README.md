PokePaste
=======
Pokepaste is a simple pastebin, with a clean interface, that supports
highlighting of the syntax created by Pokemon Showdown. This syntax has
become a standard, and is supported by multiple tools, yet no pastebin
(up until now) has had support for syntax-highlighting it.

PokePaste highlights mon names with their primary type, moves with their
type, items (where appropriate) with their associated type and more. It
also features an image preview for mons and items using art from the
Pokemon Global Link. The site is simple, standards-compliant and mobile-
friendly.

PokePaste is grateful to BrowserStack <https://www.browserstack.com> for
help testing PokePaste v3 across different browsers.

You can use PokePaste at <https://pokepast.es>.

PokePaste v3 was rewritten from the ground up in Go. Most of the
functionality required is in the standard library, with the exception of
the MySQL and Blowfish libraries. The file `config.go' is not included;
it contains the settings for URL encryption and database connection.
Example:

	package main

	var (
		dataSourceName = "username:password@/dbname"
		key            = []byte("Super Secret Key")
	)

See <https://github.com/go-sql-driver/mysql#dsn-data-source-name> for
more information on the database connection parameters.

If you plan on putting this on anything resembling the public Internet,
please consider running it under a framework or behind a reverse proxy
of some kind (nginx works well).

Licensing
======

PokePaste is licensed under the 3-clause BSD licence. In short, you can
do whatever you want with the code as long as you retain the copyright
notice (see LICENSE) and don't hold me liable.

Pokémon and Pokémon character names are trademarks of Nintendo.
