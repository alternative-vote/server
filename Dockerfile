FROM node:5

# Make working directory
RUN mkdir -p /usr/src/app/
WORKDIR /usr/src/app/

# Install source
COPY . /usr/src/app/
RUN npm install

# Start app
EXPOSE 3000
CMD ["npm", "start"]