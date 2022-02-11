

# Medulla Prototyping
_____________

### What is Medulla Prototypes

This repo is designed to be an area for code, and processes to be tested, 
and reviewed before migrating into the actual medulla git repository.
_____________

### Prototype Branch Structure

Each prototype git branch must have the format of 

`{first initial of first name:lowercase}{last name}/prototype/{name of prototype:lowercase}-{version number of prototype}`

Examples

`kbecker/prototype/medulla-deploy-process`

Alternate prototypes targeting the same feature, or functionality would be branched as


`kbecker/prototype/medulla-deploy-process-alt-2`<br>
`kbecker/prototype/medulla-deploy-process-alt-3`<br>
`kbecker/prototype/medulla-deploy-process-alt-4`<br>
_______

### How a Prototype is Implemented

Each prototype must be fully self contained runnable micro-service based program, that implements a feature, or functionality 
that is desired to be migrated into the Medulla repository. 

Each prototype must be runnable in Tye. At the root of each project should be a `tye.yaml` manifest that allows
other developers to spin up the prototype with 
`tye run --watch`

Each prototype must have a  `README.md` at the root that describes the prototype in detail, what is accomplished, 
and how to test locally.  
