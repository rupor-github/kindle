package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "kal/Godeps/_workspace/src/github.com/mattn/go-sqlite3"
	"kal/Godeps/_workspace/src/github.com/twinj/uuid"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var exts map[string]bool = map[string]bool{".azw": true, ".mobi": true, ".prc": true, ".pobi": true, ".azw3": true, ".azw6": true, ".yj": true, ".azw1": true, ".tpz": true, ".pdf": true, ".txt": true, ".html": true, ".htm": true, ".jpg": true, ".jpeg": true, ".azw2": true}
var firstRun bool

type book struct {
	uuid        string
	location    string
	collections map[string]*collection
}

type collection struct {
	uuid  string
	label string
	books map[string]*book
}

func findBook(books []book, pred func(*book) bool) int {
	for i := range books {
		if pred(&books[i]) {
			return i
		}
	}
	return -1
}

func findCollection(collections []collection, pred func(*collection) bool) int {
	for i := range collections {
		if pred(&collections[i]) {
			return i
		}
	}
	return -1
}

func cleanDB() (err error) {
	db, err := sql.Open("sqlite3", dbDSNrw)
	if err != nil {
		return
	}
	defer db.Close()

	_, err = db.Exec("DROP TABLE IF EXISTS Autolists")
	return
}

func writeDB(collections []collection) (err error) {
	if len(collections) == 0 {
		return
	}

	db, err := sql.Open("sqlite3", dbDSNrw)
	if err != nil {
		return
	}
	defer db.Close()

	now := time.Now().Unix()
	tx, err := db.Begin()
	if err == nil {
		if _, err = tx.Exec("CREATE TABLE IF NOT EXISTS Autolists (uuid PRIMARY KEY NOT NULL UNIQUE, label, path, time)"); err == nil {
			if _, err = tx.Exec("DELETE FROM Autolists"); err == nil {
				for _, c := range collections {
					// path and time are not currently used, but may become useful if and when I decide to really optimize
					_, err = tx.Exec(fmt.Sprintf("INSERT INTO Autolists(uuid, label, path, time) VALUES (\"%s\",\"%s\",\"%s\",\"%d\")", c.uuid, c.label, dbPath, now))
					if err != nil {
						break
					}
				}
			}
		}
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}
	return
}

func readDB() (collectionsOld map[string]string, collections []collection, books []book, err error) {
	db, err := sql.Open("sqlite3", dbDSNro)
	if err != nil {
		return
	}
	defer db.Close()

	getCount := func(query string) (size int, err error) {
		if err = db.QueryRow(query).Scan(&size); err == sql.ErrNoRows {
			err = nil
			size = 0
		}
		return
	}
	getItems := func(query string, reader func(rows *sql.Rows, acc int) error) (err error) {
		rows, err := db.Query(query)
		if err != nil {
			return
		}
		defer rows.Close()

		for i := 0; rows.Next(); i++ {
			if err = reader(rows, i); err != nil {
				return
			}
		}
		err = rows.Err()
		return
	}

	var size int

	if size, err = getCount("SELECT COUNT(name) FROM sqlite_master WHERE type='table' AND name = \"Autolists\""); err != nil {
		return
	}
	if size > 0 {
		collectionsOld = make(map[string]string)
		if err = getItems("SELECT uuid, label FROM Autolists",
			func(rows *sql.Rows, _ int) (err error) {
				var u, l string
				if err = rows.Scan(&u, &l); err == nil {
					collectionsOld[u] = l
				}
				return
			}); err != nil {
			return
		}
	} else {
		firstRun = true
		log.Print("This is initial program run...")
	}

	if size, err = getCount("SELECT COUNT(p_type) FROM Entries WHERE p_type = \"Collection\""); err != nil {
		return
	}
	collections = make([]collection, size, size)
	if err = getItems("SELECT p_uuid, p_titles_0_nominal FROM Entries WHERE p_type = \"Collection\"",
		func(rows *sql.Rows, acc int) (err error) {
			var u, l string
			if err = rows.Scan(&u, &l); err == nil {
				collections[acc] = collection{uuid: u, label: l, books: make(map[string]*book)}
			}
			return
		}); err != nil {
		return
	}

	if size, err = getCount("SELECT COUNT(p_type) FROM Entries WHERE p_type = \"Entry:Item\" AND p_location IS NOT NULL AND p_location LIKE '" + dbPath + "%'"); err != nil {
		return
	}
	books = make([]book, size, size)
	if err = getItems("SELECT p_uuid, p_location, p_cdeKey, p_cdeType FROM Entries WHERE p_type = \"Entry:Item\" AND p_location IS NOT NULL AND p_location LIKE '"+dbPath+"%'",
		func(rows *sql.Rows, acc int) (err error) {
			var u, l string
			var k, t sql.NullString
			if err = rows.Scan(&u, &l, &k, &t); err == nil {
				books[acc] = book{uuid: u, location: l, collections: make(map[string]*collection)}
			}
			return
		}); err != nil {
		return
	}

	// We really only need collections map on a book entry, books map on collection entry is currently unused
	if err = getItems("SELECT i_collection_uuid, i_member_uuid FROM Collections",
		func(rows *sql.Rows, acc int) (err error) {
			var cu, bu string
			if err = rows.Scan(&cu, &bu); err == nil {
				if ci, bi := findCollection(collections, func(c *collection) bool { return c.uuid == cu }), findBook(books, func(b *book) bool { return b.uuid == bu }); ci != -1 && bi != -1 {
					collections[ci].books[bu] = &books[bi]
					books[bi].collections[cu] = &collections[ci]
				} else {
					log.Printf("RI: Skipping collection %s (collection_idx: %d, book_idx: %d)", cu, ci, bi)
				}
			}
			return
		}); err != nil {
		return
	}

	return
}

func readDeviceFolders() (files map[string]string, err error) {
	checkFile := func(path string, info os.FileInfo, err1 error) error {
		if err1 == nil && !info.IsDir() {
			dir, file := filepath.Split(path)
			if strings.HasPrefix(dir, fsPath) {
				if rel := strings.TrimPrefix(dir, fsPath); len(rel) > 0 && !strings.HasPrefix(rel, "dictionaries") {
					if accepted, supported := exts[strings.ToLower(filepath.Ext(file))]; supported && accepted {
						files[file] = strings.TrimSuffix(filepath.ToSlash(rel), "/")
					}
				}
			}
		}
		return err1
	}

	files = make(map[string]string)
	err = filepath.Walk(fsPath, checkFile)
	return
}

var debug bool
var verbose bool
var help bool
var action string

func init() {
	flag.StringVar(&action, "action", "sync", "action to perform: sync|clean|none")
	flag.BoolVar(&verbose, "verbose", false, "print additional information")
	flag.BoolVar(&debug, "debug", false, "save JSON request to file, do not update database, do not write to syslog")
	flag.BoolVar(&help, "help", false, "print help information")
}

func main() {
	var err error

	start := time.Now()

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nKindle auto-lists tool (rupor)\nVersion 0.1\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
	}

	if flag.Parse(); flag.NFlag() == 0 || help {
		action = ""
	}

	// this neeeds to happen after we have command line arguments, but before anything else
	prepareLog()

	// See if we are allowed to proceed
	if config.NotActive {
		action = "none"
	}

	// Make sure only one instance of this program is running at the time
	lock := filepath.Join(os.TempDir(), "kal_lock")
	os.Remove(lock)
	flock, err := os.OpenFile(lock, os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		log.Print("Another instance of this program is already running")
		os.Exit(2)
	}
	defer flock.Close()

	switch action {
	default:
		flag.Usage()
		os.Exit(1)

	case "none":
		log.Print("Actions are disabled via configuration file. Doing nothing...")

	case "clean":
		log.Print("Removing ALL existing collections and cleaning database")

		_, collections, _, err := readDB()
		if err != nil {
			log.Fatal(err)
		}
		q := NewQueue()
		for _, c := range collections {
			q.DeleteCollection(c.uuid)
		}

		log.Printf("Sending commands, %s", time.Since(start))
		if debug {
			err = q.Dump()
		} else {
			err = q.Execute()
			// Remove our table(s) from database if any
			if err == nil {
				err = cleanDB()
			}
		}

	case "sync":
		log.Printf("Syncing collections with folders [%s]", fsPath)

		// Get all necessary information
		collectionsOurs, collections, books, err := readDB()
		if err != nil {
			log.Fatal(err)
		}
		files, err := readDeviceFolders()
		if err != nil {
			log.Fatal(err)
		}

		// Remove ALL collections on a first run, remove OUR collections otherwise.
		// ccat already took care of updating member links for any removed books and collections.
		// we just need to take special care of updating our book entries properly in case
		//   any of our books were manually added to some collections not managed by this program.
		q := NewQueue()
		for _, c := range collections {
			if _, isOur := collectionsOurs[c.uuid]; firstRun || isOur {
				q.DeleteCollection(c.uuid)
				for _, b := range c.books {
					if _, found := b.collections[c.uuid]; found {
						delete(b.collections, c.uuid)
					} else {
						log.Printf("RI: Book %s does not belong to collection %s", b.uuid, c.uuid)
					}
				}
			}
		}

		// Preparing new collection set
		collectionsNew := make([]collection, 0, 1)

		uuid.SwitchFormat(uuid.CleanHyphen)
		for file, label := range files {
			if i := findBook(books, func(b *book) bool { return b.location == filepath.ToSlash(filepath.Join(dbPath, label, file)) }); i < 0 {
				log.Printf("Invalid location for: %s, ignoring", file)
			} else {
				j := findCollection(collectionsNew, func(c *collection) bool { return c.label == label })
				if j < 0 {
					collectionsNew = append(collectionsNew, collection{uuid: strings.ToLower(uuid.NewV4().String()), label: label, books: make(map[string]*book)})
					j = len(collectionsNew) - 1
				}
				books[i].collections[collectionsNew[j].uuid] = &collectionsNew[j]
				collectionsNew[j].books[books[i].uuid] = &books[i]
			}
		}
		for _, c := range collectionsNew {
			q.InsertCollection(c.uuid, c.label)
			q.UpdateCollection(c.uuid, c.books)
		}
		for _, b := range books {
			q.UpdateBook(b.uuid, len(b.collections))
		}

		log.Printf("Sending commands, %s", time.Since(start))
		if debug {
			err = q.Dump()
		} else {
			err = q.Execute()

			// Save in the database for the next run
			if err == nil {
				err = writeDB(collectionsNew)
			}
		}
	}
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Done..., %s", time.Since(start))
	}
}
