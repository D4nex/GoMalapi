<h3 align="center">MalapiGo</h3>
<p align="center">
   A simple module that uses <i>malapi.io</i> to obtain information about <b>WinAPI Calls</b>
</p>

## Install
```
git clone https://github.com/D4nex/MalapiGo
cd MalapiGo/src/
go build -o malapi-go main.go
sudo mv malapi-go /usr/bin/
sudo chmod +x /usr/bin/malapi-go
```

## Usage
```console
d4nex@pwn:~$ malapi-go
malAPI-go - Get information about a Windows API calls simply!
          Usage: ./malapi-go <API CALL>
```
## Example
```
./malapi-go LoadLibraryA


[+] API CALL:
            LoadLibraryA
[+] DESCRIPTION:
                LoadLibraryA is used to load a specified module into the address space of the calling process. Malware commonly use this to load DLLs dynamically for evasion purposes.
[+] LIBRARY(DLL):
                Kernel32.dll
[+] TAGS:
                    Injection
                    Evasion
[+]DOCUMENTATION:   https://docs.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-loadlibrarya
```
