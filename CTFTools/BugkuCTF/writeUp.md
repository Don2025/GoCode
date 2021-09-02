# Bugku CTF

#### 1.[签到题](https://www.amanctf.com/challenges/detail/id/1.html)

关注微信公众号`Bugku`回复`flag`即可获取`flag`。

------

#### 2.[这是一张单纯的图片](https://www.amanctf.com/challenges/detail/id/2.html)

修改`.jpg`后缀为`.html`，双击打开可以在乱码的最后一行找到`key{you are right}`提交就行啦。

------

#### 3.[隐写](https://www.amanctf.com/challenges/detail/id/3.html)

`tweakpng`打开图片后会弹出提示，将图片大小修改为500x500就可以看到`BUGKU{a1e5aSA}`啦。

![](![](https://paper.tanyaodan.com/amanctf/3/2.png)

------

#### 4.[talnet](https://www.amanctf.com/challenges/detail/id/4.html)

`Wireshark`打开文件后，协议分级过滤器中输入`telnet contains "flag"`即可看到`flag`。

![](https://paper.tanyaodan.com/amanctf/4/1.png)

------

#### 5.[眼见非实]()

将`.docx`文件修改后缀为`.zip`后解压缩，用`grepWin`对解压后的文件夹进行文字搜索`flag`，可以在`眼见为实/word/document.xml`文件中找到`flag`。

![](https://paper.tanyaodan.com/amanctf/5/1.png)

------

#### 6.[啊哒](https://www.amanctf.com/challenges/detail/id/6.html)

查看图片属性可以找到相机型号`73646E6973635F32303138`，将后缀改为`.rar`可以看到压缩包中有`flag.txt`，提取需要输入解压密码，直接输入相机型号不对，16进制转字符串得到解压密码`sdnisc_2018`，解压后打开`flag.txt`看到`flag{3XiF_iNf0rM@ti0n}`。

![](https://paper.tanyaodan.com/amanctf/6/1.png)



#### 7.[又一张图片，还单纯吗](https://www.amanctf.com/challenges/detail/id/7.html)

Kali-Linux环境下无脑`foremost -i`分离图片，output目录下另一张图片就是`falg{NSCTF_e6532a34928a3d1dadd0b049d5a3cc57}`。

```bash
┌──(tyd㉿kali-linux)-[~/ctf]
└─$ foremost -i file.jpg                
Processing: file.jpg
|*|
```

------

#### 8.[猜](https://www.amanctf.com/challenges/detail/id/8.html)

题目作者用下半身出的什么🐔⑧破题？？浪费👴俩个金币下载图片，`flag{liuyifei}`。

------

#### 9.[宽带信息泄露](https://www.amanctf.com/challenges/detail/id/9.html)

用`RouterPassView`打开`conf.bin`文件后可以看到宽带用户名的值`<Username val=053700357621 />`从而得到`flag{053700357621}`。

------