# go-student-service
The basics demo for go student-service

# Go Installation
```
Download go from the source
Open ~/.zshrc
Add GOPATH and GOBIN 
Add GOBIN to PATH

export GOPATH=$HOME/go/go1.19.3
export GOBIN=$HOME/go/go1.19.3/bin
export PATH="$GOBIN:$PATH"

Install VSCODE
Install Go Extension
https://learn.microsoft.com/en-us/azure/developer/go/configure-visual-studio-code
```

# Mysql Docker Command

Mysql 

`
docker run --name go-mysql -e MYSQL_ROOT_PASSWORD=password -d -p 3306:3306 mysql
`
# Database Initializer
```
CREATE DATABASE Test;
DROP TABLE IF EXISTS Test.student;
CREATE TABLE Test.student (
  id      INT AUTO_INCREMENT NOT NULL,
  name    VARCHAR(255) NOT NULL,
  age     INT NOT NULL,
  level   INT NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO Test.student (name, age, level) VALUES
  ('Hussien Mohamed', 20, 1),
  ('Ahmed Rady', 22, 3),
  ('Wael Gamel', 21, 2),
  ('Ziad Gamal', 22, 3);
```

# First Steps To Create Program

```
go mod init example/student-service
touch main.go
```