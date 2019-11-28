// Lute - A structured markdown engine.
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under the Mulan PSL v1.
// You can use this software according to the terms and conditions of the Mulan PSL v1.
// You may obtain a copy of Mulan PSL v1 at:
//     http://license.coscl.org.cn/MulanPSL
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
// PURPOSE.
// See the Mulan PSL v1 for more details.

package test

import (
	"github.com/b3log/lute/html"
	"github.com/b3log/lute/html/atom"
	"strings"
	"testing"
)

func TestHTMLParse(t *testing.T) {
	reader := strings.NewReader("<p>foo</p>")
	htmlRoot := &html.Node{Type: html.ElementNode}
	htmlNodes, err := html.ParseFragment(reader, htmlRoot)
	if nil != err || 1 != len(htmlNodes) {
		t.Fatal("test HTML parse failed")
	}

	if atom.P != htmlNodes[0].DataAtom {
		t.Fatal("test HTML parse failed")
	}
}

func TestHTMLParse0(t *testing.T) {
	reader := strings.NewReader("<p>&& &nbsp;</p>")
	htmlRoot := &html.Node{Type: html.ElementNode}
	htmlNodes, err := html.ParseFragment(reader, htmlRoot)
	if nil != err || 1 != len(htmlNodes) {
		t.Fatal("test HTML parse failed")
	}

	if atom.P != htmlNodes[0].DataAtom {
		t.Fatal("test HTML parse failed")
	}
}