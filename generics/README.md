# generics katas

## Dev Notes.

* Simple function printing a slice of any given type without using reflection.

https://github.com/fernandoocampo/justforfun/blob/f1f3a1b495925f1bc9322ae610c3814193b3ab3c/generics/generics.go#L35-L55

* Underlying types

type definition

https://github.com/fernandoocampo/justforfun/blob/f1f3a1b495925f1bc9322ae610c3814193b3ab3c/generics/generics.go#L10-L14

type method

https://github.com/fernandoocampo/justforfun/blob/f1f3a1b495925f1bc9322ae610c3814193b3ab3c/generics/generics.go#L57-L70

* Struct types

type definition

https://github.com/fernandoocampo/justforfun/blob/f1f3a1b495925f1bc9322ae610c3814193b3ab3c/generics/generics.go#L16-L25

constructor

https://github.com/fernandoocampo/justforfun/blob/f1f3a1b495925f1bc9322ae610c3814193b3ab3c/generics/generics.go#L27-L33

* Type as constraint

if we try to use a type different from int and string we are going to get a compile error like this: `TYPE does not implement generics.Addend`.

interface to restrict types

https://github.com/fernandoocampo/justforfun/blob/41d281bfd891c32f8fe69dccc0a1c2240bbfdcc6/generics/generics.go#L10-L17

how to specify the constraint

https://github.com/fernandoocampo/justforfun/blob/41d281bfd891c32f8fe69dccc0a1c2240bbfdcc6/generics/generics.go#L85-L87