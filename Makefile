TARGET = testdom
INCLUDE = ./
GFLAGS = -I $(INCLUDE)
SOURCES = *.go
ifeq ($(GOARCH),386)
	CC = 8g
	LD = 8l
	OBJ_SUFFIX = 8
else
	CC = 6g
	LD = 6l
	OBJ_SUFFIX = 6
endif
OBJECTS = $(shell echo $(SOURCES) | sed -e 's,\.go,\.$(OBJ_SUFFIX),g')

all: $(TARGET)

$(TARGET): $(OBJECTS)
	$(LD) -o $(TARGET) $(OBJECTS)

%.$(OBJ_SUFFIX) : %.go
	$(CC) $(GFLAGS) $^

clean:
	rm -f $(TARGET) $(OBJECTS)
