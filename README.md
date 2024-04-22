Para que esto funcione utiliza una base de datos postreSQL que corre en un container en DOCKER.

Primero crear el container con la imagen de postgres en docker con lo siguiente:

docker run --name some-postgres -e POSTGRES_USER=leo -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres

El nombre del container en este ejemplo es some-postgres
Si no tenes la imagen, docker la busca e instala, luego de que este todo creado, en el mismo comando ya le pasé parametros para postres, estos son:

-e POSTGRES_USER = leo       es para crear el usuario de la BD, hay uno por defecto, pero es mejor crear uno
-e POSTGRES_PASSWORD =       creo que si no colocas un pass no te pide nada, pero es mejor agregar uno para practicar, 
                             no existen conexiones sin passwords
-p 5432:5432  este parametro es para docker, le indica con que puerto local hacer un proxy respecto al puerto publicado en la imagen de docker
-d            este parametro es para docker, indica que corra de forma detached, es decir que corra en segundo plano siempre
postgres    esto indica el nombre de la imagen que tiene que ir a buscar, esto lo obtiene de dockerhub https://hub.docker.com/_/postgres

Podes conectar la BD directo con algun gestor como DBEAVER, o usarlo directo desde la consola con docker. Hay que crear la BD gorm ppara este ejemplo.

CREATE DATABASE gorm; 

SI queres conectarte a la BD desde la consola hacer esto:
1 - docker exec -it some-postgres bash  conectarse con el bash de docker, y le indicas a que container en este caso some-postgres, it es modo interactivo.

2- psql -U leo --password   te logueas con el user de la BD, y podes ejecutar comandos de sql
3 - se puede consultar relaciones con \d o \dt o \l sino directamente podes ingresar queries de SQL. Por ej para crear la base datos gorm. Pero podes cosas como SELECT * from users (si la tabla ya estuviera creada)

------------------------------------------------------------------------------------

Las tablas se crean usando GORM, con las migraciones que te permite configurar. Esto lo ves desde los modelos, hay unas notaciones para hacer ahi. Ademas tenes que generar la conexion a la base de datos usando gorm, previo tenes que instalar el driver de la base de datos, por ej si usas mysql usa un driver si usas postgres usa otro. Esto lo encontras en la pagina de GORM.

https://gorm.io/docs/connecting_to_the_database.html

de ahi te fijas el import del driver que usa el ejemplo, en este caso   "gorm.io/driver/postgres" y usas esto para instalarlo. 
Para instalar usas el comando    go get -u gorm.io/driver/postgres   para que lo agregue al go.mod y este disponible para usar en la conexion.
s