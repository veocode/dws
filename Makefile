build:
	docker-compose up --build
	./dist/dws version

clear:
	rm -rf ./dist/*

link:
	sudo ln -s $(shell pwd)/dist/dws /usr/local/bin/dws

unlink:
	sudo rm -f /usr/local/bin/dws
