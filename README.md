# safe

rename .env_example to .env and change the data.

rename config_example.json to config.json and change the data.

Generate private key (.key)
Key considerations for algorithm "RSA" ≥ 2048-bit

```bash
openssl genrsa -out server.key 2048
```

Key considerations for algorithm "ECDSA" ≥ secp384r1
List ECDSA the supported curves (openssl ecparam -list_curves)

```bash
openssl ecparam -genkey -name secp384r1 -out server.key
```

Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)

```bash
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

for more security watch https://github.com/denji/golang-tls