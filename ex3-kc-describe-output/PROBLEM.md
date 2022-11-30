# Kubectl Describe Output

You've just run `kubectl describe po/kube-dns-cache-5v9ls` to get the output in
`output.txt`. You want to retrieve the name of the Node this pod is running on,
but you want to do it in a programmatic way.

The value you want to isolate is on line 6 -
`ip-10-31-54-255.us-west-2.compute.internal`. How would you isolate this value
using text commands?
