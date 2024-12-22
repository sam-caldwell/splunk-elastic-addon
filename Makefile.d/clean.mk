clean:
	@rm -rf data/*
	@rm -rf build
	@mkdir -p data/certs/ca
	@mkdir -p data/certs/es
	@mkdir -p data/certs/splunk
	@mkdir -p data/elastic
	@mkdir -p data/splunk/apps
	@mkdir -p data/splunk/data
	@mkdir -p build/
	@echo 'clean:ok'
