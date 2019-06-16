
FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./ /go/src/github.com/renanferr/bongo
WORKDIR /go/src/github.com/renanferr/bongo

# RUN go get github.com/bwmarrin/discordgo
RUN go get .
RUN go build

CMD if [ ${APP_ENV} = production ]; \
	then \
	bot; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi
	
EXPOSE 3000