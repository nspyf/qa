var API = "47.102.204.136:91";

document.getElementById("register").onclick = function() {
    username = document.getElementById("username").value;
    password = document.getElementById("password").value;
    passwordRep = document.getElementById("passwordRep").value;
    if (password != passwordRep) {
        alert("两次密码不一样")
        return
    }

    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");


    var raw = JSON.stringify({
        "username": username,
        "password": password
    });

    var requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        mode: 'cors'
    };

    fetch(API + "/register", requestOptions)
        .then(response => response.json())
        .then((response) => {
            console.log(response);
            if (response.status == "1") {
                alert("注册成功");
                window.location.href = "./index.html";
            } else {
                alert("请求错误:" + response.message);
            }
        })
        .catch(error => console.log('error', error));
}