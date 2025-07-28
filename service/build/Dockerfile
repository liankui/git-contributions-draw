FROM alpine:3.18.4

# We'll likely need to add SSL root certificates
# RUN apk --no-cache add ca-certificates

ARG GIT_HASH
ARG GIT_BRANCH
LABEL gitHash=$GIT_HASH gitBranch=$GIT_BRANCH

WORKDIR /

ADD cmd/git-contributions-draw-server .
CMD ["/git-contributions-draw-server"]
