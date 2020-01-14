package main

import (
	"fmt"
	"strings"

	student ".."
	"github.com/01-edu/z01"
)

func PrintList(l *student.NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func PrintList2(l *student.List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func listPushBack(l *student.NodeI, data int) *student.NodeI {
	n := &student.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func Reversee(a, b string) int {
	return strings.Compare(b, a)
}

func main() {
	fmt.Println(student.Atoi(""))
	fmt.Println(student.Atoi("-"))
	fmt.Println(student.Atoi("--123"))
	fmt.Println(student.Atoi("1"))
	fmt.Println(student.Atoi("-3"))
	fmt.Println(student.Atoi("8292"))
	fmt.Println(student.Atoi("9223372036854775807"))
	fmt.Println(student.Atoi("-9223372036854775808"))

	fmt.Println()

	fmt.Println(student.RecursivePower(-7, -2))
	fmt.Println(student.RecursivePower(-8, -7))
	fmt.Println(student.RecursivePower(4, 8))
	fmt.Println(student.RecursivePower(1, 3))
	fmt.Println(student.RecursivePower(-1, 1))
	fmt.Println(student.RecursivePower(-6, 5))

	fmt.Println()

	for i := 0; i <= 9; i++ {
		student.PrintCombN(i)
	}

	fmt.Println()

	student.PrintNbrBase(919617, "01")
	z01.PrintRune('\n')
	student.PrintNbrBase(753639, "CHOUMIisDAcat!")
	z01.PrintRune('\n')
	student.PrintNbrBase(-174336, "CHOUMIisDAcat!")
	z01.PrintRune('\n')
	student.PrintNbrBase(-661165, "1")
	z01.PrintRune('\n')
	student.PrintNbrBase(-861737, "Zone01Zone01")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	student.PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "-ab")
	z01.PrintRune('\n')
	student.PrintNbrBase(-9223372036854775808, "0123456789")
	z01.PrintRune('\n')

	fmt.Println()

	fmt.Println(student.AtoiBase("bcbbbbaab", "abc"))
	fmt.Println(student.AtoiBase("0001", "01"))
	fmt.Println(student.AtoiBase("00", "01"))
	fmt.Println(student.AtoiBase("saDt!I!sI", "CHOUMIisDAcat!"))
	fmt.Println(student.AtoiBase("AAho?Ao", "WhoAmI?"))
	fmt.Println(student.AtoiBase("thisinputshouldnotmatter", "abca?"))
	fmt.Println(student.AtoiBase("125", "0123456789"))
	fmt.Println(student.AtoiBase("uoi", "choumi"))
	fmt.Println(student.AtoiBase("bbbbbab", "-ab"))

	fmt.Println()

	str := "The earliest foundations of what would become computer science predate the invention of the modern digital computer"
	fmt.Println(student.SplitWhiteSpaces(str))
	str = "Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity,"
	fmt.Println(student.SplitWhiteSpaces(str))
	str = "aiding in computations such as multiplication and division ."
	fmt.Println(student.SplitWhiteSpaces(str))
	str = "Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment."
	fmt.Println(student.SplitWhiteSpaces(str))
	str = "Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]"
	fmt.Println(student.SplitWhiteSpaces(str))
	str = " In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,"
	fmt.Println(student.SplitWhiteSpaces(str))

	fmt.Println()

	str = "|=choumi=|which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished,"
	fmt.Println(student.Split(str, "|=choumi=|"))
	str = "!==!which!==!was!==!making!==!all!==!kinds!==!of!==!punched!==!card!==!equipment!==!and!==!was!==!also!==!in!==!the!==!calculator!==!business[10]!==!to!==!develop!==!his!==!giant!==!programmable!==!calculator,"
	fmt.Println(student.Split(str, "!==!"))
	str = "AFJCharlesAFJBabbageAFJstartedAFJtheAFJdesignAFJofAFJtheAFJfirstAFJautomaticAFJmechanicalAFJcalculator,"
	fmt.Println(student.Split(str, "AFJ"))
	str = "<<==123==>>In<<==123==>>1820,<<==123==>>Thomas<<==123==>>de<<==123==>>Colmar<<==123==>>launched<<==123==>>the<<==123==>>mechanical<<==123==>>calculator<<==123==>>industry[note<<==123==>>1]<<==123==>>when<<==123==>>he<<==123==>>released<<==123==>>his<<==123==>>simplified<<==123==>>arithmometer,"
	fmt.Println(student.Split(str, "<<==123==>>"))

	fmt.Println()
	result := student.ConvertBase("4506C", "0123456789ABCDEF", "choumi")
	fmt.Println(result)
	result = student.ConvertBase("babcbaaaaabcb", "abc", "0123456789ABCDEF")
	fmt.Println(result)
	result = student.ConvertBase("256850", "0123456789", "01")
	fmt.Println(result)
	result = student.ConvertBase("uuhoumo", "choumi", "choumi")
	fmt.Println(result)
	result = student.ConvertBase("683241", "0123456789", "0123456789")
	fmt.Println(result)

	fmt.Println()

	result2 := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	student.AdvancedSortWordArr(result2, student.Compare)
	fmt.Println(result2)
	result2 = []string{"The", "earliest", "computing", "device", "undoubtedly", "consisted", "of", "the", "five", "fingers", "of", "each", "hand"}
	student.AdvancedSortWordArr(result2, student.Compare)
	fmt.Println(result2)

	result2 = []string{"The", "word", "digital", "comesfrom", "\"digits\"", "or", "fingers"}
	student.AdvancedSortWordArr(result2, student.Compare)
	fmt.Println(result2)

	result2 = []string{"The", "earliest", "computing", "device", "undoubtedly", "consisted", "of", "the", "five", "fingers", "of", "each", "hand"}
	student.AdvancedSortWordArr(result2, Reversee)
	fmt.Println(result2)
	result2 = []string{"\"digits\"", "The", "comesfrom", "digital", "fingers", "or", "word"}
	student.AdvancedSortWordArr(result2, Reversee)
	fmt.Println(result2)

	result2 = []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	student.AdvancedSortWordArr(result2, Reversee)
	fmt.Println(result2)

	fmt.Println()
	nbits := student.ActiveBits(15)
	fmt.Println(nbits)
	nbits = student.ActiveBits(17)
	fmt.Println(nbits)
	nbits = student.ActiveBits(4)
	fmt.Println(nbits)
	nbits = student.ActiveBits(11)
	fmt.Println(nbits)
	nbits = student.ActiveBits(9)
	fmt.Println(nbits)
	nbits = student.ActiveBits(12)
	fmt.Println(nbits)
	nbits = student.ActiveBits(2)
	fmt.Println(nbits)

	fmt.Println()
	var link *student.NodeI

	link = listPushBack(link, 0)
	link = listPushBack(link, 1)
	link = listPushBack(link, 2)
	link = listPushBack(link, 3)
	link = listPushBack(link, 4)
	link = listPushBack(link, 5)
	link = listPushBack(link, 24)
	link = listPushBack(link, 25)
	link = listPushBack(link, 54)

	PrintList(link)

	link = student.SortListInsert(link, 33)
	PrintList(link)

	fmt.Println()

	var linkks *student.NodeI

	linkks = listPushBack(link, 0)
	linkks = listPushBack(link, 2)
	linkks = listPushBack(link, 18)
	linkks = listPushBack(link, 33)
	linkks = listPushBack(link, 37)
	linkks = listPushBack(link, 37)
	linkks = listPushBack(link, 39)
	linkks = listPushBack(link, 52)
	linkks = listPushBack(link, 53)
	linkks = listPushBack(link, 57)

	PrintList(linkks)

	linkks = student.SortListInsert(linkks, 53)
	PrintList(linkks)

	fmt.Println()
	var link2 *student.NodeI
	var link3 *student.NodeI

	PrintList(student.SortedListMerge(link3, link2))

	link2 = listPushBack(link2, 2)
	link2 = listPushBack(link2, 2)
	link2 = listPushBack(link2, 4)
	link2 = listPushBack(link2, 9)
	link2 = listPushBack(link2, 12)
	link2 = listPushBack(link2, 12)
	link2 = listPushBack(link2, 19)
	link2 = listPushBack(link2, 20)

	PrintList(student.SortedListMerge(link3, link2))

	link3 = listPushBack(link3, -2)
	link3 = listPushBack(link3, 9)

	PrintList(student.SortedListMerge(link3, link2))

	fmt.Println()
	link4 := &student.List{}
	link5 := &student.List{}

	fmt.Println("----normal state----")
	PrintList2(link5)
	student.ListRemoveIf(link5, 1)
	fmt.Println("------answer-----")
	PrintList2(link5)
	fmt.Println()

	fmt.Println("----normal state----")
	student.ListPushBack(link4, 98)
	student.ListPushBack(link4, 98)
	student.ListPushBack(link4, 33)
	student.ListPushBack(link4, 34)
	student.ListPushBack(link4, 33)
	student.ListPushBack(link4, 34)
	student.ListPushBack(link4, 33)
	student.ListPushBack(link4, 89)
	student.ListPushBack(link4, 33)
	PrintList2(link4)

	student.ListRemoveIf(link4, 34)
	fmt.Println("------answer-----")
	PrintList2(link4)

	fmt.Println()
	root := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	node := student.BTreeSearchItem(root, "1")
	replacement := &student.TreeNode{Data: "3"}
	root = student.BTreeTransplant(root, node, replacement)
	student.BTreeApplyInorder(root, fmt.Println)

	fmt.Println()
	root = &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	student.BTreeApplyByLevel(root, fmt.Println)

	fmt.Println()
	root = &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	node = student.BTreeSearchItem(root, "4")
	fmt.Println("Before delete:")
	student.BTreeApplyInorder(root, fmt.Println)
	root = student.BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	student.BTreeApplyInorder(root, fmt.Println)
}
