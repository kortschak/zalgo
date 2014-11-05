// Copyright ©2013 Dan Kortschak. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package zalgo implements a zalgo text io.Writer.
//
// In addition to providing the basic Zalgo text transformation, two default
// transformers are provided:
//  Corruption - writes to os.Stdout
//  Perdition  - writes to os.Stderr
// These defaults are set with what are reasonable defaults (whatever that might
// mean in this context).
//
// So for example:
//	fmt.Fprintln(zalgo.Corruption, "ZALGO!")
// will print something like the following to stdout:
//  Z̗̻̭̫A̝̼̼̯ͩͅL̠͍̮̝ͪG͓͚̣͙͈ͪO͕̍ͨ!̭̒
package zalgo

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"math/rand"
	"os"
	"time"
	"unicode/utf8"
)

var (
	up = []rune{
		0x030d, 0x030e, 0x0304, 0x0305,
		0x033f, 0x0311, 0x0306, 0x0310,
		0x0352, 0x0357, 0x0351, 0x0307,
		0x0308, 0x030a, 0x0342, 0x0343,
		0x0344, 0x034a, 0x034b, 0x034c,
		0x0303, 0x0302, 0x030c, 0x0350,
		0x0300, 0x0301, 0x030b, 0x030f,
		0x0312, 0x0313, 0x0314, 0x033d,
		0x0309, 0x0363, 0x0364, 0x0365,
		0x0366, 0x0367, 0x0368, 0x0369,
		0x036a, 0x036b, 0x036c, 0x036d,
		0x036e, 0x036f, 0x033e, 0x035b,
		0x0346, 0x031a,
	}

	down = []rune{
		0x0316, 0x0317, 0x0318, 0x0319,
		0x031c, 0x031d, 0x031e, 0x031f,
		0x0320, 0x0324, 0x0325, 0x0326,
		0x0329, 0x032a, 0x032b, 0x032c,
		0x032d, 0x032e, 0x032f, 0x0330,
		0x0331, 0x0332, 0x0333, 0x0339,
		0x033a, 0x033b, 0x033c, 0x0345,
		0x0347, 0x0348, 0x0349, 0x034d,
		0x034e, 0x0353, 0x0354, 0x0355,
		0x0356, 0x0359, 0x035a, 0x0323,
	}

	middle = []rune{
		0x0315, 0x031b, 0x0340, 0x0341,
		0x0358, 0x0321, 0x0322, 0x0327,
		0x0328, 0x0334, 0x0335, 0x0336,
		0x034f, 0x035c, 0x035d, 0x035e,
		0x035f, 0x0360, 0x0362, 0x0338,
		0x0337, 0x0361, 0x0489,
	}

	zalgoChars = func() map[rune]struct{} {
		zc := make(map[rune]struct{})
		for _, z := range up {
			zc[z] = struct{}{}
		}
		for _, z := range down {
			zc[z] = struct{}{}
		}
		for _, z := range middle {
			zc[z] = struct{}{}
		}
		return zc
	}()
)

var (
	Corruption = NewCorrupter(os.Stdout)
	Perdition  = NewCorrupter(os.Stderr)
)

var rnd *rand.Rand

func init() {
	Corruption.Up = complex(2, 0.3)
	Corruption.Middle = complex(1, 0.1)
	Corruption.Down = complex(5, 0.7)

	Perdition.Up = complex(2, 0.3)
	Perdition.Middle = complex(1, 0.1)
	Perdition.Down = complex(5, 0.7)

	rnd = rand.New(rand.NewSource(time.Now().Unix()))
}

// Zalgo alters a Corrupter based in the number of bytes written by the Corrupter
// and the current rune. If Zalgo chooses, the rune is spared.
type Zalgo func(int, rune, *Corrupter) (chosen bool)

// Corrupter implements io.Writer transforming plain text to zalgo text.
type Corrupter struct {
	Up     complex128
	Middle complex128
	Down   complex128
	Zalgo
	w io.Writer
	n int
	b []byte
}

// NewCorrupter returns a new Corrupter that writes to w.
func NewCorrupter(w io.Writer) *Corrupter {
	return &Corrupter{w: w, b: make([]byte, utf8.MaxRune)}
}

// Write writes the byte slice p to the Corrupter's underlying io.Writer
// returning the number of bytes written and any error that occurred during
// the write operation.
func (z *Corrupter) Write(p []byte) (n int, err error) {
	x := xml.NewDecoder(bytes.NewReader(p))
	_, err = x.Token()
	if err == nil {
		err = x.Skip()
	}
	if err == nil {
		return 0, errors.New("zalgo: cannot consume XML")
	}

	h := html.NewTokenizer(bytes.NewReader(p))
	var f bool
L:
	for {
		t := h.Next()
		switch t {
		case html.ErrorToken:
			if h.Err() == io.EOF {
				break L
			}
		case html.StartTagToken, html.EndTagToken, html.SelfClosingTagToken, html.DoctypeToken:
			f = true
		}
	}
	if f {
		return 0, errors.New("zalgo: cannot consume HTML")
	}

	var _n int
	n = z.n
	defer func() {
		n = z.n - n
	}()
	for _, r := range string(p) {
		if _, ok := zalgoChars[r]; ok {
			continue
		}
		z.b = z.b[:utf8.RuneLen(r)]
		utf8.EncodeRune(z.b, r)
		_n, err = z.w.Write(z.b)
		z.n += _n
		if err != nil {
			return
		}
		if z.Zalgo != nil && z.Zalgo(z.n, r, z) {
			continue
		}
		for i := real(z.Up); i > 0; i-- {
			if rnd.Float64() < imag(z.Up) {
				_n, err = fmt.Fprintf(z.w, "%c", up[rnd.Intn(len(up))])
				z.n += _n
				if err != nil {
					return
				}
			}
		}
		for i := real(z.Middle); i > 0; i-- {
			if rnd.Float64() < imag(z.Middle) {
				_n, err = fmt.Fprintf(z.w, "%c", middle[rnd.Intn(len(middle))])
				z.n += _n
				if err != nil {
					return
				}
			}
		}
		for i := real(z.Down); i > 0; i-- {
			if rnd.Float64() < imag(z.Down) {
				_n, err = fmt.Fprintf(z.w, "%c", down[rnd.Intn(len(down))])
				z.n += _n
				if err != nil {
					return
				}
			}
		}
	}
	return
}
