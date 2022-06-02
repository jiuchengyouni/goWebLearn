let e = document.getElementById('nickname');
let a = document.getElementById("password");
let b = document.getElementById("repeatpsd");
var c = document.getElementById("protocal");
console.log(c.value);
var x = 0;
var y = 0;

e.onkeyup = () => {
    console.log(e.value);
    if ((/^[A-Za-z0-9-_]+$/g.test(e.value)) && (a.value == b.value)) {
        document.getElementById("result").innerHTML = "";
        document.getElementById("username").value = String(e.value) + "@stu.xmut.edu.cn";
    }
    else if (e.value == '' || e.value == undefined || e.value == null) {
        document.getElementById("result").innerHTML = "请输入用户名";
        document.getElementById("username").value = "";
    }
    else if ((/^[A-Za-z0-9-_]+$/g.test(e.value)) && (a.value != b.value)) {
        document.getElementById("result").innerHTML = "两次输入的密码不匹配";
    }
    else {
        document.getElementById("result").innerHTML = "用户名只能包含数字，字母和下划线";
        document.getElementById("username").value = "";
    }
}



a.onkeyup = () => {
    console.log(a.value);
    if (a.value == '' || a.value == undefined || a.value == null) {
        if (e.value == '' || e.value == undefined || e.value == null) {
            document.getElementById("result").innerHTML = "请输入用户名";
        }
        else if (/^[A-Za-z0-9-_]+$/g.test(e.value)) {
            document.getElementById("result").innerHTML = "请输入密码";
            b.style.borderColor='#FF0000';
        }
        else {
            document.getElementById("result").innerHTML = "用户名只能包含数字，字母和下划线";
        }
    }
    else if (a.value != b.value) {
        document.getElementById("result").innerHTML = "两次输入的密码不匹配";
        b.style.borderColor = '#FF0000';
    }
    else if (a.value == b.value) {
        y += 1;
        if (y == 1) {
            document.getElementById("result").innerHTML = "请勾选同意用户协议";
        }
        document.getElementById("result").innerHTML = "";
        repeatpsd.style.borderColor = '#D3D3D3'
    }
    b.onkeyup = () => {
        console.log(b.value);
        if (a.value == b.value) {
            if (e.value == '' || e.value == undefined || e.value == null) {
                document.getElementById("result").innerHTML = "请输入用户名";
                repeatpsd.style.borderColor = '#D3D3D3';
            }
            else if (/^[A-Za-z0-9-_]+$/g.test(e.value)) {
                repeatpsd.style.borderColor = '#D3D3D3';
                document.getElementById("result").innerHTML = "";
                y += 1;
                if (y == 1) {
                    document.getElementById("result").innerHTML = "请勾选同意用户协议";
                }
                document.getElementById("protocal").onclick = function () {
                    x += 1;
                    if (x % 2 == 1 && a.value == b.value) {
                        document.getElementById("result").innerHTML = "";
                    }
                    else if (x % 2 == 0 && a.value == b.value) {
                        document.getElementById("result").innerHTML = "请勾选同意用户协议";
                    }
                }
            }
            else {
                repeatpsd.style.borderColor = '#D3D3D3';
                document.getElementById("result").innerHTML = "用户名只能包含数字，字母和下划线";
            }
        }
        else if (a.value != b.value) {
            if (e.value == '' || e.value == undefined || e.value == null) {
                document.getElementById("result").innerHTML = "请输入用户名";
            }
            else if (/^[A-Za-z0-9-_]+$/g.test(e.value)) {
                document.getElementById("result").innerHTML = "两次输入的密码不匹配";
                repeatpsd.style.borderColor = '#FF0000';
            }
            else {
                document.getElementById("result").innerHTML = "用户名只能包含数字，字母和下划线";
            }
        }
    }
}
