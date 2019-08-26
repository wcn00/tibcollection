# 	tibcollection - Activity

This repo is a collections contribution for the legacy flogo engine located in the github.com/TIBCOSoftware repo.
It was created so that it could be used on Flogo Enterprise and TCI until such time that equivalent functionality could be built into the standard palatte.

## License
tibcollection is licensed under a BSD-type license. See [LICENSE](LICENSE) for license text.


The objective if this contribution is to allow a flow to aggregate the results of array operations in either the main or sub flows such that all results are then presented in one place.  This permits the results of several nested array operations to be rolled up into a single or multiple arrays of objects.

There are three operations:
APPEND   If no key is provided a uniqueue key will be generated.  This allows for the collection to be used safely by multiple threads.  If no object is provided only the key will be returned, however if an object is provided it will be appended to the collection.  If a key is provided the object will be appended to the named collection.  

GET    Get will return an array of all the objects in the collection for the associated key.  It is an error to call GET without a key.

DELETE   Delete will delete the collection associated with the provided key and free its memory.  It is not necessary to call delete within the same flow that the collection is created in, however some mechanism must be in place to delete the collection or it will consume memory indefinitely.

Objects can be any mappable value.  Mixing different objects in the same collection will make it difficult to iterate over them in a useful way, but it is not prohibited.

