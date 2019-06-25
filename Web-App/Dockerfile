FROM node:12-stretch
MAINTAINER Valentin Marlier <valentin.marlier@ynov.com>

ENV MONGO_URI=mongodb://mongo:27018/monitoring
ENV TOKEN_KEY=yourtokenkey
ENV PORT=8081
ENV ADDRESS=0.0.0.0

WORKDIR /opt/web-app

COPY myapp/ ./

RUN chmod +x ./backend/build.sh
RUN chmod +x ./frontend/build.sh

RUN npm install --unsafe-perm

RUN node backend/scripts/addAdmin.js

EXPOSE 8081
CMD npm start