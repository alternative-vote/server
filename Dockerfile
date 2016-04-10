FROM node:5

# Make working directory
RUN mkdir -p /usr/local/server
WORKDIR /usr/local/server

# Install source
COPY . /usr/local/server
RUN npm install

# Start app
EXPOSE 3000
CMD ["npm", "start"]