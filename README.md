# CredSniffer
Small GO tool to find hard-coded credentials on local computer or network share

## Install
go build sniffer.go

## Usage
sniffer -dir=\\\\host\share -ext=ps1 -pattern=password,Password,AsSecureString

### Specify a directory
sniffer -dir=\\\\host\share 

### Specify extension
sniffer -ext=txt

### Specify patterns
sniffer -pattern="My Random Pattern",Token,token

!!sniffer is case sensitive

## ToDo
 - Add multi-threading

## License

MIT
