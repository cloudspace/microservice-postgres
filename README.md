# Microservice-postgres (Go)
A service that runs postgres queries written in Go.


## Input
Command
./main [db_host] [db_port] [db_user] [db_password] [db_name] [query]

## Output
JSON string

Example:
query=SELECT * FROM feeds LIMIT 1
{"result":"[{"approved":true,"created_at":"2014-11-06T19:53:24.239813Z","default":false,"etag":"qqHU7VQHtNatrqA/YlVDJvQGVV8","id":10,"last_modified_at":"2014-11-10T13:04:51Z","last_parsed_at":"2014-11-10T14:58:23.947517Z","name":"Pinch My Salt","next_parse_at":"2014-11-10T23:30:23.947786Z","parse_backoff_level":7,"parser_options":null,"process_end":"2014-11-10T14:58:23.98856Z","process_start":null,"processing":false,"scheduled":false,"site_url":"http://pinchmysalt.com","summary":"Food, Recipes, and Photography","updated_at":"2014-11-10T14:58:23.94935Z","url":"http://feeds.feedburner.com/pinchmysalt"}]"}

## To build

Building Dockerfile: docker build -t johnleeroy/microservice-postgres

Running Dockerfile: docker run -ti johnleeroy/microservice-postgres ./main [...]
