# 1. Generate CA's primary key and self-signed certificate

openssl req -x509 -newkey rsa:4096 -days 365 -keyout ca-key.pem -out ca-cert.pem -subj "/C=RU/ST=Moscow/L=Moscow/O=Datacademy/OU=Education/CN=datacademy.net/emailAddress=sergej.isaeff@gmail.com"

openssl x509 -in ca-cert.pem -noout -text

# 2. Generate web server's private key and sertificate signing request

openssl req -newkey rsa:4096 -keyout server-key.pem -out server-req.pem -subj "/C=RU/ST=RF/L=Moscow/O=Datacademy/OU=Education/CN=datacademy.net/emailAddress=sergej.isaeff@gmail.com"

# 3. Use CA's primary key to sign sertificate signing reqzest and get back the signed certificate

openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem

openssl x509 -in server-cert.pem -noout -text
