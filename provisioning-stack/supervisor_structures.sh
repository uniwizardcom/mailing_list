#!/bin/bash


tee /etc/supervisor/conf.d/10-application-api-worker.conf << EOF
[program:application-api-worker]
process_name=%(program_name)s_%(process_num)02d
directory=/var/applications/mailing_list
command=/var/applications/mailing_list/application-api
autostart=true
autorestart=true
stopasgroup=true
killasgroup=true
user=root
numprocs=1
redirect_stderr=true
stdout_logfile=/var/applications/mailing_list/logs/supervisor.log
stopwaitsecs=3600
EOF

supervisorctl reread
supervisorctl update
supervisorctl start application-api-worker:*
