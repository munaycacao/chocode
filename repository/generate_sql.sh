dbml2sql model.dbml --mysql -o model.sql && rm dbml-error.log
cd ..
# TODO: make sure SQLC generation works after DBML generation by implementing a db migration tool
sqlc generate && echo "Updated SQLC generated files"
