# Finn's Geburtstag
### Karte zu Finn's 18. Geburtstag

# Mit Docker laufen lassen
Das Programm soll als Docker-Container laufen

## Docker container bauen
```bash
$ docker build .
```
## Programm in einen Docker-Container verpacken
```bash
$ docker images # alle docker images anzeigen
$ docker save <container-id> -o FinnsGeburtstag18 # als tar-Archiv speichern
```

## Den Container ausführen
```bash
$ docker load < FinnsGeburtstag18 # in das docker system laden
$ docker images # Verifizieren
$ docker run <container-id> # laufen lassen
```
