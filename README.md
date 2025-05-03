# ssh-diddy-simulator

**2 men - 2 cigarettes - 7 hours.**

A terminal-based multiplayer snake game over SSH, written in pure Go.  
Runs in your terminal. Looks like a school project. Plays like a fever dream.

---

## 🚀 Overview

You SSH into the server.  
You're assigned a random colored letter.  
You are now a snake.

Eat food (ϖ), dodge other snakes, and try not to implode.  
If you collide with yourself — that’s *Darwinism*.  
If another snake is longer — that’s **C**apitalism.  
If you press `l` — you **literally delete yourself**.

---

## 🧠 Usage

### Commands

```bash
go run .
```

Or deploy it somewhere weird and invite strangers to suffer.

<sup>*We definitely wouldn't do that ourselves... [🐍🐍🐍](do_not.md)*<sup>

Then from another terminal:

```bash
ssh -p 8080 localhost
```

### Controls

WASD to move.
`l` to **rage quit**.

---

## ⚙️ How It Works

* `gliderlabs/ssh`: handles real-time SSH sessions like a chatroom for worms.
* ANSI escape codes: to draw rainbow puke in your terminal.
* Mutexes: to prevent the snakes from tearing each other (and your CPU) apart.
* A `DickSize` field on food objects: we don’t talk about that.

The game loop ticks every 100ms.
Players are tracked via IP, symbolized by a single letter, and rendered with ANSI colors.

---

## 🗑️ Naveed

Naveed is the kind of guy who:

* **Writes code without testing it,**
* **Commits directly to `main`,**
* And still says **"it works on my machine."**

> [!CAUTION]
> DO NOT let Naveed near this codebase.
> 
> **Do not code review with him. Do not merge his branches. Do not speak his variable names aloud.**
>
> We are still undoing what he called "aesthetic concurrency".
>
> We may never recover...

---

## ©️ License

MIT.
Use it, break it, remix it.
Just don’t let Naveed near it again.

Never again...

---

## 🧑‍💻 Contributors <sup>(and naveed)</sup>

- [@mush1e](https://github.com/mush1e)
- [@danqzq](https://github.com/danqzq)
- ~~Naveed~~

---

## 📜 Quotes

> "This code has more race conditions than Mario Kart." - *mush1e*

> "This is the only Go project where dying is a feature, not a failure." - *danqzq*

> "I gave the food a `DickSize` attribute ironically. Now it’s the most stable part of the system." - *mush1e*

> "You may control the game logic, the rendering, even the physics - but you will never control Naveed" - *Sun Tzu probably*

> "`H E L P`" - *deadlocked goroutine*

> "It’s built in Go, but it runs like JavaScript having a panic attack." - *mush1e*

> "I benchmarked the game’s performance by how loudly my laptop fan screamed." - *danqzq*

> "Just why" - *Ken Thompson*

> "This is not failure. This is persistence after meaning has vanished." - *Alan Turing upon viewing `ssh-diddy-simulator`*

<br><br><br>

# "We're sorry" - dev team
