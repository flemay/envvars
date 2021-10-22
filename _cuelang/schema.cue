import "list"

#Tag: {
	name:  string
	desc?: string
}
tags: [...#Tag]
#tagNames: [ for tag in tags {tag.name}]
_checkTagsUniqueness: list.UniqueItems(#tagNames) & true

#Envvar: {
	name:     string
	desc?:    string
	optional: bool | *true
	example?: string
	tags?: [...string]
	_checkTagsExist: [ for t in tags {list.Contains(#tagNames, t) & true}]
	_checkTagsUniqueness: list.UniqueItems(tags) & true
}

envvars: [...#Envvar]
#envvarNames: [ for envvar in envvars {envvar.name}]
_checkEnvvarNamesUniqueness: list.UniqueItems(#envvarNames) & true
