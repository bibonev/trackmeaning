FROM iron/go:dev
WORKDIR /app

ENV SRC_DIR=/go/src/gitlab.com/boyan.bonev/intern-hackday/api/go
ENV GOOGLE_APPLICATION_CREDENTIALS=/go/src/gitlab.com/boyan.bonev/intern-hackday/api/go/google-credentials.json

ADD . $SRC_DIR

# Build it: 
RUN go get -v github.com/gin-gonic/gin; 
RUN go get -v cloud.google.com/go/translate; 
RUN go get -v cloud.google.com/go/language/apiv1; 
RUN go get -v google.golang.org/genproto/googleapis/cloud/language/v1; 
RUN cd $SRC_DIR; go build main.go; cp main /app/

ENTRYPOINT ["./main"]