package main

import "testing"

// func TestLongestPalindrome(t *testing.T) {
// 	got := longestPalindrome("abccccdd")
// 	want := 7

// 	if got != want {
// 		t.Errorf("got %v, wanted %v", got, want)
// 	}
// }

// func TestLongestPalindrome2(t *testing.T) {
// 	got := longestPalindrome("a")
// 	want := 1

// 	if got != want {
// 		t.Errorf("got %v, wanted %v", got, want)
// 	}
// }

func TestLongestPalindrome3(t *testing.T) {
	input := "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"
	got := longestPalindrome(input)
	want := 983

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
