# Financiera Compartamos Challenge - backend

El siguiente Api Rest que cuenta con las funcionalidades de CRUD , validaciones del body request, modular aplicando patrones de diseño como inyecion de dependencia.
En la carpeta assets encontrart un archivo curl para probarlos endpoints.

Aplicaciones necesarias
- Docker desktop
- go v20

1. Duplicar el archivo ".env.template" y renombrar ".env"

2. Para levantar el api ejecutar los siguientes comandos si cuenta con make ejecutar los comando make sino los normales.
```bash
db_up o docker-compose up -d
```
3. Ejecutar el comando para descargar las dependencias
```bash
go mod download
```
4. Ejecutar el comando para arrancar el servidor

```bash
make run o go run./...
```