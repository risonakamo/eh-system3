# remaining stuff
- date time needs to be formatted better
- sort album list by last update time

## done
- need redirect at / page
- viewer mode not yet working

# optimising
- test out if multithread album info get, since that seems to be the slowest
- try out increasing fiber server threads to see if improves something
- results caching so dont have to continuously scan

# new ideas
- random order mode for albums in list
- think about custom "combined" albums - select a bunch of things then display all those at the same time (normally this only works recursively)
- keyboard shortcut to download image with hashed filename (filename is based on location or something, and is hashed so downloading it multiple times will have the same value (unless its filepath changes))

# other
- thumbnail generation
    - can use node generation for now