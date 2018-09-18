## AWS SDK

```
pip3 install boto3
aws configure
// Set credentials
python3
```

## Dynamo DB
  Primary key
    Partition key - ex. itemID
    Sort key      - ex. SongID, AlbumID, EventID
  Point-in-time recovery (PITR) - Automatic backups of table data.
  Encryption (at rest) - uses AWS Key Management Service. Can only be enabled at table creation. Can't be disabled.
  Time to live (TTL) - Delete items when they expire. No extra cost by enabling TTL. Per item basis.



## Lambda

```python
def handler_name(event, context):   # event: some data passed by the event(trigger)
    ...                             # context: runtime info
    return some_value
```
### Invoke lambda function using SDK

```
import boto3

client = boto3.client('lambda')

response = client.invoke(
    FunctionName='string',
    InvocationType='Event'|'RequestResponse'|'DryRun',
    LogType='None'|'Tail',
    ClientContext='string',
    Payload=b'bytes'|file,
    Qualifier='string'
)
```
Invocation Type
  Event: asyncronous
  RequestResponse: syncronous
  DryRun: authorization required before running

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

### EB CLI Commands

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

## Solutions to Errors

### Bundle install fails
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

### Error when bundle install Devise
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


## AWS IoT Configuration


Sending Message from clients
Supported protocols : MQTT, HTTP, Websocket
https://docs.aws.amazon.com/iot/latest/developerguide/protocols.html

HTTP - use curl to emulate a button press
```
curl --tlsv1.2 --cacert root-CA.crt --cert 4b7828d2e5-certificate.pem.crt --key 4b7828d2e5-private.pem.key -X POST -d "{ \"serialNumber\": \"G030JF053216F1BS\", \"clickType\": \"SINGLE\", \"batteryVoltage\": \"2000mV\" }" "https://a1pn10j0v8htvw.iot.us-east-1.amazonaws.com:8443/topics/iotbutton/virtualButton?qos=1"
```
--tlsv1.2 (MUST)
curl installed with OpenSSL and use version 1.2 of TLS.

--cacert <filename>
CA certificate file name to verify the peer.

--cert <filename>
The client certificate filename.

--key <filename>
The private key filename.

-X POST
The type of request.

-d <data>
POST data you want to publish

"https://..."
The REST API endpoint for the device.
At AWS IoT console choose Registry to expand your choices.
Parameters: port, topic, the quality of service in a query string (?qos=1).



## Accessing Database(Postgresql) from AWS CLI

List RDS instances
```
aws rds describe-db-instances
```

Connect to the DB

```
psql \
   --host=<DB instance endpoint> \
   --port=<port> \
   --username=<master user name> \
   --password \
   --dbname=<database name>
```
You will be prompted for the password.
* remove \

## AWS IoT
### IoT devices and shadow RESTful API

Shadow RESTful API endpoint URL
```
https://ENDPOINT/things/THING_NAME/shadow
```
Get endpoint url
```
aws iot describe-endpoint
```
List things from CLI
```
aws iot list-things
```

### Jobs

Create new job
```
import boto3

client = boto3.client('iot')
response = client.create_job(
    jobId='string',                            // Specify an ID.
    targets=[                                  // Device ARN (Amazon Resource Names)
        'string',
    ],
    documentSource='string',                   // Document link at S3.
    document='string',                         // Document title
    description='string',                      // Short description of the job
    presignedUrlConfig={
        'roleArn': 'string',                   // Role ARN defined in IAM.
        'expiresInSec': 123
    },
    targetSelection='CONTINUOUS'|'SNAPSHOT',   // See the description below for the detail.
    jobExecutionsRolloutConfig={
        'maximumPerMinute': 123
    },
    documentParameters={                       // Parameters for the job document.
        'string': 'string'
    }
)
```

- Job: is a remote operation. Created and managed using the Jos HTTPS API, AWS CLI or AWS SDKs
- Job document: JSON document. Containing descriptions of the remote operations to be performed by the devices. Stored in S3 bucket
- Target: target devices or groups of devices
- Protocols: MQTT/HTTP SigV4/HTTP TLS
- Job execution: target starts an execution of a job by downloading the job document
- Snapshot job: a single job, not continuous
- Continuous job: a continuous job
- Rollouts:
- Presigned URLs: for time-limited access to data in job document.

### Message Broker

- Publish & Subscribe broker service
- Sends and receives messages to and from AWS IoT
- A client can send messages, specifying a topic URL, all the clients subscribing to the topic receives the message
- Protocols: MQTT/HTTP/MQTT+Websocket

### Worlflow

First, a device has to subscribe to topics in order for the device to start performing jobs.

Use AWS SDK or AWS CLI.

1. Device comes online and subscribe to ```notify-next``` topic.
2. Call the ```DescribeJobExecution``` MQTT API with jobId ```$next```
3. Call the ```UpdateJobExecution``` MQTT API
* ```StartNextPendingJobExecution``` does step 2 and 3 at once.
4. Perform the actions in the job document by ```UpdateJobExecution```
5. ```DescribeJobExecution``` to monitor job execution
6. Call the ```UpdateJobExecution``` MQTT API to update the job execution status

### Amazon RDS Security

By default network access is turned off to a DB instance.

Related Security Groups:
- VPC Security Groups => Between DB instances and EC2 instances inside a VPC. Enables specific source to access a DB instance in a VPC.

- EC2 Security Groups => Access to EC2

- DB Security Groups (Legacy) => EC2-Classic DB instances outside of the VPC (EC2-Classic is the old version of EC2-VPC, new customers after 2013-12-04 will only use EC2-VPC)

1. specify rules in a security group (about IP address range, port, EC2 security group)

#### VPC Security Groups

Set Inboud Rules & Outbound Rules, Tags (optional).

Enables specific source to access a DB instance in a VPC.
By specifying:
- A range of addresses (ex. 203.0.113.0/24) or
- Another VPC security group.

Specify a port from which each range of address accesses DB instances.
Need SSH access to instances in the VPC?
=> Allow access to TCP port 22 for the specified range of addresses.










