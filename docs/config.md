# Configuración de CliGO

CliGO puede ser configurado mediante un archivo JSON ubicado en `~/.cligo/config.json`.

## Crear archivo de configuración

```bash
# Crear el directorio
mkdir -p ~/.cligo

# Crear el archivo de configuración
cat > ~/.cligo/config.json << EOF
{
  "language": "es"
}
EOF
```

## Opciones disponibles

### language

Establece el idioma predeterminado de la aplicación.

**Valores válidos:**
- `"es"` - Español
- `"en"` - English

**Ejemplo:**

```json
{
  "language": "es"
}
```

## Prioridad de configuración de idioma

CliGO determina el idioma a usar en el siguiente orden de prioridad:

1. **Flag `--lang`**: Si se especifica el flag al ejecutar el comando
   ```bash
   cligo --lang en
   ```

2. **Archivo de configuración**: Si existe `~/.cligo/config.json` con la opción `language`

3. **Variable de entorno**: Lee las variables `LANG` o `LANGUAGE` del sistema

4. **Idioma por defecto**: Español (es)

## Ejemplos de uso

### Configurar español como predeterminado

```bash
mkdir -p ~/.cligo
echo '{"language":"es"}' > ~/.cligo/config.json
```

### Configurar inglés como predeterminado

```bash
mkdir -p ~/.cligo
echo '{"language":"en"}' > ~/.cligo/config.json
```

### Verificar configuración actual

```bash
# Ver el archivo de configuración
cat ~/.cligo/config.json

# Ver qué idioma se está usando
cligo version
```

### Cambiar temporalmente el idioma

Aunque tengas un idioma configurado, puedes cambiarlo temporalmente con el flag:

```bash
# Configurado en español, pero usar inglés una vez
cligo --lang en

# Configurado en inglés, pero usar español una vez
cligo --lang es
```

## Ubicación del archivo

- **Linux/macOS**: `~/.cligo/config.json`
- **Windows**: `%USERPROFILE%\.cligo\config.json`

## Resetear configuración

Para volver a la configuración por defecto, simplemente elimina el archivo:

```bash
rm ~/.cligo/config.json
```

O elimina todo el directorio de configuración:

```bash
rm -rf ~/.cligo
```