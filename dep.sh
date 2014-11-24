#ugly staff this should replace with a dep tool like godep
echo "resolving depdencies"

go get github.com/go-sql-driver/mysql
go get github.com/go-xorm/xorm

echo "resolve depdencies complete"