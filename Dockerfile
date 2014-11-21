FROM google/golang
WORKDIR /gopath
ADD . /gopath

RUN ./dep.sh

CMD ["bash"]

#ENTRYPOINT ["/bin/bash", "-c"]
