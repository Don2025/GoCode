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

查看图片属性可以找到相机型号`73646E6973635F32303138`，将后缀改为`.rar`可以看到压缩包中有`flag.txt`，提取需要输入解压密码，直接输入相机型号不对，仔细观察发现这一串是16进制字符串。

```go
package main

import "encoding/hex"

// hexToStr converts Hex code to text string.
func hexToStr(num string) string {
    str, _ := hex.DecodeString(num)
    return string(str)
}

func main() {
    println(hexToStr("73646E6973635F32303138"))
}
```

16进制转字符串得到解压密码`sdnisc_2018`，解压后打开`flag.txt`看到`flag{3XiF_iNf0rM@ti0n}`。

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
data:image/jpg;base64,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
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

#### 21.[爆照](https://www.amanctf.com/challenges/detail/id/21.html)

`Kali-Linux`环境下用`binwalk`扫描图片。

![](https://paper.tanyaodan.com/amanctf/21/1.png)

接着`foremost -i file.jpg`在`output`中可以看到分离出了一张图片和一个`.zip`压缩包。

在`Windows`环境下打开这个`.zip`文件，里面有很多文件，全部放进`Kali-Linux`中发现全是图片，`binwalk`的运行结果如下：

```
┌──(tyd㉿kali-linux)-[~/ctf]
└─$ binwalk 8          

DECIMAL       HEXADECIMAL     DESCRIPTION
--------------------------------------------------------------------------------
0             0x0             PC bitmap, Windows 3.x format,, 303 x 300 x 8

┌──(tyd㉿kali-linux)-[~/ctf]
└─$ binwalk 88

DECIMAL       HEXADECIMAL     DESCRIPTION
--------------------------------------------------------------------------------
0             0x0             JPEG image data, JFIF standard 1.01
30            0x1E            TIFF image data, big-endian, offset of first image directory: 8

┌──(tyd㉿kali-linux)-[~/ctf]
└─$ binwalk 888

DECIMAL       HEXADECIMAL     DESCRIPTION
--------------------------------------------------------------------------------
0             0x0             JPEG image data, JFIF standard 1.01
30            0x1E            TIFF image data, big-endian, offset of first image directory: 8

┌──(tyd㉿kali-linux)-[~/ctf]
└─$ binwalk 8888

DECIMAL       HEXADECIMAL     DESCRIPTION
--------------------------------------------------------------------------------
0             0x0             JPEG image data, JFIF standard 1.01
30            0x1E            TIFF image data, big-endian, offset of first image directory: 8
10976         0x2AE0          Zip archive data, at least v2.0 to extract, compressed size: 644, uncompressed size: 1202, name: 1509126368.png
11760         0x2DF0          End of Zip archive, footer length: 22

┌──(tyd㉿kali-linux)-[~/ctf]
└─$ binwalk 88888

DECIMAL       HEXADECIMAL     DESCRIPTION
--------------------------------------------------------------------------------
0             0x0             PC bitmap, Windows 3.x format,, 303 x 300 x 8

┌──(tyd㉿kali-linux)-[~/ctf]
└─$ binwalk 888888

DECIMAL       HEXADECIMAL     DESCRIPTION
--------------------------------------------------------------------------------
0             0x0             PC bitmap, Windows 3.x format,, 303 x 300 x 8

┌──(tyd㉿kali-linux)-[~/ctf]
└─$ binwalk 8888888

DECIMAL       HEXADECIMAL     DESCRIPTION
--------------------------------------------------------------------------------
0             0x0             PC bitmap, Windows 3.x format,, 303 x 300 x 8       

┌──(tyd㉿kali-linux)-[~/ctf]
└─$ binwalk 88888888

DECIMAL       HEXADECIMAL     DESCRIPTION
--------------------------------------------------------------------------------
0             0x0             PC bitmap, Windows 3.x format,, 303 x 300 x 8
```

88，888，8888这三个文件可疑，`88.jpg`直接把二维码写脸上了，用`QR_Research`扫描出来是`bilibili`，得到了第一个信息。

`888.jpg`的属性中详情信息备注了`c2lsaXNpbGk=`，`base64`解码后得到`silisili`，这是第二个信息。

`foremost 8888`可以分离出一个`.zip`文件，打开里面有张二维码图片，用`QR_Research`扫描出来是`panama`，第三个信息到手。

根据题意拼接一下就能得到`flag{bilibili_silisili_panama}`。

------

#### 68.[滑稽](https://www.amanctf.com/challenges/detail/id/68.html)

越来越快的滑稽脸搞起眼睛好痛啊艹，`view-source`查看源码，可以直接找到`flag{e0ba515db4cea63ac32e3f005b87a7fe}`。

```html
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<meta name="viewport" content="width=device-width,hight=device-hight,minimum-scale=1.0,maximum-scale=1.0,ser-scalable=none"/>
<title>BugkuCTF-WEB1</title>

<style type="text/css">
  body { 
    margin: 0; 
    padding: 0; 
    position: relative; 
    background-image: url(images/xh.jpg); 
    background-position: center; 
    /*background-repeat: no-repeat;*/ 
    width: 100%; height: 100%; 
    background-size: 100% 100%; 
 }
</style>

</head>
<body id="body" onLoad="init()">
<!flag{e0ba515db4cea63ac32e3f005b87a7fe}>
<script type="text/javascript" src="js/ThreeCanvas.js"></script>
<script type="text/javascript" src="js/Snow.js"></script>

<script type="text/javascript">
	var SCREEN_WIDTH = window.innerWidth;//
	var SCREEN_HEIGHT = window.innerHeight;
	var container;
	var particle;//粒子
	var camera;
	var scene;
	var renderer;
	var starSnow = 1;
	var particles = []; 
	var particleImage = new Image();
	//THREE.ImageUtils.loadTexture( "img/ParticleSmoke.png" );
	particleImage.src = 'images/funny.png';

	function init() {
		//alert("message3");
		container = document.createElement('div');//container：画布实例;
		document.body.appendChild(container);
		camera = new THREE.PerspectiveCamera( 60, SCREEN_WIDTH / SCREEN_HEIGHT, 1, 10000 );
		camera.position.z = 1000;
		//camera.position.y = 50;
		scene = new THREE.Scene();
		scene.add(camera);
		renderer = new THREE.CanvasRenderer();
		renderer.setSize(SCREEN_WIDTH, SCREEN_HEIGHT);
		var material = new THREE.ParticleBasicMaterial( { map: new THREE.Texture(particleImage) } );
			//alert("message2");
		for (var i = 0; i < 500; i++) {
			//alert("message");
			particle = new Particle3D( material);
			particle.position.x = Math.random() * 2000-1000;
			
			particle.position.z = Math.random() * 2000-1000;
			particle.position.y = Math.random() * 2000-1000;
			//particle.position.y = Math.random() * (1600-particle.position.z)-1000;
			particle.scale.x = particle.scale.y =  1;
			scene.add( particle );
			particles.push(particle); 
		}
		container.appendChild( renderer.domElement );
		//document.addEventListener( 'mousemove', onDocumentMouseMove, false );
		document.addEventListener( 'touchstart', onDocumentTouchStart, false );
		document.addEventListener( 'touchmove', onDocumentTouchMove, false );
		document.addEventListener( 'touchend', onDocumentTouchEnd, false );
		setInterval( loop, 1000 / 30 );
	}
	
	var touchStartX;
	var touchFlag = 0;//储存当前是否滑动的状态;
	var touchSensitive = 80;//检测滑动的灵敏度;
	//var touchStartY;
	//var touchEndX;
	//var touchEndY;
	function onDocumentTouchStart( event ) {
		if ( event.touches.length == 1 ) {
			event.preventDefault();//取消默认关联动作;
			touchStartX = 0;
			touchStartX = event.touches[ 0 ].pageX ;
			//touchStartY = event.touches[ 0 ].pageY ;
		}
	}

	function onDocumentTouchMove( event ) {
		if ( event.touches.length == 1 ) {
			event.preventDefault();
			var direction = event.touches[ 0 ].pageX - touchStartX;
			if (Math.abs(direction) > touchSensitive) {
				if (direction>0) {touchFlag = 1;}
				else if (direction<0) {touchFlag = -1;};
				//changeAndBack(touchFlag);
			}	
		}
	}

	function onDocumentTouchEnd (event) {
		// if ( event.touches.length == 0 ) {
		// 	event.preventDefault();
		// 	touchEndX = event.touches[ 0 ].pageX ;
		// 	touchEndY = event.touches[ 0 ].pageY ;
		// }这里存在问题
		var direction = event.changedTouches[ 0 ].pageX - touchStartX;
		changeAndBack(touchFlag);
	}

	function changeAndBack (touchFlag) {
		var speedX = 25*touchFlag;
		touchFlag = 0;
		for (var i = 0; i < particles.length; i++) {
			particles[i].velocity=new THREE.Vector3(speedX,-10,0);
		}
		var timeOut = setTimeout(";", 800);
		clearTimeout(timeOut);
		var clearI = setInterval(function () {
			if (touchFlag) {
				clearInterval(clearI);
				return;
			};
			speedX*=0.8;
			if (Math.abs(speedX)<=1.5) {
				speedX=0;
				clearInterval(clearI);
			};
			for (var i = 0; i < particles.length; i++) {
				particles[i].velocity=new THREE.Vector3(speedX,-10,0);
			}
		},100);
	}

	function loop() {
		for(var i = 0; i<particles.length; i++){
			var particle = particles[i]; 
			particle.updatePhysics(); 
			with(particle.position)
			{
				if((y<-1000)&&starSnow) {y+=2000;}
				if(x>1000) x-=2000; 
				else if(x<-1000) x+=2000;
				if(z>1000) z-=2000; 
				else if(z<-1000) z+=2000;
			}			
		}
		camera.lookAt(scene.position); 
		renderer.render( scene, camera );
	}
</script>
</body>
</html>
```

------

#### 69.[计算器](https://www.amanctf.com/challenges/detail/id/69.html)

进入靶机后直接`F12`查看源代码里面的源，点开`js`，打开里面的`code.js`里面躺着`flag`。

------

#### 70.[GET](https://www.amanctf.com/challenges/detail/id/70.html)

靶机给出的信息如下：

```php
$what=$_GET['what'];
echo $what;
if($what=='flag')
echo 'flag{****}';
```

直接`?what=flag`无脑访问http://114.67.246.176:15470/?what=flag。

------

#### 71.[POST](https://www.amanctf.com/challenges/detail/id/71.html)

靶机给出的信息如下：

```php
$what=$_POST['what'];
echo $what;
if($what=='flag')
echo 'flag{****}';
```

直接无脑`curl`就完事了。

```bash
> curl -X POST --data "what=flag" 114.67.246.176:16240
$what=$_POST['what'];<br>
echo $what;<br>
if($what=='flag')<br>
echo 'flag{****}';<br>

flagflag{fb03903b94321f1dc330c44fcdd0cd42}
```

------

#### 72.[矛盾](https://www.amanctf.com/challenges/detail/id/72.html)

靶机给出的信息如下：

```php
$num=$_GET['num'];
if(!is_numeric($num))
{
echo $num;
if($num==1)
echo 'flag{**********}';
}
```

`PHP`是弱类型语言，直接无脑`?num=1a`就行了，等号后面的字符串第一位是数字1，之后随便加上非数字即可拿到`flag`。

------

#### 73.[alert](https://www.amanctf.com/challenges/detail/id/73.html)

这个题会不停地`alert`弹出提示框，靶机的代码如下：

```html
<html>  
<head>  
<title>BKCTF-WEB6</title>  
<script language="javascript">   
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
alert("flag就在这里");  
alert("来找找吧"); 
<!-- &#102;&#108;&#97;&#103;&#123;&#101;&#48;&#48;&#97;&#97;&#101;&#48;&#56;&#97;&#51;&#48;&#99;&#52;&#100;&#100;&#102;&#55;&#98;&#52;&#102;&#98;&#52;&#52;&#49;&#99;&#52;&#100;&#53;&#54;&#54;&#54;&#99;&#125; --></script> 
</head>  
</html>  
```

直接无脑`view-source:http://114.67.246.176:13088/`在最下面可以看到注释，这是`Unicode`编码，可以用Go语言自建函数将其转换成`ASCII`码就可以得到`flag{e00aae08a30c4ddf7b4fb441c4d5666c}`。

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// UnicodeToASCII converts Unicode to ASCII
func UnicodeToASCII(text string) string {
	//Remove the last semicolon, then cut it into a string array according to the semicolon
	arr := strings.Split(strings.Trim(text, ";"), ";")
	var ans string
	for _, x := range arr {
		n, _ := strconv.Atoi(strings.Trim(x, "&#"))
		ans += fmt.Sprintf("%c",n)
	}
	return ans
}

func main() {
    text := "&#102;&#108;&#97;&#103;&#123;&#101;&#48;&#48;&#97;&#97;&#101;&#48;&#56;&#97;&#51;&#48;&#99;&#52;&#100;&#100;&#102;&#55;&#98;&#52;&#102;&#98;&#52;&#52;&#49;&#99;&#52;&#100;&#53;&#54;&#54;&#54;&#99;&#125;"
	println(UnicodeToASCII(text))
}
```

------

#### 74.[你必须让他停下](https://www.amanctf.com/challenges/detail/id/74.html)

用`BurpSuite`收到的包发送到`Repeater`再发送查看`Response`，可以直接看到`flag{877af93bba7b315cb3af1cfc208878d2}`。

![](https://paper.tanyaodan.com/amanctf/74/1.png)

------

#### 75.[eval](https://www.amanctf.com/challenges/detail/id/75.html)



------

#### 76.[变量1](https://www.amanctf.com/challenges/detail/id/76.html)

靶机的信息如下：

```php
flag In the variable ! <?php
error_reporting(0);       //关闭php错误显示
include "flag1.php";      //引入flag1.php文件代码
highlight_file(__file__);
if(isset($_GET['args'])){  //通过get方式传递args变量才能执行
    $args = $_GET['args'];
    //这个正则表达式的意思是匹配任意[A-Za-z0-9_]的字符，就是任意大小写字母和0到9以及下划线组成
    if(!preg_match("/^\w+$/",$args)){
        die("args error!");
    }
    eval("var_dump($$args);"); //eval中引号里面的php代码按正常的php代码被执行
}
?>
```

`$args`很有可能是一个数组，那么应该想到的就是超全局变量`$GLOBALS`，全局变量的值在这个超级全局变量里面是一个键值，可以通过变量名在`$GLOBALS`找到相对应的值。所以这题通过构造一个`GET`参数，直接传`GET`一个全局变量即可，无脑访问`http://123.206.87.240:8004/index1.php?args=GLOBALS`。

------

#### 77.[头等舱](https://www.amanctf.com/challenges/detail/id/77.html)

`Burp Suite`抓包后`Repeater`发送请求就能在`Response`的响应头中看到`flag{1cc4cfd3b610b5c18b7202717e246240}`。

------

#### 78.[网站被黑](https://www.amanctf.com/challenges/detail/id/78.html)

`Kali-Linux`环境下使用命令行`dirsearch -u http://114.67.246.176:11241/`扫描网站的后台路径。

![](https://paper.tanyaodan.com/amanctf/78/1.png)

访问`http://114.67.246.176:11241/shell.php`可以看到一个需要输入密码的登录框，随便试了几个密码错误啦。

![](https://paper.tanyaodan.com/amanctf/78/2.png)

使用`Burp Suite`抓包后`Send to Intruder`，导入常用字典后开始爆破，得到密码`hack`，登录后得到`flag`。

![](https://paper.tanyaodan.com/amanctf/78/3.png)

------

#### 79.[本地管理员](https://www.amanctf.com/challenges/detail/id/79.html)

把源代码中的`dGVzdDEyMw==`进行`base64`解码后得到`test123`，输入`admin`和`test123`后提示需要本地浏览，添加`X-Forwarded-For: 127.0.0.1`即可拿到`flag`。

![](https://paper.tanyaodan.com/amanctf/79/1.png)

------

#### 80.[源代码](https://www.amanctf.com/challenges/detail/id/80.html)

查看源代码发现靶机的信息如下：

```html
<html>
<title>BUGKUCTF-WEB13</title>
<body>
<div style="display:none;"></div>
<form action="index.php" method="post" >
看看源代码？<br>
<br>
<script>
var p1 = '%66%75%6e%63%74%69%6f%6e%20%63%68%65%63%6b%53%75%62%6d%69%74%28%29%7b%76%61%72%20%61%3d%64%6f%63%75%6d%65%6e%74%2e%67%65%74%45%6c%65%6d%65%6e%74%42%79%49%64%28%22%70%61%73%73%77%6f%72%64%22%29%3b%69%66%28%22%75%6e%64%65%66%69%6e%65%64%22%21%3d%74%79%70%65%6f%66%20%61%29%7b%69%66%28%22%36%37%64%37%30%39%62%32%62';
var p2 = '%61%61%36%34%38%63%66%36%65%38%37%61%37%31%31%34%66%31%22%3d%3d%61%2e%76%61%6c%75%65%29%72%65%74%75%72%6e%21%30%3b%61%6c%65%72%74%28%22%45%72%72%6f%72%22%29%3b%61%2e%66%6f%63%75%73%28%29%3b%72%65%74%75%72%6e%21%31%7d%7d%64%6f%63%75%6d%65%6e%74%2e%67%65%74%45%6c%65%6d%65%6e%74%42%79%49%64%28%22%6c%65%76%65%6c%51%75%65%73%74%22%29%2e%6f%6e%73%75%62%6d%69%74%3d%63%68%65%63%6b%53%75%62%6d%69%74%3b';
eval(unescape(p1) + unescape('%35%34%61%61%32' + p2));
</script>

<input type="input" name="flag" id="flag" /> 
<input type="submit" name="submit" value="Submit" />
</form>
```

将`p1`+`%35%34%61%61%32`+`p2`拼接而成的字符串进行URL解码可以得到一串字符串`67d709b2b54aa2aa648cf6e87a7114f1`，输入后点击`Submit`即可得到真实的`flag`：`flag{2aeef9f603806011a66df8360330b0a3}`。

------
