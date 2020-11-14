#
# Makefile
#
VERSION = snapshot
GHRFLAGS =
.PHONY: build release

default: build

build:
	goxc -d=pkg -pv=$(VERSION) -bc="linux,windows,darwin,!386"

release:
	ghr  -u ashsmith  $(GHRFLAGS) v$(VERSION) pkg/$(VERSION)
