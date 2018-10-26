# Kray

Kray is a Work in Progress project that is designed to be cloud agnostic in performing the following:
  - Create a new Kubernetes Namespace (DONE)
  - Generate a DNS Record in the given environment (TODO)
  - Create a new Database instance to be used for the development environment (TODO)

# Motivation

To achieve consistency and simplicity as I progress in my personal/professional adventures.

# What is a UDE?

A UDE (User Development Environment) is simply put a collection of resources which establish a baseline
for what is required in order for a collection of applications to run on various cloud platforms.

# Caveats

This project is extremely opinionated and this is because this is a Proof of Concept towards standardizing
some of my own practices into a reusable application that I can use personally and professionally.

What you consider a baseline for what is required for a collection of applications to run on various cloud
platforms is entirely up to you. This is me. You do you ;)

# Testing

## The Minimum Standard

*80%* Coverage (preferably 100% happy path)

## TLDR

1. `testing` package is your friend.
1. `https://github.com/stretchr/testify` assertions and requirements are more than welcome.
1. table driven tests are often worthwhile.
1. tests will dictate how the production code is structured.
1. the minimum standard of coverage will be upheld.
1. we have decided against the use of mocking libraries

## Unit Testing

Most practical ideas have been stolen from Mitchell Hashimoto's very comprehensive talk on advanced testing in Go [video][1].


### Tests must be written before code reaches master

Because tests can't be written later down the line. The way in which tests are written dictates the way in which the go we are testing is written.

Any code that doesn't have suitable testing is subject to being thrown away and written from scratch with tests. In fact it is urged that is done as soon as possible.

### I like testify

Simple comparisons are good enough to test with. However, it can get tedious and inconsistent to write our own failure messages. `assert` and `require` reduce the noise in a test and provide nicely formatted default failure messages. Plus it works very well with the standard libraries.

## Mocking

Currently our stance is simple. We don't use any third party libraries for mocking. Instead we favor hand-rolled mocks, dummies, noops, fakes and spies.

Here's an example:

Say that we have an interface which when supplied with a "container ID", it is intended to stop that container from running. The interface for this functionality might look like this:

```go
type ContainerStopper interface {
    Stop(containerID string) error

}
```

#
```

Instead of turning to a mocking library, all sorts of hand-rolled tricks can be employed to swap this behavior out.

```go
type spyContainerStopper struct {
    id     string
    err    error

}

func (s spyContainerStopper) Stop(id string) error {
    s.id = id
    return n.err

}
```

The first example above is the simple spy-like implementation which captures the input `id` on calls to `Stop(...)`. It also returns the error configured on the struct in the `err` field.

```go
type containerStopperFunc func(string) error

func (c containerStopperFunc) Stop(s string) error { return c(s)  }
```

The second allows for simple anonymous functions to be used in a test case as an implementation. Note that this is often only useful for smaller interfaces, often single function (but your interfaces should be small anyway). The following demonstrates an anonymous function which achieves the same ends as the struct implementation.

```go
func TestSomething(t *testing.T) {
    var (
        capturedID string
        spy        ContainerStopper = containerStopperFunc(func(id string) error {
                   capturedID = id
                   return errors.New("something went wrong")
                                        
        })
    )

    //...

}
```

If you are dealing with a big pesky interface, it is recommended that you use something like [impl](https://github.com/josharian/impl) to generate a skeleton and speed things up.

# Contributing

Fork it, PR it. I always welcome ideas and hope if nothing else this provides someone some inspiration of what not to do ;)
