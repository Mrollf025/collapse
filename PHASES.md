# Collapse — Development Phases
### Kanban Task Breakdown

Each phase has a clear done condition. Do not start the next phase until the done condition is met. Resist adding features from later phases into earlier ones.

---

## PHASE 0 — Walking Skeleton
**Done when:** You can SSH in, walk between rooms, and see it render in a Bubbletea layout. Nothing else.

### Tasks

- [ ] Initialize Go module, set up project structure
- [ ] Add Wish dependency, wire up basic SSH server
- [ ] Create hardcoded room graph (5-10 rooms, no database)
- [ ] Create minimal Character struct (name only)
- [ ] Implement movement between rooms via exit commands
- [ ] Basic Bubbletea layout: text pane + input bar
- [ ] Game loop receives input, updates state, re-renders
- [ ] Multiple SSH connections handled independently
- [ ] Verify two players can be in the same room simultaneously

---

## PHASE 1 — Persistence
**Done when:** Your character and the world survive a server restart.

### Tasks

- [ ] Add modernc.org/sqlite dependency
- [ ] Design and create initial schema (characters, rooms, exits)
- [ ] Implement character save and load
- [ ] Implement world load from SQLite on startup
- [ ] Basic character creation flow on first login
- [ ] Session management — reconnect restores your character
- [ ] Automated SQLite backup script (rolling 24hr window to network drive)
- [ ] Seed script for initial room graph

---

## PHASE 2 — World Shape
**Done when:** The city service list works, one outdoor zone has spatial geometry, and one dungeon has a coordinate grid.

### Tasks

- [ ] Implement town-as-list navigation (`go smithy`, `go inn`)
- [ ] Room types: town node, outdoor zone, dungeon cell
- [ ] Outdoor zone: internal cell geometry (10m per cell)
- [ ] Outdoor zone: exits with distance and direction
- [ ] Dungeon: [x,y] coordinate display in room header
- [ ] hjkl / arrow key movement in dungeon and outdoor zones
- [ ] Basic map pane rendering (ASCII, outdoor and dungeon modes)
- [ ] Text pane describes spatial info in parseable text (screen reader safe)
- [ ] Ping action returns structured environment readout
- [ ] Room header shows room name and coordinates where applicable
- [ ] Seed one city district (service list), one outdoor zone, one 3-room dungeon

---

## PHASE 3 — Character and Skills
**Done when:** Characters have stats, a background, and a small skill set that meaningfully differentiates them.

### Tasks

- [ ] Character struct expanded: stats, background, skills
- [ ] BRP-adjacent stat system (percentile skills)
- [ ] Small fixed skill list (melee, ranged, medicine, mechanics, scavenging, stealth)
- [ ] Background archetypes that set starting skill spread
- [ ] Character creation flow: name, background, review
- [ ] Skill checks: pass/fail with degree of success
- [ ] Display character sheet in a dedicated pane or command
- [ ] Local config file for display preferences (~/.config/collapse/config.toml)

---

## PHASE 4 — Combat
**Done when:** You can find an enemy, engage them, and resolve the fight with meaningful tactical decisions.

### Tasks

- [ ] NPC/enemy struct and basic AI
- [ ] Contact detection: enemies visible at range in outdoor zones
- [ ] Tactical combat loop: tick-based, outdoor mode
- [ ] Ranged attack: roll against skill, check range and cover
- [ ] Suppression mechanic: pinned state affects enemy actions
- [ ] Cover system: terrain objects grant defense bonus
- [ ] Movement in combat: costs action, can't fire same tick
- [ ] Close quarters combat loop: dungeon/indoor mode
- [ ] Melee attack: roll against skill, adjacency required
- [ ] Health, wounds, death
- [ ] Loot on enemy death
- [ ] Flee mechanic
- [ ] Keybindings for common combat actions (no command typing mid-fight)
- [ ] Combat log readable at tick pace
- [ ] Death consequence: respawn at last safe location, gear retained

---

## PHASE 5 — Economy Foundation
**Done when:** A player can work an offline job, produce goods, and sell them at a shop.

### Tasks

- [ ] Item system: base items, quality tiers
- [ ] Shop system: NPC shops with dynamic inventory
- [ ] Offline job system: log out at job location, produce basic goods while offline
- [ ] Production cap: stops after X hours offline
- [ ] Profit cut: player earns share of what their specific goods sold for
- [ ] Raw material input: shops require materials to produce goods
- [ ] Player vendor: set prices, sell quality goods directly
- [ ] Logout location tracking: wilderness, town, inn, job site
- [ ] Rested bonus: inn and town logout rewards
- [ ] Passive skill tick: one queued skill gains one point after inn logout + X hours

---

## PHASE 6 — Factions and Settlements
**Done when:** The three city factions exist, settlements can be built, and faction pressure on settlements is felt.

### Tasks

- [ ] Faction system: Cincinnati, Lexington, Louisville
- [ ] Faction reputation per character
- [ ] Settlement system: players can claim and name a location
- [ ] Settlement services: build a smithy, a market, an inn
- [ ] Settlement production tied to offline economy
- [ ] Faction influence: cities exert pressure on nearby settlements
- [ ] Faction missions: quests that shift reputation and territorial control
- [ ] Player alignment: independent, allied, or vassal to a faction
- [ ] NPC patrols: faction soldiers in faction-controlled zones
- [ ] Caravan system: goods move between nodes on a schedule
- [ ] Caravan guard job: log out as guard, earn pay and XP, safe death

---

## PHASE 7 — The Weird Fringe
**Done when:** Weird zones exist, strange creatures inhabit them, and the mutation system has a first pass.

### Tasks

- [ ] Weird zone room type: different visual language, different text tone
- [ ] Cryptid creature roster: Appalachian folklore, wrong-but-recognizable
- [ ] Mutation system first pass: exposure mechanic, first mutation tier
- [ ] Mutations as combat abilities: ranged equivalent, weird science flavor
- [ ] Weird zone hazards: radiation exposure, environmental damage
- [ ] Deep weird zone: old plant site or flooded mine as endgame dungeon
- [ ] Lore fragments: found text, journals, radio transmissions
- [ ] Mutation visibility: other players can see your mutations

---

## BACKLOG — Future Phases

These are real ideas, not yet scoped into phases.

- Downloadable client with tileset support (community tile packs)
- Browser terminal delivery via xterm.js + WebSocket
- Sound hooks for downloadable client
- PvP consent system for settlement warfare
- Hex map of wider region, active battle hex generation
- Discord integration for out-of-game notifications
- Accessibility audit and screen reader testing pass
- Public alpha, player feedback loop

---

## Done Conditions Summary

| Phase | Done When |
|---|---|
| 0 | SSH in, walk between rooms, Bubbletea renders |
| 1 | Character and world survive restart |
| 2 | City list, outdoor geometry, dungeon grid all work |
| 3 | Characters have stats and meaningful skills |
| 4 | Full combat loop, tactical and close quarters |
| 5 | Offline jobs, shops, player economy working |
| 6 | Factions, settlements, caravans |
| 7 | Weird zones, mutations, cryptid creatures |
