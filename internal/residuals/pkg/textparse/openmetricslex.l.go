// Code generated by golex. DO NOT EDIT.

// Copyright 2018 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package textparse

import (
	"fmt"
)

// Lex is called by the parser generated by "go tool yacc" to obtain each
// token. The method is opened before the matching rules block and closed at
// the end of the file.
func (l *openMetricsLexer) Lex() token {
	if l.i >= len(l.b) {
		return tEOF
	}
	c := l.b[l.i]
	l.start = l.i

yystate0:

	switch yyt := l.state; yyt {
	default:
		panic(fmt.Errorf(`invalid start condition %d`, yyt))
	case 0:	// start condition: INITIAL
		goto yystart1
	case 1:	// start condition: sComment
		goto yystart5
	case 2:	// start condition: sMeta1
		goto yystart25
	case 3:	// start condition: sMeta2
		goto yystart27
	case 4:	// start condition: sLabels
		goto yystart30
	case 5:	// start condition: sLValue
		goto yystart35
	case 6:	// start condition: sValue
		goto yystart39
	case 7:	// start condition: sTimestamp
		goto yystart43
	case 8:	// start condition: sExemplar
		goto yystart50
	case 9:	// start condition: sEValue
		goto yystart55
	case 10:	// start condition: sETimestamp
		goto yystart61
	}

	goto yystate0	// silence unused label error
	goto yystate1	// silence unused label error
yystate1:
	c = l.next()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '#':
		goto yystate2
	case c == ':' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate4
	}

yystate2:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == ' ':
		goto yystate3
	}

yystate3:
	c = l.next()
	goto yyrule1

yystate4:
	c = l.next()
	switch {
	default:
		goto yyrule8
	case c >= '0' && c <= ':' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate4
	}

	goto yystate5	// silence unused label error
yystate5:
	c = l.next()
yystart5:
	switch {
	default:
		goto yyabort
	case c == 'E':
		goto yystate6
	case c == 'H':
		goto yystate10
	case c == 'T':
		goto yystate15
	case c == 'U':
		goto yystate20
	}

yystate6:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'O':
		goto yystate7
	}

yystate7:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'F':
		goto yystate8
	}

yystate8:
	c = l.next()
	switch {
	default:
		goto yyrule5
	case c == '\n':
		goto yystate9
	}

yystate9:
	c = l.next()
	goto yyrule5

yystate10:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'E':
		goto yystate11
	}

yystate11:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'L':
		goto yystate12
	}

yystate12:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'P':
		goto yystate13
	}

yystate13:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == ' ':
		goto yystate14
	}

yystate14:
	c = l.next()
	goto yyrule2

yystate15:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'Y':
		goto yystate16
	}

yystate16:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'P':
		goto yystate17
	}

yystate17:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'E':
		goto yystate18
	}

yystate18:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == ' ':
		goto yystate19
	}

yystate19:
	c = l.next()
	goto yyrule3

yystate20:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'N':
		goto yystate21
	}

yystate21:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'I':
		goto yystate22
	}

yystate22:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'T':
		goto yystate23
	}

yystate23:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == ' ':
		goto yystate24
	}

yystate24:
	c = l.next()
	goto yyrule4

	goto yystate25	// silence unused label error
yystate25:
	c = l.next()
yystart25:
	switch {
	default:
		goto yyabort
	case c == ':' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate26
	}

yystate26:
	c = l.next()
	switch {
	default:
		goto yyrule6
	case c >= '0' && c <= ':' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate26
	}

	goto yystate27	// silence unused label error
yystate27:
	c = l.next()
yystart27:
	switch {
	default:
		goto yyabort
	case c == ' ':
		goto yystate28
	}

yystate28:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '\n':
		goto yystate29
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate28
	}

yystate29:
	c = l.next()
	goto yyrule7

	goto yystate30	// silence unused label error
yystate30:
	c = l.next()
yystart30:
	switch {
	default:
		goto yyabort
	case c == ',':
		goto yystate31
	case c == '=':
		goto yystate32
	case c == '}':
		goto yystate34
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate33
	}

yystate31:
	c = l.next()
	goto yyrule13

yystate32:
	c = l.next()
	goto yyrule12

yystate33:
	c = l.next()
	switch {
	default:
		goto yyrule10
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate33
	}

yystate34:
	c = l.next()
	goto yyrule11

	goto yystate35	// silence unused label error
yystate35:
	c = l.next()
yystart35:
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate36
	}

yystate36:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate37
	case c == '\\':
		goto yystate38
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate36
	}

yystate37:
	c = l.next()
	goto yyrule14

yystate38:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate36
	}

	goto yystate39	// silence unused label error
yystate39:
	c = l.next()
yystart39:
	switch {
	default:
		goto yyabort
	case c == ' ':
		goto yystate40
	case c == '{':
		goto yystate42
	}

yystate40:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate41
	}

yystate41:
	c = l.next()
	switch {
	default:
		goto yyrule15
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate41
	}

yystate42:
	c = l.next()
	goto yyrule9

	goto yystate43	// silence unused label error
yystate43:
	c = l.next()
yystart43:
	switch {
	default:
		goto yyabort
	case c == ' ':
		goto yystate45
	case c == '\n':
		goto yystate44
	}

yystate44:
	c = l.next()
	goto yyrule17

yystate45:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '#':
		goto yystate47
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '\x1f' || c == '!' || c == '"' || c >= '$' && c <= 'ÿ':
		goto yystate46
	}

yystate46:
	c = l.next()
	switch {
	default:
		goto yyrule16
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate46
	}

yystate47:
	c = l.next()
	switch {
	default:
		goto yyrule16
	case c == ' ':
		goto yystate48
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate46
	}

yystate48:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '{':
		goto yystate49
	}

yystate49:
	c = l.next()
	goto yyrule18

	goto yystate50	// silence unused label error
yystate50:
	c = l.next()
yystart50:
	switch {
	default:
		goto yyabort
	case c == ',':
		goto yystate51
	case c == '=':
		goto yystate52
	case c == '}':
		goto yystate54
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate53
	}

yystate51:
	c = l.next()
	goto yyrule23

yystate52:
	c = l.next()
	goto yyrule21

yystate53:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate53
	}

yystate54:
	c = l.next()
	goto yyrule20

	goto yystate55	// silence unused label error
yystate55:
	c = l.next()
yystart55:
	switch {
	default:
		goto yyabort
	case c == ' ':
		goto yystate56
	case c == '"':
		goto yystate58
	}

yystate56:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate57
	}

yystate57:
	c = l.next()
	switch {
	default:
		goto yyrule24
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate57
	}

yystate58:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate59
	case c == '\\':
		goto yystate60
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate58
	}

yystate59:
	c = l.next()
	goto yyrule22

yystate60:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate58
	}

	goto yystate61	// silence unused label error
yystate61:
	c = l.next()
yystart61:
	switch {
	default:
		goto yyabort
	case c == ' ':
		goto yystate63
	case c == '\n':
		goto yystate62
	}

yystate62:
	c = l.next()
	goto yyrule26

yystate63:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate64
	}

yystate64:
	c = l.next()
	switch {
	default:
		goto yyrule25
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate64
	}

yyrule1:	// #{S}
	{
		l.state = sComment
		goto yystate0
	}
yyrule2:	// HELP{S}
	{
		l.state = sMeta1
		return tHelp
		goto yystate0
	}
yyrule3:	// TYPE{S}
	{
		l.state = sMeta1
		return tType
		goto yystate0
	}
yyrule4:	// UNIT{S}
	{
		l.state = sMeta1
		return tUnit
		goto yystate0
	}
yyrule5:	// "EOF"\n?
	{
		l.state = sInit
		return tEofWord
		goto yystate0
	}
yyrule6:	// {M}({M}|{D})*
	{
		l.state = sMeta2
		return tMName
		goto yystate0
	}
yyrule7:	// {S}{C}*\n
	{
		l.state = sInit
		return tText
		goto yystate0
	}
yyrule8:	// {M}({M}|{D})*
	{
		l.state = sValue
		return tMName
		goto yystate0
	}
yyrule9:	// \{
	{
		l.state = sLabels
		return tBraceOpen
		goto yystate0
	}
yyrule10:	// {L}({L}|{D})*
	{
		return tLName
	}
yyrule11:	// \}
	{
		l.state = sValue
		return tBraceClose
		goto yystate0
	}
yyrule12:	// =
	{
		l.state = sLValue
		return tEqual
		goto yystate0
	}
yyrule13:	// ,
	{
		return tComma
	}
yyrule14:	// \"(\\.|[^\\"\n])*\"
	{
		l.state = sLabels
		return tLValue
		goto yystate0
	}
yyrule15:	// {S}[^ \n]+
	{
		l.state = sTimestamp
		return tValue
		goto yystate0
	}
yyrule16:	// {S}[^ \n]+
	{
		return tTimestamp
	}
yyrule17:	// \n
	{
		l.state = sInit
		return tLinebreak
		goto yystate0
	}
yyrule18:	// {S}#{S}\{
	{
		l.state = sExemplar
		return tComment
		goto yystate0
	}
yyrule19:	// {L}({L}|{D})*
	{
		return tLName
	}
yyrule20:	// \}
	{
		l.state = sEValue
		return tBraceClose
		goto yystate0
	}
yyrule21:	// =
	{
		l.state = sEValue
		return tEqual
		goto yystate0
	}
yyrule22:	// \"(\\.|[^\\"\n])*\"
	{
		l.state = sExemplar
		return tLValue
		goto yystate0
	}
yyrule23:	// ,
	{
		return tComma
	}
yyrule24:	// {S}[^ \n]+
	{
		l.state = sETimestamp
		return tValue
		goto yystate0
	}
yyrule25:	// {S}[^ \n]+
	{
		return tTimestamp
	}
yyrule26:	// \n
	{
		l.state = sInit
		return tLinebreak
		goto yystate0
	}
	panic("unreachable")

	goto yyabort	// silence unused label error

yyabort:	// no lexem recognized

	return tInvalid
}
