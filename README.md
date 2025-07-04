# TaskWarrior Notes CLI

TaskWarrior Notes CLI (`twn`) to narzędzie wiersza poleceń
do zarządzania notatkami powiązanymi z zadaniami TaskWarrior.
Pozwala synchronizować metadane zadań z notatkami,
pobierać ścieżki do notatek
oraz wykonywać inne operacje związane z notatkami zadań.

## Instalacja

Wymagania:

- Go 1.24+

Aby zbudować i zainstalować narzędzie:

```sh
make install
```

Lub ręcznie:

```sh
go build -o twn .
cp twn $(go env GOPATH)/bin/
```

## Budowanie

Aby zbudować binarkę:

```sh
make build
```

Aby uruchomić bezpośrednio:

```sh
make run
```

## Licencja

MIT
