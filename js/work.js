var API = "http://47.102.204.136:91";

function loadDemo() {
    demoObj = document.getElementById("demo");

    var requestOptions = {
        method: 'GET',
        mode: 'cors'
    };

    fetch(API + "/information?user=" + GetUrlParam("user"), requestOptions)
        .then(response => response.json())
        .then((response) => {
            console.log(response);
            if (response.data[0] == undefined) {
                demoObj.innerText = "还没有人给TA提问";
            } else {
                demoObj.innerText = JSON.stringify(response.data, null, 2);
            }

        })
        .catch(error => console.log('error', error));
}

function GetUrlParam(paraName) {
    var url = document.location.toString();
    var arrObj = url.split("?");
    if (arrObj.length > 1) {
        var arrPara = arrObj[1].split("&");
        var arr;
        for (var i = 0; i < arrPara.length; i++) {
            arr = arrPara[i].split("=");
            if (arr != null && arr[0] == paraName) {
                return arr[1];
            }
        }
        return "";
    } else {
        return "";
    }
}

document.getElementById("ask").onclick = function() {
    content = document.getElementById("content").value;

    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    var raw = JSON.stringify({
        "username": GetUrlParam("user"),
        "data": content
    });

    var requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        mode: 'cors'
    };

    fetch(API + "/question", requestOptions)
        .then(response => response.json())
        .then((response) => {
            console.log(response);
            if (response.status == "1") {
                loadDemo();
                alert("提问成功");
            } else {
                alert("请求错误:" + response.message);
            }
        })
        .catch(error => console.log('error', error));
}

document.getElementById("respond").onclick = function() {
    option = document.getElementById("option").value;
    content = document.getElementById("content").value;

    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");
    myHeaders.append("Token", localStorage.getItem("aqToken"));

    var raw = JSON.stringify({
        "id": option,
        "data": content
    });

    var requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        mode: 'cors'
    };

    fetch(API + "/user/answer", requestOptions)
        .then(response => response.json())
        .then((response) => {
            console.log(response);
            if (response.status == "1") {
                loadDemo();
                alert("回复成功")
            } else {
                alert("请求错误:" + response.message);
            }
        })
        .catch(error => console.log('error', error));
}

loadDemo()

shareObj = document.getElementById("share");
shareObj.innerText = window.location.toString();