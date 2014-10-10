# Kindle PW(1/2) Auto-lists
======

## Credits

* This would not be possible without outstanding work of multiple individuals from **Kindle Developer's Corner** on mobileread.com. I am new Kindle user (got my PW2 couple of months back) to know them all, so I will just mention a few here (in no particular order): ixtab, geekmaster, NiLuJe, knc1, eureka, twobob.
* This particular program was heavily influenced by relatively recent barsanuphe's work [Librarian & LibrarianSync: CLI ebook manager & Kindle collections builder](http://www.mobileread.com/forums/showthread.php?t=245691)

## Rationale

This program (Kindle Auto Lists) implements (IMHO) the only critical piece of functionality completely missing from Kindle firmware - it automatically creates user collections, based on books directory structure.
Going through mobileread forums I was able to find multiple scripts/programs which to some degree do what I need, barsanuphe's LibrarianSync being really close, but have some dependencies and features which I would like to avoid.

## Requrements

* Jailbroken Kindle PW2 (should work on PW1, but I do not have the device to check this claim)

## Installation
### This is **work-in-progress**

Download latest rupor_X.X.7z archive from [here](https://github.com/rupor-github/kindle/tree/master/package). Inside you will find proper installation and
uninstallation .bin files. Just follow [regular procedure](http://www.amazon.com/gp/help/customer/display.html?nodeId=201307490).

Install contains handy [KUAL](http://www.mobileread.com/forums/showthread.php?t=203326) extension (Collection from folders) which under normal circumstances you would not need,
but may find useful in case of problems.

At the first run program will remove ALL existing collections and will create new set, based on existing folder structure.
On sequential runs it will only operate only with collections it creates, preserving and maintaining other collections.

At any time you could use "showlog | grep KAL" to see what it does.

### Command line

* -action=sync|clean where "clean" will delete ALL collections and clean program's action history from database
* -debug program will save CCAL command set to file and would not update database. In addition all program output will go to screen, rather than in syslog

### Config file (JSON)

* relRoot - root directory (inside "documents" folder to synchronize)
* notActive - if true, program will immediately exit. This allows user to stop synchronizing folders without editing /etc/lipc-daemon-events.conf

### Performance
Currently it takes Kindle CCAL about a minute to update ~300 collections with ~1200 books:
```
141010:200102 KAL[7053]: 2014/10/10 20:01:02 Syncing collections with folders [/mnt/us/documents/]
141010:200126 KAL[7053]: 2014/10/10 20:01:26 Sending 1975 commands, 24.523070169s
141010:200214 KAL[7053]: 2014/10/10 20:02:14 Processed - 1975 changes...
141010:200218 KAL[7053]: 2014/10/10 20:02:18 Done..., 1m15.759366259s
```
I would not expect this to work for librarians with many thousands books on devices...

