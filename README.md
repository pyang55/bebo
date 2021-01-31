# bebo
Venmo Scheduler

## Requirements
- Docker
- Venmo Token(https://api.venmo.com/v1/oauth/authorize?client_id=1112&scope=access_profile,make_payments&response_type=code)

## Build and Installation
- Fill out bebo.env file in local directory with time zone and venmo token
- docker-compose build
- docker-compose up -d
- Go to your browser localhost:8080

More features coming soon
UI is utilizing https://github.com/alseambusher/crontab-ui
