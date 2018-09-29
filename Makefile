app_version=$(shell git rev-parse HEAD 2> /dev/null | sed "s/\(.*\)/\1/")
env=development
debug=false
profile=default

compile:
	docker-compose -f tests/app/docker-compose.yml run app go build -o ./dist/main

test: create_build_folder compile
	docker-compose run ansible ansible-playbook tests/test.yml \
	-i tests/inventory -e "app_version=$(app_version)" \
	-e env=$(env) -e debug=$(debug) \
	-e profile=$(profile) ;