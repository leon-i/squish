# squish  
CLI image optimizer, inspired by [squoosh.app](https://squoosh.app/) - outputs optimized images to a folder in your current directory so you don't need to do drag/drop/download stuff  
  
built with Go + Cobra + mozjpeg
  
## Installation  
squish requires Go and mozjpeg to run its commands  
```  
# follow steps to config GOPATH  
brew instlall go  
  
# follow steps to config path
brew install mozjpeg
```  
  
If you have go + mozjpeg installed...  
```
git clone https://github.com/leon-i/squish.git  
  
cd squish  
  
go install squish  
```  
  
## Commands & Flags  
### base  
```  
# displays welcome message
squish  
  
# displays available commands + their descriptions
squish --help  
squish -h
```  
  
### all  
```
# optimizes all images in current directory using mozjpeg  
# defaults: quality = 75, output directory = 'squished'
squish all  
  
# sets output image quality to given integer (1 - 100)  
squish all --quality <number>  
squish all -q <number>  
  
# sets output directory name to given string  
squish all --destination <string>  
squish all -d <string>  
```  
  
### only  
```
# optimizes explicitly named images in current directory using mozjpeg  
# note: if the image name contains spaces, wrap in quotes  
# defaults: quality = 75, output directory = 'squished'
squish only image_1.png image_2.jpg  
squish only image_1.png "spaced image name.jpg"
  
# sets output image quality to given integer (1 - 100)  
squish only image_1.png --quality <number>  
squish only image_1.png -q <number>  
  
# sets output directory name to given string  
squish only image_1.png --destination <string>  
squish only image_1.png -d <string>  
```
