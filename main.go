package main

import (
	"fmt"
	"strings"

	_ "github.com/gohugoio/hugo/helpers"

	"github.com/russross/blackfriday/v2"
	_ "github.com/russross/blackfriday/v2"
	// "github.com/KinoshitaYstr/blackfriday/v2"
	// _ "github.com/KinoshitaYstr/blackfriday/v2"
)

// https://github.com/russross/blackfriday/blob/acedacffef10e8f1943455d4dd8bb8f4dfe8d0f8/block_test.go#L968
func goHTMLByTestPerfomattedHTML() {
	testData1 := `
<!-- テストデータだよ！！！ -->

<img data-action="display" style="display: none;"/>

<img src="https://qiita-image-store.s3.amazonaws.com/0/45617/015bd058-7ea0-e6a5-b9cb-36a4fb38e59c.png">

ネコ（猫）は、狭義には食肉目ネコ科ネコ属に分類されるリビアヤマネコ（ヨーロッパヤマネコ）が家畜化されたイエネコ（家猫、Felis silvestris catus）に対する通称である。

イヌ（犬）と並ぶ代表的なペットとして日本を含め世界中で広く飼われている。>_<

<!-- <img data-action="erase" style="display: none;"/> -->

<!-- <img data-action="hidden" style="display: none;"/> -->

<!-- <img data-action="display" style="display: none;"/> -->

<img src="https://qiita-image-store.s3.amazonaws.com/0/45617/015bd058-7ea0-e6a5-b9cb-36a4fb38e59c.png">

より広義には、ヤマネコやネコ科動物全般を指すこともある

by wikipedia

<img data-action="erase" style="display: none;"/>

<img data-action="hidden" style="display: none;"/>

![Qiita](https://qiita-image-store.s3.amazonaws.com/0/45617/015bd058-7ea0-e6a5-b9cb-36a4fb38e59c.png "Qiita")

<div class="aa"
id="ok">
sasasadasda
</div>

<form method="post" action="/">
<select name="example">
<option value="サンプル1">サンプル1</option>
<option value="サンプル2">サンプル2</option>
<option value="サンプル3">サンプル3</option>
</select>
</form>

	`

	// https://github.com/russross/blackfriday/blob/acedacffef10e8f1943455d4dd8bb8f4dfe8d0f8/helpers_test.go#L66
	var extentions blackfriday.Extensions
	extentions = 0
	var HTMLFlag blackfriday.HTMLFlags
	HTMLFlag = blackfriday.UseXHTML
	HTMLFlag |= blackfriday.Smartypants
	HTMLFlag |= blackfriday.SmartypantsFractions
	HTMLFlag |= blackfriday.SmartypantsDashes
	HTMLFlag |= blackfriday.SmartypantsLatexDashes
	// HTMLFlag = blackfriday.CommonHTMLFlags
	// 動いてないかも？？？？
	// よくわからぬ
	// https://github.com/russross/blackfriday/blob/v2/inline_test.go
	// 使い方
	var referenceOverride blackfriday.ReferenceOverrideFunc

	var HTMLRendererParams blackfriday.HTMLRendererParameters
	HTMLRendererParams.Flags = HTMLFlag
	renderer := blackfriday.NewHTMLRenderer(HTMLRendererParams)
	fmt.Println(string(blackfriday.Run([]byte(testData1), blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(extentions), blackfriday.WithRefOverride(referenceOverride))))
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(string(blackfriday.Run([]byte("<div>Oha!</div>"), blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(extentions), blackfriday.WithRefOverride(referenceOverride))))
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	for i, str := range strings.Split(testData1, "\n") {
		// for (str[0] == ' ' || str[0] == '\t') && len(str) >= 2 {
		// 	str = str[1:]
		// }
		fmt.Println(i)
		fmt.Println(str)
		fmt.Println(string(blackfriday.Run([]byte(str), blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(extentions), blackfriday.WithRefOverride(referenceOverride))))
	}

	// 最初にタブあると、その文がコードで表記されちゃう
	// 	testData2 := `
	// <!-- テストデータだよ！！！ -->
	// # aaaaaaaaaa
	// <div class="test">
	// <p>aaa</p>
	// <img data-action="display" style="display: none;"/>
	// <img src="https://qiita-image-store.s3.amazonaws.com/0/45617/015bd058-7ea0-e6a5-b9cb-36a4fb38e59c.png">
	// ネコ（猫）は、狭義には食肉目ネコ科ネコ属に分類されるリビアヤマネコ（ヨーロッパヤマネコ）が家畜化されたイエネコ（家猫、Felis silvestris catus）に対する通称である。
	// イヌ（犬）と並ぶ代表的なペットとして日本を含め世界中で広く飼われている。>_<
	// <!-- <img data-action="erase" style="display: none;"/> -->
	// <!-- <img data-action="hidden" style="display: none;"/> -->
	// <!-- <img data-action="display" style="display: none;"/> -->
	// <img src="https://qiita-image-store.s3.amazonaws.com/0/45617/015bd058-7ea0-e6a5-b9cb-36a4fb38e59c.png">
	// より広義には、ヤマネコやネコ科動物全般を指すこともある
	// by wikipedia
	// <img data-action="erase" style="display: none;"/>
	// <img data-action="hidden" style="display: none;"/>
	// ![Qiita](https://qiita-image-store.s3.amazonaws.com/0/45617/015bd058-7ea0-e6a5-b9cb-36a4fb38e59c.png "Qiita")
	// <div class="aa"
	// id="ok">
	// sasasadasda
	// </div>
	// <form method="post" action="/">
	// <select name="example">
	// <option value="サンプル1">サンプル1</option>
	// <option value="サンプル2">サンプル2</option>
	// <option value="サンプル3">サンプル3</option>
	// </select>
	// </form>
	// <form method="post" action="/">
	// <select name="example">
	// <option value="サンプル1">サンプル1</option>
	// <option value="サンプル2">サンプル2</option>
	// <option value="サンプル3">サンプル3</option>
	// </select>
	// </form>
	// </div>`

	testData2 := `
Simple block on one line:

<div>foo</div>

And nested without indentation:

<img data-action="display" style="display: none;"/>

![Qiita](https://qiita-image-store.s3.amazonaws.com/0/45617/015bd058-7ea0-e6a5-b9cb-36a4fb38e59c.png "Qiita")

ネコ（猫）は、狭義には食肉目ネコ科ネコ属に分類されるリビアヤマネコ（ヨーロッパヤマネコ）が家畜化されたイエネコ（家猫、Felis silvestris catus）に対する通称である。

イヌ（犬）と並ぶ代表的なペットとして日本を含め世界中で広く飼われている。>_<

<form method="post" action="/">

<select name="example">

<option value="サンプル1">サンプル1</option>

<option value="サンプル2">サンプル2</option>

<option value="サンプル3">サンプル3</option>

</select>

</form>

<div>
<div>
<div>
foo
</div>
<div style=">"/>
</div>
<div>bar</div>
</div>
`

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	for i, str := range strings.Split(testData2, "\n") {
		fmt.Println(i)
		fmt.Println(str)
		fmt.Println(string(blackfriday.Run([]byte(str), blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(extentions), blackfriday.WithRefOverride(referenceOverride))))
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(string(blackfriday.Run([]byte(testData2), blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(extentions), blackfriday.WithRefOverride(referenceOverride))))
	// fmt.Println(string(blackfriday.Run([]byte(testData2), blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(extentions))))
	// fmt.Println(string(blackfriday.Run([]byte("<div>"+testData2+"</div>\n"), blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(extentions), blackfriday.WithRefOverride(referenceOverride))))
	// fmt.Println(string(blackfriday.Run([]byte("<div>(=^・・^=)</div>\n"), blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(extentions), blackfriday.WithRefOverride(referenceOverride))))
	// fmt.Println(string(blackfriday.Run([]byte(`<img src="https://qiita-image-store.s3.amazonaws.com/0/45617/015bd058-7ea0-e6a5-b9cb-36a4fb38e59c.png">`+"\n"), blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(extentions), blackfriday.WithRefOverride(referenceOverride))))
	// fmt.Println(string(blackfriday.Run([]byte(`<div><img src="https://qiita-image-store.s3.amazonaws.com/0/45617/015bd058-7ea0-e6a5-b9cb-36a4fb38e59c.png"></div>`+"\n"), blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(extentions), blackfriday.WithRefOverride(referenceOverride))))
	// fmt.Println(string(blackfriday.Run([]byte(`![Qiita](https://qiita-image-store.s3.amazonaws.com/0/45617/015bd058-7ea0-e6a5-b9cb-36a4fb38e59c.png "Qiita")`+"\n"), blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(extentions), blackfriday.WithRefOverride(referenceOverride))))
	// fmt.Println(string(blackfriday.Run([]byte(`<div>![Qiita](https://qiita-image-store.s3.amazonaws.com/0/45617/015bd058-7ea0-e6a5-b9cb-36a4fb38e59c.png "Qiita")</div>`+"\n"), blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(extentions), blackfriday.WithRefOverride(referenceOverride))))
}

func main() {
	// fmt.Println("Hello World")

	// fmt.Println(testData1)

	// result1 := string(blackfriday.Run([]byte(testData1)))

	// fmt.Println(result1)
	goHTMLByTestPerfomattedHTML()
}
