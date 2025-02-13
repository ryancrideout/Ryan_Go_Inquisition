# Concurrency

Resource: https://gobyexample.com/

Things to talk about:
- Threading in Go (Goroutines!)
    - Goroutines vs Threading (Check)
    - Channels!
    - Synchronization
    - Wait Groups
    - Worker Pools
    - Mutex (Mutual Exclusion)
    - AND THERE'S STILL MORE!
- Go is really good at concurrency
    - Go versus other languages
        - Performance
        - Ease of use
- What is low level threading?
- Email example
    - Should be done in Go AND PHP (and maybe Python)
        - Limit the number of threads to the number of CPUs to make it fair.
    - Should maybe include everything else mentioned (Goroutines, channels, etc.)

## Goroutines versus Threads
What is a thread?
- The abstract answer: A thread is an execution context, which is all the information a CPU needs to execute a stream of instructions.
    - Huh?

- In many cases, a thread is a part of a process.
    - Okay so then what is a process?
        - Essentially a process is an instance of a computer program, executed by one or many threads.
            - Examples of computer programes: 1998's StarCraft, a Python script, the PolicyR website, etc.

- Okay so in terms of a process, a thread is a lightweight process.
    - A thread belongs to ONE process.
    - You can (loosely) think of a process as a bunch of threads.
    - https://stackoverflow.com/questions/5201852/what-is-a-thread-really

- Do threads REALLY run in parallel?
    - They absolutely can, IF you're running on a multiple core processing system.
    - If you only have one core, then no.

- This is a very high level, simplified explanation. The rabbit hole definitely exists though.


What is a Goroutine?
- A Goroutine is NOT a thread, but they provide similar functionality.
- In Go, EVERY Go program starts with at least one goroutine, you can't avoid it.

Okay but what's the difference?
- Goroutines are managed by the go runtime, threads are managed by the kernel.
    - Because Goroutines are managed by the go runtime, it's easier for threads to communicate with each other than it is for threads to talk to each other.
    - I should explain this with an analogy.
- Goroutines are NOT hardware dependant, but threads are.
- Goroutines are CHEAPER than threads, and have a faster startup time.
- More in the weeds, but memory initialized for kernel threads doesn't shrink. For Goroutines, memory initialized can shrink.
    - Money analogy - if you give someone $1000 to buy something, the thread will use the whole $1000 where the Goroutine will only take what is needed.
- https://durgeshatal1995.medium.com/understanding-the-differences-between-go-routines-and-threads-b631068d4fdd 
