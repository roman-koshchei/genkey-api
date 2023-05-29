# Genkey API

![GitHub](https://img.shields.io/github/license/roman-koshchei/genkey-api?color=%2300add8)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/roman-koshchei/genkey-api?color=%2300add8)
![GitHub Release Date](https://img.shields.io/github/release-date/roman-koshchei/genkey-api?color=%2300add8)
![GitHub last commit](https://img.shields.io/github/last-commit/roman-koshchei/genkey-api?color=%2300add8)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/roman-koshchei/genkey-api?color=%2300add8)

API for [genkey analyzer](https://github.com/semilin/genkey). Get layout analysis from api as a json. Used Gin framework. Yeah

## Rules

Input url must be formatted with url encoder, which replace ';' to '%3B'. Or just do it yourself. In javascript use encodeUriComponent()


## Path '/divided/'

Query parameters:
- `topKeys` = string of **top row keys** from left to right
- `homeKeys` = string of **home row keys** from left to right
- `botKeys` = string of **bottom row keys** from left to right
- `topFingers` = string of **top row fingers** from left to right
- `homeFingers` = string of **home row fingers** from left to right
- `botFingers` = string of **bottom row fingers** from left to right

QWERTY example:

<https://genkey-api.up.railway.app/divided/?topKeys=qwertyuiop&homeKeys=asdfghjkl%3B%27&botKeys=zxcvbnm,./&topFingers=0123344567&homeFingers=01233445677&botFingers=0123344567>



## Path '/together/'

Query parameters: 
- `keys` = string for all keys from first on top row to last on bottom row
- `fingers` = string for all fingers from first on top row to last on bottom row

QWERTY example:

<https://genkey-api.up.railway.app/together/?keys=qwertyuiop[]\asdfghjkl%3B%27zxcvbnm,./&fingers=0123344567777012334456770123344567>
