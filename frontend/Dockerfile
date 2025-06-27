# Dockerfile (frontend)
FROM node:20-alpine

WORKDIR /app

COPY . .

RUN npm install
RUN npm run build

# Utilise un serveur web minimal pour servir le build
RUN npm install -g serve
CMD ["serve", "-s", "dist"]

EXPOSE 3000
