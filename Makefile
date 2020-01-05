build:
	cd jsonlinter &&\
	go build -o linter main.go

build-container: check-image
	docker build -t $(image) .

run-app: build-container check-port
	docker run -d -p $(port):8400 $(image)

clean: check-image
	docker rm -f $$(docker ps -a -f ancestor=$(image):latest -q) &&\
	docker rmi $(image)

check-image:
ifndef image
	@echo $(error Image name is not defined)
endif

check-port:
ifndef port
	@echo $(error External Port is not defined)
endif