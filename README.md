# filesmv
引数で指定した拡張子のファイルを全て探して，拡張子名のディレクトリに移動させる


## Install

```
brew tap mitubaEX/filesmv

brew install filesmv
```

## Usage

```
filesmv <filename extention>
```


## Example

```
$ ls
README.md  main.go

$ filesmv go

$ tree
.
├── README.md
└── go
    └── main.go

1 directory, 2 files

```


