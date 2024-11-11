# githooks

## Instalación

Para instalar `gohtpasswd`, puedes descargar el binario desde el siguiente enlace:

- [Descargar gohtpasswd](https://github.com/njavilas2015/gohtpasswd/releases/download/1.0.1/gohtpasswd)

```bash
curl -s https://api.github.com/repos/njavilas2015/gohtpasswd/releases/latest | jq -r .assets[0].browser_download_url | wget -i -
```


Después de descargar, asegúrate de que el binario sea ejecutable y mueve el archivo a un directorio en tu `PATH`:

```bash
chmod +x gohtpasswd
sudo mv gohtpasswd /usr/local/bin/
```

## Uso
Para usar `gohtpasswd`, puedes ejecutar el siguiente comando en la terminal:

```bash
gohtpasswd --username njavilas --password holaconfig --print
```

## Hooks del Lado del Cliente
Estos hooks se ejecutan en el repositorio local y permiten personalizar el flujo de trabajo antes o después de acciones como hacer un commit, crear un mensaje, hacer un push, entre otros.

- [] **applypatch-msg**: Se ejecuta después de aplicar un parche (git am), permitiendo editar o rechazar el mensaje del commit.
pre-applypatch: Se ejecuta antes de aplicar un parche, útil para ejecutar pruebas antes de aplicarlo.
- [ ] **post-applypatch**: Se ejecuta después de aplicar un parche exitosamente.
- [x] **pre-commit**: Se ejecuta antes de crear un commit, comúnmente usado para ejecutar pruebas o verificaciones de código.
- [] **prepare-commit-msg**: Se ejecuta antes de que se presente el editor para el mensaje de commit, permite modificar el mensaje predeterminado.

- [x] **commit-msg**: Se ejecuta después de que se escribe el mensaje de commit, permitiendo validar o modificar el mensaje.

- [x] **post-commit**: Se ejecuta después de hacer un commit, comúnmente usado para notificaciones o acciones posteriores al commit.

- [x] **pre-rebase**: Se ejecuta antes de realizar un rebase, permite abortar o personalizar el proceso.

- [x] **post-checkout**: Se ejecuta después de un git checkout, útil para cambiar configuraciones específicas de la rama.

- [x] **post-merge**: Se ejecuta después de un merge, se usa para restaurar archivos o ejecutar comandos posteriores a la fusión.
- [x] **pre-push**: Se ejecuta antes de un push, permite realizar verificaciones o cancelar el push si no se cumplen ciertas condiciones.
- [] **pre-auto-gc**: Se ejecuta antes de que Git realice una recolección de basura automática, útil para limpiar o preparar datos.
- [x] **post-rewrite**: Se ejecuta después de comandos que reescriben el historial (como git commit --amend y git rebase), permite realizar acciones en los commits reescritos.

## Contacto
Si tienes alguna pregunta o necesitas soporte, no dudes en contactarme:

Nombre: Javier Avila
Email: [njavilas2015@gmail.com]
GitHub: njavilas2015

## Apóyame con un café ☕️

Si te gusta mi trabajo y quieres apoyarme, ¡puedes invitarme a un café! 😊

[![Buy Me a Coffee](https://img.buymeacoffee.com/button-api/?text=Buy%20Me%20a%20Coffee&emoji=coffee&slug=tu_nombre&button_colour=FF5F5F&font_colour=ffffff&font_family=Cookie)](https://buymeacoffee.com/njavilas
)
