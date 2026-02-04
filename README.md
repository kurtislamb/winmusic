# winmusic

A CLI which can be installed in Windows SubSystem For Linux To Control Windows Music, the tool leverages windows default keep mappings. This is useful for users who run tiny keyboards without all the extra keys!

## Build and Install Build then Add the tool to usr/bin
```make
make build-install
```

## Build - Create the tool in output
```make
make build
```

## Install  Add the tool to usr/bin
```make
make install
```

## Usage Examples
```bash
## Get all Commands
winmusic help 

## Play
winmusic p

## Stop
winmusic Stop

## Increase Volume by 5 increments of 2
winmusic volumeup 5
```