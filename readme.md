# xtool

![](screenshot.png)


## build in win


```bash
rsrc -manifest win.mainfest -ico icon/icon.ico -o rsrc.syso

go build -ldflags="-H windowsgui"
```

## auto gen menu

see `gen.go` & `menu.yaml`

```bash
go generate
```

### support type

- link
- clipboard


## reference

- https://github.com/akavel/rsrc
- https://github.com/getlantern/systray


