// The MIT License (MIT)
//
// Copyright (c) 2017 Sheikh Humaid AlQassimi
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var fileName, final, name string
	fmt.Print("Enter file name (G1M): ")
	fmt.Scanln(&fileName)
	f, err := os.Open("./" + fileName)
	if err == nil {
		defer f.Close()
		buf := make([]byte, 1)
		n, berr := f.Read(buf)
		i := 0
		for berr == nil {
			i++
			if (i >= 61) && (i <= 68) {
				name = name + identifyByteToString(buf[0:n])
			}
			if i > 86 {
				final = final + identifyByteToString(buf[0:n])
			}
			n, berr = f.Read(buf)
		}
	} else {
		fmt.Println("Invalid file!")
		return
	}
	fmt.Println("Program name: " + name)
	fmt.Println(identifyAllCode(final))
}

func identifyAllCode(code string) string {
	code = strings.Replace(code, "[U+00F7][U+0000]", "If ", -1)
	code = strings.Replace(code, "[U+00F7][U+0001]", "Then ", -1)
	code = strings.Replace(code, "[U+00F7][U+0002]", "Else ", -1)
	code = strings.Replace(code, "[U+00F7][U+0003]", "IfEnd", -1)
	code = strings.Replace(code, "[U+00F7][U+0004]", "For ", -1)
	code = strings.Replace(code, "[U+00F7][U+0005]", " To  ", -1)
	code = strings.Replace(code, "[U+00F7][U+0006]", " Step ", -1)
	code = strings.Replace(code, "[U+00F7][U+0007]", "Next ", -1)
	code = strings.Replace(code, "[U+00F7][U+0008]", "While ", -1)
	code = strings.Replace(code, "[U+00F7][U+0009]", "WhileEnd", -1)
	code = strings.Replace(code, "[U+00F7][U+0010]", "Locate ", -1)
	code = strings.Replace(code, "[U+00F7][U+000A]", "Do", -1)
	code = strings.Replace(code, "[U+00F7][U+000B]", "LpWhile ", -1)
	code = strings.Replace(code, "[U+00F7][U+000C]", "Return", -1)
	code = strings.Replace(code, "[U+00F7][U+000D]", "Break", -1)
	code = strings.Replace(code, "[U+00F7][U+000E]", "Stop", -1)
	code = strings.Replace(code, "[U+00F7][U+009E]", "Menu ", -1)
	code = strings.Replace(code, "[U+00F7][U+0018]", "ClrText", -1)
	code = strings.Replace(code, "[U+00F7][U+0019]", "ClrGraph", -1)
	code = strings.Replace(code, "[U+00F7][U+001A]", "ClrList ", -1)
	code = strings.Replace(code, "[U+00F9][U+001E]", "ClrMat ", -1)
	code = strings.Replace(code, "[U+00F9][U+003E]", "ClrVct ", -1)
	code = strings.Replace(code, "[U+00F9][U+008C]", "BinomialPD(", -1)
	code = strings.Replace(code, "[U+00F9][U+008D]", "BinomialCD(", -1)
	code = strings.Replace(code, "[U+007F][U+008F]", "GetKey", -1)

	code = strings.Replace(code, "[U+00E6][U+0097]", "🡥", -1)
	code = strings.Replace(code, "[U+00E6][U+0096]", "🡤", -1)
	code = strings.Replace(code, "[U+00E6][U+0092]", "🡡", -1)
	code = strings.Replace(code, "[U+00E6][U+00A4]", "●", -1)

	code = strings.Replace(code, "[U+00E2]", "Lbl ", -1)
	code = strings.Replace(code, "[U+00EC]", "Goto ", -1)
	code = strings.Replace(code, "[U+00E9]", "Isz", -1)
	code = strings.Replace(code, "[U+00E8]", "Dsz", -1)
	code = strings.Replace(code, "[U+00ED]", "Prog", -1)
	code = strings.Replace(code, "[U+0086]", "√", -1)
	code = strings.Replace(code, "[U+00B8]", "xROOT", -1)
	code = strings.Replace(code, "[U+008B]", "²", -1)
	code = strings.Replace(code, "[U+0095]", "log ", -1)
	code = strings.Replace(code, "[U+0085]", "ln ", -1)
	code = strings.Replace(code, "[U+0081]", "sin ", -1)
	code = strings.Replace(code, "[U+00DB]", "Rad", -1)
	code = strings.Replace(code, "[U+00DA]", "Deg", -1)
	code = strings.Replace(code, "[U+00BB]", "ꜚ", -1)

	code = strings.Replace(code, "[U+0082]", "cos ", -1)
	code = strings.Replace(code, "[U+0092]", "[cos^-1] ", -1)
	code = strings.Replace(code, "[U+0083]", "tan ", -1)
	code = strings.Replace(code, "[U+0028]", "(", -1)
	code = strings.Replace(code, "[U+0029]", ")", -1)
	code = strings.Replace(code, "[U+002C]", ",", -1)
	code = strings.Replace(code, "[U+00A9]", "×", -1)
	code = strings.Replace(code, "[U+00B9]", "÷", -1)
	code = strings.Replace(code, "[U+0089]", "+", -1)
	code = strings.Replace(code, "[U+0099]", "-", -1)
	code = strings.Replace(code, "[U+0011]", "≠", -1)
	code = strings.Replace(code, "[U+0087]", "‒", -1) // (-)
	code = strings.Replace(code, "[U+007F]P", "[i]", -1)
	code = strings.Replace(code, "[U+00A5]", "[e^]", -1)
	code = strings.Replace(code, "[U+0013]", "⇒", -1)
	code = strings.Replace(code, "[U+00C2]", "x̄", -1)
	code = strings.Replace(code, "[U+00E6]Q", "δ", -1)
	code = strings.Replace(code, "[U+0010]", "≤", -1)
	code = strings.Replace(code, "[U+0012]", "≤", -1)

	code = strings.Replace(code, "[U+00A8]", "^", -1)
	code = strings.Replace(code, "[U+000F]", "ϵ", -1)
	code = strings.Replace(code, "[U+00D0]", "π", -1)
	code = strings.Replace(code, "[U+00CD]", "ř", -1)
	code = strings.Replace(code, "[U+00CE]", "θ", -1)
	code = strings.Replace(code, "[U+000E]", "->", -1)
	code = strings.Replace(code, "[U+000C]", "𐏐\n", -1)
	code = strings.Replace(code, "[U+000D]", "\n", -1)

	code = strings.Replace(code, "[U+00F7]M", "", -1) // End of file
	code = strings.Replace(code, "[U+0000]IM", "", -1)
	code = strings.Replace(code, "[U+0000]", "", -1)
	code = strings.Replace(code, "[U+0020]", " ", -1)
	code = strings.Replace(code, "[U+00E7]", "[SML]", -1)
	return code
}

func identifyByteToString(sym []byte) string {
	unicode := fmt.Sprintf("%U", sym)
	var validChar = regexp.MustCompile(`[a-zA-Z0-9\.\?\=\"\:\-\[\]\|]`)
	if validChar.MatchString(fmt.Sprintf("%s", sym)) {
		return fmt.Sprintf("%s", sym)
	}
	return unicode
}
