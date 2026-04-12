# Collapse
### A TUI-MUD served over SSH

---

## The Premise

Y2K happened. The cascade failure didn't just crash computers — it crashed civilization. In the power vacuum that followed, warlords and rogue generals fought over the remnants. The wars damaged nuclear plants across the midwest. Forty years of slow radiation leakage into the watershed, the soil, the deep places has done its quiet work.

It is 2039. You are in what used to be southern Ohio and northern Kentucky.

---

## Design Philosophy

**Human-centric.** No fantasy races. No aliens. Mutations exist but they are scars the world left on people, not a character archetype. The drama is human-scaled.

**Heavy sim, light weird.** The core game is survival, scavenging, trade, and faction conflict. The weird stuff is earned by geography — it lives in the hills, the deep mines, the river caves near old plant sites. Players who never go there never encounter it.

**Cramped by design.** One city to start. A small number of zones. Every room matters. No filler content. No skill bloat. A skill you take changes how you play.

**The interface surfaces options.** Players should never need to memorize a command list. Context-sensitive input, keybindings, visible actions. The interface teaches by showing what is possible in each moment.

**SSH-first.** The game is a purpose-built TUI delivered over SSH. Not a telnet MUD with a nicer client. Not a web app pretending to be a terminal. The client and server are the same codebase. The player's terminal is the game.

---

## Setting

**Geography:** The triangle of Cincinnati, Lexington, and Louisville as faction city-states. The Appalachian foothills and river systems as natural constraints and trade corridors. Later expansion into surrounding states.

**Three content tiers:**

1. **The human world** — settlements, scavenging, trade, faction politics. Open ground. Tactical ranged combat. This is where most of the game lives.

2. **The weird fringe** — areas downriver from old plants, deep Appalachian interior, abandoned industrial zones. Mutation begins here. Things are wrong but recognizable.

3. **The deep weird** — old plant sites, flooded mines, places nobody comes back from unchanged. Full strangeness. Endgame content.

**The three factions:** Cincinnati, Lexington, Louisville. Each shaped by who won their local wars forty years ago. Each wants to dominate the surrounding settlements. Players can align, play them against each other, or build independent.

---

## Technical Stack

| Layer | Technology |
|---|---|
| Language | Go |
| SSH delivery | Charm/Wish |
| TUI framework | Bubbletea + Lipgloss |
| Database | SQLite via modernc.org/sqlite |
| Config | Local TOML file (~/.config/collapse/config.toml) |

**Key architectural decisions:**

- Game logic is pure Go with no transport dependency. Tested independently.
- SQLite stores all game state. No flat files for world data.
- Player display preferences live in local config, not on the server character record.
- The map pane and text pane carry the same information. The map is an accelerator, not the only source of truth. Screen readers work.
- Backups are automated SQLite snapshots to network storage, rolling 24-hour window.

---

## World Structure

**Towns and cities** are service lists, not navigable spaces. `go smithy`, `go inn`, `go market`. Fast, honest, no mandatory traversal.

**Outdoor zones** have internal spatial geometry. A cell is 10 meters. Exits have distance and direction. `treeline 40m N`, `ravine mouth 25m E`. The roguelike map renders this spatially. Text output describes it for screen readers.

**Dungeons and buildings** use a coordinate grid. [x,y] displayed in the room header. Roguelike movement on arrow keys or hjkl. You know where you are.

---

## Combat

**Two modes, one engine:**

*Tactical (outdoor):* Rooms have internal geometry. Distance matters. Cover matters. Suppression mechanics. Fire and movement. A cell is 10m, movement and firing are tradeoffs per tick. Ranged dominant.

*Close quarters (dungeon/indoor):* The tactical map collapses to a single cell. Distance is zero. Combat resolves in the room. Melee and mutations dominant.

**Tick-based.** Combat is slow enough to read. Decisions matter. Input is keybinding-driven so you never break flow to type a command.

**BRP-adjacent mechanics.** Percentile skills. No classes. Competency-based. Combat is lethal. Preparation and positioning matter more than stats.

---

## Character

No fantasy races. Background and skills differentiate characters. A skill you take changes how you play, not just a number.

**Three combat expressions:**
- Melee
- Ranged
- Mutations (rare, weird zones only, like spells but grounded in weird science)

**Progression:** Active play for most advancement. Offline skill ticks for passive gains — log out at the right location, cap at one point per session. Slow but meaningful for dedicated players.

---

## Economy

**Supply chains are real.** Raw materials must be sourced, moved, and processed. Nothing appears from nowhere.

**Offline jobs:** Log out at a job location to produce basic tier goods while offline. Cap after X hours to prevent exploitation. Log back in to collect your cut of what sold.

**Player vendors:** Active crafting produces quality goods. Players set their own prices, build reputations.

**Caravans:** Move goods between nodes. Players can log out as caravan guards — flat rate pay, XP, death respawns at origin with no gear loss. Caravan success and failure has downstream economic consequences.

**Logout tiers:**
- Wilderness: disconnect, no penalty, no reward
- Town: rested bonus after X hours offline
- Inn: better rested bonus, passive skill tick toward one queued skill

---

## Accessibility

The map is never the only source of spatial information. Every spatial fact the map shows is also present as parseable text output. Distance, cover, contacts, exits — all in structured text. Screen readers work. Blind players can play.

---

## What This Is Not

- Not a telnet MUD with a nicer client
- Not a fantasy game with humans swapped in
- Not a feature-complete MMO at launch
- Not designed for old MUD players specifically
- Not trying to save the MUD genre

It is a terminal game that learned the right lessons from MUDs.

---

## Name

**Collapse**

The event. The setting. The tone. The game.

`ssh play.collapse.game`
