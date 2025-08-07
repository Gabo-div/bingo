# BinGOD!

> El bingo por defecto hecho por y para el pueblo Venezolano.

Un monorepo para una aplicación de Bingo full-stack, construida con un frontend en React (Vite) y un backend en Go, utilizando gRPC para la comunicación.

## 📂 Estructura del Proyecto

Este es un monorepo que utiliza `pnpm workspaces` y `Turborepo`.

-   `apps/web`: La aplicación frontend construida con React y Vite.
-   `apps/server`: El servidor principal de la aplicación en Go.
-   `apps/auth`: Servicio de autenticación en TypeScript.
-   `packages/database`: Contiene las migraciones de la base de datos, queries SQL y el código Go generado para el acceso a datos.
-   `packages/protobuf`: Define los contratos gRPC (Protobuf) para la comunicación entre servicios.

## 🚀 Instalación y Uso

Sigue estos pasos para poner en marcha el entorno de desarrollo.

### Prerrequisitos

Es necesario tener instalado lo siguiente en tu sistema:

-   [**pnpm**](https://pnpm.io/installation)
-   [**Go**](https://go.dev/doc/install)
-   [**air**](https://github.com/air-verse/air) para live-reloading en Go.
-   [**Docker**](https://www.docker.com/get-started)

### Pasos de Instalación

1.  **Instalar dependencias de pnpm:**
    ```sh
    pnpm i
    ```

2.  **Preparar dependencias de Go:**
    Este comando instalará las dependencias de los módulos de Go en el workspace.
    ```sh
    pnpm run prep
    ```

3.  **Configurar variables de entorno:**
    Este script creará los archivos `.env` necesarios para los diferentes servicios a partir de los ejemplos.
    ```sh
    ./env-cpy.sh
    ```
    *Nota: Puede que necesites dar permisos de ejecución al script (`chmod +x env-cpy.sh`).*

4.  **Iniciar la base de datos con Docker:**
    ```sh
    docker-compose up -d
    ```

### Ejecución

Para iniciar todos los servicios en modo de desarrollo (con hot-reloading):

```sh
pnpm run dev
```

Esto ejecutará el frontend y el backend simultáneamente gracias a Turborepo.