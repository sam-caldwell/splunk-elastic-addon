certs: cert/ca cert/elastic cert/splunk
	@echo "done: ok"

cert/ca:
	@mkdir -p data/certs || true
	@openssl genrsa -out data/certs/ca/ca.key 4096
	@openssl req -x509 \
				-new \
				-nodes \
				-key data/certs/ca/ca.key \
				-sha256 \
				-days 365 \
				-out data/certs/ca/ca.crt \
				-subj "/CN=Local CA"

cert/elastic:
	@openssl genrsa -out data/certs/es/elasticsearch.key 2048
	@openssl req -new \
				-key data/certs/es/elasticsearch.key \
				-out data/certs/es/elasticsearch.csr \
				-subj "/CN=elasticsearch"
	@openssl x509 -req \
				 -in data/certs/es/elasticsearch.csr \
				 -CA data/certs/ca/ca.crt \
				 -CAkey data/certs/ca/ca.key \
				 -CAcreateserial \
				 -out data/certs/es/elasticsearch.crt \
				 -days 3650 \
				 -sha256

cert/splunk:
	@openssl genrsa -out data/certs/splunk/splunk.key 2048
	@openssl req -new \
				-key data/certs/splunk/splunk.key \
				-out data/certs/splunk/splunk.csr \
				-subj "/CN=splunk"
	@openssl x509 -req \
				 -in data/certs/splunk/splunk.csr \
				 -CA data/certs/ca/ca.crt \
				 -CAkey data/certs/ca/ca.key \
				 -CAcreateserial \
				 -out data/certs/splunk/splunk.crt \
				 -days 3650 \
				 -sha256


