# web-app

# Build it and run it

## Requirements

- Node v10.15.0 (npm v6.4.1)
- MongoDB (see below for basic mongodb setup in a Docker container)

```bash
# Create an attachable docker network
sudo docker network create $network_name --attachable

# Then run a new container
sudo docker run --name $container_name -p $host_port:27017 -v $volume_path:/data/db --network $network_name -d --restart=on-failure mongo:$version

# You can add basic authentication with environment tags
sudo docker run --name $container_name -p $host_port:27017 -v $volume_path:/data/db --env MONGO_INITDB_ROOT_USERNAME="root" --env MONGO_INITDB_ROOT_PASSWORD="toor" --network $network_name -d --restart=on-failure mongo:$version
```

## Production

### Build Setup

1. Install dependencies and build front end app

```bash
# For Windows, run : `npm config set script-shell "C:\\Program Files\\git\\bin\\bash.exe"`, "C:\\Program Files\\git\\bin\\bash.exe" being the path to your git executable
npm install
```

2. Define your environment variables and start MongoDB

```bash
# MongoDB URL
MONGO_URI=mongodb://X.X.X.X:27017/collectionName

# Token Key for authentication
TOKEN_KEY=yourtokenkey
```

3. Run the script to add the admin in DB

```bash
node backend/scripts/addAdmin.js
```

4. Run backend server

```bash
# Run backend server
npm start
```

### Troubleshooting

If you have an error during the `npm install` step, try one of the following:

1. execute the script with Yarn :

```bash
yarn install
```

2. run the following commands manually instead of executing the script :

```bash
# backend
cd backend && npm ci
# frontend
cd frontend && npm ci && node ./build/build.js
```

## Local

### Environment variables

- In local you need to create these files:

##### /backend/.env

```bash
# MongoDB URL
MONGO_URI=mongodb://0.0.0.0:27017/monitoring

# Token Key for authentication
TOKEN_KEY=ih4veN0id3a
```

##### /frontend/.env

IMPORTANT: Only needed in local environment

```bash
# URL of API
BASE_API=http://localhost:8080
PORT=8082

```

### Local backend Setup

Back end developement build with hot reloading

```bash
# go to front end folder
cd ../backend

# start server (and serve prod front files on localhost:8080 )
npm start

# activate hot reload with nodemon
npm run dev
```

### Local frontend Setup

Front end developement build with hot reloading

```bash
# go to front end folder
cd frontend

# serve with hot reload at localhost:8082
npm run dev
```

Front end production build

```bash
# build for production with minification
npm run build

# build for production and view the bundle analyzer report
npm run build --report
```

## CI/CD Informations

This web-app can be customized for testing, development and production use.

##### ./Dockerfile && ./Dockerfile-dev

line 4 to 7:
- ENV MONGO_URI=mongodb://mongo:27017/monitoring -> the mongodb URL
- ENV TOKEN_KEY=yourtokenkey -> user token (can be any string)
- ENV PORT=8081 -> change the port where the server will listen
- ENV ADDRESS=0.0.0.0 -> change the address where the server will listen

###### ⚠ BEST PRACTICE:
```
For the MONGO_URI:
if you have specified you're mongodb container with a network and you link your web-app built container to the mongodb network, you can specify mongo:27017 instead of a real IP address.
```

##### ./.gitlab-ci.yml

line 50 and 71:
- [...] "docker run --name $web-app-container-name -p $host-port:$port-defined-in-Dockerfile --link $mongodb-container-name:mongo --network $mongodb-network -d --restart=on-failure $registry/$image"
