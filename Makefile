include $(GOROOT)/src/Make.inc

TARG=xml/dom
GOFILES=\
	core_interfaces.go \
	node.go \
	attr.go \
	document.go \
	element.go \
	text.go \
	nodelists.go \
	namednodemap.go \
	dom.go

include $(GOROOT)/src/Make.pkg
