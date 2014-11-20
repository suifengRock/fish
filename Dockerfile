FROM google/golang
WORKDIR /gopath
ADD . /gopath

RUN go get github.com/go-sql-driver/mysql
Run go get github.com/go-xorm/xorm
RUN go get github.com/PuerkitoBio/goquery
RUN go get github.com/go-martini/martini

CMD ["bash"]

#ENTRYPOINT ["/bin/bash", "-c"]
