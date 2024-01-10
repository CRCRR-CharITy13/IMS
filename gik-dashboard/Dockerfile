# Use an official Node.js runtime as the base image
FROM node:19.6.0-alpine as build
# Set the working directory within the container
WORKDIR /gik-dashboard

# Copy package.json and package-lock.json to the container
COPY package*.json ./

# Install dependencies
RUN npm install --force

# Copy the rest of the application code to the container
COPY . .

# Build the React app
RUN npm run build 

# Use Nginx as the production server
FROM nginx:alpine

# Copy the build React app to Nginx's web server directory
COPY --from=build /gik-dashboard/build /usr/share/nginx/html
# Expose port 80 for the Nginx server
EXPOSE 80

# Start Nginx when the container runs
CMD ["nginx", "-g", "daemon off;"]