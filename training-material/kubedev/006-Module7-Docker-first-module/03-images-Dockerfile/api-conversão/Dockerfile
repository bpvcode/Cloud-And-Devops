FROM node

WORKDIR app/src

# VER EXPLICAÇÃO EM `../README.md` (## Criação de um Dockerfile - NOTAS)
COPY package*.json ./

RUN npm install

COPY . .

# Document the exposed port by the application
EXPOSE 8080

# Este comando não será executado na criação da imagem, mas sim na criação do container
# Por essa razão, uso o CMD em vez do RUN
CMD ["node", "index.js"]