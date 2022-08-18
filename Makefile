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

uninstall:
	$(RM) $(DESTDIR)$(PREFIX)/bin/thesaurus

.DEFAULT_GOAL := all

.PHONY: all thesaurus clean install uninstall
