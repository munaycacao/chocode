dbml2sql model.dbml --mysql -o model.sql && rm dbml-error.log
cd ..
sqlc generate && echo "Updated SQLC generated files"