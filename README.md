Aplicación de Chat cliente-servidor. El servido se construye en **Go** y el cliente en **Vala**.

El servidor es concurrente; la comunicación entre él y los usuarios se hace mediante Sockets utilizando el protocolo JSON.  

El proyecto utiliza **[Taskfile](https://taskfile.dev/)**  para automatizar la construcción de los programas.


## Requisitos Previos

Asegúrate de contar con los siguientes paquetes y herramientas en tu sistema operativo:

### Entorno Go
* [Go](https://go.dev/) (versión 1.22 o superior).

### Entorno Vala y Dependencias del Sistema
* Compilador **`valac`**.
* Bibliotecas  de **GLib / Gio** y **Json-Glib**.



## Gestión con Taskfile
El archivo `Taskfile.yml` incluido en la raíz del repositorio define las tareas de construcción.

### Tareas disponibles

| Comando | Descripción |
| :--- | :--- |
| `task` | Compila tanto el servidor en Go como el cliente en Vala |
| `task build-servidor` | Compila el servidor Go (`bin/Servidor`) |
| `task build-cliente` | Compila el cliente Vala (`bin/Cliente`) |
| `task run-servidor` | Inicia el servidor (puerto `5050` por defecto o personalizado) |
| `task run-cliente` | Inicia el cliente |
| `task clean` | Elimina la carpeta de binarios `bin/` |

---

## Compilación y Ejecución

### Compilación Completa

Para compilar tanto el servidor como el cliente en un solo paso ejecuta:

```bash
task
```
Esto generará los ejecutables correspondientes dentro del directorio `bin/`:
- **`bin/Servidor`**: Servidor en Go.
- **`bin/Cliente`**: Cliente en Vala.

### Compilación Individual
Para compilar uno solo de los programas:
```bash
# Solo el servidor Go
task build-servidor

# Solo el cliente Vala
task build-cliente
```
---

### Ejecución

El servidor requiere que se especifique el **puerto de escucha** como argumento de línea de comandos.

### Iniciar el Servidor 


```bash
# Usa el puerto 5050 por defecto:
task run-servidor

# Ó puedes especificar un puerto diferente :
task run-servidor -- 9000
```

El cliente requiere que se especifiquen la **direción ip** y el **puerto**, en ese orden, como argumentos de línea de comandos.

### 2. Iniciar el Cliente 
```bash
# Especifica la IP y el puerto a los cuales conectarse:
task run-cliente -- 127.0.0.1 5050
```

