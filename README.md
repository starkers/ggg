# ggg

**g**ood **g**it **g**etter


Written because I like to `git clone` and `cd` in a single command



# Example

In this example we:
1. Clone a https:// git repo and automatically `cd` into its location.
2. Next we clone a `git@` git repo and automatically arrive in its current working dir.
3. Repeat command 1, and we are back where its code resides


![example](/media/example-zsh.png)



## FAQ:

### Can I clone to a different path?

Yes, set `path` in: `~/.config/ggg.toml`

### Why?

I used to use [this monstrosity](https://github.com/starkers/homedirectory/blob/a8f4e95dd5bd6eb857e30935396e51a442acd619/home/aliases#L105-L159) of a zsh/bash alias for years, but after migrating to fish I missed it... Also, it would be nice to make something more "robust"


---

# installation

1. Download `ggg` and ensure the file is executable

### golang:
If you have go: `go get -u github.com/starkers/ggg`

### cli:
download a binary from a release from: https://github.com/starkers/ggg/releases/latest

(don't forget to `chmod +x` the file after downloading it)

You can place the binary anywhere **so long as it is inside your** `$PATH`

In this case I'm going to store mine as `$HOME/.bin/ggg`

EG:

```
# make a directory if it doesn't exist
GGG_BIN=~/.bin/ggg/ggg-bin

# download it
wget https://github.com/release/latest -O "${GGG_BIN}"

# make it executable
chmod +x "${GGG_BIN}"
```

2: add Aliases to your various shell configs.. EG:


### 2.1 bash / zsh

add something like this to your `~/.profile`, `~/.bashrc` or `~/.zshrc` (whichever u use really):

```
if [[ -o interactive ]]; then
  if hash ggg 2>/dev/null; then
    eval "$(ggg hook zsh)"
  else
    echo 'WARN: "ggg" binary not found in your PATH'
  fi
fi
```

### 2.2 fish

Add something like this to a fish conf

(for example: `~/.config/fish/conf.d/ggg.fish`)

```
# ggg
# vi: ft=fish

if status is-interactive
  if command -s ggg > /dev/null
    ggg hook fish | source
  else
    echo 'WARN: "ggg" binary not found in PATH'
  end
end

```



# Configure

on first run, **ggg** creates a config file called `~/.config/ggg.toml`

you can update settings such as your base path for git cloneing there

