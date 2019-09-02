# collection

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

# Installation

To install the collection contribution open the "Extensions" page and click "Upload".  By default the wizard expects to load the contribution from an online repository and to do that enter:
    github.com/wcn00/tibcollection/activity/tibcollection
in the repository URL field and click "Import".
The import operation can be quite time consuming so be patient.

If you have modified the connector locally and want to use or test that modification, zip up the source tree of the repository starting at the activity folder:
/tmp/activity
└── tibcollection
    ├── activity.go
    ├── activity.json
    └── activity_test.go

Use the upload zip option to install the contribution from your own source.

To uninstall the contributon click the "Tibco Legacy Collection API" tile and select the trashcan icon.



# Samples
In the samples folder is the file "TibCollection.json" which is a flow which illustrates all the operations of the collection.  It is meant to be tested using a rest tool of some sort.
When testing use post to send the following json object to the server at:
    http://localhost:4343/testcollection/?key=somekey
The paylad should look like:
{
	"name": "walter",
	"age": "45",
	"eyecolor": "blue"
	
}

This is because the mapper depends on the format of the payload.  You can remove the mapper and reconfigure the app to work with arbitrary data but you will have to tolerate the activities being marked in error.  (this may change in future releases of flogo)

Happy trails!
