# Desafio de Gridcube

## Desafio

Crear 2 contenedores utilizando docker compose en su versión 3, el primer contenedor debe ser utilizado para generar una conexión vía API al servicio de Instagram para autorizar la autenticación desde el contenedor a un usuario especifico, luego de autorizar el ingreso a Instagram, desde el segundo contenedor se debe publicar una foto aleatoria en la cuenta del usuario con los datos de acceso obtenidos a través del primer contenedor. (es mandatorio usar el sistema de DNS embebido que ofrece docker a través de docker compose para comunicar los contenedores).

## Solución

### Problemas

Trate de utilizar el API de Instagram y de Facebook (Graph API), ya que la de instagram al parecer será deprecada, sin tener resultados positivos. Esto lo expliqué en extenso en otro email (privado).

### Workaround

Busque si existia alguna librería que ya hiciera todo el manejo de sesión y de subida de archivos a Instagram, por suerte, encontre [goinsta](https://github.com/ahmdrz/goinsta), esta librería maneja todo lo relacionado con la sesión de instagram y además maneja la subida de imágenes.


## Ahora si, la solución

Cree dos contenedores

- `auth` que es el encargado de recibir las credenciales de un usuario y además un caption.
- `publisher` el cual se encarga de recibir estos datos, autenticar al usuario y buscar una imagen aleatorea desde [loremflickr](https://loremflickr.com), que luego es posteada a la cuenta de instagram.

Para esta solución use como lenguaje [Golang](https://golang.org/) para hacer unos pequeños microservicios.
Estos servicios se comunican usando [GRPC](https://grpc.io/), `auth` tiene una interfaz publica REST con un método `POST` que recibe por medio de un JSON los datos del usuario junto con su caption de la siguiente forma:

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
- Endpoint: /photo

### Servicio Publisher

- Puerto: 8091
- Protocolo: GRPC
- Método: UploadImage

### ¿Cómo correr el código?

```sh
$ make run
```

luego en otra terminal

```sh
$ curl -X POST -H "Content-Type:application/json" http://localhost:8090/photo -d '{"username": "---", "password": "---", "caption": "caption de prueba"}'
```

Si todo sale bien un OK debería desplegarse en la terminal y en la cuenta de instagram la imagen aleatoria.

