#!/bin/bash

rabbitmqctl add_user un St1r.o2n,gP3a?5ssword
rabbitmqctl set_user_tags un administrator
rabbitmqctl set_permissions -p "/" "un" ".*" ".*" ".*"

# /etc/rabbitmq/rabbitmqadmin declare queue --vhost=/ name=queue_name durable=true
# /etc/rabbitmq/rabbitmqadmin delete queue name=queue_name

# for list of queue: rabbitmq-queues
/etc/rabbitmq/rabbitmqadmin declare queue --vhost=/ name=mailing_messages durable=true
