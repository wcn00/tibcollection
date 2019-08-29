# fecollection

This repo is a collections contribution for Tibco Flogo.
Its purpose is to allow one to accumulate operations performed during array processing such that they are available after the array is processed and to parent and child flows.

## License
tibcollection is licensed under a BSD-type license. See [LICENSE](LICENSE) for license text.

There are three operations:
APPEND   If no key is provided a uniqueue key will be generated.  This allows for the collection to be used safely by multiple threads.  If no object is provided only the key will be returned, however if an object is provided it will be appended to the collection.  If a key is provided the object will be appended to the named collection.  

GET    Get will return an array of all the objects in the collection for the associated key.  It is an error to call GET without a key.

DELETE   Delete will delete the collection associated with the provided key and free its memory.  It is not necessary to call delete within the same flow that the collection is created in, however some mechanism must be in place to delete the collection or it will consume memory indefinitely.

Objects can be any mappable value.  Mixing different objects in the same collection will make it difficult to iterate over them in a useful way, but it is not prohibited.
Because Flogo Studio is designed to map values between activities using schema based metadata it will flag the use of hidden fields as being in error but the runtime will work correctly.  For instance while iterating over the results of a "get" operation using the folowing mapping:

    $iteration[value].name

Will produce an error because Studio does not have metadata for the objects in the collection.  This can be overcome by using the mapper activity to give the objects in the collection some structure based on a json schema.

The samples folder has a project.json file which you can import to illustrate the above technique.

