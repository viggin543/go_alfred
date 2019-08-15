package main

import "strings"

type String struct {
	Str string
}

func (str String) withoutWhitespace()  String {
	return String{strings.ReplaceAll(str.Str," ","")}
}

func (str String) noCloseTag()  String {
	return String{strings.ReplaceAll(str.Str,"\" />","")}
}

func (str String) inUserHomeDir()  String {
	return String{strings.ReplaceAll(str.Str,"<option value=\"$USER_HOME$","~")}
}

func (str String) outsideUserHomeDir()  String {
	return String{strings.ReplaceAll(str.Str,"<option value=\"","")}
}