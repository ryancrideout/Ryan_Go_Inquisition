FROM node:22-alpine

# Change the working directory to install bootstrap
WORKDIR /opt/go_react/react/react-app
RUN npm install bootstrap@5.2.3
RUN npm install react-router-dom
RUN npm install classnames --save
RUN npm install -D sass

EXPOSE 5173

WORKDIR /opt/go_react/react/react-app