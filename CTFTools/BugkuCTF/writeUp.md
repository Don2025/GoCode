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

#### 10.[隐写2](https://www.amanctf.com/challenges/detail/id/10.html)

将`welcome.jpg`文件修改成`welcome.rar`打开后查看`提示.jpg`知道了`flag.rar`的解压密码是三位数。先在`Kali-Linux`上`sudo apt-get install fcrackzip`，所有参数如下：

> `fcrackzip -h`查看可用参数
>
> - [-b|--brute-force] use brute force algorithm 爆破模式
> - [-D|--dictionary] use a dictionary 使用一个自定义字典
> - [-B|--benchmark] execute a small benchmark 基准测试
> -  [-c|--charset characterset]  use characters from charset 使用指定的字符集 一般'aA1!'
> - [-h|--help] show this message 显示消息
> - [--version] show the version of this program 显示版本
> -  [-V|--validate] sanity-check the algorithm 健全性检查
> - [-v|--verbose] be more verbose 更详细
> - [-p|--init-password string]  use string as initial password/file 指定开始字符
> - [-l|--length min-max] check password with length min to max 限制密码长度
> - [-u|--use-unzip] use unzip to weed out wrong passwords 使用unzip清除错误的密码
> - [-m|--method num] use method number "num" (see below) 指定破解类型
> -  [-2|--modulo r/m] only calculcate 1/m of the password 只计算 1/m 的密码

`-b`爆破模式，`-c '1'`设置字符集只使用数字0-9，`-u`清除错误的密码，组合起来就是命令行`fcrackzip -b -c '1' -l 3 -u flag.rar`，得到`flag.rar`的解压密码是871。

```bash
┌──(tyd㉿kali-linux)-[~/ctf]
└─$ fcrackzip -b -c '1' -l 3 -u flag.rar

PASSWORD FOUND!!!!: pw == 871
```

解压缩`flag.rar`后得到一张图片`3.jpg`，用`WinHex`打开图片后可以在最后看到`f1@g{eTB1IEFyZSBhIGhAY2tlciE=}`，直接提交不对，把花括号中的字符串进行base64解码，得到`f1@g{y0u Are a h@cker!}`。

------

#### 11.[多种方法解决](https://www.amanctf.com/challenges/detail/id/11.html)

解压缩后得到一个`.exe`文件，修改后缀为`.txt`后打开可以看到如下信息：

```
data:image/jpg;base64,iVBORw0KGgoAAAANSUhEUgAAAIUAAACFCAYAAAB12js8AAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAArZSURBVHhe7ZKBitxIFgTv/396Tx564G1UouicKg19hwPCDcrMJ9m7/7n45zfdxe5Z3sJ7prHbf9rXO3P4lLvYPctbeM80dvtP+3pnDp9yF7tneQvvmcZu/2lf78zhU+5i9yxv4T3T2O0/7eud68OT2H3LCft0l/ae9ZlTo+23pPvX7/rwJHbfcsI+3aW9Z33m1Gj7Len+9bs+PIndt5ywT3dp71mfOTXafku6f/2uD09i9y0n7NNd2nvWZ06Ntt+S7l+/68MJc5O0OSWpcyexnFjfcsI+JW1ukpRfv+vDCXOTtDklqXMnsZxY33LCPiVtbpKUX7/rwwlzk7Q5JalzJ7GcWN9ywj4lbW6SlF+/68MJc5O0OSWpcyexnFjfcsI+JW1ukpRfv+vDCXOTWE7a/i72PstJ2zfsHnOTpPz6XR9OmJvEctL2d7H3WU7avmH3mJsk5dfv+nDC3CSWk7a/i73PctL2DbvH3CQpv37XhxPmJrGctP1d7H2Wk7Zv2D3mJkn59bs+nDA3ieWEfdNImylJnelp7H6bmyTl1+/6cMLcJJYT9k0jbaYkdaansfttbpKUX7/rwwlzk1hO2DeNtJmS1Jmexu63uUlSfv2uDyfMTWI5Yd800mZKUmd6Grvf5iZJ+fW7PjzJ7v12b33LSdtvsfuW75LuX7/rw5Ps3m/31rectP0Wu2/5Lun+9bs+PMnu/XZvfctJ22+x+5bvku5fv+vDk+zeb/fWt5y0/Ra7b/ku6f71+++HT0v+5l3+tK935vApyd+8y5/29c4cPiX5m3f5077emcOnJH/zLn/ar3d+/flBpI+cMDeNtJkSywn79BP5uK+yfzTmppE2U2I5YZ9+Ih/3VfaPxtw00mZKLCfs00/k477K/tGYm0baTInlhH36iSxflT78TpI605bdPbF7lhvct54mvWOaWJ6m4Z0kdaYtu3ti9yw3uG89TXrHNLE8TcM7SepMW3b3xO5ZbnDfepr0jmlieZqGd5LUmbbs7onds9zgvvU06R3TxPXcSxPrW07YpyR1pqTNKUmdKUmdk5LUaXzdWB/eYX3LCfuUpM6UtDklqTMlqXNSkjqNrxvrwzusbzlhn5LUmZI2pyR1piR1TkpSp/F1Y314h/UtJ+xTkjpT0uaUpM6UpM5JSeo0ft34+vOGNLqDfUosN7inhvUtJ+ybRtpMd0n39Goa3cE+JZYb3FPD+pYT9k0jbaa7pHt6NY3uYJ8Syw3uqWF9ywn7ppE2013SPb2aRnewT4nlBvfUsL7lhH3TSJvpLunecjWV7mCftqQbjSR1puR03tqSbkx/wrJqj7JPW9KNRpI6U3I6b21JN6Y/YVm1R9mnLelGI0mdKTmdt7akG9OfsKzao+zTlnSjkaTOlJzOW1vSjelPWFbp8NRImylJnWnL7r6F7zN3STcb32FppUNTI22mJHWmLbv7Fr7P3CXdbHyHpZUOTY20mZLUmbbs7lv4PnOXdLPxHZZWOjQ10mZKUmfasrtv4fvMXdLNxndYWunQlFhutHv2W42n+4bds7wl3VuuskSJ5Ua7Z7/VeLpv2D3LW9K95SpLlFhutHv2W42n+4bds7wl3VuuskSJ5Ua7Z7/VeLpv2D3LW9K97avp6GQ334X3KWlz+tukb5j+hO2/hX3Ebr4L71PS5vS3Sd8w/Qnbfwv7iN18F96npM3pb5O+YfoTtv8W9hG7+S68T0mb098mfcP0Jxz/W+x+FPethvUtN2y/m7fwnvm1+frzIOklDdy3Gta33LD9bt7Ce+bX5uvPg6SXNHDfaljfcsP2u3kL75lfm68/D5Je0sB9q2F9yw3b7+YtvGd+bb7+vCEN7ySpMzXSZrqL3bOcsN9Kns4T2uJRk6TO1Eib6S52z3LCfit5Ok9oi0dNkjpTI22mu9g9ywn7reTpPKEtHjVJ6kyNtJnuYvcsJ+y3kqfzxNLiEUosJ+xTYvkudt9yg3tqpM2d5Cf50mKJEssJ+5RYvovdt9zgnhppcyf5Sb60WKLEcsI+JZbvYvctN7inRtrcSX6SLy2WKLGcsE+J5bvYfcsN7qmRNneSn+RLK5UmbW4Sywn7lOzmhH3a0u7ZN99hadmRNjeJ5YR9SnZzwj5taffsm++wtOxIm5vEcsI+Jbs5YZ+2tHv2zXdYWnakzU1iOWGfkt2csE9b2j375jtcvTz+tuX0vrXF9sxNkjrTT+T6rvyx37ac3re22J65SVJn+olc35U/9tuW0/vWFtszN0nqTD+R67vyx37bcnrf2mJ75iZJneknUn+V/aWYUyNtpqTNqZE2UyNtGlvSjTsT9VvtKHNqpM2UtDk10mZqpE1jS7pxZ6J+qx1lTo20mZI2p0baTI20aWxJN+5M1G+1o8ypkTZT0ubUSJupkTaNLenGnYnl6TujO2zP3DTSZkp2c8L+0xppM32HpfWTIxPbMzeNtJmS3Zyw/7RG2kzfYWn95MjE9sxNI22mZDcn7D+tkTbTd1haPzkysT1z00ibKdnNCftPa6TN9B2uXh5/S9rcbEk37jR2+5SkzpSkzo4kdaavTg6/JW1utqQbdxq7fUpSZ0pSZ0eSOtNXJ4ffkjY3W9KNO43dPiWpMyWpsyNJnemrk8NvSZubLenGncZun5LUmZLU2ZGkzvTVWR/e0faJ7Xdzw/bMKbGc7PbNE1x3uqNtn9h+Nzdsz5wSy8lu3zzBdac72vaJ7Xdzw/bMKbGc7PbNE1x3uqNtn9h+Nzdsz5wSy8lu3zzBcsVewpyS1LmTWG7Y3nLCPm1JN05KLP/D8tRGzClJnTuJ5YbtLSfs05Z046TE8j8sT23EnJLUuZNYbtjecsI+bUk3Tkos/8Py1EbMKUmdO4nlhu0tJ+zTlnTjpMTyP/R/i8PwI//fJZYb3Jvv8Pd/il+WWG5wb77D3/8pflliucG9+Q5//6f4ZYnlBvfmO1y9PH7KFttbfhq+zySpMyVtbr7D1cvjp2yxveWn4ftMkjpT0ubmO1y9PH7KFttbfhq+zySpMyVtbr7D1cvjp2yxveWn4ftMkjpT0ubmO1y9ftRg9y0n7FPD+paTtk9O71sT13Mv7WD3LSfsU8P6lpO2T07vWxPXcy/tYPctJ+xTw/qWk7ZPTu9bE9dzL+1g9y0n7FPD+paTtk9O71sT1/P7EnOTWG5wb5LUmRptn3D/6b6+eX04YW4Syw3uTZI6U6PtE+4/3dc3rw8nzE1iucG9SVJnarR9wv2n+/rm9eGEuUksN7g3SepMjbZPuP90X9+8PpwwN0mb72pYfzcn1rf8NHwffXXWhxPmJmnzXQ3r7+bE+pafhu+jr876cMLcJG2+q2H93ZxY3/LT8H301VkfTpibpM13Nay/mxPrW34avo++OuvDCXOT7OZGu7e+5YT9XYnlhH36DlfvfsTcJLu50e6tbzlhf1diOWGfvsPVux8xN8lubrR761tO2N+VWE7Yp+9w9e5HzE2ymxvt3vqWE/Z3JZYT9uk7XL1+1GD3LX8avt8klhu2t5yc6F+/68OT2H3Ln4bvN4nlhu0tJyf61+/68CR23/Kn4ftNYrlhe8vJif71uz48id23/Gn4fpNYbtjecnKif/3+++HTnub0fd4zieUtvLfrO1y9PH7K05y+z3smsbyF93Z9h6uXx095mtP3ec8klrfw3q7vcPXy+ClPc/o+75nE8hbe2/Udzv9X+sv/OP/881/SqtvcdpBh+wAAAABJRU5ErkJggg==
```

不要纠结这么base64怎么转图片，直接打开`Sublime Text`新建`.html`文件，然后把上面这一串扔`<img>`标签的`src`中就能看到图片，用`QR_Research`扫描二维码就能得到`KEY{dca57f966e4e4e31fd5b15417da63269}`。

```html
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Flag</title>
</head>
<body>
	<img src="data:image/jpg;base64,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">
</body>
</html>
```

------

#### 12.[闪的好快](https://www.amanctf.com/challenges/detail/id/12.html)

使用`stegSolve`打开文件，`Analyse->Frame Browser`可以看到一共有18个图层，逐一进行扫码，将扫码结果拼接成字符串可以得到`SYC{F1aSh_so_f4sT}`。

------

#### 13.[come_game](https://www.amanctf.com/challenges/detail/id/13.html)

血压上来了，玩到第五关可以看到`flag`就是`SYC{6E23F259D98DF153}`。

------

#### 14.[白哥的鸽子](https://www.amanctf.com/challenges/detail/id/14.html)

先用`WinHex`打开这张图片，搜索不到`flag`，但是可以在最后看到一行可疑的字符串`fg2ivyo}l{2s3_o@aw__rcl@`，栅栏密码解密后可以在第5栏得到`flag{w22_is_v3ry_cool}`。

------

#### 15.[linux](https://www.amanctf.com/challenges/detail/id/15.html)

先解压`unzip`解压后可以看到有一个压缩包`1.tar.gz`。

```
tar -zxvf 1.tar.gz
```

继续解压可以得到一个`flag`文件，`cat flag`即可得到`key{feb81d3834e2423c9903f4755464060b}`。

------

#### 16.[隐写3](https://www.amanctf.com/challenges/detail/id/16.html)

使用`tweakpng`查看图片可以看到图片的高度被修改过，将图片增高可以看到`flag{He1l0_d4_ba1}`。

------

#### 17.[做个游戏](https://www.amanctf.com/challenges/detail/id/17.html)

用`jd-gui`打开`heiheihie.jar`后直接`Search`查找`flag`，可以看到`flag{RGFqaURhbGlfSmlud2FuQ2hpamk=}`，base64解密即可。

```go
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	decodeBytes, _ := base64.StdEncoding.DecodeString("RGFqaURhbGlfSmlud2FuQ2hpamk=")
	flag := "flag{" + fmt.Sprintf("%s", string(decodeBytes)) + "}"
	println(flag)
}
```

最后可以得到`flag{DajiDali_JinwanChiji}`。

------

#### 18.[想蹭网先解开密码](https://www.amanctf.com/challenges/detail/id/18.html)

Kali-Linux环境下使用命令 `crunch 11 11 -t 1391040%%%% >> bugku18.txt`创建密码字典。

![](https://paper.tanyaodan.com/amanctf/18/1.png)

使用命令 `aircrack-ng -w bugku18.txt wifi.cap`进行爆破，输入数字`3`，即可拿到手机号码`13910407686`，由题意可知`flag`是`flag{13910407686}`。

![](https://paper.tanyaodan.com/amanctf/18/2.png)

------

#### 19.[Linux2](https://www.amanctf.com/challenges/detail/id/19.html)

解压缩后用`WinHex`打开这个`brave`文件，搜索`KEY`即可找到`KEY{24f3627a86fc740a7f36ee2c7a1c124a}`。

------

#### 20.[细心的大象](https://www.amanctf.com/challenges/detail/id/20.html)

图片`1.jpg`的属性中备注了一串字符串，直接当密码输入不对。

![](https://paper.tanyaodan.com/amanctf/20/1.png)

```go
package main

import (
	"encoding/base64"
)

func main() {
	pwd, _ := base64.StdEncoding.DecodeString("TVNEUzQ1NkFTRDEyM3p6")
	println(string(pwd))
}
```

`base64`解码后拿到`.rar`压缩包的密码`MSDS456ASD123zz`，解压缩后拿到另一张图片`2.jpg`。

用`tweakpng`打开图片后，修改图片高度即可看到`BUGKU{a1e5aSA}`。

![](https://paper.tanyaodan.com/amanctf/20/2.png)

------
