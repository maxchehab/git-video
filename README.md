# 🎥 Git-Video 🎥 
A tool for linking Youtube videos in GitHub markdown files with a **single** link.

## 📖 Instructions 📖
Using Git-Video is very simple. Just append your Youtube video to `git-video.maxchehab.com`. 

For example:
```markdown
[![super sick video](http://git-video.maxchehab.com/youtube.com/watch?v=y6120QOlsfU)](http://git-video.maxchehab.com/youtube.com/watch?v=y6120QOlsfU)
```
will show up as: 

[![super sick video](http://git-video.maxchehab.com/youtube.com/watch?v=y6120QOlsfU)](http://git-video.maxchehab.com/youtube.com/watch?v=y6120QOlsfU)

The webservice only really needs the `v=videoID` query so the following is also acceptable:
```markdown
[![super sick video](http://git-video.maxchehab.com?v=y6120QOlsfU)](http://git-video.maxchehab.com?v=y6120QOlsfU)
```
## ⚡ How to set this up on your own. ⚡
Make sure you have installed [golang](https://golang.org/doc/install).

If you have the dependencies for this project, pop open a terminal (`ctrl-alt-t` for all you nasty neck-beards) and get ready to copy-pasta some commands.

Download the project:
```bash
$ git clone https://github.com/maxchehab/git-video.git
$ cd git-video
```

To build and run the executable just...
```bash
$ ./run
```
<h3 align="center">Here, have a gopher.</h3>
<p align="center">
  <img src="https://media.giphy.com/media/hM87DMnls5oZy/giphy.gif" width="450px" />
</p>

