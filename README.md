# Pokedex CLI 🔴
 
A lightweight, interactive command-line Pokedex built in Go. This tool allows you to traverse the Pokemon world, discover local inhabitants, and manage your data—all without leaving your terminal.
 
This project was developed as part of the [Boot.dev](http://Boot.dev) curriculum to master Golang, network programming, and efficient memory management. While the core foundation follows the course, I have extended the application by architecting and implementing several custom features from scratch.
 
## 🚀 Features
 
- **Interactive REPL:** A responsive "Read-Eval-Print Loop" for a seamless user experience.
- **PokéAPI Integration:** Real-time data fetching from the PokéAPI.
- **Custom Local Caching:** A high-performance, thread-safe cache built from scratch to minimize network latency and API load.
- **World Exploration:** Paginated navigation through different regions and location areas.
## 🛠 Installation
 
Make sure you have Go installed (version 1.18 or higher recommended).
 
1. Clone the repository:
```bash
git clone https://github.com/yourusername/pokedexcli.git
cd pokedexcli
```
 
2. Build the executable:
```bash
go build -o pokedex
```
 
3. Run the program:
```bash
./pokedex
```
 
## 🎮 Usage
 
Once the REPL is running, you can use the following commands:
 
| Command | Description |
| --- | --- |
| `help` | Displays a help message with all available commands. |
| `map` | Displays the next 20 location areas in the Pokemon world. |
| `mapb` | Displays the previous 20 location areas. |
| `pokedex` | Display the pokemons you have in your pokedex. |
| `explore` | Displays pokemons in an specific area |
| `catch` | Try to catch a pokemon |
| `inspect` | Display the info of an already caught pokemon |
| `pokedex` | Display the pokemons you have in your pokedex |
| `battle` | Battle and get a new pokemon. |
| `run` | Just run and continue with your easy life. |
| `exit` | Gracefully exits the Pokedex CLI. |

## ⚙️ Technical Highlights
 
This project focuses on back-end fundamentals and Go-specific patterns:
 
- **Concurrency:** Utilizes Go Routines for background cache reaping (automatic cleanup of expired entries).
- **Thread Safety:** Uses `sync.Mutex` to ensure the cache remains safe during concurrent read/write operations.
- **Data Handling:** Efficient JSON unmarshaling into custom Go structs.
- **Extensibility:** Architecture designed to easily add new commands or improve on the ones already built for more fun and a better user experience.

## 🔮 Future Improvements & Roadmap

To evolve this CLI I would like to implement the following features:

* **Enhanced Battle Experience**: Transition from basic stat-checking to a turn-based combat system with move selection and status effects.
* **Persistent Progression**:
    * **Leveling System**: Gain experience points (XP) from battles to level up your Pokemon.
    * **Data Persistence**: Save user progress, Pokedex collection, and Pokemon levels to a local file (JSON or SQLite) so progress isn't lost on `exit`.
* **Advanced Capture Mechanics**:
    * **Strategic Weakening**: Implement the ability to battle a Pokemon first to lower its HP, making it significantly easier to catch.
    * **Specialized Pokéballs**: Introduction of Great Balls, Ultra Balls, and Master Balls with varying catch-rate multipliers.
* **Dynamic Exploration**:
    * **Traveling System**: A "travel" mechanic where moving between regions takes time or triggers random encounters based on the terrain.
    * **Biomes**: Finding specific Pokemon types based on the "environment" the user is currently exploring.
* **Resource Management (Stamina/Energy)**:
    * An energy system that limits how many actions (battling or catching) a user can take before needing to "rest" or use an item.