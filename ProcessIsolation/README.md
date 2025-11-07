# 🧩 Step 1 — Process Isolation with Linux Namespaces (Building a Docker-like Runtime in Go)

> In this tutorial, we’ll learn how Docker isolates processes using Linux **namespaces**,  
> and we’ll implement our own mini-container runtime in **Go** that runs a process inside isolated namespaces.

---

## 🧠 1. What is Process Isolation?

When you run a normal process on Linux (like `/bin/bash`), it shares everything with the host:
- Same hostname
- Same process list
- Same filesystem
- Same network stack

But inside a **container**, the process believes it’s running on its own system:
- It sees only its own processes
- It has its own hostname
- It can have a private network interface

This illusion is provided by the **Linux kernel** using **namespaces**.

---

## ⚙️ 2. Linux Namespaces

Namespaces partition kernel resources so that one set of processes sees one set of resources,  
and another set sees a different set.

Each namespace type isolates a specific resource:

| Namespace | Flag | Isolates | Example |
|------------|------|-----------|----------|
| **PID** | `CLONE_NEWPID` | Process IDs | Container sees its own PID 1 |
| **UTS** | `CLONE_NEWUTS` | Hostname and domain | Container can change its hostname |
| **Mount (MNT)** | `CLONE_NEWNS` | Filesystem mounts | Container can mount without affecting host |
| **Network (NET)** | `CLONE_NEWNET` | Network interfaces, routes | Container gets its own `eth0`, IP, etc. |
| **IPC** | `CLONE_NEWIPC` | Shared memory, semaphores | Container IPC is private |
| **User** | `CLONE_NEWUSER` | UID/GID mappings | Root inside, unprivileged outside |

Namespaces are created by the `clone()` system call:

```c
clone(CLONE_NEWPID | CLONE_NEWUTS | CLONE_NEWNS | CLONE_NEWNET | CLONE_NEWIPC, ...);
```

This is what Docker (and all container runtimes) do under the hood.

---

## 🧩 3. Visual Understanding

When you run:

```bash
docker run ubuntu bash
```

Docker internally does something like this:

```
┌───────────────────────────┐
│ Host (dockerd)            │
│ ───────────────────────── │
│ clone() new namespaces    │
└──────────────┬────────────┘
               │
               ▼
      ┌──────────────────────┐
      │ Container Process    │
      │ PID namespace → #1   │
      │ Hostname → isolated  │
      │ Network → veth0      │
      └──────────────────────┘
```

The process runs in an isolated environment — but still uses the **same Linux kernel**.

---

## 💻 4. Implementation in Go

We’ll now write a Go program that runs a process in new namespaces.

### 📁 Setup

```bash
mkdir mydocker
cd mydocker
go mod init mydocker
```

### 🧾 File: `main.go`

```go
package main

import (
    "fmt"
    "os"
    "os/exec"
    "syscall"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: mydocker run <command>")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "run":
        run()
    default:
        fmt.Println("Unknown command:", os.Args[1])
    }
}

func run() {
    fmt.Printf("Running %v in new namespaces...\n", os.Args[2:])

    // Command to run inside the container
    cmd := exec.Command(os.Args[2], os.Args[3:]...)

    // Connect container process I/O to host terminal
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    // Apply namespace isolation
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Cloneflags: syscall.CLONE_NEWUTS |
            syscall.CLONE_NEWPID |
            syscall.CLONE_NEWNS |
            syscall.CLONE_NEWNET |
            syscall.CLONE_NEWIPC,
    }

    if err := cmd.Run(); err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}
```

---

## ▶️ 5. Running the Container

Compile and run:
```bash
sudo go run main.go run /bin/bash
```

Now you’re inside a **containerized shell**.

Try:
```bash
hostname
ps aux
```

You’ll notice that:
- PID numbers are different
- The hostname can be changed independently

Inside:
```bash
hostname mycontainer
exit
```

Outside:
```bash
hostname   # still unchanged
```

---

## 🧠 6. How It Works Internally

Let’s break down what happens:

| Go Function | Underlying Mechanism | Description |
|--------------|----------------------|--------------|
| `exec.Command()` | Spawns a child process | Equivalent to `fork() + exec()` |
| `SysProcAttr.Cloneflags` | Calls `clone()` syscall with flags | Creates new namespaces |
| `cmd.Run()` | Executes and waits for completion | Starts the new isolated process |

So this simple Go program:
- Calls the **Linux kernel’s clone syscall**
- Creates **new PID, UTS, MNT, NET, and IPC namespaces**
- Starts `/bin/bash` inside them

This is effectively the **core of Docker**.

---

## 🧩 7. Experiments

1. **Custom Hostname**

   Add inside `run()`:
   ```go
   syscall.Sethostname([]byte("suryansh-container"))
   ```
   Run again and check `hostname` inside.

2. **PID Isolation Demo**
   Add before `cmd.Run()`:
   ```go
   fmt.Printf("Parent PID: %d\n", os.Getpid())
   ```
   Inside container:
   ```bash
   ps
   ```
   → You’ll see your process is PID 1 inside the container.

3. **Namespace-by-Namespace Testing**

   Try enabling only one flag at a time:
   ```go
   syscall.CLONE_NEWUTS              // Only hostname isolated
   syscall.CLONE_NEWUTS | CLONE_NEWPID // PID + hostname isolated
   ```
   Observe what changes.

---

## 🧭 8. Summary

| You Learned | Description |
|--------------|-------------|
| **Linux namespaces** | Isolate resources per process |
| **clone() syscall** | Foundation of all containers |
| **Process isolation in Go** | `SysProcAttr.Cloneflags` |
| **Hands-on demo** | Your first mini container runtime |

✅ You’ve now built a basic **container sandbox** from scratch in Go —  
without using Docker or any external runtime.

---

## 🚀 Next Step: Step 2 — Filesystem Isolation

In the next tutorial, we’ll learn:
- How Docker changes what “/” means using **`chroot`** and **`pivot_root`**
- How **mount namespaces** isolate filesystems
- How to create a **custom root filesystem** using Alpine Linux
- How to enter it using Go

---

**Next → Step 2 — Filesystem Isolation: chroot + pivot_root**
