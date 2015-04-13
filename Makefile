
compose:
	@docker-compose up

client:
	@PGPASSWORD=sloth psql -h 192.168.59.103 -U tj -p 5432

.PHONY: client compose