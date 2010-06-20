NAME := europa
GOC := 6g
GOL := 6l
RM := rm -f
SRCS := object.go msg.go call.go block.go number.go

.SUFFIXES: .go .6

all: $(NAME)

$(NAME): build_pkg
	$(GOC) main.go
	$(GOL) -o $(NAME) main.6

build_pkg: $(SRCS)
	$(GOC) -o $(NAME).6 $(SRCS)

clean:
	$(RM) main.6 $(NAME).6 $(NAME)
