db-up:
	/Users/ryan/go/bin/sql-migrate up -config=dbconfig.yml -env=development

db-down:
	/Users/ryan/go/bin/sql-migrate down -config=dbconfig.yml -env=development

db-status:
	/Users/ryan/go/bin/sql-migrate status -config=dbconfig.yml -env="development"