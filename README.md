# The "mailing_list" project
A microservice that will store customers and send them emails based on mailing ID

## Examples:

### Subscribe new messages for mailing

`curl -X POST 172.16.10.10:8080/api/messages -d '{"email":"jan.kowalski@example.com","title":"Interview","content":"simple text","mailing_id":1, "insert_time": "2020-04-24T05:42:38.725412916Z”}'`

`curl -X POST 172.16.10.10:8080/api/messages -d '{"email":"jan.kowalski@example.com","title":"Interview","content":"simple text","mailing_id":2, "insert_time": "2020-04-24T05:42:38.725412916Z}'`

`curl -X POST 172.16.10.10:8080/api/messages -d '{"email":"jan.kowalski@example.com","title":"Interview","content":"simple text","mailing_id":3, "insert_time": "2020-04-24T05:42:38.725412916Z}'`

### Deleting subcribe from mailing

`curl -X DELETE 172.16.10.10:8080/api/messages/{id}`

### The sending action to subscribe to all (mailing_id = -1) or only from specific group (mailing_id > 0)

` curl -X POST 172.16.10.10:8080/api/messages/send -d '{"mailing_id":1}'`

*********************************************************************************************

## Building VM and provisioning:

### create/start VM's
1) `provisioning-machines/main$ vagrant up`

### install needed services on VM's (all must be started inside destiny VM)
1) `provisioning-stack/base.sh`
2) `provisioning-stack/postgresql_server_install.sh`
3) `provisioning-stack/postgresql_structures.sh`
5) `provisioning-stack/rabbitmq_server_install.sh`
6) `provisioning-stack/rabbitmq_structures.sh`
4) `provisioning-stack/redis_server_install.sh`

### GUI for services
1) RabbitMQ: http://172.16.10.10:15672/#/

*********************************************************************************************


