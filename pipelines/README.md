# Pipelines

Everything at the command line is *composable*. Take the output from one command
and pass it to another command. Take the output from that command and pass it to
another command.

For example, the `wc` program counts characters, lines, or words.

To count the number of words in Moby Dick:

```
wc wc/example.txt
```

Or the number of lines:

```
wc -l wc/example.txt
```

But `wc` is also set up to read from the left side of a pipe. You can use `cat`
to print the contents of a file or files to stdout, and then pipe the result to
`wc`:

```
cat wc/example.txt | wc -l
```

You can keep going! For example if you wanted to see how many lines in Moby Dick
refer to Ishmael:

```
cat wc/example.txt | grep Ishmael | wc -l
```

By default output gets printed to the terminal. But you can also redirect to
a file using the `>` operator:

```
cat wc/example.txt | grep Ishmael > ishmael-references.txt
```

## Stdout vs Stderr

When you run a program the output goes to two different people

- One person prints out the normal results (stdout)
- One person prints out any errors (stderr)

Two streams, by default, both of them get printed to the command line. So if you
try to do word count on a file that doesn't exist:

wc badfile.txt
wc: badfile.txt: No such file or directory (os error 2)

That message gets printed to stderr.

Messages printed to stderr won't go through a pipe.

```
cat wc/badfile.txt | grep Ishmael > ishmael-references.txt
```

Will end up with an empty file.
