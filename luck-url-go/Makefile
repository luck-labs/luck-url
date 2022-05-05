all: install run

install: clean gen_deploy build
	mkdir -p deploy/bin
	mv ./$(PROJECT_NAME) deploy/bin/

gen_deploy:
	mkdir -p deploy/log/$(PROJECT_NAME)
	mkdir -p deploy/conf
	go mod vendor
	cp -rf conf/* deploy/conf/
	cp -rf conf/$(LUCK_URL_ENV)/* deploy/conf/$(LUCK_URL_ENV)/

clean:
	rm -rf deploy/log/$(PROJECT_NAME)
	rm -rf deploy/conf
	rm -rf deploy/bin
	rm -rf vendor

build:
	go build -o $(PROJECT_NAME)

run:
	./deploy/bin/luck-url-go
	#nohup ./deploy/bin/luck-url-go & >> ./deploy/log/setup.log 2>&1 &
