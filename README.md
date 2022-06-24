# Genkey API

API for [genkey analyzer](https://github.com/semilin/genkey) made by @semilin. Get layout analysis from api as a json. Used Gin framework.


## Rules

Input url must be formatted with url encoder, which replace ';' to '%3B'. Or just do it yourself.


## Path '/divided/'

Query parameters:
- topKeys = string of `top row keys` from left to right
- homeKeys = string of `home row keys` from left to right
- botKeys = string of `bottom row keys` from left to right
- topFingers = string of `top row fingers` from left to right
- homeFingers = string of `home row fingers` from left to right
- botFingers = string of `bottom row fingers` from left to right

QWERTY example:
`/divided/?topKeys=qwertyuiop&homeKeys=asdfghjkl%3B%27&botKeys=zxcvbnm,./&topFingers=0123344567&homeFingers=01233445677&botFingers=0123344567`


## Path '/'

Query parameters: 
- keys = string for all keys from first on top row to last on bottom row
- fingers = string for all fingers from first on top row to last on bottom row

QWERTY example:
`/?keys=qwertyuiop[]\asdfghjkl%3B%27zxcvbnm,./&fingers=0123344567777012334456770123344567`