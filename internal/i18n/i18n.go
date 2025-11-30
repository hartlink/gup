package i18n

import (
	"gup/internal/config"
	"os"
	"strings"
)

// Language representa el idioma actual
type Language string

const (
	Spanish Language = "es"
	English Language = "en"
)

var currentLanguage Language = Spanish

// Translations contiene todas las traducciones
var translations = map[Language]map[string]string{
	Spanish: {
		// Mensajes generales
		"app.name":        "GUP",
		"app.description": "Herramienta CLI para automatizar el mantenimiento de tu servidor Ubuntu",
		"app.welcome":     "¡Bienvenido a GUP!",

		// Comandos
		"cmd.help":         "Usa 'gup --help' para ver todos los comandos disponibles.",
		"cmd.available":    "Comandos disponibles:",
		"cmd.update.desc":  "Actualizar lista de paquetes",
		"cmd.demo.desc":    "Ver demostración de la interfaz",
		"cmd.version.desc": "Ver versión del programa",
		"cmd.permissions":  "La mayoría de comandos usarán sudo automáticamente si es necesario",

		// Update command
		"update.short":           "Actualiza la lista de paquetes del sistema",
		"update.long":            "Ejecuta 'apt update' para actualizar la lista de paquetes disponibles\ndesde los repositorios configurados en el sistema Ubuntu.\nEl comando usará sudo automáticamente si es necesario.",
		"update.no_perms":        "Este comando requiere permisos de administrador.",
		"update.use_sudo":        "Por favor, ejecuta el comando con 'sudo gup update'",
		"update.description":     "Actualizando lista de paquetes",
		"update.checking_sudo":   "Verificando permisos de administrador...",
		"update.sudo_required":   "Se requieren permisos de administrador",
		"update.partial_success": "Lista de paquetes actualizada con algunas advertencias (algunos repositorios pueden tener errores)",

		// Upgrade command
		"upgrade.short":          "Actualiza los paquetes instalados",
		"upgrade.long":           "Ejecuta 'apt upgrade' para instalar las versiones más recientes de los paquetes instalados.",
		"upgrade.description":    "Actualizando paquetes del sistema",
		"upgrade.success":        "Paquetes actualizados correctamente",
		"upgrade.prompt_restart": "¿Deseas reiniciar el servidor ahora? (s/n):",

		// Install command
		"install.short":            "Instala nuevos paquetes",
		"install.long":             "Ejecuta 'apt install' para instalar uno o más paquetes.",
		"install.description":      "Instalando paquetes",
		"install.success":          "Paquetes instalados correctamente",
		"install.preparing":        "Se instalarán los siguientes paquetes: %v",
		"install.updating_first":   "Primero se actualizará la lista de paquetes...",
		"install.error.not_found":  "No se pudo encontrar el paquete '%s'",
		"install.error.permission": "Permiso denegado. Ejecuta el comando con sudo",
		"install.error.network":    "Error de red al intentar instalar paquetes",
		"install.error.dependency": "Error de dependencias al instalar paquetes",
		"install.error.unknown":    "Error desconocido al instalar paquetes",

		// Demo command
		"demo.short":       "Ejecuta una demostración de la interfaz de GUP",
		"demo.long":        "Ejecuta una demostración que muestra cómo se ve la interfaz de Bubble Tea\nsin requerir permisos de administrador. Útil para probar el funcionamiento.",
		"demo.description": "Ejecutando demostración de actualización de paquetes",
		"demo.output":      "Simulando apt update...",
		"demo.success":     "Paquetes actualizados correctamente",

		// Version command
		"version.short":     "Muestra la versión de GUP",
		"version.long":      "Muestra la versión actual de GUP y información de compilación.",
		"version.tool":      "Herramienta CLI para mantenimiento de servidor Ubuntu",
		"version.developed": "Desarrollado con ❤️  usando Go, Cobra y Bubble Tea",

		// Restart command
		"restart.short":     "Reinicia el servidor",
		"restart.long":      "Reinicia el servidor ejecutando 'reboot now'.",
		"restart.warning":   "⚠️  ADVERTENCIA: Esta acción reiniciará el servidor inmediatamente",
		"restart.confirm":   "¿Estás seguro que deseas reiniciar el servidor? (s/n):",
		"restart.cancelled": "Reinicio cancelado",
		"restart.executing": "Reiniciando servidor...",

		// Root command
		"root.short": "CLI para automatizar tareas de mantenimiento de servidor Ubuntu",
		"root.long":  "GUP (Go Update) es una herramienta de línea de comandos diseñada para\nautomatizar las tareas comunes de mantenimiento de servidores Ubuntu.\nUtiliza una interfaz interactiva con Bubble Tea para una mejor experiencia de usuario.",

		// Flags
		"flag.verbose": "Salida detallada",
		"flag.lang":    "Idioma de la interfaz (es/en)",

		// UI Messages
		"ui.title":     "GUP - Herramienta de Mantenimiento Ubuntu",
		"ui.preparing": "Preparando",
		"ui.executing": "Ejecutando",
		"ui.command":   "Comando",
		"ui.success":   "Comando ejecutado exitosamente!",
		"ui.error":     "Error al ejecutar el comando",
		"ui.output":    "Salida:",
		"ui.continue":  "Presiona Enter para continuar o 'q' para salir",
		"ui.no_output": "(sin salida)",
		"ui.truncated": "... (salida truncada)",

		// Errors
		"error.interface": "Error ejecutando la interfaz",
		"error.executing": "Ejecutando comando",
	},
	English: {
		// General messages
		"app.name":        "GUP",
		"app.description": "CLI tool to automate Ubuntu server maintenance",
		"app.welcome":     "Welcome to GUP!",

		// Commands
		"cmd.help":         "Use 'gup --help' to see all available commands.",
		"cmd.available":    "Available commands:",
		"cmd.update.desc":  "Update package list",
		"cmd.demo.desc":    "View interface demo",
		"cmd.version.desc": "View program version",
		"cmd.permissions":  "Most commands will use sudo automatically if needed",

		// Update command
		"update.short":           "Updates the system package list",
		"update.long":            "Runs 'apt update' to update the list of available packages\nfrom the repositories configured on the Ubuntu system.\nThe command will use sudo automatically if needed.",
		"update.no_perms":        "This command requires administrator permissions.",
		"update.use_sudo":        "Please run the command with 'sudo gup update'",
		"update.description":     "Updating package list",
		"update.checking_sudo":   "Checking administrator permissions...",
		"update.sudo_required":   "Administrator permissions required",
		"update.partial_success": "Package list updated with some warnings (some repositories may have errors)",

		// Upgrade command
		"upgrade.short":          "Upgrades installed packages",
		"upgrade.long":           "Runs 'apt upgrade' to install the newest versions of all installed packages.",
		"upgrade.description":    "Upgrading system packages",
		"upgrade.success":        "Packages upgraded successfully",
		"upgrade.prompt_restart": "Do you want to restart the server now? (y/n):",

		// Install command
		"install.short":            "Installs new packages",
		"install.long":             "Runs 'apt install' to install one or more packages.",
		"install.description":      "Installing packages",
		"install.success":          "Packages installed successfully",
		"install.preparing":        "The following packages will be installed: %v",
		"install.updating_first":   "Package list will be updated first...",
		"install.error.not_found":  "Unable to locate package '%s'",
		"install.error.permission": "Permission denied. Run the command with sudo",
		"install.error.network":    "Network error while trying to install packages",
		"install.error.dependency": "Dependency error while installing packages",
		"install.error.unknown":    "Unknown error while installing packages",

		// Demo command
		"demo.short":       "Run a demonstration of GUP's interface",
		"demo.long":        "Runs a demonstration showing how the Bubble Tea interface looks\nwithout requiring root permissions. Useful for testing functionality.",
		"demo.description": "Running package update demonstration",
		"demo.output":      "Simulating apt update...",
		"demo.success":     "Packages updated successfully",

		// Version command
		"version.short":     "Shows GUP version",
		"version.long":      "Shows the current GUP version and build information.",
		"version.tool":      "CLI tool for Ubuntu server maintenance",
		"version.developed": "Developed with ❤️  using Go, Cobra and Bubble Tea",

		// Restart command
		"restart.short":     "Restart the server",
		"restart.long":      "Restarts the server by running 'reboot now'.",
		"restart.warning":   "⚠️  WARNING: This action will restart the server immediately",
		"restart.confirm":   "Are you sure you want to restart the server? (y/n):",
		"restart.cancelled": "Restart cancelled",
		"restart.executing": "Restarting server...",

		// Root command
		"root.short": "CLI to automate Ubuntu server maintenance tasks",
		"root.long":  "GUP (Go Update) is a command-line tool designed to\nautomate common Ubuntu server maintenance tasks.\nIt uses an interactive Bubble Tea interface for a better user experience.",

		// Flags
		"flag.verbose": "Verbose output",
		"flag.lang":    "Interface language (es/en)",

		// UI Messages
		"ui.title":     "GUP - Ubuntu Maintenance Tool",
		"ui.preparing": "Preparing",
		"ui.executing": "Executing",
		"ui.command":   "Command",
		"ui.success":   "Command executed successfully!",
		"ui.error":     "Error executing command",
		"ui.output":    "Output:",
		"ui.continue":  "Press Enter to continue or 'q' to quit",
		"ui.no_output": "(no output)",
		"ui.truncated": "... (output truncated)",

		// Errors
		"error.interface": "Error running interface",
		"error.executing": "Executing command",
	},
}

// SetLanguage establece el idioma actual
func SetLanguage(lang string) {
	lang = strings.ToLower(strings.TrimSpace(lang))
	switch lang {
	case "en", "english":
		currentLanguage = English
	case "es", "spanish", "español":
		currentLanguage = Spanish
	default:
		// Intentar detectar del sistema
		currentLanguage = detectSystemLanguage()
	}
}

// GetLanguage retorna el idioma actual
func GetLanguage() Language {
	return currentLanguage
}

// T (Translate) obtiene la traducción para una clave
func T(key string) string {
	if trans, ok := translations[currentLanguage][key]; ok {
		return trans
	}
	// Fallback a español si no existe la clave
	if trans, ok := translations[Spanish][key]; ok {
		return trans
	}
	return key
}

// detectSystemLanguage detecta el idioma del sistema
func detectSystemLanguage() Language {
	lang := os.Getenv("LANG")
	if lang == "" {
		lang = os.Getenv("LANGUAGE")
	}

	lang = strings.ToLower(lang)
	if strings.Contains(lang, "en") {
		return English
	}

	return Spanish
}

// Init inicializa el sistema de traducción
func Init() {
	// Primero intentar cargar desde configuración
	cfg, err := config.Load()
	if err == nil && cfg.Language != "" {
		SetLanguage(cfg.Language)
		return
	}

	// Si no hay configuración, detectar idioma del sistema automáticamente
	currentLanguage = detectSystemLanguage()
}
