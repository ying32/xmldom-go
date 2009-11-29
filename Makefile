include $(GOROOT)/src/Make.$(GOARCH)

TARG=xml/dom
GOFILES=\
	dom.go

include $(GOROOT)/src/Make.pkg
