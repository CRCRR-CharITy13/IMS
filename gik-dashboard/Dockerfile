# pull official base image
FROM node:16-alpine

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY package.json ./
COPY yarn.lock ./
RUN yarn install 

RUN yarn add global serve

EXPOSE 3000

# add app
COPY . .

ARG REACT_APP_API_URL

RUN sh .docker/create_env_file.sh REACT_APP_API_URL=$REACT_APP_API_URL

RUN yarn build

# start app
CMD serve -s build