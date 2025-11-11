"# Tic-Tac-Toe 🎮

Ein elegantes, funktionsreiches Tic-Tac-Toe-Spiel in Go mit mehrsprachiger Unterstützung, verschiedenen Schwierigkeitsgraden und anpassbaren Anzeigemodi.

## Features ✨

### 🎯 Spielmodi
- **Einzelspieler**: Spiele gegen den Computer mit 3 Schwierigkeitsgraden
- **Mehrspieler**: Spiele gegen einen Freund am selben Computer

### 🤖 KI-Schwierigkeitsgrade
- **Easy**: Zufällige Züge - perfekt für Anfänger
- **Medium**: Intelligente Strategie - blockt Gewinnzüge und versucht zu gewinnen
- **Hard**: Unschlagbar - verwendet den Minimax-Algorithmus für perfektes Spiel

### 🌍 Mehrsprachig
- Deutsch 🇩🇪
- Englisch 🇬🇧

### 🎨 Anzeigemodi
- **Normal Printer**: Klassische, einfache Textdarstellung
- **Fancy Printer**: Stilvolle ASCII-Art-Darstellung (powered by [go-figure](https://github.com/common-nighthawk/go-figure))

## Installation 📦

### Voraussetzungen
- Go 1.25.0 oder höher

### Installation
```bash
# Repository klonen
git clone https://github.com/WilliamAkaWill/tic-tac-toe.git
cd tic-tac-toe

# Abhängigkeiten installieren
make install
# oder
go mod tidy
```

## Verwendung 🚀

### Spiel starten
```bash
# Mit Make
make run

# Oder direkt mit Go
go run .
```

### Spiel kompilieren
```bash
# Erstellt ausführbare Datei 'tictactoe'
make build

# Anschließend ausführen
./tictactoe
```

## Spielanleitung 🎲

1. **Sprache wählen**: Wähle zwischen Deutsch (0) oder Englisch (1)
2. **Spielmodus wählen**: Einzelspieler (1) oder Mehrspieler (2)
3. **Schwierigkeitsgrad wählen** (nur Einzelspieler): Easy (0), Medium (1) oder Hard (2)
4. **Anzeigemodus wählen**: Normal (0) oder Fancy (1)
5. **Spielen**: Gib die Position ein (1-9), um dein Zeichen zu platzieren

### Spielfeld-Positionen
```
 1 | 2 | 3
-----------
 4 | 5 | 6
-----------
 7 | 8 | 9
```

## Projektstruktur 📁

```
tic-tac-toe/
├── errors/          # Fehlerbehandlung
│   └── errors.go
├── game/            # Spiellogik
│   ├── game.go
│   └── helper.go
├── helper/          # Hilfsfunktionen
│   ├── helper.go
│   └── helper_test.go
├── language/        # Mehrsprachige Unterstützung
│   ├── english.go
│   ├── german.go
│   └── helper.go
├── minimax/         # KI-Algorithmus
│   ├── minimax.go
│   ├── minimax_test.go
│   └── data_sets.go
├── player/          # Spieler-Implementierungen
│   ├── computer.go
│   └── human.go
├── print/           # Anzeigemodi
│   ├── fancy_print.go
│   ├── normal_print.go
│   └── printer.go
├── shared/          # Gemeinsame Datentypen
│   ├── convert.go
│   ├── convert_test.go
│   ├── game.go
│   └── types.go
├── validate/        # Validierung und Spielzustandsprüfung
│   ├── validate.go
│   ├── validate_test.go
│   └── data_sets.go
├── main.go          # Einstiegspunkt
├── Makefile         # Build-Skripte
└── README.md
```

## Tests 🧪

Das Projekt enthält umfassende Unit-Tests für alle Kernfunktionen.

```bash
# Alle Tests ausführen
make test

# Oder mit Go
go test ./...

# Tests mit Coverage
go test -cover ./...

# Detaillierte Coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Testabdeckung
- Minimax-Algorithmus (45+ Tests)
- Validierung und Spielzustand (54+ Tests)
- Hilfsfunktionen (30+ Tests)
- Konvertierungsfunktionen (52+ Tests)

## Technische Details 🔧

### Minimax-Algorithmus
Der Hard-Modus verwendet den klassischen Minimax-Algorithmus:
- Bewertet alle möglichen Spielzüge rekursiv
- Geht von optimalem Gegner-Spiel aus
- Bevorzugt schnellere Siege und spätere Niederlagen durch Tiefengewichtung
- Basiert auf: [Wikipedia - Minimax](https://en.wikipedia.org/wiki/Minimax)

### Spielzustands-Validierung
- Prüft alle 8 Gewinnkombinationen (3 Zeilen, 3 Spalten, 2 Diagonalen)
- Erkennt Unentschieden-Situationen
- Validiert Eingaben und verhindert ungültige Züge

## Make-Befehle 🛠️

| Befehl | Beschreibung |
|--------|--------------|
| `make run` | Startet das Spiel |
| `make build` | Kompiliert das Programm |
| `make test` | Führt alle Tests aus |
| `make install` | Installiert Abhängigkeiten |
| `make all` | Führt alle obigen Befehle aus |

## Abhängigkeiten 📚

- [go-figure](https://github.com/common-nighthawk/go-figure) - ASCII Art Rendering für Fancy Printer

## Lizenz 📄

Dieses Projekt ist Open Source.

## Autor ✍️

**WilliamAkaWill**
- GitHub: [@WilliamAkaWill](https://github.com/WilliamAkaWill)

## Danksagungen 🙏

- Vielen Dank an [common-nighthawk](https://github.com/common-nighthawk) für die [go-figure](https://github.com/common-nighthawk/go-figure) Library

## Screenshots 📸

### Normal Printer
```
 X | O | X
-----------
 O | X | 
-----------
   |   | O
```

### Fancy Printer
```
   __  __    _     ____      _      ___
   \ \/ /   | |   |___ \    | |    / _ \
    \  /    | |     __) |   | |   | | | |
    /  \    | |    / __/    | |   | |_| |
   /_/\_\   | |   |_____|   | |    \___/
            |_|             |_|
--------------------------------------------
    _  _       _    __  __    _      __
   | || |     | |   \ \/ /   | |    / /_
   | || |_    | |    \  /    | |   | '_ \
   |__   _|   | |    /  \    | |   | (_) |
      |_|     | |   /_/\_\   | |    \___/
              |_|            |_|
--------------------------------------------
     ___      _      ___      _      ___
    / _ \    | |    ( _ )    | |    / _ \
   | | | |   | |    / _ \    | |   | (_) |
   | |_| |   | |   | (_) |   | |    \__, |
    \___/    | |    \___/    | |      /_/
```

---

Viel Spaß beim Spielen! 🎉
" 
