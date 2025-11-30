# GUP - Herramienta CLI para Mantenimiento de Servidor Ubuntu

GUP (Go Update) es una herramienta de línea de comandos desarrollada en Go que utiliza Bubble Tea y Cobra para automatizar tareas comunes de mantenimiento en servidores Ubuntu.

## 🚀 Características

- **Interfaz Interactiva**: Utiliza Bubble Tea para una experiencia de usuario moderna y atractiva
- **Comandos Estructurados**: Implementado con Cobra para un manejo robusto de comandos y argumentos
- **Mínimas Dependencias**: Diseñado para usar las menos dependencias posibles
- **Binario Standalone**: Genera un binario ejecutable independiente
- **Multiidioma (i18n)**: Soporte completo para Español e Inglés, con detección automática del idioma del sistema

## 📦 Instalación

### Prerrequisitos
- Go 1.21 o superior
- Sistema Ubuntu (para las funciones de mantenimiento)
- Permisos de administrador (sudo) para comandos del sistema

### Compilar desde el código fuente

```bash
# Clonar el repositorio
git clone https://github.com/hartlink/gup.git
cd gup

# Descargar dependencias
go mod tidy

# Compilar el binario
make build

# O compilar directamente con Go
go build -o gup main.go
```

### Desplegar en el Servidor

Después de compilar, puedes desplegar el binario en tu servidor Ubuntu:

```bash
# Copiar el binario a tu servidor
scp build/gup usuario@tu-servidor:/tmp/

# En el servidor, moverlo a una ruta del sistema
ssh usuario@tu-servidor
sudo mv /tmp/gup /usr/local/bin/
sudo chmod +x /usr/local/bin/gup

# Verificar la instalación
gup version
```

Alternativamente, compila directamente en el servidor:

```bash
# En el servidor
git clone https://github.com/hartlink/gup.git
cd gup
make build
sudo cp build/gup /usr/local/bin/
```

## 🛠️ Uso

### Comandos Disponibles

#### `gup update`
Actualiza la lista de paquetes del sistema ejecutando `apt update`.

```bash
gup update
```

#### `gup upgrade`
Actualiza todos los paquetes instalados a sus últimas versiones. Ejecuta `apt update` automáticamente primero.

```bash
gup upgrade
```

#### `gup install`
Instala uno o más paquetes. Ejecuta `apt update` automáticamente primero.

```bash
gup install <paquete1> [paquete2] [...]

# Ejemplo
gup install nginx postgresql
```

> **Nota**: Todos los comandos usarán `sudo` automáticamente si es necesario.

**Opciones:**
- `-v, --verbose`: Muestra información detallada durante la ejecución
- `-l, --lang`: Selecciona el idioma (es/en)

### Idiomas

GUP detecta automáticamente el idioma de tu sistema, pero puedes cambiarlo manualmente:

```bash
# Español
gup --lang es

# English  
gup --lang en

# Aplicable a cualquier comando
gup version --lang es
gup update --lang en
```

Para más información sobre el sistema de traducción, consulta [docs/i18n.md](docs/i18n.md).

### Configuración Permanente

Puedes establecer tu idioma preferido de forma permanente:

```bash
# Crear archivo de configuración
mkdir -p ~/.gup
echo '{"language":"es"}' > ~/.gup/config.json
```

Ver más opciones en [docs/config.md](docs/config.md).

### Ejemplos

```bash
# Actualizar la lista de paquetes
gup update

# Actualizar todos los paquetes
gup upgrade

# Instalar paquetes específicos
gup install nginx
gup install postgresql redis-server

# Actualizar con salida detallada
gup update --verbose

# Ver ayuda
gup --help
gup update --help
```

## 🔧 Desarrollo

### Estructura del Proyecto

```
cli_go/
├── main.go              # Punto de entrada principal
├── cmd/                 # Comandos de Cobra
│   ├── root.go          # Comando raíz
│   ├── update.go        # Comando update
│   ├── upgrade.go       # Comando upgrade
│   ├── install.go       # Comando install
│   ├── demo.go          # Comando demo
│   └── version.go       # Comando version
├── internal/            # Código interno de la aplicación
│   ├── ui.go           # Interfaz de Bubble Tea
│   ├── apt/            # Gestión de paquetes APT
│   │   └── apt.go      # Lógica Update/Upgrade/Install
│   ├── i18n/           # Sistema de internacionalización
│   │   └── i18n.go     # Traducciones ES/EN
│   └── config/         # Sistema de configuración
│       └── config.go   # Gestión de configuración
├── docs/               # Documentación
│   ├── i18n.md        # Guía de internacionalización
│   └── config.md      # Guía de configuración
├── build/              # Binario compilado
│   └── gup            # Ejecutable
├── go.mod              # Módulo de Go
├── go.sum              # Checksums de dependencias
├── Makefile            # Comandos de construcción
├── .gitignore         # Archivos ignorados
├── README.md           # Versión en inglés
└── README_ES.md        # Este archivo
```

### Comandos de Desarrollo

```bash
# Compilar
make build

# Limpiar binarios
make clean

# Ejecutar en modo desarrollo
make dev

# Instalar dependencias
make deps
```

### Agregar Nuevos Comandos

1. Crear un nuevo archivo en `cmd/` (ej: `cmd/upgrade.go`)
2. Implementar el comando usando Cobra
3. Usar la interfaz de Bubble Tea desde `internal/ui.go`
4. Registrar el comando en `cmd/root.go`

## 📝 Roadmap

- [x] `gup upgrade` - Actualizar paquetes del sistema
- [x] `gup install` - Instalar paquetes
- [ ] `gup cleanup` - Limpiar paquetes innecesarios
- [ ] `gup status` - Mostrar estado del sistema
- [ ] `gup logs` - Ver logs del sistema
- [ ] `gup services` - Gestionar servicios systemd
- [ ] Configuración personalizable
- [ ] Logging avanzado
- [ ] Tests unitarios

## 🌍 Soporte de Idiomas

Este proyecto está disponible en:
- 🇺🇸 [English](README.md)
- 🇲🇽 [Español](README_ES.md) (este archivo)

## 🤝 Contribuciones

Las contribuciones son bienvenidas. Por favor:

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/nueva-funcionalidad`)
3. Commit tus cambios (`git commit -am 'Agregar nueva funcionalidad'`)
4. Push a la rama (`git push origin feature/nueva-funcionalidad`)
5. Abre un Pull Request

## 📄 Licencia

Este proyecto está licenciado bajo GPLv3. Ver el archivo `LICENSE` para más detalles.

## ⚠️ Advertencias

- **Permisos de Root**: Muchos comandos requieren permisos de administrador
- **Compatibilidad**: Diseñado específicamente para sistemas Ubuntu/Debian
- **Uso Responsable**: Siempre revisa los comandos antes de ejecutarlos en producción