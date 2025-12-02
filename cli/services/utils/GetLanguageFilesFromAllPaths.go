package utils

import (
	"fmt"
	"path/filepath"
)

// ----------------------HASHMAP----------------------------------
const ArraySize = 7 // .jsx, .js, .java, .cpp, .go, .py, .tsx

type HashTable struct {
	array [ArraySize]*bucket
}

func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].del(key)
}

func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("already exists")
	}
}

func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil { //very interesting for loop acting as a while loop
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *bucket) del(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
		} else {
			previousNode = previousNode.next
		}
	}
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func hash(key string) int {
	var sum = 0
	for _, char := range key {
		sum += int(char)
	}
	return sum
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result

}

// -------------------------GetLanguageFilesFromAllPaths-----------------------------------
func GetLanguageFilesFromAllPaths(paths []string) []string {
	var supportedLangs = []string{
		".jsx",
		".py",
		".tsx",
		".go",
		".cpp",
		".java",
		".js",
	}

	var languageFiles = []string{}

	table := Init()
	for _, fileType := range supportedLangs {
		table.Insert(fileType)
	}

	for _, path := range paths {
		extension := filepath.Ext(path)
		if table.Search(extension) {
			languageFiles = append(languageFiles, path)
		}
		continue
	}
	return languageFiles
}
