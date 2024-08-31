package main

import (
	"fmt"
	"os"
)




func main()  {
	args, err:= ParseArgs()
    if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tracks, err := ParseCueFile(args.filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tracks = FilterTitles(tracks, args.whitelist)	
	for _, t := range tracks {
		fmt.Println(t.DisplayName())
	}
}
