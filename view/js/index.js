var API = "http://47.102.204.136:91";
//var API = "http://127.0.0.1:91";

document.getElementById("login").onclick = function() {
    username = document.getElementById("username").value;
    password = document.getElementById("password").value;

    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    var raw = JSON.stringify({
        "username": username,
        "password": password
    });

    var requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw
    };

    fetch(API + "/login", requestOptions)
        .then(response => response.json())
        .then((response) => {
            if (response.status == "1") {
                localStorage.setItem("qaToken", response.data.token)
                localStorage.setItem("qaUsername", username)
                alert("登录成功");
                window.location.href = "./user?user=" + username;
            } else {
                alert("请求错误:" + response.message);
            }
        })
        .catch(error => console.log('error', error));
}

function tokenVerify() {
    if (localStorage.getItem("qaToken") == null || localStorage.getItem("qaUsername") == null) {
        return;
    }

    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");
    myHeaders.append("Token", localStorage.getItem("qaToken"));

    var requestOptions = {
        method: 'GET',
        headers: myHeaders,
    };

    fetch(API + "/verification", requestOptions)
        .then(response => response.json())
        .then((response) => {
            if (response.status == "1") {
                alert("欢迎回来！");
                window.location.href = "./user?user=" + localStorage.getItem("qaUsername");
            } else {
                alert("令牌失效：" + response.message + ".请重新登录");
                return;
            }
        })
        .catch(error => console.log('error', error));
}

function main() {
    tokenVerify();
}

main();