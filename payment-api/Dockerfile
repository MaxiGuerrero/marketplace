# BUILDER STAGE
FROM node:lts-alpine AS builder

# Checking version node
RUN node --version
# Select workdir
WORKDIR /build

# Copy code
COPY . /build

# Build application 
RUN npm install
RUN npm run build

# LAUNCHER SERVER STAGE
# Image
FROM node:lts-alpine
ENV DIR_SWAGGER='./dist/src/shared/docs/swagger.yml'

# Select workdir
WORKDIR /usr/src/app

# Checking version node
RUN node --version

# Add artifacts from builder image
COPY  --chown=node:node --from=builder /build/dist dist
COPY  --chown=node:node --from=builder /build/package.json package.json
COPY  --chown=node:node --from=builder /build/tsconfig.json tsconfig.json

RUN npm install --omit=dev

RUN ls .
RUN ls ./dist/
# Run server
ENTRYPOINT ["npm", "run", "start:prod"]

