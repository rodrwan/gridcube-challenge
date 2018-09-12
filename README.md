# Desafio de Gridcube

## Desafio

Crear 2 contenedores utilizando docker compose en su versión 3, el primer contenedor debe ser utilizado para generar una conexión vía API al servicio de Instagram para autorizar la autenticación desde el contenedor a un usuario especifico, luego de autorizar el ingreso a Instagram, desde el segundo contenedor se debe publicar una foto aleatoria en la cuenta del usuario con los datos de acceso obtenidos a través del primer contenedor. (es mandatorio usar el sistema de DNS embebido que ofrece docker a través de docker compose para comunicar los contenedores).

## Solución

### Problemas

Trate de utilizar el API de Instagram y de Facebook, ya que la de instagram al parecer será deprecada, sin tener resultados positivos.

### Work around

Busque en google y gracias a algunos consejos de Moisés, me puse a buscar si existia alguna librería que ya hiciera todo el manejo de sesión y de subida de archivos a Instagram, es así como llegue a [goinsta](https://github.com/ahmdrz/goinsta) la cual maneja todo lo relacionado con las sesión de instagram y además maneja la subida de imágenes.

Luego de esto manos a la obra.

## Ahora si, la solución

Son dos contenedores uno que se llama `auth` que es el encargado de recibir las credenciales de un usuario y además un caption, si son mandatorios.
Y otro contenedor llamado `publisher`  el cual se encarga de recibir estos datos, autenticar al usuario y elegir una imagen random desde [loremflickr](https://loremflickr.com) la que es posteada a la cuenta de instagram.

Para esta solución use Golang para hacer unos pequeños microservicios. Estos se comunican atraves de una conexión con [GRPC](https://grpc.io/) y `auth` tiene una interfaz publica REST con un método `POST` que recibe por medio de un JSON los datos del usuario junto con su caption de la siguiente forma:

```json
{
    "username": "...",
    "password": "...",
    "caption": "Algo bonito sobre la imagen random",
}
```

### Servicio Auth

- Puerto: 8090
- Protocolo: HTTP
- Método: POST

### Servicio Publisher

- Puerto: 8091
- Protocolo: GRPC
- Método: UploadImage