# stock_exchange_api

## Requerimientos

* Clonar este proyecto en GOPATH/src/github.com/api/stock_exchange-api

## Dependencias

Las dependencias estan agregadas con el go.mod, en caso de agregar una nueva dependencia realizar un go get

## Configuracion

Lo primero que necesita hacer es cargar las variables de entorno, estas se pueden cargar en el archivo raiz del servidor o temporal por proyecto, para ejecutarlas de manera temporal se debe crear un archivo config.sh, cargar las variables en el archivo y ejecutar la siguiente sentencia (Linux o Mac)

	$ source config/config.sh

Las variables de entorno que se deben cargar son:

````
export DBNAME=stock_exchange
export DBUSERNAME=root
export DBUSERPASS=123456
export DBHOST=localhost
export DBPORT=3306
export JWTKEY=NsABfTzgPFnby8lx
export ENCRYPT_KEY=bc3478c930c16bd2214345158256aca1
````

## Arquitectura

### controllers
* Contiene los endpoints y tiene como objetivo recibir la peticion con sus respectivos parametros, validarlos y enviarlos al helper para continuar el proceso

### db
* Contiene la conexion a la base de datos

### helpers
* Carpeta que contiene la logica del proyecto; realiza la conexion con los querys en repository, valida la informacion con los utils y devuelve una respuesta al controller para que esta pueda ser consumida

### httpmodels
* Estructuras para el manejo de datos en el proyecto tales como el request o response de los endpoints

### utils
* Funciones generales que se usan para validaciones o ajustes, entre estos esta el metodo de encriptar o desencriptar la contraseña

### main.go
* Archivo principal que compila el proyecto y divide las peticiones o endpoints.

## Arrancar el proyecto

Golang es un proyecto compilado por lo cual se debe ejecutar la siguiente sentencia para que el proyecto se ejecute como desarrollador:

	$ go run main.go

Para ejecutar el proyecto y crear el archivo compilado como ambiente de produccion se debe ejecutar la siguiente sentencia:

	$ go build main.go -o stock-exchange

Al ingresar a [http://127.0.0.1:5000](http://127.0.0.1:5000) podra ver el proyecto corriendo.