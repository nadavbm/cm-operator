# cm-operator

before start, make sure that the following version of docker, go and kubectl are installed correctly on your machine:

```
$ Docker version 20.10.12, build e91ed57

$ Client Version: version.Info{Major:"1", Minor:"23", GitVersion:"v1.23.5", GitCommit:"c285e781331a3785a7f436042c65c5641ce8a9e9", GitTreeState:"clean", BuildDate:"2022-03-16T15:58:47Z", GoVersion:"go1.17.8", Compiler:"gc", Platform:"linux/amd64"}
Unable to connect to the server: dial tcp 192.168.39.252:8443: connect: no route to host

$ go version go1.17.8 linux/amd64
```

set your go environment:
```
export GOROOT=/usr/local/go
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOBIN
export GOPATH=$HOME/go
export GO111MODULE=on
```

to create this operator:
```
operator-sdk init --domain example.com --repo github.com/nadavbm/cm-operator
```

create initial api resource
```
operator-sdk create api --group opconfigmap --version v1alpha1 --kind OpConfigMap --resource --controller
```

run:
```
make generate
```
to generate crd bases:
```
make manifests
```
it will create the crd under bases

upload container to docker hub:
```

```