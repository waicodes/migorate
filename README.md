# migorate

**migorate** is a tiny, no‑CLI SQL migration runner for Go.  
It executes `.sql` files in order, records which versions have been applied, and blocks modified files with a built‑in **SHA‑256 hash check**—keeping your schema history safe without any external tools.

---

## ✨  Why use migorate?

| Feature | Description |
|---------|-------------|
| **Plain SQL** | Write migrations in pure SQL—no DSL, no code generation. |
| **Version tracking** | Each file is logged in a `migrations` table so it runs only once. |
| **Tamper detection** | SHA‑256 hash is stored; if a file is edited after apply, migorate aborts. |
| **Zero binaries** | Just `go get`, import, and call `migorate.Run`. |
| **DB‑agnostic** | Works with any database driver implementing `database/sql` (Postgres, MySQL, SQLite, etc.). |

---

## 📦 Installation

```bash
go get github.com/waicodes/migorate
```

## 📁 Project structure

```bash
your-app/
└── migrations/
    ├── 001_create_users.sql
    └── 002_add_email_column.sql
```
You can change directory for save file migration.

## 🧪 Example

```bash
db, err := sql.Open("postgres", dsn)
if err != nil {
    log.Fatal("Failed to connect to DB:", err)
}

if err := migorate.Run(db, "migrations"); err != nil {
    log.Fatal("Migration failed:", err)
}
```
You can change db connnection driver, this only example hot to use **migorate**