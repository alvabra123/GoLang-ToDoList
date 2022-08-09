# GoLang-ToDoList

1. Install MySQL Server
```
sudo apt install mysql-server
```
2. Check MySQL Service
```
sudo systemctl status mysql
sudo systemctl start mysql.service
```
3. Create MySQL Database
```
sudo mysql
CREATE DATABASE GoDB;
USE GoDB;
```
4. Create MySQL Table
```
CREATE TABLE ToDoList (
  ID INT(4) UNSIGNED NOT NULL AUTO_INCREMENT,
  Assignee VARCHAR(100) NOT NULL,
  Assigner VARCHAR(100) NOT NULL,
  Task VARCHAR(10000) NOT NULL,
  Start VARCHAR(100) NOT NULL,
  Finish VARCHAR(100) NOT NULL,
  Status VARCHAR(100) NOT NULL,
  PRIMARY KEY (ID)
);
```
5. Download and Install Go
```
https://go.dev/doc/install
```
6. Install Go MySQL Driver
```
go get -u github.com/go-sql-driver/mysql
```
7. Edit main.go Based on Your MySQL Settings
```
func dbConn() (db *sql.DB) {
  dbDriver := "mysql"
  dbUser := "root"
  dbPass := ""
  dbName := "GoDB"
  db,err := sql.Open(dbDriver,dbUser+":"+dbPass+"@/"+dbName)
  if err != nil {panic(err.Error())}
  return db
}
```
8. Run main.go
```
go run main.go
```
9. Open Browser and Type URL
```
http://localhost:8080
```
10. Result
![alt text](https://github.com/alvabra123/GoLang-ToDoList/blob/main/ToDoList/assets/result.png)
