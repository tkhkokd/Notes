
##Deploying WebApp to AWS using Elastic Beanstalk.

### Install Elastic Beanstalk Command Line.

```shell
$ brew update
$ brew install awsebcli
$ eb --version
```
at the repository's directory,
```
$ eb init
$ eb deploy ENVIRONMENT_NAME
// ENVIRONMENT_NAME can be found at the AWS console.
```

## EB CLI Commands

```shell
$ eb init -i  // Change region
$ eb printenv // Print ENV variables set
$ eb setenv   // Set ENV variables
$ eb deploy   // Deploy
$ eb open     // Open the endpoint
```

Setting ENV Variable
```shell
$ eb setenv
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

#### Bundle install fails
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
Solution1: Upgrade the instance memory

Solution2: Locally package gem dependency
=>
```
$ bundle package
$ ls vendor/cache

=> lists gem list
$ git add vendor/cache
$ git ci -m "added cache gems"
```

#### Error when bundle install Devise
```
 Bundler will use `/tmp/bundler/home/webapp' as your home directory temporarily.
  rake aborted!
  Devise.secret_key was not set. Please add the following to your Devise initializer:

    config.secret_key = 'e31c2202e452210388d0563f9b224ae19e778a8c56a7c48ee82df6d283e6587fa9910b349ea1ba90b8acb9df372da762b2629533685cd51709d6f3fb3b1f93dc'

  Please ensure you restarted your application after installing Devise or setting the key.
  /var/app/ondeck/config/routes.rb:2:in `block in <top (required)>'
  /var/app/ondeck/config/routes.rb:1:in `<top (required)>'
  /var/app/ondeck/config/environment.rb:5:in `<top (required)>'
  /opt/rubies/ruby-2.4.4/bin/bundle:23:in `load'
  /opt/rubies/ruby-2.4.4/bin/bundle:23:in `<main>'
  Tasks: TOP => environment
  (See full trace by running task with --trace) (Executor::NonZeroExitStatus)
```

```
$ eb setenv SECRET_KEY_BASE=$(rails secret) -e Testenv
```



