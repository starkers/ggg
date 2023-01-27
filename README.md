# ggg

**g**ood **g**it **g**etter


Written because I like to `git clone` and `cd` in a single command.

I also like to clome my git repo's into a heirarchical directory structure which.

EG:

`https://github.com/org-name/repo.git` ==  `$HOME/src/github.com/org-name/repo` 



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

# installation steps:

## 1. get the binary

### via download
download a binary from a release from: https://github.com/starkers/ggg/releases/latest

(don't forget to `chmod +x` the file after downloading it)

You can place the binary anywhere **so long as it is inside your** `$PATH`

In this case I'm going to store mine as `$HOME/.bin/ggg`


2: add Aliases to your various shell configs.. EG:


### compile and install it with go
If you have go: `go install github.com/starkers/ggg@latest` should install it for you under your `$GOPATH/bin`

the `ggg` binary should be built now


## 2. configure your shell

In order to work you need to tell your shell about `ggg`.

> TIP: Don't forget to open a fresh terminal/shell session **after** modifying your settings


### bash / zsh

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

### fish

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

On first run, **ggg** automatically creates a config file called `~/.config/ggg.toml`.

You can update settings (such as your base path) in the config file.

# Extending

If you have any feature requests, ideas for improvements please feel free to raise a PR or Issue.

# License
Apache-2.0