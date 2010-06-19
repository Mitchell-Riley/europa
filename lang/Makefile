NAME := europa
GOC := 6g
GOL := 6l
RM := rm -f
SRCS := $(wildcard *.go)

.SUFFIXES: .go .6

all: $(NAME)

$(NAME): build
	$(GOL) -o $(NAME) $(NAME).6

build: $(SRCS)
	$(GOC) -o $(NAME).6 $(SRCS)

clean:
	$(RM) $(NAME).6 $(NAME)
