include $(GOROOT)/src/Make.$(GOARCH)

TARG=xml/dom
GOFILES=\
	core_interfaces.go \
	dom.go

include $(GOROOT)/src/Make.pkg
