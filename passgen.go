package main

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"fmt"
	mrand "math/rand"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app        = kingpin.New("passgen", "A password generate in Go application.")
	noVaildNum = app.Flag("noVaildNum", "No valid Number in a generating password.").Short('N').Bool()
	noVaildSym = app.Flag("noVaildSym", "No valid Symbol in a generating password.").Short('S').Bool()
	noVaildUnr = app.Flag("noVaildUnr", "No valid Unreadable character in a generating password.").Short('R').Bool()
	number     = app.Flag("number", "Number of generating password.").Default("1").Short('n').Int()
	length     = app.Flag("length", "Length of generating password.").Default("31").Short('l').Int()
)

const (
	hom = iota
	inc
	att
	sha
	dol
	par
	hed
	and
	sta
	le1
	ri1
	und
	hyp
	pls
	equ
	le2
	ri2
	le3
	ri3
	bas
	seq
	sec
	dqu
	squ
	lar
	sma
	col
	cam
	que
	sla
	pip
)

const (
	a = iota
	b
	c
	d
	e
	f
	g
	h
	i
	j
	k
	//	l
	m
	n
	o
	p
	q
	r
	s
	t
	u
	v
	w
	x
	y
	z
	A
	B
	C
	D
	E
	F
	G
	H
	I
	J
	K
	L
	M
	N
	//	O
	P
	Q
	R
	S
	T
	U
	V
	W
	X
	Y
	Z
	l
	O
)

type SymbolMap map[int]rune
type AlphaMap map[int]rune
type NumberMap map[int]rune

type passMap struct {
	SymbolMap SymbolMap
	AlphaMap  AlphaMap
	NumberMap NumberMap
}

func newSymbolMap() SymbolMap {
	return SymbolMap{
		hom: '~',
		inc: '!',
		att: '@',
		sha: '#',
		dol: '$',
		par: '%',
		hed: '^',
		and: '&',
		sta: '*',
		le1: '(',
		ri1: ')',
		und: '_',
		hyp: '-',
		pls: '+',
		equ: '=',
		le2: '{',
		ri2: '}',
		le3: '[',
		ri3: ']',
		bas: '¥',
		seq: ':',
		sec: ';',
		dqu: '"',
		squ: '`',
		lar: '>',
		sma: '<',
		col: ',',
		cam: '.',
		que: '?',
		sla: '/',
		//	pip: '|',
	}
}

func newAlphaMap() AlphaMap {
	return AlphaMap{
		a: 'a',
		b: 'b',
		c: 'c',
		d: 'd',
		e: 'e',
		f: 'f',
		g: 'g',
		h: 'h',
		i: 'i',
		j: 'j',
		k: 'k',
		//	l:	'l',
		m: 'm',
		n: 'n',
		o: 'o',
		p: 'p',
		q: 'q',
		r: 'r',
		s: 's',
		t: 't',
		u: 'u',
		v: 'v',
		w: 'q',
		x: 'x',
		y: 'y',
		z: 'z',
		A: 'A',
		B: 'B',
		C: 'C',
		D: 'D',
		E: 'E',
		F: 'F',
		G: 'G',
		H: 'H',
		I: 'I',
		J: 'J',
		K: 'K',
		L: 'L',
		M: 'M',
		N: 'N',
		//	O:	'O',
		P: 'P',
		Q: 'Q',
		R: 'R',
		S: 'S',
		T: 'T',
		U: 'U',
		V: 'V',
		W: 'W',
		X: 'X',
		Y: 'Y',
		Z: 'Z',
	}
}

func newNumberMap() NumberMap {
	return NumberMap{
		0: '0',
		1: '1',
		2: '2',
		3: '3',
		4: '4',
		5: '5',
		6: '6',
		7: '7',
		8: '8',
		9: '9',
	}
}

func newpassMap() *passMap {
	return &passMap{
		SymbolMap: newSymbolMap(),
		AlphaMap:  newAlphaMap(),
		NumberMap: newNumberMap(),
	}
}

func seedInit() (seed int64) {
	err := binary.Read(cryptorand.Reader, binary.LittleEndian, &seed)
	if err != nil {
		panic(err)
	}
	return
}

func numRoll(max int) (dice int) {
	mrand.Seed(seedInit())
	dice = mrand.Intn(max) + 1
	return
}

func (pm *passMap) diceRoll(numroll func(int) int, num int) (str []rune) {
	for count := 0; count < *length; count++ {
		dice := numroll(num)
		var n rune
		var nn int
		if dice <= 2 && !*noVaildNum {
			for {
				nn = numRoll(len(pm.NumberMap))
				if *noVaildUnr && nn == 0 {
					continue
				}
				break
			}
			n = pm.NumberMap[nn]
		} else if dice <= 4 && !*noVaildSym {
			n = pm.SymbolMap[numRoll(len(pm.SymbolMap))]
		} else {
			n = pm.AlphaMap[numRoll(len(pm.AlphaMap))]
		}
		str = append(str, n)
	}
	return
}

func main() {
	app.HelpFlag.Short('h')
	kingpin.MustParse(app.Parse(os.Args[1:]))

	pm := newpassMap()
	if !*noVaildUnr {
		pm.SymbolMap[pip] = '|'
		pm.AlphaMap[l] = 'l'
		pm.AlphaMap[O] = 'O'
	}

	for count := 0; count < *number; count++ {
		fmt.Println(string(pm.diceRoll(numRoll, 10)))
	}
}
