## Stage breakdown for Class Inheritance

1. Class Hierarchy
	- Basic subclass
	- Super and sub classes inside local scope
	- parent in global, child in local
	- cycle: `class Oops < Oops {}`
	- parent is not a class
		- not callable
		- callable
2. Inheriting methods
	- check type of subclass 
	- ~~function which expects super should accept sub~~ not statically typed can't check 
	- access inherited fields
		- non callable
		- callable
3. Super methods
	- super.method()
	- sub class method with same name gets precedence (**overriding**)
	- super should return 1st class function
	- super should be called from superclass ;)
4. Invalid super
	- call super.func() from a class with no superclass
	- super.func() outside
	- can't have bare super
			- without class
			- also without . & prop
5. Misc
	- inherit constructors
	- Var resolving
		- class.cook
		- super.cook
		- this.cook