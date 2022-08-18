# thesaurus
# See LICENSE for copyright and license details.
.POSIX:

include config.mk

all: thesaurus

thesaurus:
	$(GO) build $(GOFLAGS)

clean:
	$(RM) thesaurus

install: all
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f thesaurus $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/thesaurus
	mkdir -p $(DESTDIR)$(PREFIX)/share/thesaurus
	chmod 755 $(DESTDIR)$(PREFIX)/share/thesaurus
	cp -f words.txt $(DESTDIR)$(PREFIX)/share/thesaurus

uninstall:
	$(RM) $(DESTDIR)$(PREFIX)/bin/thesaurus
	$(RM) $(DESTDIR)$(PREFIX)/share/thesaurus/words.txt

.DEFAULT_GOAL := all

.PHONY: all thesaurus clean install uninstall
