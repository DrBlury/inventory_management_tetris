# Start from a Node.js image.
FROM node:18-alpine

# Set the working directory.
WORKDIR /app

# # Copy package.json and package-lock.json for installing Node.js dependencies.
# COPY package.json package-lock.json ./

# Install Node.js dependencies.
RUN npm install -g prettier

# Install prettier plugins
RUN npm install --save-dev prettier-plugin-go-template
RUN npm install --save-dev prettier-plugin-sql

# Run the format-all task.
CMD prettier --write "**/*.{yaml,yml,json,md,sql,templ}"