
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
* RDS_HOSTNAME has to be updated after rebuilding the environment.


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

### Solutions to Errors

#### Bundle install fails => Upgrade the instance memory
```
An error occurred while installing nokogiri (1.6.8.1), and Bundler cannot
continue.
Make sure that `gem install nokogiri -v '1.6.8.1'` succeeds before bundling.

Installing nokogiri 1.6.8.1 with native extensions

Gem::Ext::BuildError: ERROR: Failed to build gem native extension.

    current directory: /opt/rubies/ruby-2.3.1/lib/ruby/gems/2.3.0/gems/nokogiri-1.6.8.1/ext/nokogiri
/opt/rubies/ruby-2.3.1/bin/ruby -r ./siteconf20161205-3296-zuwti1.rb extconf.rb
Cannot allocate memory - /opt/rubies/ruby-2.3.1/bin/ruby -r ./siteconf20161205-3296-zuwti1.rb extconf.rb 2>&1
```






