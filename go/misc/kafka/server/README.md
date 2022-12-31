# Links

https://kafka.apache.org/downloads

# VSCode extension
Tools for Apache Kafka

# Commands

brew install kafka # // /usr/local/opt/kafka
brew services restart kafka

## List all topics
/usr/local/opt/kafka/bin/kafka-topics --bootstrap-server=localhost:9092 --list

## List all consumer groups
/usr/local/opt/kafka/bin/kafka-consumer-groups --bootstrap-server=localhost:9092 --list

## Create a new topic
/usr/local/opt/kafka/bin/kafka-topics --bootstrap-server=localhost:9092 --topic=bondhello --create

## Produce data
/usr/local/opt/kafka/bin/kafka-console-producer --bootstrap-server=localhost:9092 --topic=DepositFundEvent

## Consume data
/usr/local/opt/kafka/bin/kafka-console-consumer --bootstrap-server=localhost:9092 --topic=bondhello --group=myconsumer1
/usr/local/opt/kafka/bin/kafka-console-consumer --bootstrap-server=localhost:9092 --topic=bondhello --group=myconsumer2

## Consume data by pattern
/usr/local/opt/kafka/bin/kafka-console-consumer --bootstrap-server=localhost:9092 --include="OpenAccountEvent|OpenAccountEvent|WithdrawFundEvent|CloseAccountEvent" --group=log

## Describe topic

/usr/local/opt/kafka/bin/kafka-topics --bootstrap-server=localhost:9092 --topic=bondhello --describe