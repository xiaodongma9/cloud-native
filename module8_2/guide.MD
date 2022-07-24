### Install ingress controller

```sh
kubectl create -f deployment.yaml
```

### Generate key-cert

```sh
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout cloudnative.key -out cloudnative.crt -subj "/CN=cloudnative.com/O=cloudnative" -addext "subjectAltName = DNS:cloudnative.com"
```

### Create secret

```sh
kubectl create secret tls cloudnative-tls --cert=./tls.crt --key=./tls.key
```

### Create a ingress

```sh
kubectl create -f ingress.yaml
```

### Test the result

```sh
curl -H "Host: cloudnative.com" https://10.109.204.181 -v -k
```