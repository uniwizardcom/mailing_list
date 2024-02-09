#!/bin/bash

# for list of queue: rabbitmq-queues
/etc/rabbitmq/rabbitmqadmin declare queue --vhost=/ name=mailing_group_subscription durable=true
/etc/rabbitmq/rabbitmqadmin declare queue --vhost=/ name=mailing_group_send durable=true
/etc/rabbitmq/rabbitmqadmin declare queue --vhost=/ name=mailing_group_delete durable=true
