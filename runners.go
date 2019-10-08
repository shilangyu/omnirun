package main

// This is only a stringified json with some predefined runners
var runnersJSON = `
[
	{
		"exts": ["js"],
		"run": ["node $or_file"]
	},
	{
		"exts": ["py"],
		"run": ["python $or_file"]
	},
	{
		"exts": ["go"],
		"run": ["go run $or_file"]
	},
	{
		"exts": ["bf"],
		"run": ["brainfuck $or_file"]
	},
	{
		"exts": ["cpp"],
		"run": ["g++ $or_file -o abc123.exe", "./abc123.exe"]
	},
	{
		"exts": ["c"],
		"run": ["gcc $or_file -o abc123.exe", "./abc123.exe"]
	}
]
`
