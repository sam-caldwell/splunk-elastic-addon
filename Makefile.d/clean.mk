clean:
	@rm -rf data/*
	@mkdir -p data/certs/ca
	@mkdir -p data/certs/es
	@mkdir -p data/certs/splunk
	@mkdir -p data/elastic
	@mkdir -p splunk/{apps,data}/
	@echo 'clean:ok'