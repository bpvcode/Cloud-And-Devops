# -------------------------- PERSONAL PROJECT INFO --------------------------

## Project Steps:

**1º - Build the image based on the [Dockefile](Dockerfile):**

```bash
docker build -t bpvcode/angular-app:v1 .
```

**2º - Run the container based on the image created in the last step:**
  * `-d` -> Modo daemon - Correr container em backgroud (Sem fazer parar o terminal)
  * `-p` -> Port binding - Mapear o porto `8080` do `localhost` com o porto `80` do container

```bash
docker run -d -p 8080:80 bpvcode/angular-app:v1
```

**3º - [OPTIONAL] Check if container is running:**

```bash
docker container ls
```

Output:

```bash
CONTAINER ID   IMAGE            COMMAND                  CREATED          STATUS          PORTS                  NAMES
1d0838b668f5   angular-app:v1   "/docker-entrypoint.…"   18 minutes ago   Up 18 minutes   0.0.0.0:8080->80/tcp   amazing_ride
```

**4º - App is running in [browser]("http://localhost:8080") - `http://localhost:8080`**

# -------------------------- DEFAULT ANGULAR PROJECT INFO --------------------------

## AngularTemplateApplicationImageMultistage

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 11.2.0.

## Development server

Run `ng serve` for a dev server. Navigate to `http://localhost:4200/`. The app will automatically reload if you change any of the source files.

## Code scaffolding

Run `ng generate component component-name` to generate a new component. You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.

## Build

Run `ng build` to build the project. The build artifacts will be stored in the `dist/` directory. Use the `--prod` flag for a production build.

## Running unit tests

Run `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io).

## Running end-to-end tests

Run `ng e2e` to execute the end-to-end tests via [Protractor](http://www.protractortest.org/).

## Further help

To get more help on the Angular CLI use `ng help` or go check out the [Angular CLI Overview and Command Reference](https://angular.io/cli) page.
