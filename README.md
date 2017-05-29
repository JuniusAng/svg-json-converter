# svg-json-converter
Go script to convert svg to json for project that use react-native-svg-uri

# what it does
As we know, android cannot bundle anything beside XML and PNG file as resource, 
so if your react-native project uses react-native-svg-uri module, you will need to convert the svg into something allowed.
In this case this script will find and convert any .svg file and wrapped it in json format
  
# how to use
you can build the source directly or use the provided binary for each OS 

to make things faster, please specify the PROJECT_PATH to your JS src folder.

run this in terminal : svg-json-converter [PROJECT_PATH]

you can also add it into your scripts part in your project's package.json

example 
"scripts":{
    "convert-svg": "[PATH_TO_SVG_JSON_CONVERTER]/svg-json-converter src/",
} 
