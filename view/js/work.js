var API = "http://47.102.204.136:91";

function loadDemo() {
    //demoObj = document.getElementById("demo");
    infoObj = document.getElementById("information");

    var requestOptions = {
        method: 'GET'
    };

    fetch(API + "/information?user=" + GetUrlParam("user"), requestOptions)
        .then(response => response.json())
        .then((response) => {
            console.log(response);
            if (response.status == "1") {
                if (response.data[0] == undefined) {
                    demoObj.innerText = "还没有人给TA提问";
                } else {
                    //demoObj.innerText = JSON.stringify(response.data, null, 2);
                    infoObj.innerHtml = "";

                    questionLen = response.data.length;
                    for (i = 0; i < questionLen; i++) {
                        newSpan = document.createElement("span");
                        newSpan.className = "question";
                        newSpan.value = response.data[i].QuestionID;
                        newSpan.innerText = response.data[i].Question;
                        infoObj.appendChild(newSpan);

                        answerLen = response.data[i].Answer.length;
                        for (j = 0; j < answerLen; j++) {
                            newSpan = document.createElement("span");
                            newSpan.className = "answer";
                            //newSpan.value = response.data[i].Answer[j];
                            newSpan.innerText = response.data[i].Answer[j];
                            infoObj.appendChild(newSpan);

                            newBr = document.createElement("br");
                            infoObj.appendChild(newBr);
                        }
                    }

                }
            } else {
                alert("请求错误:" + response.message);
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
        body: raw
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
    myHeaders.append("Token", localStorage.getItem("qaToken"));

    var raw = JSON.stringify({
        "id": option,
        "data": content
    });

    var requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw
    };

    fetch(API + "/user/answer", requestOptions)
        .then(response => response.json())
        .then((response) => {
            console.log(response);
            if (response.status == "1") {
                loadDemo();
                alert("回复成功")
            } else {
                alert("请求错误:" + response.message + ".请尝试重新登录");
                localStorage.removeItem("qaToken");
                localStorage.removeItem("qaUsername");
                window.location.href = "./";
            }
        })
        .catch(error => console.log('error', error));
}

document.getElementById("exist").onclick = function() {
    localStorage.removeItem("qaToken");
    localStorage.removeItem("qaUsername");
    window.location.href = "./";
}

document.getElementById("copy").onclick = function() {
    shareObj = document.getElementById("shareCopy");
    shareObj.select();
    document.execCommand("Copy");
    alert("复制成功");
}

function main() {
    loadDemo();
    shareTextObj = document.getElementById("shareText");
    shareTextObj.innerText = window.location.toString();
    shareCopyObj = document.getElementById("shareCopy");
    shareCopyObj.innerText = window.location.toString();
}

main();