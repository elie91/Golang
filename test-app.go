package main

import "testing"


func testUpdateArticle(t *testing.T) {
	channelInt, channelOut := make(chan UpdateArticle), make(chan string)
	go manageUpdate(channelInt, channelOut)
	var testArticle UpdateArticle
	testArticle.id = 1
	testArticle.currentPrice = 4000
	channelInt <- testArticle
	res := <-channelOut
	if res != "Article was updated" {
		t.Errorf("Update failed")
	}
}
