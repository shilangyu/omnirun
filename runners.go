package main

// predefined runners
var initialRunners = []Runner{
	{Exts: []string{"js"}, Run: []string{"node $or_file"}},
	{Exts: []string{"go"}, Run: []string{"go run $or_file"}},
	{Exts: []string{"py"}, Run: []string{"python $or_file"}},
	{Exts: []string{"bf"}, Run: []string{"brainfuck $or_file"}},
	{Exts: []string{"cpp"}, Run: []string{"g++ $or_file -o abc123.exe", "./abc123.exe"}},
	{Exts: []string{"c"}, Run: []string{"gcc $or_file -o abc123.exe", "./abc123.exe"}},
}
