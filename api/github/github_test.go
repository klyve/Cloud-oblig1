package githubapi

import "testing"

func TestJsonPrimaryParsing(t *testing.T) {
	file := "./json/primary.json"
	resp := ReadPrimaryFile(file)
	if resp.Name != "git" {
		t.Error("Expected resp name to be git got", resp.Name)
	}
	if resp.Owner.Login != "git" {
		t.Error("Expected owner to be git got ", resp.Owner)
	}
}

func TestReadLanguagesFile(t *testing.T) {
	file := "./json/languages.json"
	match := make([]string, 17)
	match[0] = "C"
	match[1] = "Shell"
	match[2] = "Perl"
	match[3] = "Tcl"
	match[4] = "Python"
	match[5] = "C++"
	match[6] = "Makefile"
	match[7] = "Emacs Lisp"
	match[8] = "JavaScript"
	match[9] = "M4"
	match[10] = "Roff"
	match[11] = "Perl 6"
	match[12] = "Go"
	match[13] = "CSS"
	match[14] = "PHP"
	match[15] = "Assembly"
	match[16] = "Objective-C"

	resp := ReadLanguagesFile(file)

	for i := range match {
		if resp.Languages[i] != match[i] {
			t.Fail()
		}
	}
}

func TestJsonContributorsParsing(t *testing.T) {
	file := "./json/contributors.json"
	resp := ReadContributorsFile(file)
	if resp.Contributions != 18497 {
		t.Error("Expected resp name to be 18497 got", resp.Contributions)
	}
	if resp.Login != "gitster" {
		t.Error("Expected Committer to be gitster got ", resp.Login)
	}
}

func TestCreateErrorCode(t *testing.T) {
	code := CreateErrorCode(400, "Hello world")
	if code.Code != 400 {
		t.Error("Error code is not 400")
	}
	if code.Message != "Hello world" {
		t.Error("Message is not correct")
	}
}

func TestCombineJson(t *testing.T) {
	primary := ReadPrimaryFile("./json/primary.json")
	language := ReadLanguagesFile("./json/languages.json")
	contrib := ReadContributorsFile("./json/contributors.json")
	match := make([]string, 17)
	match[0] = "C"
	match[1] = "Shell"
	match[2] = "Perl"
	match[3] = "Tcl"
	match[4] = "Python"
	match[5] = "C++"
	match[6] = "Makefile"
	match[7] = "Emacs Lisp"
	match[8] = "JavaScript"
	match[9] = "M4"
	match[10] = "Roff"
	match[11] = "Perl 6"
	match[12] = "Go"
	match[13] = "CSS"
	match[14] = "PHP"
	match[15] = "Assembly"
	match[16] = "Objective-C"

	combined := CombineJsonData(primary, language, contrib)
	for i := range match {
		if combined.Languages[i] != match[i] {
			t.Fail()
		}
	}
	if combined.Committer != "gitster" {
		t.Fail()
	}
	if combined.Contributions != 18497 {
		t.Error()
	}

	if combined.Name != "git" {
		t.Error()
	}
	if combined.Owner != "git" {
		t.Error()
	}

}
