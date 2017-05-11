# safe

create .env file

MYSQL_ROOT_PASSWORD=databaserootpassword
MYSQL_DATABASE=databasename
MYSQL_USER=databaseuser
MYSQL_PASSWORD=databaseuserpassword


create config.json

{
  "host": "172.28.0.23:3306",
  "user": "root",
  "password": databaserootpassword,
  "database": databasename
}


Generate private key (.key)
# Key considerations for algorithm "RSA" ≥ 2048-bit
openssl genrsa -out server.key 2048
# Key considerations for algorithm "ECDSA" ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl ecparam -genkey -name secp384r1 -out server.key

Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

for more security watch https://github.com/denji/golang-tls