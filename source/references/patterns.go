package references

import "regexp"

var PseudoElementsPtrn = regexp.MustCompile(`(?i)(code segment)|(data segment)|(stack segment)|(\.code segment)|(\.data segment)|(\.stack segment)|(\.model small)|(\.data)|(\.stack)|(\.code)|byte ptr|word ptr`)
var BaseConstantPtrn = regexp.MustCompile(`\w$`)
var StringPtrn = regexp.MustCompile(`(^")|(^')`)
var ConstantPtrn = regexp.MustCompile(`(^\d)|(^")|(^')`)
var FullStringPtrn = regexp.MustCompile(`("+(.)+")|('+(.)+')`)
var DupPtrn = regexp.MustCompile(`(?i)(DUP)+\(+(.*)+\)`)
var SymbolPtrn = regexp.MustCompile(`^\w\D`)
var BadQuotesPtrn = regexp.MustCompile(`"[^"\n]*(?:$|[^"]+$)`)
var EndQuotesPtrn = regexp.MustCompile(`"$|'$`)

/* Segment Patterns */
var CodeSegmentPtrn = regexp.MustCompile(`(?i)(code segment)|(\.code segment)|(\.code)`)
var StackSegmentPtrn = regexp.MustCompile(`(?i)(stack segment)|(\.stack segment)|(\.stack)`)
var DataSegmentPtrn = regexp.MustCompile(`(?i)(data segment)|(\.data segment)|(\.data)`)
