<p align="center">
  <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" width="120" alt="Go logo" />
</p>

# Farms Project
<p align="center">
  API for managing farmers, farms, and harvests.<br>
</p>

  Backend - Built with <a href="https://go.dev/" target="_blank">Golang</a>.  <br>
  [Frontend](https://github.com/samuel-prates/farmer-microfrontend/blob/master/README.md) - A microfrontend application built with **React 19**, **TypeScript**, **Redux Toolkit**, **Styled Components** and **Vite**.

---

## Description

Project design to manage farmers with theres farms.

## Features

- **Farmer**, **Farm**, and **Harvest** management
- **DTO** validation with class-validator and class-transformer
- **Swagger** (OpenAPI) documentation
- **Modular architecture**
- ⚛️ **React 19** with TypeScript
- 🗂️ **Redux Toolkit** for state management
- 🎨 **Styled Components** for CSS-in-JS
- 🧪 **Jest** and **React Testing Library** for unit and integration tests
- 🧬 **Atomic Design** component structure
- ⚡ **Vite** for fast development and builds
- 🧩 Ready for microfrontend architecture

## Makefile

This project includes a `Makefile` to simplify common development tasks.  
You can use the following commands:

| Command                | Description                               |
|------------------------|-------------------------------------------|
| `make submodule-init`  | import git submodules                     |
| `make submodule-update`| update git submodules                     |
| `make build-project`   | Install all dependencies                  |
| `make docker-up`       | Start Docker containers                   |
| `make docker-down`     | Stop and remove Docker containers/volumes |
| `make docker-restart`  | Rebuild and restart Docker containers     |
| `make docker-erase`    | Remove containers and image               |

**Usage example:**

```bash
make submodule-init
```

## Running the project

```bash
make submodule-init
make docker-up
```

## API Documentation

After running the project, access the Swagger UI at:

```
http://localhost:3000/api
```

## Testing

```bash
# unit tests
make test-project
```

## Technologies

- [NestJS](https://nestjs.com/)
- [TypeScript](https://www.typescriptlang.org/)
- [Swagger](https://swagger.io/)
- [class-validator](https://github.com/typestack/class-validator)
- [class-transformer](https://github.com/typestack/class-transformer)
- [React](https://react.dev/)
- [Redux Toolkit](https://redux-toolkit.js.org/)
- [Styled Components](https://styled-components.com/)
- [Vite](https://vitejs.dev/)
- [Jest](https://jestjs.io/)
- [React Testing Library](https://testing-library.com/docs/react-testing-library/intro/)



## License

MIT