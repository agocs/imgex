Imgex
=====

An index for image similarity, an image comparison algorithm, and a Python implementation thereof.

What's the big idea?
====================

An image is a bunch of pixels, and two images are the same if all of their pixels are the same. Seems easy, but what if one image has been modified? How can we compare the similarity between two images?


How does it work?
=================

Determining the Imgex
---------------------

Much like Soundex determines the wordy qualities of a string, Imgex determines the looky qualities of an image (these are technical terms). We can use Imgex to store a small representation of an image, and we can use the representation of two images to determine if the images are materially similar or not. 

### Determining the Imgex of an image

- Choose an *n*
  - *n* represents the precision of the Imgex. 
  - Higher values for *n* are more precise, but require *n*^2 space and comparison time.
- Normalize the image.
  - Turn image greyscale.
  - Adjust brightness until image is 50% gray.
- Divide the image into *n* by *n* segments.
- For each segment, calculate a 95% confidence interval of the average pixel brightness.
- Store the Imgex as an *n* by *n* matrix of condifence intervals.


Comparing two images using Imgex
--------------------------------

Let's say you have two images, and have the *n* by *n* Imgex calculated for each image (that is, the *n* value for both Imgexes is the same). For each segment, 

- Calculate the overlap of the confidence intervals (see code sample below).
- Overlap value approaches 1 as the segments become similar.
- Negative overlap values indicate dissimilar segments

Perform this for corresponding segment for the two images, generating an overlap matrix. If the overlap matrix consists largely of positive values approaching 1, the two images are similar. If the matrix consists of negative values, the two images are dissimilar. If the two images are sufficently similar, increase *n* and compare again until the images become dissimilar or *n* becomes sufficently high for your purposes.

```
# confidence intervals are for the same segment of two images.
# intervals are expressed as (low,high)
# overlap is the amount of the two confidence intervals that overlap

def find_overlap_area(ci1, ci2):
  ol_low = max(n[0],m[0])
	ol_high = min(n[1],m[1])
	ol = (ol_low,ol_high)
	return ol
  
def find_range(ci):
  return ci[1]-ci[0]
  
def calculate_overlap(ci1, ci2):
  ol = find_overlap_area(ci1, ci2)
  range_ci1 = find_range(ci1)
  range_ci2 = find_range(ci2)
  range_ol = find_range(ol)
  return 2*float(range_ol)/(float(range_n)+float(range_m))
  
#segments are similar
confidence_interval_1 = (125,131)
confidence_interval_2 = (129, 133)
calculate_overlap(confidence_interval_1, confidence_interval_2) #0.4

#segments have the same value
confidence_interval_1 = (125,131)
confidence_interval_2 = (125,131)
calculate_overlap(confidence_interval_1, confidence_interval_2) #1.0

#segments are wildly different
confidence_interval_1 = (12,13)
confidence_interval_2 = (250,253)
calculate_overlap(confidence_interval_1, confidence_interval_2) #-118.5
```


Limitations
===========

This does not necessarially rate inverted or rotated images as being similar.
