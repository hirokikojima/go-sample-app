FROM node:10.16.3

ENV NODE_ENV=development
ENV HOST 0.0.0.0

# Make work directory
RUN mkdir /app
WORKDIR /app

# Copy app
COPY ./client .

# Build the nuxt app
RUN npm i
RUN npm run build

# Command to run the executable
EXPOSE 3000
CMD ["npm", "start"]