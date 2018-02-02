# raymond

This is a fork of `aymerick/raymond`, a Golang implementation of Handlebars 3.0, with non-breaking API changes.

## What changed?

### Return errors instead of panicking

The following functions now return `error` instead of triggering a `panic`.

- `RegisterHelper() error`
- `RegisterHelpers() error`
- `RegisterPartial() error`
- `RegisterPartials() error`
- `Template.RegisterHelper() error`
- `Template.RegisterHelpers() error`
- `Template.RegisterPartial() error`
- `Template.RegisterPartials() error`

### Allow previously registered helpers and partials to be removed

The following functions were added to the API.

- `RemovePartial()`
- `RemoveHelper()`
- `Template.RemovePartial()`
- `Template.RemoveHelper()`

### Use golang/dep to manage dependencies

The raymond subpackages are vendored along with the YAML dependency.

## Should you use this fork?

As of now (Feb 2018), upstream has multiple pull requests waiting to be merged with similar functionality.
If they are merged by the time you read this, then you should probably use the upstream repo.

These are the pending pull requests you should check:

- [RegisterHelper and RegisterHelpers do not panic](https://github.com/aymerick/raymond/pull/18)
- [Added methods to remove global partials](https://github.com/aymerick/raymond/pull/16)
- [added the ability to remove a registered helper](https://github.com/aymerick/raymond/pull/9)

## Future changes

One of my projects heavily depends on `raymond` and I'll likely be making future changes that
fit my specific use-case. Some off the top of my head are:

- Encapsulating globals so multiple raymond instances can be ran.
- Returning human readable errors for handlebars syntax issues.

## Installing

```
$ go get github.com/ef2k/raymond
```

## Documentation

Raymond already has excellent documentation.

You can find it here: [Read the Original Raymond Docs](./original_README.md)
