package main

func main() {
	// TODO extract path from the first arg or infer cwd if no args

	// TODO start cli animation (ansi escape chars (kaomoji might be fun to animate)) while the search run

	// TODO build a tree representation 3 field (name, size, childern), dirs should init nil size

	// TODO reverse populate size on dirs as sum of children

	// TODO get terminal dimentions, figure out line count

	// TODO figure out top x relevant items to display to match line count

	// TODO render it out bby (sofar probably as name:size in human readable format: horizontal barchart (figure out how to signify folder membership))
}
