# Introduction

`brass` is a lib concerned with persisting a map[string][]inteface of data. 
This map is maintained by an instace of a `bank`, and each bank can have a number
of `rolls` which are slices of data.  Each `roll` holds a reference to the
containing bank so at any point the entire bank (including the referenced `roll`)
can be serialized to disk via the `bank` api.

## License

See license file.

The use and distribution terms for this software are covered by the
[Eclipse Public License 1.0][EPL-1], which can be found in the file 'license' at the
root of this distribution. By using this software in any fashion, you are
agreeing to be bound by the terms of this license. You must not remove this
notice, or any other, from this software.


[EPL-1]: http://opensource.org/licenses/eclipse-1.0.txt

