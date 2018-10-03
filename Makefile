app_version=$(shell git rev-parse HEAD 2> /dev/null | sed "s/\(.*\)/\1/")
env=development
debug=false
profile=mwas
test=true

compile:
	docker-compose -f tests/app/docker-compose.yml run app go build -o ./dist/main

test:  compile
	docker-compose run ansible ansible-playbook tests/deploy.yml \
	-i tests/inventory -e "app_version=$(app_version)" \
	-e env=$(env) -e debug=$(debug) -e test=$(test) \
	-e profile=$(profile) ;