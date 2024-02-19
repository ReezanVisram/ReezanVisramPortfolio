FROM node:21-alpine as build

WORKDIR /app
COPY . .
RUN npm i
RUN npm run build


FROM node:21-alpine

WORKDIR /app

COPY --from=build /app/package*.json ./
COPY --from=build /app/build ./

EXPOSE 3000
CMD ["node", "./index.js"]