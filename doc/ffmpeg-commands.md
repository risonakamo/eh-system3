tall aspect
```
ffmpeg.exe -i 1.mp4 -vf thumbnail,scale=100:-2,crop=100:100 -frames:v 1 out.jpg
```

wide aspect
```
ffmpeg.exe -i 1.mp4 -vf thumbnail,scale=-2:100,crop=100:100 -frames:v 1 out.jpg
```

both aspect
```
ffmpeg.exe -i 1.mp4 -vf thumbnail,scale='if(gt(iw,ih),-2,100)':'if(gt(iw,ih),100,-2)',crop=100:100 -frames:v 1 out.jpg
```

can spam `-i <file`, but need corresponding `1.jpg` at the end. basically, if can generate all the output names together can gen all in one command