FROM node:18.11.0

WORKDIR /home/node/app

COPY package.json yarn.lock ./
RUN yarn install

COPY . .

CMD if [ "$ENV" = "development" ]; then \
        yarn start; \
    else \
        yarn build && \
        yarn start -p 3100; \
    fi