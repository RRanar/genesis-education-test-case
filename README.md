## Setup ##
For starting this application you need to have correct .env file that will be use for
Docker run command
In project you can find .env.example file that can help fill all needed environment variables
for starting project
For start project first build docker image
```bash
docker build -t sample-image-name .
```
and next start image
```bash
docker run --rm -it -p your-port:container-port --env-file your-env-file sample-iamge-name
```