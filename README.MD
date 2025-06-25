# 🛠️ USB Boot Tool del Adeptus Mechanicus

**Invoca el Espíritu Máquina y convierte tu USB en un portador sagrado de sistemas operativos.**  
Clona imágenes ISO de Linux o Windows a memorias USB desde la terminal, con estilo y ritual.

## 🚀 Características

- 📦 Sin dependencias externas (solo Go estándar)
- 🐧 Soporta imágenes ISO de Linux
- 🪟 Detecta imágenes de Windows (modo experimental)
- 🧠 Detección de discos removibles
- 🔥 Proceso de transferencia con animaciones y confirmaciones

## ⚙️ Requisitos

- Go 1.18 o superior
- Linux (soporte multiplataforma en desarrollo)
- Permisos de `sudo` para escribir en dispositivos USB

## 🧪 Instalación

```bash
git clone https://github.com/tuusuario/usb-boot-tool.git
cd usb-boot-tool
go build -o usbtool
```

## 🧙 Uso

```bash
sudo ./usbtool
```

## 📜 Instrucciones

Te guiará para seleccionar una ISO (~/Documentos/ISO)

Luego detectará dispositivos USB disponibles

Confirmarás antes de iniciar la purificación

Y el Espíritu Máquina hará el resto...

## 🛑 Advertencia

⚠️ Este proceso elimina todo el contenido del USB seleccionado. Asegúrate de elegir el dispositivo correcto.

## 🧩 Estructura del proyecto

```
.
├── main.go
├── utils/
│   ├── io.go
│   ├── prompts.go
│   └── scan.go
└── detection/
    └── iso.go
```

## 🧿 Licencia

Este proyecto es un tributo al Omnissiah.
Libre para uso personal, educativo o ritualístico.

```
"Bendito sea el bit, maldito sea el bug."
```
