#mysql 
Install
```go
go get github.com/aprchen/db

```
---
How to use
```go
// before get data need use function LoadConfiguration to link mysql
err := db.Mysql().LoadConfiguration(db.MysqlMessageFromEnv());
// then get data like this:
rows, err := db.Mysql().Master.DB().Query(cond, vals...)
```
---
Recommended Use
- [sqlBuilder](https://github.com/didi/gendry.git) 