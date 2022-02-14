
# Overview

`ttg` (`t`op `t`ri`g`rams) is a command line program that when given one text file path or more returns a list of the 100 most common three word sequences. This in the context of a coding challenge for the SSEDT position at NR.

# How To Run

There are 2 ways to feed data into the program. In the first one, the program accepts as arguments a space-separated list of one or more file paths:

```sh
go run . file1.txt file2.txt fileN.txt
```

or, compiled, as:

```sh
./ttg file1.txt file2.txt fileN.txt
```

For example, you can run the solution by using some test text files available in the `testdata` folder:

```sh
./ttg testdata/mobydick.txt testdata/don_quixote.txt testdata/odyssey_of_homer.txt
```

The second one is through `stdin`, just stream text data into it similar to regular *nix-based command line tools:

```sh
cat file1.txt | go run .
```
or, compiled, as:

```sh
cat file1.txt | ./ttg
```

In the second form, even this below will work, though the input will be treated as a single stream (stdin) containing the the merged content of all the text files piped into the solution:

```sh
cat file1.txt file2.txt fileN.txt | ./ttg
```

# What To Do Next

Given additional time, these are potential items where improvement could be required:

- Use a golang channel to allow goroutines communicate they are finished. Goroutines implemented, though working efficiently, do not currently have unit testing in place.
- Parameterized N, the number of most common 3-word sentences, to allow users specify a different number of 3-word sentences results.
- Add code to allow the program to identify itself, its service account, machine running on, etc., at runtime.
- Try more succint/elegant regular expressions options inside the data `CleanUp` function code.
- In terms of unit testing:
    - Implement mocking of stdin input.
    - Add more test cases, including identification of edge ones.
- Measure solution performance (speed, RAM).
- Implement unicode chars support.
- Research root cause of bug reported below.

# Bugs

## Approximate Results

When compared to searching in a text editor for a 3-word sentence string that you know is in the text, the results are approximate but not exact. Let me try to explain.

For example, if running the solution with the text file called `don_quixote.txt` (which is included in the `testdata` folder) as input to the program:

```sh
./ttg testdata/don_quixote.txt
```

you will notice that the top result shows as follows:

```sh
"said don quixote" - 450
```
However, when opening the same `testdata/don_quixote.txt` in a text editor and using the editor's string search feature to search for `said don quixote`, you will get a different number of occurrences; in this text file case, 426.

Something similar happens for the case of `testdata/odyssey_of_homer.txt`. The solution reports 81 occurrences of `odysseus of many`, the editor finds 80 of them.

Some possible explanation to this behavior could be the way some punctuaction signs or other symbols like, apostrophe or hypen, are currently handled by this solution; but again some more time would be needed to find the underlying root cause.

## Missing trigrams between buffered data chunks

While buffering chunks of file data, we are missing combining the last two words of the current chunk of stream data with the first word of the next chunk of stream data.

# Additional Notes

## Docker

With the code in this repo, you can use the following `dockerfile` content to build a docker image in your local host:

```docker
FROM golang:1.17-alpine

WORKDIR /ttg

COPY . /ttg

USER root

RUN apk add bash --no-cache bash

CMD ["/bin/sh" "-c" "tail -f /dev/null"]
```

Once you clone this repo, go to its main project folder (ttg), and having `docker` installed, run:

```sh
docker build . <DOCKER_IMAGE:version>
```

For example:

```sh
docker build . moalf/ttg:latest
```

To run your image locally and connect to it, execute the following:

```sh
docker run -it moalf/ttg:latest bash
```

Once you get to the `bash` prompt inside that local container, all the commands already mentioned in the `HOW TO RUN` section above, should be easily and successfully run in the container as well.

If you want to skip building a local container for any reason, pull my `moalf/ttg:latest` image from docker hub, by doing this:

```sh
docker pull moalf/ttg
```