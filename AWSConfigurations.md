
##Deploying WebApp to AWS using Elastic Beanstalk.

### Install Elastic Beanstalk Command Line.

```shell
brew update
brew install awsebcli
eb --version
```
at the repository's directory,
```
eb init
eb deploy ENVIRONMENT_NAME
// ENVIRONMENT_NAME can be found at the AWS console.
```

## EB CLI Commands

```shell
eb init -i  // Change region
eb printenv // Print ENV variables set
eb setenv   // Set ENV variables
eb deploy   // Deploy
eb open     // Open the endpoint
```

Setting ENV Variable
```shell
eb setenv
RDS_DB_NAME=ebdb
RDS_USERNAME=***********
RDS_PASSWORD=*********** RDS_HOSTNAME=***********.rds.amazonaws.com // The endpoint
RDS_PORT=****
```
* ENV variables have to be updated after rebuilding the environment.


In Rails App ```config/database.yml```
```yml
default: &default
  adapter: postgresql
  encoding: unicode
・・・(中略)
production:
  <<: *default
  database: <%= ENV['RDS_DB_NAME'] %>
  username: <%= ENV['RDS_USERNAME'] %>
  password: <%= ENV['RDS_PASSWORD'] %>
  host: <%= ENV['RDS_HOSTNAME'] %>
  port: <%= ENV['RDS_PORT'] %>
```






