# Go performance/features test

Trying some operations with Go and measure approx time :)

### querying (...)

\* array.go : array-like positional manipulation of the selection.
- Eq()
- First()
- Get()
- Index...()
- Last()
- Slice()

 \* expand.go : methods that expand or augment the selection's set.

- Add...()
- AndSelf()
- Union(), which is an alias for AddSelection()

\* filter.go : filtering methods, that reduce the selection's set.

End()
Filter...()
Has...()
Intersection(), which is an alias of FilterSelection()
Not...()
* iteration.go : methods to loop over the selection's nodes.

Each()
EachWithBreak()
Map()
* manipulation.go : methods for modifying the document

After...()
Append...()
Before...()
Clone()
Empty()
Prepend...()
Remove...()
ReplaceWith...()
Unwrap()
Wrap...()
WrapAll...()
WrapInner...()
* property.go : methods that inspect and get the node's properties values.

Attr*(), RemoveAttr(), SetAttr()
AddClass(), HasClass(), RemoveClass(), ToggleClass()
Html()
Length()
Size(), which is an alias for Length()
Text()
* query.go : methods that query, or reflect, a node's identity.

Contains()
Is...()
* traversal.go : methods to traverse the HTML document tree.

Children...()
Contents()
Find...()
Next...()
Parent[s]...()
Prev...()
Siblings...()
* type.go : definition of the types exposed by goquery.

Document
Selection
Matcher
* utilities.go : definition of helper functions (and not methods on a *Selection) that are not part of jQuery, but are useful to goquery.

NodeName
OuterHtml