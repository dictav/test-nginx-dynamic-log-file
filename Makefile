all:
	docker-compose up --build
	@echo "PLEASE access to http://localhost:8080/log"
	@echo "And check /mylog directory"
