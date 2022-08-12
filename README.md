# Core API of Crane
core api of crane.



# DEV GUIDE
clone the project to your $GOPATH. following command will generate crd yamls and files in the project directory.
```
make update
make verify
```


# INSTALL CRD
```

# install crd
kubectl create -f artifacts/deploy/

```

# INSTALL CRD ONLINE

You can find other versions base on the branch or tags in dist branch.

[Click here to view the early version.](https://fastly.jsdelivr.net/gh/gocrane/api@dist/)

```bash

kubectl create -f https://raw.githubusercontent.com/gocrane/api/dist/main/all.yaml

```