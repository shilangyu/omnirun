package main

// predefined runners
var initialRunners = []Runner{
	{Exts: []string{"js"}, Run: []string{"node $or_file"}},
	{Exts: []string{"go"}, Run: []string{"go run $or_file"}},
	{Exts: []string{"py"}, Run: []string{"python $or_file"}},
	{Exts: []string{"bf"}, Run: []string{"brainfuck $or_file"}},
	{Exts: []string{"cpp"}, Run: []string{"g++ $or_file -o a.exe", "./a.exe", "rm a.exe"}},
	{Exts: []string{"c"}, Run: []string{"gcc $or_file -o a.exe", "./a.exe", "rm a.exe"}},
	{Exts: []string{"rs"}, Run: []string{"rustc $or_file -o a.exe", "./a.exe", "rm a.exe"}},
	{Exts: []string{"fish"}, Run: []string{"fish $or_file"}},
	{Exts: []string{"ts"}, Run: []string{"ts-node $or_file"}},
}
