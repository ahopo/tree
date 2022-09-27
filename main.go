package main

import (
	"strings"

	"github.com/ahopo/ccmd"
	"github.com/ahopo/ccmd/git"
)

func main() {
	// g := ccmd.InitGit()
	// g.SetRepo("https://gitlab.mynt.xyz/cluster5/gaws.git")
	// g.SetRootFolder(".")
	// g.SetSuperBranch("master")
	t := Init("https://gitlab.mynt.xyz/cluster5/gaws.git", ".", "master")
	t.Clone()
}

type Tree struct {
	g git.Config
}

//initialize
func Init(repo string, rootfolder string, superBranch string) *Tree {
	t := new(Tree)
	g := ccmd.InitGit()
	g.SetRepo(repo)
	g.SetRootFolder(rootfolder)
	g.SetSuperBranch(superBranch)
	t.g = g
	return t
}

//clone
func (t *Tree) Clone() {
	t.g.Clone().Branch("").Exec()
}

//checkout branch
func (t *Tree) CheckoutBranch(branchname string) {
	_, e := t.g.Checkout().Branch(branchname).Exec()
	if e != nil {
		panic(e)
	}
}

//checkout tag
func (t *Tree) CheckoutTag(tagname string) {
	_, e := t.g.Checkout().Tag(tagname).Exec()
	if e != nil {
		panic(e)
	}
}

//get branches
func (t *Tree) GetBranches() []string {
	o, e := t.g.GetAllBranchs().Exec()
	if e != nil {
		panic(e)
	}
	return strings.Split(o, "\n")
}

//get rags
func (t *Tree) GetTags() []string {
	o, e := t.g.GetAllBranchs().Exec()
	if e != nil {
		panic(e)
	}
	return strings.Split(o, "\n")
}

//get latest tag
func (t *Tree) GetlatestTags() string {
	return t.GetTags()[(len(t.GetTags()) - 2)]
}
