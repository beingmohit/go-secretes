FROM golang

ADD https://github.com/golang/dep/releases/download/v0.5.3/dep-linux-amd64 /usr/bin/dep

RUN chmod +x /usr/bin/dep

RUN mkdir -p $GOPATH/src/github.com/beingmohit/go-secretes

WORKDIR $GOPATH/src/github.com/beingmohit/go-secretes

COPY . ./

RUN dep ensure 

CMD go run main.go 2>&1 | tee /var/log/app.log && tail -f /dev/null