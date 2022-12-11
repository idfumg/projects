gen-cert:
	openssl genrsa -out key.pem && \
	openssl req -new -key key.pem -out cert.pem && \
	openssl req -x509 -days 365 -key key.pem -in cert.pem -out certificate.pem

clean:
	rm cert.pem certificate.pem key.pem