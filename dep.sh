#ugly staff this should replace with a dep tool like godep
echo "resolving depdencies"

go get github.com/go-sql-driver/mysql
go get github.com/go-xorm/xorm
go get github.com/PuerkitoBio/goquery
go get github.com/go-martini/martini

echo "resolve depdencies complete"