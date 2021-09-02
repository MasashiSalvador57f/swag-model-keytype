# swag-model-keytype

### Problem to solve 
* Need to define keytype for TypeScript interface
* `keyof Foo` generate Union type but some IDEs don't support completion for union element
* generate keyType for swagger defined models. 

### How to use

Install
```
go install github.com/MasashiSalvador57f/swag-model-keytype
# or 
# get latest binary from https://github.com/MasashiSalvador57f/swag-model-keytype/releases
```

Generate Key Type
```
swag-model-keytype -f <your swagger file path> -o <output file name>
```

### example

```
swag-model-keytype -f sample/swagger.yaml -o KeyTypes.ts
```

```typescript
// auto generated file DO NOT EDIT.
const ErrorKey = {
	Code:"code",
	Message:"message",
	} as const;
export type ErrorKey = typeof ErrorKey[keyof typeof ErrorKey]

const PetKey = {
	Id:"id",
	Name:"name",
	Tag:"tag",
	} as const;
export type PetKey = typeof PetKey[keyof typeof PetKey]
```
