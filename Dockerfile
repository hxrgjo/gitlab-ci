FROM scratch
MAINTAINER ray
ADD /artifacts/test-ci /
ENTRYPOINT ["/test-ci"]
CMD  version
