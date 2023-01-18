FROM alpine:3.17
COPY tekton-test /app/tekton-test
ENTRYPOINT ["/app/tekton-test"]
