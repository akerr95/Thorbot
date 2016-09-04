# Use an empty docker container
FROM centurylink/ca-certs

# Copy the local binary, and the configurationfile
# Before docker build, run go build
ADD thorbot /
ADD config /config

# Run command by default when the container starts.
CMD ["/thorbot"]