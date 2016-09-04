
#### Thorbot
###### Who could be more worthy to handle your discord channel then Thor himself?

##### Prerequisites
This package uses https://github.com/Masterminds/glide for dependency management.

##### Build instruction
1. Get dependencies with `go get` or `glide install`
2. Rename config.template.json to config.json and fill your app credentials
3. Run the command below to create a statically linked Go binary
```
$ CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w'
```
4. Build your docker container
```
docker build -t Thorbot .
```
5. Run the docker container
```
docker run -d Thorbot
```