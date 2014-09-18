# Kindle PW2 related (experimental and probably temporary)
======

## I do programing on Windows (64 bits), so this work may help you to setup cross environment for Go language **targetting Kindle PW2**, including complete usage of *cgo*.

*go_1.3.1* contains command files to setup cross compilation environment to build *Go* itself. 
*cross* has self extracting archive with complete build of Linaro 20140811 toolchain tuned for Kindle PW2 (eglibc 2.12, binutils 2.24, etc) hosted 
on Windows 64 bits. Please, *note* that self extractor will need to create symbolic links, so it has to be run with admin privileges!

In order to have full working Go on Windows 64 bits, targeting **Windows 32/64 bits and Linux ARM** you will need to obtain and setup following tools:

* [mingw compilers for 32 and 64 bits](http://win-builds.org) 
* mingw64 hosted cross-compiler for arm (armkpw2.exe here)
* [Git](http://msysgit.github.com)
* [SVN](http://www.sliksvn.com)
* [Mercurial](http://mercurial.selenic.com)
* [Bazaar](http://wiki.bazaar.canonical.com/WindowsDownloads)

My command files assume that all of the above has been installed under single directory:
```
GOTOOLS
|
+---armkpw2
+---Bazaar
+---Git
+---Mercurial
+---mingw
+---mingw64
\---SlikSvn
```

Download and unpack [Go 1.3.1 source distribution](https://storage.googleapis.com/golang/go1.3.1.src.tar.gz). 
Apply `go_131_windows.patch` to sources.

Let's assume you have it in `c:\go` and my command files and necessary tools in `c:\gotools`. 
Then issuing following commands will build development environment:

```
c:
cd c:\go\src

ren .\make.bat .\make.orig
copy c:\gotools\make.cmd .

c:\gotools\buildgo_cross_amd64.cmd windows amd64
.\make.cmd
c:\gotools\buildgo_cross_amd64.cmd windows 386
.\make.cmd --no-clean
c:\gotools\buildgo_cross_amd64.cmd linux arm
.\make.cmd --no-clean

```

Now you could use `buildgo_(lin|win)_(arm|x86|x64).cmd` files as a templates (some paths hard-coded) to setup host environment for proper target during 
your development.