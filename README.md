# try_go_test

## build image

```sh
docker image build --tag try_go_test .
```

## create container

```sh
docker run -d -it -v 【src dir】:/go/src/try_go_test --name try_go_test try_go_test
```
