DOCKER=$(shell which docker)

mysql_create:
	${DOCKER} run --name vuttr --network host -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE=main -d mysql

mysql_delete:
	${DOCKER} rm -f vuttr