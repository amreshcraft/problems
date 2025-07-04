# GoProFoundation

## 📂 **1. Formatting & Number Conversion – `fmt` & Format Specifiers**
- [x] Difference between `%v`, `%+v`, `%#v` with example outputs.
- [x] How to format integers as hexadecimal, binary, octal using `fmt.Sprintf`.
- [x] Convert float to string with precision control.
- [x] Create a custom string representation using `Stringer` interface.
- [x] Format using width and left/right alignment.
- [x] Format a number with leading zeros.
- [ ] Format large numbers using commas (like 1,000,000).
- [ ] Print struct with and without field names.
- [ ] Explain `%t` format for boolean.
- [ ] Print pointer address and value in Go.
- [ ] Show difference between `%q` and `%s` when printing strings.
- [ ] How to suppress printing newline in `fmt.Print`.
- [ ] Write a custom format for a date using `fmt`.
- [ ] Use `%e` and `%E` for scientific notation printing.
- [ ] Format a complex number in Go.

---

## 📂 **2. Crypto & Hashing – `crypto` Package**
- [x] Generate SHA-256 hash of a string.
- [x] Generate MD5 hash of a file.
- [x] Hash a password using bcrypt.
- [x] Verify bcrypt password.
- [ ] Generate random secure tokens.
- [ ] Generate HMAC using SHA-512.
- [ ] Explain difference between SHA-256 and HMAC-SHA-256.
- [ ] Generate cryptographically secure random numbers.
- [ ] Implement simple RSA encryption/decryption.
- [ ] What is salting? How to add salt to passwords.
- [ ] Difference between encoding/base64 and hashing.
- [ ] Build a small login system with password hashing using bcrypt.

---

## 📂 **3. Struct, Methods, Embedding**
- [ ] Implement method receivers (pointer vs value) with examples.
- [ ] Create a struct with embedded fields and access them.
- [ ] Build a method chain using pointer receivers.
- [ ] Explain zero values in structs with custom example.
- [ ] Compare two structs deeply.
- [ ] Implement a struct with a custom constructor function.
- [ ] How to pass struct by reference to a function.
- [ ] Design a polymorphic system using struct embedding.
- [ ] Implement deep copy of a struct.
- [ ] Build a linked list using structs.
- [ ] Build a simple in-memory database using structs.
- [ ] Struct tags: explain and parse using `reflect`.
- [ ] Add validation using struct tags.
- [ ] Implement caching using map + struct.
- [ ] Difference between anonymous field and named field in embedding.

---

## 📂 **4. Interface Deep Dive**
- [ ] Build your own Reader interface.
- [ ] Create a polymorphic logging system using interface.
- [ ] Difference between concrete types vs interface types in Go.
- [ ] Explain type assertion with practical example.
- [ ] Build a system using empty interface and type switches.
- [ ] Write a function that accepts any data type using interface{}.
- [ ] Build an in-memory repository using interface.
- [ ] Build a mock service using interface for testing.
- [ ] Implement io.Reader and io.Writer for your own struct.
- [ ] Create interface segregation example in Go.
- [ ] Write a decorator using interface.
- [ ] Build a composite interface and use it.
- [ ] Interface vs struct: When to use what.
- [ ] Write code where one interface satisfies another.
- [ ] Implement a FileStorage interface with local and in-memory storage.
- [ ] Implement a Logger interface with zap and fmt backends.
- [ ] Build a service layer interface for a URL shortener.
- [ ] Build an interface for external API calls.
- [ ] Use interface to create a strategy pattern.
- [ ] Explain empty interface vs generic in Go 1.18.

---

## 📂 **5. Design Patterns in Go**
- [x] Build Singleton using sync.Once.
- [ ] Implement Factory Pattern in Go.
- [ ] Implement Abstract Factory Pattern.
- [ ] Build Observer Pattern using channels.
- [ ] Implement Decorator Pattern with middleware example.
- [ ] Build Strategy Pattern using interface.
- [ ] Implement Adapter Pattern in Go.
- [ ] Implement Command Pattern.
- [ ] Build a Producer-Consumer system using channels (Pipeline pattern).
- [ ] Implement a Builder Pattern in Go.
- [ ] Explain Prototype Pattern in Go.
- [ ] Create a Service Locator Pattern using maps.
- [ ] Implement Chain of Responsibility Pattern.
- [ ] Build Proxy Pattern using interface.
- [ ] Build Template Pattern using function callbacks.

---

## 📂 **6. Time Package Mastery**
- [ ] Parse string to time.Time using different layouts.
- [ ] Format time in custom layouts.
- [ ] Implement timeout using `time.After`.
- [ ] Build a scheduler that runs every X seconds.
- [ ] Measure function execution time.
- [ ] Convert timestamps to local timezones.
- [ ] Difference between `time.Now()` and `time.UTC()`.
- [ ] Calculate date difference in days, months, years.
- [ ] Add/subtract durations from current time.
- [ ] Build a retry system using `time.Sleep` and exponential backoff.

---

## 📂 **7. Generics in Go**
- [ ] Build a generic stack.
- [ ] Build a generic repository.
- [ ] Write a generic function to compare any two values.
- [ ] Build a generic min/max function.
- [ ] Build a generic queue.
- [ ] Build a generic map to slice converter.
- [ ] Implement generic filtering on slices.
- [ ] Build a type-safe generic cache.
- [ ] Implement generic logger that accepts any data type.
- [ ] Use generics to build a type-safe producer-consumer system.

---

## 📂 **8. Bufio, File Reading, Writing, OS Packages**
- [ ] Read large file line by line using bufio.Scanner.
- [ ] Read entire file into memory.
- [ ] Write string to a file using bufio.Writer.
- [ ] Append data to a file.
- [ ] Use os/exec to run system commands (like ls, cat).
- [ ] Read and list all files in a directory.
- [ ] Check file permissions and modify them.
- [ ] Create and delete files using Go.
- [ ] Implement grep like functionality using bufio.
- [ ] Use os.Stat to check file info.
- [ ] Write logs to files using os and bufio.
- [ ] Compress and decompress files.
- [ ] Build a file watcher using fsnotify.
- [ ] Build a CLI to copy files using bufio.
- [ ] Handle file read/write errors properly.

---

## 📂 **9. Logging Best Practices**
- [ ] Implement zap logger with structured logs.
- [ ] Setup log levels: Debug, Info, Warn, Error.
- [ ] Create log rotation for large files.
- [ ] Add request ID and trace ID in logs.
- [ ] Log to file and stdout simultaneously.
- [ ] Implement logging middleware in HTTP server.
- [ ] Implement JSON formatted logs.

---

## 📂 **10. Goroutines & Concurrency Mastery**
- [ ] Build a worker pool with goroutines.
- [ ] Use sync.WaitGroup to wait for multiple goroutines.
- [ ] Difference between buffered and unbuffered channels.
- [ ] Implement timeout using context.WithTimeout.
- [ ] Build a pipeline system using channels.
- [ ] Deadlock example and fix it.
- [ ] Race condition example and fix it.
- [ ] Implement fan-out, fan-in pattern.
- [ ] Build a publish-subscribe system using channels.
- [ ] Build concurrent downloader.
- [ ] Use sync.Mutex to protect shared resource.
- [ ] Use sync.Map safely across goroutines.
- [ ] Use select statement to multiplex channels.
- [ ] Build a rate limiter using channels.
- [ ] Use context for cancellation across goroutines.
- [ ] Build a batch processor using goroutines.
- [ ] Implement parallel file processing.
- [ ] Build a chat room backend using goroutines and channels.
- [ ] Use sync.Cond to signal waiting goroutines.
- [ ] Use sync.Once for singleton initialization.
- [ ] Use sync.Pool for memory optimization.
- [ ] Build a real-time notification system.
- [ ] Goroutine leak example and how to prevent it.
- [ ] Use errgroup for handling multiple goroutines.
- [ ] Difference between goroutine and thread.

---

## 📂 **11. Networking – net Package**
- [ ] Build a TCP server and TCP client.
- [ ] Build a UDP server and client.
- [ ] Implement an HTTP server using net/http without any framework.
- [ ] Build a TCP file transfer server.
- [ ] Build a TCP chat server using goroutines.
- [ ] Implement HTTP client with timeout and retries.
- [ ] Build custom router using net/http.
- [ ] Build concurrent HTTP downloader.

---

## 📂 **12. System Level – OS, OS/Exec, Sys Packages**
- [ ] Run system commands and capture output using os/exec.
- [ ] Handle OS signals like SIGINT, SIGTERM gracefully.
- [ ] Use syscall to read environment variables.
- [ ] Implement cross-platform file system operations.
- [ ] Build a CLI tool using Go.
- [ ] Fetch system memory and CPU info.
- [ ] Implement graceful shutdown in HTTP server.
- [ ] Build a simple process monitor in Go.

---
