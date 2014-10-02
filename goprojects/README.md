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
### Eventually I am planning proper Kindle's "update" file, however currently this is **work-in-progress**

* Unpack archive at the root of your kindle books drive (I am assuming that program is in /mnt/us/kal/bin and config file is /mnt/us/kal/config)
* Add line "fullScanFinish    com.lab126.scanner    /mnt/us/kal/bin/kal -action=sync" to /etc/lipc-daemon-events.conf on your device
* reboot your device

At the first run program will remove ALL existing collections and will create new set, based on existing folder structure.
On sequential runs it will only operate only with collections it creates, preserving and maintaining other collections.

You could use "showlog | grep KAL" to see what it does.

### Command line

* -action=sync|clean where "clean" will delete ALL collections and clean program's action history from database
* -debug program will save CCAL command set to file and would not update database. In addition all program output will go to screen, rather than in syslog
* -verbose will print some additional information

### Config file (JSON)

* relRoot - root directory (inside "documents" folder to synchronize)
* notActive - if true, program will immediately exit. This allows user to stop synchronizing folders without editing /etc/lipc-daemon-events.conf

### Performance
Currently it takes Kindle CCAL about 4 seconds to update ~30 collections with ~100 books:
```
141002:185358 KAL[12302]: 2014/10/02 18:53:58 Syncing collections with folders [/mnt/us/documents/mybooks/]
141002:185359 KAL[12302]: 2014/10/02 18:53:59 Sending 179 commands, 140.092792ms
141002:185401 KAL[12302]: 2014/10/02 18:54:01 Processed - 179 changes...
141002:185402 KAL[12302]: 2014/10/02 18:54:02 Done..., 3.82349459s
```
I would not expect this to work for librarians with thousands of books on devices...

